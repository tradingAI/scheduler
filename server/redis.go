package server

import "fmt"

const (
	CREATE_JOB     = "create_job"
	STOP_JOB       = "stop_job"
	DESTORY_RUNNER = "destory"
)

func genRedisKey(prefix string, value string) string {
	return fmt.Sprintf("%s_%s", prefix, value)
}
