.PHONY: build clean

WAILS=${GOPATH}/bin/wails
WAILS_VERSION=1.8.0
BUILD_DIR=build/github.com/qaware

build:
	go get -u github.com/wailsapp/wails/cmd/wails@v${WAILS_VERSION}
	${WAILS} build -f -x linux/amd64
	${WAILS} build -f -x windows/amd64 && mv ${BUILD_DIR}/dev-tool-kit-windows-*-amd64.exe ${BUILD_DIR}/dev-tool-kit-windows-amd64.exe
	${WAILS} build -f -x darwin/amd64 && mv ${BUILD_DIR}/dev-tool-kit-darwin-*-amd64 ${BUILD_DIR}/dev-tool-kit-darwin-amd64
	grep -oP '[1-9]+\.[0-9]+\.[0-9]+' main.go > ${BUILD_DIR}/version

clean:
	rm -rf build/
	rm -rf frontend/build/
	rm -f dev-tool-kit*