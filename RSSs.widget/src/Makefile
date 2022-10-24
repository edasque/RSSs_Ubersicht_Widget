all: run macos

macos: main.go
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build  -o outputFeed.macos.amd64 main.go
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build  -o outputFeed.macos.arm64 main.go
	lipo -create -output outputFeed outputFeed.macos.amd64 outputFeed.macos.arm64

run: main.go
	go run main.go

clean:
	rm -f $(OUT)
