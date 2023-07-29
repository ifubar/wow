package storage

import (
	"github.com/ifubar/wow/entities"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTaskStorage_MarkSolved(t *testing.T) {

	storage := &TaskStorage{
		cache: cache.New(cache.NoExpiration, time.Millisecond*10),
	}

	task := entities.Task{
		InputPrefix: "some prefix",
	}
	err := storage.MarkSolved(task, time.Second)
	assert.Nil(t, err)

	err = storage.MarkSolved(task, time.Second)
	t.Log(err)
	assert.Error(t, err)

	time.Sleep(time.Second)
	err = storage.MarkSolved(task, time.Second)
	assert.Nil(t, err)
}
