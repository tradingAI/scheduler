package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/tradingAI/scheduler/server"
)

func main() {
	flag.Parse()
	flag.Set("logtostderr", "true")

	runServer()

}

func runServer() {
	// load config
	conf, err := server.LoadConf()
	if err != nil {
		glog.Fatal(err)
	}

	// new Server
	s, err := server.New(conf)
	if err != nil {
		glog.Fatal(err)
	}
	defer s.Free()

	// start server
	err = s.StartOrDie()
	if err != nil {
		glog.Fatal(err)
	}
}
