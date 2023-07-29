package entities

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

const alg = "sha256"
const difficult = 3

type ctxKey int

var taskKey ctxKey

type Task struct {
	Alg                     string
	InputPrefix             string
	MustContainLeadingZeros int
}

func (t Task) ToCtx(parent context.Context) context.Context {
	return context.WithValue(parent, taskKey, t)
}

func (t Task) FromCtx(ctx context.Context) (Task, bool) {
	u, ok := ctx.Value(taskKey).(Task)
	return u, ok
}

func (t Task) SolutionValid(solution []byte) bool {
	if !strings.HasPrefix(string(solution), t.InputPrefix) {
		return false
	}
	hash := t.hash(solution)
	leadingZerosNum := fmt.Sprintf(fmt.Sprintf("%%0%dd", t.MustContainLeadingZeros), 0)
	return strings.HasPrefix(hash, leadingZerosNum)
}

func (t Task) hash(solution []byte) string {
	hash := sha256.New()
	hash.Write(solution)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func NewTask() Task {
	return Task{
		Alg:                     alg,
		InputPrefix:             uuid.New().String(),
		MustContainLeadingZeros: difficult,
	}
}
