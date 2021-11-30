all: build-all

build-all: build-linux build-windows build-mac build-web
build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/GoBlock-Studio_linux .
build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/GoBlock-Studio_windows.exe .
build-mac:
	#DOESN'T CURRENTLY WORK: GOOS=darwin GOARCH=amd64 go build -o bin/GoBlock-Studio_mac .
build-web:
	GOOS=js GOARCH=wasm go build -o bin/web/GoBlock-Studio_web.wasm .
clean:
	go clean
	rm -rf ./bin/*
