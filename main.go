package main

import (
    "bufio"
    "fmt"
    "os"
    "./song"
    "./songqueue"
    "strings"
    "sync"
)

func JukeBox(s songqueue.SongQueue) {
    // if no song is playing, find a song and play it
    // find song in queue with highest votes that has not been played 
    // in the last 30 minutes
    // play it
    // s.UpdateLastPlayed(time.Now())

    for _, element := range s.Songs {
        element.SongInfo()
    }
    // if song is playing, poll the queue for the next one to be ready
}

func Guest(songName string) {
    // Find wait time with random
    // A rate of 0.5 means vote cast every 2 seconds
    // A rate of 2 means vote cast twice per second
    fmt.Println(songName)
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
        go Guest(element[0])
    }

    for {
    }
}

