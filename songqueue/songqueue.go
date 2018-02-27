package songqueue

import (
    "fmt"
    "../song"
    "sort"
    "time"
    "sync"
)

type SongQueue struct {
    Songs []song.Song
    Lock sync.RWMutex
}


// initialize the list of songs
// Loop over song names and durations (store in file that we read?)
func (s *SongQueue) Init() {
    songName, songDuration := "tmp", "3m40s"
    s.AddSong(songName, songDuration)
}

// Creates a Song object that we then insert into SongQueue
// params:
//  songName = "Name of Song"
//  songDuration = "3m40s"
func (s *SongQueue) AddSong(songName, songDuration string) {
    dur, _ := time.ParseDuration(songDuration)
    newSong := song.Song{
        Name: songName,
        Votes: 0,
        Duration: dur,
        LastPlayed: time.Time{},
    }
    // lock the queue
    s.Songs = append(s.Songs, newSong)
    // unlock the queue
}

func (s SongQueue) SortQueue() {
    // lock the queue
    s.Lock.Lock()
    sort.Sort(song.ByVotes(s.Songs))
    // Unlock the queue
    s.Lock.Lock()
}

// Return index of song in the queue that has the highest votes
// AND has not been played in the last 30 minutes
func (s *SongQueue) GetNextSong() *song.Song {
    now := time.Now()
    s.SortQueue()
    var requested *song.Song
    // read lock queue
    for _, element := range s.Songs {
        difference := now.Sub(element.LastPlayed)
        if difference.Minutes() > 30 {
            requested = &element
            break
        }
    }
    // read unlock queue
    return requested
}

// Find a song in the queue by name
func (s *SongQueue) GetSong(songName string) *song.Song {
    var requested *song.Song

    // read lock queue
    for _, element := range s.Songs {
        if element.Name == songName {
            requested = &element
            break
        }
    }
    // read unlock queue
    return requested
}

// Vote for a song
func (s *SongQueue) VoteForSong(name string) {
    // lock the queue
    song := s.GetSong(name)
    if song != nil {
       song.VoteForSong()
    } else {
        fmt.Println("Could not find song and therefore unable to register vote.")
    }
    // unlock the queue
}

