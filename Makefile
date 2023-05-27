build:
	go build -o bin/hlsvalidator cmd/hlsvalidator/hlsvalidator.go

install:
	go install cmd/hlsvalidator/hlsvalidator.go
