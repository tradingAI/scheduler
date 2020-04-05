package client

import (
	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"google.golang.org/grpc"
)

type Client struct {
	Client pb.SchedulerClient
	Conn   *grpc.ClientConn
}

func New() (client Client) {
	conn, err := grpc.Dial("localhost:8889", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		glog.Fatal(err)
	}

	client.Conn = conn
	client.Client = pb.NewSchedulerClient(conn)

	return
}

func (c *Client) Free() (err error) {
	err = c.Conn.Close()
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
