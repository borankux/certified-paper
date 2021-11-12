live:
	CompileDaemon -command="./certified-paper" -exclude-dir="./web"

test:
	go test ./...  -coverprofile=out/coverage.out -json > out/test-report.out

sonar:test
	sonar-scanner