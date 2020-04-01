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

	r, err := c.RegisterRunner(context.Background(), &pb.RegisterRunnerRequest{RunnerId: "66666"})
	if err != nil {
		glog.Fatal(err)
	}
	glog.Infof("Resp: %v", r.GetOk())
}
