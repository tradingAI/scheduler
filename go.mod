module github.com/tradingAI/scheduler

go 1.13

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/jinzhu/gorm v1.9.12
	github.com/lib/pq v1.1.1
	github.com/minio/minio-go/v6 v6.0.50
	github.com/tradingAI/go v0.0.0-20200401080831-5aa0707a30db
	github.com/tradingAI/proto/gen/go/common v0.0.0-00010101000000-000000000000 // indirect
	github.com/tradingAI/proto/gen/go/model v0.0.0-00010101000000-000000000000 // indirect
	github.com/tradingAI/proto/gen/go/scheduler v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.28.0
)

replace github.com/tradingAI/proto/gen/go/common => ../proto/gen/go/common

replace github.com/tradingAI/proto/gen/go/model => ../proto/gen/go/model

replace github.com/tradingAI/proto/gen/go/scheduler => ../proto/gen/go/scheduler
