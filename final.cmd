go build -ldflags "-s -w" -o ./hauler.exe
upx --brute hauler.exe