.PHONY: build css protobuf dev-app dev-css dist

build:
	go build .

init:
	yarn install

css:
	yarn build:css

protobuf:
	protoc --proto_path=. --go_out=paths=source_relative:. ei/ei.proto

dev-app: build
	echo EggLedger | DEV_MODE=1 entr -r ./EggLedger

dev-css:
	yarn dev:css

dist: css protobuf dist-windows dist-mac dist-linux

dist-windows: init css protobuf
	./build-windows.sh

dist-linux: init css protobuf
	./build-linux.sh

dist-mac: init css protobuf
	./build-macos.sh

dist-mac-arm: init css protobuf
	./build-macos-arm.sh

