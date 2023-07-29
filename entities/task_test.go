package entities

import (
	"strings"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestTask_SolutionValid(t *testing.T) {
	assert := assert.New(t)

	task := Task{
		InputPrefix:             "35b4b08f-e553-4782-9838-8c27a7b4eb98",
		MustContainLeadingZeros: 3,
	}
	solution := "35b4b08f-e553-4782-9838-8c27a7b4eb98:try:484"
	t.Log(task.hash([]byte(solution)))

	assert.True(task.SolutionValid([]byte(solution)))
	assert.False(task.SolutionValid([]byte("solution")))
	assert.True(strings.HasPrefix(task.hash([]byte(solution)), "000"))

}
