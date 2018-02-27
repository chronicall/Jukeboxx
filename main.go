package main

import (
    "bufio"
    "fmt"
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
        time.Sleep(time.Second * 4)
        // Find next song to play
        requested := s.GetNextSong()
        // If we find a song, update it's statistics and play it
        if requested >= 0 {
            s.Songs[requested].Play()
            // Sleep for the duration of the song being played
            //duration := s.Songs[requested].GetDuration()
            //time.Sleep(duration)
        } else {
            fmt.Println("No song found :C")
            fmt.Println("This will be handled by playing a random song")
        }

        fmt.Println()
        fmt.Println()
        // Print out info about the songs in the jukebox
        // Maybe not required for final handin, more for debug
        // PRINTS OUT UNSORTED QUEUE AFTER A SONG IS PLAYED
        s.Lock.RLock()
        for _, element := range s.Songs {
            element.SongInfo()
        }
        s.Lock.RUnlock()
    }
    // if song is playing, poll the queue for the next one to be ready??
}

// Our goroutine for guests to vote for songs
// Currently just vote for all songs 3 times
func Guest(songName string, s songqueue.SongQueue, lambda float32) {
    // Vote for the same song 3 times
    i := 1
    for i < 4 {
        s.VoteForSong(songName)
        i = i + 1
    }
    // TODO: Find wait time with random
    // A rate of lambda = 0.5 means vote cast ON AVERAGE every 2 seconds
    // A rate of lambda = 2 means vote cast ON AVERAGE twice per second
    // Note from Marcel: use ExpFloat64()/lambda
}

func main() {
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
    for _, element := range songListSlice {
        go Guest(element[0], s, 1.1)
    }
    // endless loop until the jukebox is shut down
    // Maybe add some shutdown sequence/quit symbol
    // Would maybe block.. idk
    for {
    }
}

