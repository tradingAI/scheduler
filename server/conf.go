package server

import (
	"os"
	"strconv"
	"time"

	"github.com/golang/glog"
	pg "github.com/tradingAI/go/db/postgres"
	minio "github.com/tradingAI/go/s3/minio"
)

type Conf struct {
	DB        pg.DBConf
	Minio     minio.MinioConf
	Scheduler SchedulerConf
}

type SchedulerConf struct {
	Port              int
	ShutdownWaitInSec time.Duration
}

// LoadConf load config from env
func LoadConf() (conf Conf, err error) {
	dbReconnectSec, err := strconv.Atoi(os.Getenv("SCHEDULER_POSTGRES_RECONNECT_SEC"))
	if err != nil {
		glog.Error(err)
		return
	}

	dbPort, err := strconv.Atoi(os.Getenv("SCHEDULER_POSTGRES_PORT"))
	if err != nil {
		glog.Error(err)
		return
	}

	port, err := strconv.Atoi(os.Getenv("SCHEDULER_PORT"))
	if err != nil {
		glog.Error(err)
		return
	}

	minioPort, err := strconv.Atoi(os.Getenv("SCHEDULER_MINIO_PORT"))
	if err != nil {
		glog.Error(err)
		return
	}

	minioSecure, err := strconv.ParseBool(os.Getenv("SCHEDULER_MINIO_SECURE"))
	if err != nil {
		glog.Error(err)
		return
	}

	conf = Conf{
		DB: pg.DBConf{
			Database:     os.Getenv("SCHEDULER_POSTGRES_DB"),
			Username:     os.Getenv("SCHEDULER_POSTGRES_USER"),
			Password:     os.Getenv("SCHEDULER_POSTGRES_PASSWORD"),
			Port:         dbPort,
			Host:         os.Getenv("SCHEDULER_POSTGRES_HOST"),
			ReconnectSec: time.Duration(dbReconnectSec) * time.Second,
		},
		Minio: minio.MinioConf{
			AccessKey: os.Getenv("SCHEDULER_MINIO_ACCESS_KEY"),
			SecretKey: os.Getenv("SCHEDULER_MINIO_SECRET_KEY"),
			Host:      os.Getenv("SCHEDULER_MINIO_HOST"),
			Port:      minioPort,
			Secure:    minioSecure,
		},
		Scheduler: SchedulerConf{
			Port:              port,
			ShutdownWaitInSec: 10 * time.Second,
		},
	}

	if err = conf.Validate(); err != nil {
		glog.Error(err)
		return
	}

	return
}

func (c *Conf) Validate() (err error) {
	if err = c.DB.Validate(); err != nil {
		glog.Info(err)
		return
	}

	if err = c.Minio.Validate(); err != nil {
		glog.Error(err)
		return
	}
	return
}
