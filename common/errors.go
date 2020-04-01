package common

import (
	"fmt"
)

type errScope string

const (
	errUnknown errScope = ""
	errRunner           = "runner_error"
)

var (
	// Runner error
	ErrEmptyRunnerID = makeError(errRunner, "runner id is empty")
	ErrNilRunner     = makeError(errRunner, "runner is nil")
)

func makeError(scope errScope, msg ...string) error {
	return fmt.Errorf("[%s]: %s", scope, msg)
}
