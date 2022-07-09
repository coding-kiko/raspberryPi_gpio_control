compile_rpi:
	env GOOS=linux GOARCH=arm GOARM=5 go build api/main.go 