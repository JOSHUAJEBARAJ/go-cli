package repository

import (
	"fmt"
	"sync"

	"joshua.com/pomo/pomodro"
)

type inMemoryRepo struct {
	sync.RWMutex
	intervals []pomodro.Interval
}

func NewInMemoryRepo() *inMemoryRepo {
	return &inMemoryRepo{
		intervals: []pomodro.Interval{},
	}
}

func (r *inMemoryRepo) Create(i pomodro.Interval) (int64, error) {
	r.Lock()
	defer r.Unlock()
	i.ID = int64(len(r.intervals)) + 1
	r.intervals = append(r.intervals, i)
	return i.ID, nil
}

func (r *inMemoryRepo) Update(i pomodro.Interval) error {
	r.Lock()
	defer r.Unlock()
	if i.ID == 0 {
		return fmt.Errorf("%w %d", pomodro.ErrInvalidId, i.ID)
	}
	r.intervals[i.ID-1] = i
	return nil
}

func (r *inMemoryRepo) ById(id int64) (pomodro.Interval, error) {
	r.Lock()
	defer r.RUnlock()
	i := pomodro.Interval{}
	if id == 0 {
		return i, fmt.Errorf("%w %d", pomodro.ErrInvalidId, id)
	}
	i = r.intervals[id-1]
	return i, nil
}
