install:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
	npm install

build: clean
	buf generate

clean:
	rm -rf ../backend/gen
	rm -rf ../frontend/gen
