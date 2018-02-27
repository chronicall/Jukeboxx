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

func JukeBox(s songqueue.SongQueue) {
    // if no song is playing, find a song and play it
    // find song in queue with highest votes that has not been played 
    // in the last 30 minutes
    // play it
    // s.UpdateLastPlayed()
    time.Sleep(time.Second * 4)
    s.Lock.RLock()
    for _, element := range s.Songs {
        element.SongInfo()
    }
    // if song is playing, poll the queue for the next one to be ready
}

func Guest(songName string, s songqueue.SongQueue) {
    // Find wait time with random
    // A rate of 0.5 means vote cast every 2 seconds
    // A rate of 2 means vote cast twice per second
    i := 1
    for i < 4 {
        s.VoteForSong(songName)
        fmt.Printf("Voted for song: %s\n", songName)
        i = i + 1
    }
}

func main() {
    fmt.Println("This is the jukeboxxxxxx")

    songList, err := os.Open("songlist.txt")
    defer songList.Close()
    if err != nil {
        panic(err)
    }

    var songListSlice [][]string
    songScanner := bufio.NewScanner(songList)
    for songScanner.Scan() {
        line := songScanner.Text()
        split := strings.Split(line, ",")
        songListSlice = append(songListSlice, split)
    }

    s := songqueue.SongQueue{[]song.Song{}, sync.RWMutex{}}
    s.Init(songListSlice)

    // Main go routine that is the jukebox that plays the songs
    go JukeBox(s)

    // Make goroutines where each guest has a favourite song on the list
    // Maybe change it to more songs and an x number of guests who has a random
    // favourite song
    for _, element := range songListSlice {
        go Guest(element[0], s)
    }

    for {
    }
}

