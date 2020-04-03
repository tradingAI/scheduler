package main

import (
	"context"
	"flag"

	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	flag.Set("logtostderr", "true")

	conn, err := grpc.Dial("localhost:8889", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		glog.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewSchedulerClient(conn)

	runner := pb.Runner{
		Id:                 "6666",
		Status:             pb.RunnerStatus_IDLE,
		CpuCoreNum:         4,
		CpuUtilization:     0.1,
		GpuNum:             0,
		GpuUtilization:     0,
		Memory:             2 * 1024 * 1024 * 1024,
		AvailableMemory:    1 * 1024 * 1024 * 1024,
		GpuMemory:          0,
		AvailableGpuMemory: 0,
		Token:              "admin",
	}

	resp, err := c.HeartBeat(context.Background(), &pb.HeartBeatRequest{Runner: &runner})
	if err != nil {
		glog.Fatal(err)
	}
	glog.Infof("Resp: %v", resp.GetOk())
}
