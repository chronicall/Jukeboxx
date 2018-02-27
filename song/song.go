package song

import (
    "fmt"
    "time"
    "sync"
)

type Song struct {
    Name string
    Votes int
    Duration time.Duration
    LastPlayed time.Time
    Lock sync.RWMutex
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
    durMinutes := s.Duration.Truncate(time.Second).String()
    hours, minutes, seconds := s.LastPlayed.Clock()
    year, month, day := s.LastPlayed.Date()
    fmt.Printf("Title: %s - Duration: %s\nVotes: %d - Last played: %02d-%v-%d, %02d:%02d.%02d\n", s.Name, durMinutes, s.Votes, day, month, year, hours, minutes, seconds)
}

func (s *Song) UpdateLastPlayed() {
    s.Lock.Lock()
    s.LastPlayed = time.Now()
    s.Lock.Unlock()
}

func (s *Song) VoteForSong() {
    s.Lock.Lock()
    fmt.Printf("Song: %s --- Before voting increment: %d\n", s.Name, s.Votes)
    s.Votes++
    fmt.Printf("Song: %s --- After voting increment: %d\n", s.Name, s.Votes)
    s.Lock.Unlock()
}
