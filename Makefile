protofiles:
	protoc ./proto/*.proto --go_out=. --go-grpc_out=.


BASE_OUTPUT=/home/anorak/Desktop/Projects/PKr-cli/PKr-cli/PKr-base/PKr-base.exe
CLI_OUTPUT=/home/anorak/Desktop/Projects/PKr-cli/PKr-cli/PKr-cli.exe

TEST_DEST=/home/anorak/Desktop/Projects/PKr-test-runs/$(TEST)
TEST=Test-8

build2test: clean build copy

build: clean
	@echo Building the Go file...
	go build
	cd PKr-base &&  go build

copy: build
	@echo Copying the executable to the destination...
	@copy $(BASE_OUTPUT) $(TEST_DEST)
	@copy $(CLI_OUTPUT) $(TEST_DEST)

clean:
	@echo Cleaning up...
	del $(CLI_OUTPUT)
	del $(BASE_OUTPUT)
	del $(TEST_DEST)\PKr-base.exe
	del $(TEST_DEST)\PKr-cli.exe

# Just for trying doesnt work [orginally for automating testing]
docker_build:
	docker build -t pkr-cli .

docker_compose:
	docker-compose up --build

get_new_base:
	go get github.com/ButterHost69/PKr-Base@latest

.PHONY: build2test build copy clean protofiles

