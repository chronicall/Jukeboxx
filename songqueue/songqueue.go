package songqueue

import (
    "fmt"
    "../song"
    "sort"
    "time"
    "sync"
)

// The struct for our queue of songs
// Has a slice of songs and a Read Write mutex
type SongQueue struct {
    Songs []song.Song
    Lock sync.RWMutex
}

// SongQueue.Init(songList)
//      initialize the list of songs
func (s *SongQueue) Init(songList [][]string) {
    // Loop over songList:
    //      song = [name, duration]
    for _, song := range songList {
        // Add the song to the queue.
        s.AddSong(song[0], song[1])
        fmt.Printf("Added song:\n\t%s --- %s\n", song[0], song[1])
    }
}

// SongQueue.AddSong(songName, songDuration)
//      Creates a Song object that we then insert into SongQueue
//      params:
//          songName = "Name of Song"
//          songDuration = "3m40s"
func (s *SongQueue) AddSong(songName, songDuration string) {
    // Parse songDuration string into time.Duration
    dur, _ := time.ParseDuration(songDuration)
    // initialize a song with the songName, 0 votes, duration and
    // null value of time.Time
    newSong := song.Song{
        Name: songName,
        Votes: 0,
        Duration: dur,
        LastPlayed: time.Time{},
    }
    // Lock the queue while we add a new song to it
    // Maybe not needed as we only do this once?
    s.Lock.Lock()
    s.Songs = append(s.Songs, newSong)
    s.Lock.Unlock()
}

// SongQueue.SortQueue()
//      Sorts the queue by song votes in decreasing order.
func (s SongQueue) SortQueue() {
    s.Lock.Lock()
    // See song.go for implementation of functions used by sort.Sort()
    sort.Sort(song.ByVotes(s.Songs))
    s.Lock.Unlock()
}

// SongQueue.GetNextSong()
//      Return index of song in the queue that has the highest votes
//      AND has not been played in the last 30 minutes
func (s *SongQueue) GetNextSong() int {
    // Get the time NOW, when we start lookin for next song
    now := time.Now()
    // and sort the queue by votes
    s.SortQueue()
    songIndex := -1

    // Loop over all songs in queue, which is sorted by votes
    // if we find a song that has not been played in 30 minutes or more,
    // we store that index and break out of the loop
    s.Lock.RLock()
    for index, element := range s.Songs {
        difference := now.Sub(element.LastPlayed)
        if difference.Minutes() > 30 {
            songIndex = index
            break
        }
    }
    s.Lock.RUnlock()

    // Return the index of the song to be played, if no song is found
    // we return -1
    return songIndex
}

// SongQueue.GetSong(songName)
//      Find a song in the queue by name
//      Same logic as in GetNextSong() otherwise
func (s *SongQueue) GetSong(songName string) int {
    songIndex := -1

    s.Lock.RLock()
    for index, element := range s.Songs {
        if element.Name == songName {
            songIndex = index
            break
        }
    }
    s.Lock.RUnlock()
    return songIndex
}

// SongQueue.VoteForSong(songName)
//      Vote for a certain song by name
func (s *SongQueue) VoteForSong(songName string) {
    s.Lock.RLock()
    // Find the song we want to vote for
    requested := s.GetSong(songName)
    if requested >= 0 {
        // Cast the vote if the song was found
        s.Songs[requested].VoteForSong()
    } else {
        fmt.Println("Could not find song and therefore unable to register vote.")
    }
    s.Lock.RUnlock()
}

