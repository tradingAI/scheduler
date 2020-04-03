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
	ErrEmptyRunnerID    = makeError(errRunner, "runner id is empty")
	ErrNilRunner        = makeError(errRunner, "runner is nil")
	ErrEmptyRunnerToken = makeError(errRunner, "runner token is empty")

	// Job error
	ErrNilCreateJobRequest          = makeError(errJob, "create job request is nil")
	ErrNilStopJobRequest            = makeError(errJob, "stop job request is nil")
	ErrInvalidCreateJobRequestInput = makeError(errJob, "input of create job request is invalid")
	ErrNilRemoveJobRequest          = makeError(errJob, "remove job request is nil")
)

func makeError(scope errScope, msg ...string) error {
	return fmt.Errorf("[%s]: %s", scope, msg)
}
