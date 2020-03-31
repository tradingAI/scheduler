package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"github.com/minio/minio-go/v6"
	pg "github.com/tradingAI/go/db/postgres"
	minio2 "github.com/tradingAI/go/s3/minio"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"google.golang.org/grpc"
)

type Servlet struct {
	Conf  Conf
	DB    *gorm.DB
	Minio *minio.Client
}

func New(conf Conf) (s *Servlet, err error) {
	// make server
	s = &Servlet{
		Conf: conf,
	}

	// Init db client
	s.DB, err = pg.NewPostgreSQL(conf.DB)
	if err != nil {
		glog.Error(err)
		return
	}

	// Init s3 client
	s.Minio, err = minio2.NewMinioClient(s.Conf.Minio)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func (s *Servlet) Free() {
	if err := s.DB.Close(); err != nil {
		glog.Error(err)
		return
	}

	return
}

func (s *Servlet) StartOrDie() (err error) {
	grpcServer := grpc.NewServer()
	pb.RegisterSchedulerServer(grpcServer, s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Conf.Scheduler.Port))
	if err != nil {
		glog.Errorf("failed to listen: %v", err)
		return
	}

	go func() {
		glog.Infof("gRPC listenning on port %d", s.Conf.Scheduler.Port)
		err = grpcServer.Serve(lis)
		if err != nil {
			glog.Errorf("failed to serve: %v", err)
			return
		}
	}()

	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(c)

	// Block until we receive our signal.
	<-c
	grpcServer.GracefulStop()
	s.Free()
	glog.Infof("shutting down")
	glog.Flush()
	os.Exit(0)

	return
}

func (s *Servlet) DestoryRunner(ctx context.Context, req *pb.DestoryRunnerRequest) (resp *pb.DestoryRunnerResponse, err error) {

	return
}

func (s *Servlet) CreateJob(ctx context.Context, req *pb.CreateJobRequest) (resp *pb.CreateJobResponse, err error) {

	return
}

func (s *Servlet) StopJob(ctx context.Context, req *pb.StopJobRequest) (resp *pb.StopJobResponse, err error) {

	return
}

func (s *Servlet) RegisterRunner(ctx context.Context, req *pb.RegisterRunnerRequest) (resp *pb.RegisterRunnerResponse, err error) {

	return
}

func (s *Servlet) HeartBeat(ctx context.Context, req *pb.HeartBeatRequest) (resp *pb.HeartBeatResponse, err error) {

	return
}
