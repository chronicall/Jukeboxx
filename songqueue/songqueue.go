package songqueue

import (
    "github.com/chronicall/jukeboxx/song"
    //"sort"
    "time"
)

type SongQueue struct {
    Songs []song.Song
}
/*
// Methods to enable sorting the queue by votes
func (s SongQueue) Len() int {
    return len(s.songs)
}
func (s SongQueue) Swap(i, j int) {
    s.songs[i], s.songs[j] := s.songs[j], s.songs[i]
}
func (s SongQueue) Less(i, j int) bool {
    return s.songs[i].Votes > s.songs[j].Votes
}
*/
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
    requested := song.Song{
        Name: songName,
        Votes: 0,
        Duration: dur,
        LastPlayed: time.Time{},
    }
    s.Add(requested)
}

// Return index of song in the queue that has the highest votes
// AND has not been played in the last 30 minutes
func (s SongQueue) GetNextSong() int {
    now := time.Now()
    for index, element := range s.Songs {
        difference := now.Sub(element.LastPlayed)
        if difference.Minutes() > 30 {
            return index
        }
    }
    return -1
}

// Find a song in the queue by name
func (s SongQueue) GetSong(songName string) int {
    for index, element := range s.Songs {
        if element.Name == songName {
            return index
        }
    }
    return -1
}

func (s *SongQueue) Add(song song.Song) {
    // lock the queue
    s.Songs = append(s.Songs, song)
    // unlock the queue
}

// Vote for a song
func (s *SongQueue) VoteForSong(name string) {
    // lock the queue
    song := s.GetSong(name)
    s.Songs[song].VoteForSong()
    // unlock the queue
}

