all: windows linux

windows:
	GOOS=windows GOARCH=amd64 go build -o renamer-${GOOS}-${GOARCH}.exe .;
	GOOS=windows GOARCH=386 go build -o renamer-${GOOS}-${GOARCH}.exe .

linux:
	GOOS=linux GOARCH=amd64 go build -o renamer-${GOOS}-${GOARCH} .;
	GOOS=linux GOARCH=386 go build -o renamer-${GOOS}-${GOARCH} .
