BINARY = renamer

all: windows linux

windows:
	GOOS=windows GOARCH=amd64 go build -o ${BINARY}-windows-amd64.exe .;
	GOOS=windows GOARCH=386 go build -o ${BINARY}-windows-386.exe .

linux:
	GOOS=linux GOARCH=amd64 go build -o ${BINARY}-linux-amd64 .;
	GOOS=linux GOARCH=386 go build -o ${BINARY}-linux-386 .
