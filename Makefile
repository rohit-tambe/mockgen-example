mock:
	rm -rf mocks
	go install github.com/golang/mock/mockgen@1.6.0
	go generate -tags=tool mockgen.go

coverage:
	go test -v -coverprofile cover.out ./YOUR_CODE_FOLDER/...
	go tool cover -html=cover.out -o cover.html