set -e
cd "$(dirname "$0")"

cd ..
git clone https://github.com/tradingAI/proto.git
cd scheduler
make proto
go test -v ./...
