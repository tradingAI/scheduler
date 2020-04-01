package common

import (
	"fmt"
)

type errScope string

const (
	errUnknown errScope = ""
	errRunner           = "runner_error"
	errJob              = "job_error"
)

var (
	// Runner error
	ErrEmptyRunnerID       = makeError(errRunner, "runner id is empty")
	ErrNilRunner           = makeError(errRunner, "runner is nil")
	ErrNilCreateJobRequest = makeError(errJob, "create job request is nil")
	ErrNilStopJobRequest   = makeError(errJob, "stop job request is nil")
	ErrEmptyJobID          = makeError(errJob, "job id is empty")
)

func makeError(scope errScope, msg ...string) error {
	return fmt.Errorf("[%s]: %s", scope, msg)
}
