live:
	CompileDaemon -command="./paperstamp" -exclude-dir="./web"

test:
	go test ./...  -coverprofile=out/coverage.out -json > out/test-report.out

sonar:test
	sonar-scanner