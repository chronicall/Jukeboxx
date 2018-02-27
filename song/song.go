package song

import (
    "fmt"
    "time"
    "sync"
)

// Struct that keeps track of our songs, their votes,
// duration and when they were last played. Each song
// has a Read and Write mutex for mutual exclusion of
// shared variables.
type Song struct {
    Name string
    Votes int
    Duration time.Duration
    LastPlayed time.Time
    Lock sync.RWMutex
}

// Methods to enable sorting the queue by votes
type ByVotes []Song

// Used by sort.Sort(ByVotes([]Song))
// Returns length of slice
func (v ByVotes) Len() int {
    return len(v)
}
// Swaps elements in the slice (after finding which is smaller
// using Less)
func (v ByVotes) Swap(i, j int) {
    v[i], v[j] = v[j], v[i]
}
// Returns a bool, saying which of two element vote variables are smaller
func (v ByVotes) Less(i, j int) bool {
    return v[i].Votes > v[j].Votes
}

// Song.GetDuration()
//      Returns the duration of a song to be used for printing
//      and calculations
func (s Song) GetDuration() time.Duration {
    return s.Duration
}

// Song.SongInfo()
//      Prints out information on a song
//      Name, votes, duration and when it was last played
//      Last played format: DD-Month-YYYY, HH:MM.SS
func (s Song) SongInfo() {
    durMinutes := s.Duration.Truncate(time.Second).String()
    hours, minutes, seconds := s.LastPlayed.Clock()
    year, month, day := s.LastPlayed.Date()
    fmt.Printf("Title: %s - Duration: %s\nVotes: %d - Last played: %02d-%v-%d, %02d:%02d.%02d\n", s.Name, durMinutes, s.Votes, day, month, year, hours, minutes, seconds)
}

// Song.VoteForSong()
//      Increments the vote count for a song
//      Thread safe thanks to locks
//      Called when a guest votes for a song
func (s *Song) VoteForSong() {
    s.Lock.Lock()
    s.Votes++
    s.Lock.Unlock()
}

// Song.Play()
//      Plays a song that has been selected
//      Resets votes and update last played time
func (s *Song) Play() {
    fmt.Printf("Now Playing: %s, for %s\n", s.Name, s.Duration.Truncate(time.Second).String())
    s.Lock.Lock()
    s.Votes = 0
    s.LastPlayed = time.Now()
    s.Lock.Unlock()
}

