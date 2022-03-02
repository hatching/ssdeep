all: cmd/ssdeep/main.go ssdeep.go score.go
	go install ./cmd/ssdeep

test:
	go test ./...
