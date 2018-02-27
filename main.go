package main

import (
    "fmt"
    "github.com/chronicall/jukeboxx/song"
    "github.com/chronicall/jukeboxx/songqueue"
)

func JukeBox() {
    // if no song is playing, find a song and play it
    // find song in queue with highest votes that has not been played 
    // in the last 30 minutes
    // play it
    // s.UpdateLastPlayed(time.Now())

    // if song is playing, poll the queue for the next one to be ready
}

func Guest(songName string) {
    // Find wait time with random
    // A rate of 0.5 means vote cast every 2 seconds
    // A rate of 2 means vote cast twice per second
}

func main() {
    s := songqueue.SongQueue{[]song.Song{}}
    s.Init()

    go JukeBox()

    fmt.Println("This is the jukeboxxxxxx")
}
