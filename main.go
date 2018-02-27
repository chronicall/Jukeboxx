package main

import (
    "bufio"
    "fmt"
    "flag"
    "math/rand"
    "os"
    "./song"
    "./songqueue"
    "strings"
    "sync"
    "time"
)

// Our main go routine that handles finding the next song to play in the queue.
func JukeBox(s songqueue.SongQueue) {
    for {
        time.Sleep(time.Second * 2)
        fmt.Println()
        fmt.Println()

        s.Lock.RLock()
        // Sort the queue before printing out info about the songs
        s.SortQueue()
        // Print out info about the songs in the jukebox
        // Maybe not required for final handin, more for debug
        // PRINTS OUT UNSORTED QUEUE AFTER A SONG IS PLAYED
        for _, element := range s.Songs {
            element.SongInfo()
        }
        s.Lock.RUnlock()
        // Find next song to play, we will always find a song because of how
        // SongQueue.GetNextSong() is implemented, if no song is found we get
        // a random song to play
        requested := s.GetNextSong()

        // Play the song and update it's statistics
        s.Songs[requested].Play()

        // Sleep for the duration of the song being played
        //duration := s.Songs[requested].GetDuration()
        //time.Sleep(duration)

        fmt.Println()

    }
}

// Our goroutine for guests to vote for songs
// Currently just vote for all songs 3 times
func Guest(songName string, s songqueue.SongQueue, lambda float64) {
    // Vote for the same song 3 times
    for{
        s.VoteForSong(songName)
        time.Sleep(time.Second * time.Duration((rand.ExpFloat64() / lambda)))
    }
    // TODO: Find wait time with random
    // A rate of lambda = 0.5 means vote cast ON AVERAGE every 2 seconds
    // A rate of lambda = 2 means vote cast ON AVERAGE twice per second
    // Note from Marcel: use ExpFloat64()/lambda
}

func main() {
    guests := flag.Int("guests", 12, "an integer, the number of guests")

    flag.Parse()

    fmt.Println("This is the jukeboxxxxxx")

    // Open the file songlist.txt
    songList, err := os.Open("songlist.txt")
    // This makes it so that the file handle is closed eventually
    defer songList.Close()
    // PANIC if we get an error
    if err != nil {
        panic(err)
    }

    // 2 dimensional slice of songs:
    //      [[songName1, songDuration1],
    //       [songName2, songDuration2]
    //       .
    //       .]
    var songListSlice [][]string

    // Scanner to read over the file
    songScanner := bufio.NewScanner(songList)
    for songScanner.Scan() {
        // Get next line token in the file
        line := songScanner.Text()
        // split the line from file into array [name, duration]
        split := strings.Split(line, ",")
        // append to slice
        songListSlice = append(songListSlice, split)
    }

    // Initialize our song queue
    s := songqueue.SongQueue{[]song.Song{}, sync.RWMutex{}}
    s.Init(songListSlice)

    // Main go routine that is the jukebox that plays the songs
    // Start that here
    go JukeBox(s)

    // Make goroutines where each guest has a favourite song on the list
    // Maybe change it to more songs and an x number of guests who has a random
    // favourite song
    listLen := len(songListSlice)
    for i := 0; i < *guests; i++ {
        lambda := (rand.Float64() * 3) + 0.2
        go Guest(songListSlice[rand.Intn(listLen)][0], s, lambda)
    }

    // endless loop until the jukebox is shut down
    // Maybe add some shutdown sequence/quit symbol
    // Would maybe block.. idk
    for {
    }
}

