BINARY = renamer
GOOS = windows
GOARCH = amd64

all: windows linux

windows:
	GOOS=windows GOARCH=amd64 go build -o ${BINARY}-${GOOS}-${GOARCH}.exe .;
	GOOS=windows GOARCH=386 go build -o ${BINARY}-${GOOS}-${GOARCH}.exe .

linux:
	GOOS=linux GOARCH=amd64 go build -o ${BINARY}-${GOOS}-${GOARCH} .;
	GOOS=linux GOARCH=386 go build -o ${BINARY}-${GOOS}-${GOARCH} .
