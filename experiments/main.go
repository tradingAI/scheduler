package main

import (
	"flag"

	"github.com/golang/glog"

	"github.com/tradingAI/scheduler/experiments/client"
)

func main() {
	flag.Parse()
	flag.Set("logtostderr", "true")

	client := client.New()

	runnerID := "66666"

	err := client.RegisterRunner(runnerID)
	if err != nil {
		glog.Fatal(err)
	}

	err = client.CreateJob(runnerID)
	if err != nil {
		glog.Fatal(err)
	}

}
