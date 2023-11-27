build:
	docker build -t sequencer .

run:
	docker run -p 3000:3000 -p 3001:3001 sequencer

test:
	go test ./...

.PHONY: build run