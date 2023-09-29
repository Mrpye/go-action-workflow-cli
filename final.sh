export PATH=$PATH:/usr/local/go/bin
CGO_ENABLED=0  GOARCH=386 GOOS=linux go build -buildvcs=false -ldflags "-s -w" -o ./hauler
./upx --brute hauler