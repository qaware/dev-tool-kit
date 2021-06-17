.PHONY: build build-debug clean

WAILS=$$(go env GOPATH)/bin/wails
WAILS_VERSION=1.16.4-pre2

build:
	go get -u github.com/wailsapp/wails/cmd/wails@v${WAILS_VERSION}
	${WAILS} build -f -x linux/amd64
	${WAILS} build -f -x windows/amd64 && mv ${BUILD_DIR}/dev-tool-kit-windows-*-amd64.exe ${BUILD_DIR}/dev-tool-kit-windows-amd64.exe
	#${WAILS} build -f -x darwin/amd64 && mv ${BUILD_DIR}/dev-tool-kit-darwin-*-amd64 ${BUILD_DIR}/dev-tool-kit-darwin-amd64

build-debug:
	go get -u github.com/wailsapp/wails@v${WAILS_VERSION}
	go get -u github.com/wailsapp/wails/cmd/wails@v${WAILS_VERSION}
	${WAILS} build -f -d

clean:
	rm -rf build/
	rm -rf frontend/build/
	rm -f dev-tool-kit*
