mocks:
	cd ./store/mocks/; go generate;
	cd ./service/mocks/; go generate;

build:
	go build cmd/main.go

run:
	cd cmd; go run main.go
