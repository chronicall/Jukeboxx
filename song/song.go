package song

import (
    "fmt"
    "time"
)

type Song struct {
    Name string
    Votes int
    Duration time.Duration
    LastPlayed time.Time
}

func (s Song) SongInfo() {
    durMinutes, durSeconds := s.Duration / 60, s.Duration % 60
    hours, minutes, seconds := s.LastPlayed.Clock()
    year, month, day := s.LastPlayed.Date()
    fmt.Printf("Title: %s - Duration: %02d:%02d\nVotes: %d - Last played: %02d-%v-%d, %02d:%02d.%02d\n", s.Name, durMinutes, durSeconds, s.Votes, day, month, year, hours, minutes, seconds)
}

func (s *Song) UpdateLastPlayed(t time.Time) {
    // lock song?
    s.LastPlayed = t
    // unlock song
}

func (s *Song) VoteForSong() {
    // lock song?
    s.Votes++
    // unlock song
}
