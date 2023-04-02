fmt:
	go fmt ./...

test:
	go test -v ./...

vet:
	go vet ./...

verify:
	go mod verify

tidy:
	go mod tidy

staticcheck:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...

lint:
	go install golang.org/x/lint/golint@latest
	golint ./...

check: test verify vet staticcheck lint
clean: fmt tidy