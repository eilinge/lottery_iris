
.PHONE: test, sonar, build, clean, run

app := cmpp_service

commit := $(shell git rev-parse HEAD)
commit_flag := -X main.Commit=$(commit)
build_time := $(shell date "+%Y-%m-%d.%H:%M:%S")
build_flag := -X main.Build=$(build_time)

test: export GOFLAGS=-mod=vendor
test:
	go test -cover -race ./...

sonar: export GOFLAGS=-mod=vendor
sonar:
	go test -coverprofile=out/cov.out
	sonar-scanner

build: export CGO_ENABLED=0
build:
	go build -mod=vendor -v -ldflags '$(commit_flag) $(build_flag)' -o out/$(app)

deploy_test: export GOOS=linux
deploy_test: build
	ansible-playbook -i scripts/hosts scripts/deploy.yaml

clean:
	rm -rf out

run:
	cd web/ && go run main.go
