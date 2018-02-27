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

type ByVotes []Song

// Methods to enable sorting the queue by votes
func (v ByVotes) Len() int {
    return len(v)
}
func (v ByVotes) Swap(i, j int) {
    v[i], v[j] = v[j], v[i]
}
func (v ByVotes) Less(i, j int) bool {
    return v[i].Votes > v[j].Votes
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
