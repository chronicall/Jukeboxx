package songqueue

import (
    "fmt"
    "math/rand"
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
        FirstVote: time.Time{},
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
// Assumes that the queue has been locked
func (s SongQueue) SortQueue() {
    // See song.go for implementation of functions used by sort.Sort()
    sort.Sort(song.ByVotes(s.Songs))
}

// SongQueue.GetNextSong()
//      Return index of song in the queue that has the highest votes
//      AND has not been played in the last 30 minutes
func (s *SongQueue) GetNextSong() int {
    // Get the time NOW, when we start lookin for next song
    now := time.Now()
    // and sort the queue by votes
    // Set up values to use below
    songIndex := -1
    foundSong := false

    // Loop over all songs in queue, which is sorted by votes
    s.Lock.Lock()
    s.SortQueue()

    for index, element := range s.Songs {
        lastPlayedDifference:= now.Sub(element.LastPlayed)
        firstVoteDifference := now.Sub(element.FirstVote)
        if !foundSong {
            if lastPlayedDifference.Minutes() > 30 {
                // if we find a song, set the bool value to true and set the index
                // of the song
                foundSong = true
                songIndex = index
            }
        }
        // Keep iterating to check for "rotting" songs that have received a vote
        // but have not been played in 1 hour or more
        // ensures that every song is eventually played
        // this time value can be changed
        if element.Votes > 0 && firstVoteDifference.Hours() > 1 {
            songIndex = index
            break
        }
    }
    s.Lock.Unlock()

    // If no song has been found above, pick a random song.
    if songIndex == -1 {
        songIndex = rand.Intn(len(s.Songs))
    }
    return songIndex
}

// SongQueue.GetSong(songName)
//      Find a song in the queue by name
//      Same logic as in GetNextSong() otherwise
// Only called from within SongQueue.VoteForSong where the queue
// has already been locked
func (s *SongQueue) GetSong(songName string) int {
    songIndex := -1

    for index, element := range s.Songs {
        if element.Name == songName {
            songIndex = index
            break
        }
    }
    return songIndex
}

// SongQueue.VoteForSong(songName)
//      Vote for a certain song by name
func (s *SongQueue) VoteForSong(songName string) {
    s.Lock.Lock()
    // Find the song we want to vote for
    requested := s.GetSong(songName)
    if requested >= 0 {
        // Cast the vote if the song was found
        s.Songs[requested].VoteForSong()
    } else {
        fmt.Println("Could not find song and therefore unable to register vote.")
    }
    s.Lock.Unlock()
}

