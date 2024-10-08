package stats

import (
	"math"
	"sync"
)

// Calculate the mean, variance, and standard deviation using Welford's method.
// https://en.wikipedia.org/wiki/Algorithms_for_calculating_variance
type Stats struct {
	count int64
	mean  float64
	M2    float64
	mu    sync.Mutex
}

func NewStats() *Stats {
	return &Stats{}
}

func (s *Stats) Add(value float64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.count++
	delta := value - s.mean
	s.mean += delta / float64(s.count)
	delta2 := value - s.mean
	s.M2 += delta * delta2
}

func (s *Stats) Mean() float64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.mean
}

func (s *Stats) Variance() float64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.count < 2 {
		return 0
	}
	return s.M2 / float64(s.count-1)
}

func (s *Stats) StdDev() float64 {
	return math.Sqrt(s.Variance())
}
