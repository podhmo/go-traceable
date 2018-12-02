run:
	rm -rf output
	go install -v cmd/...
	mkdir -p output
	TRACE=output go-run-httptrace httptrace/_example/github/main.go
