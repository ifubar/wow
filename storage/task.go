package storage

import (
	"github.com/ifubar/wow/entities"
	"github.com/patrickmn/go-cache"
	"time"
)

const expInterval = time.Minute

type TaskStorage struct {
	cache *cache.Cache
}

func NewTask() *TaskStorage {
	storage := cache.New(cache.NoExpiration, expInterval)
	return &TaskStorage{
		cache: storage,
	}
}

func (t *TaskStorage) MarkSolved(task entities.Task, duration time.Duration) error {
	return t.cache.Add(task.InputPrefix, task, duration)
}
