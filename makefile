all: 		gotool
				@go build -v .

clean:
				rm -f apiserver
				find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}

gotool:
				gofmt -w .
				go tool vet . |& grep -v vendor;true

ca:
				openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Balalalala Company/OU=IT/CN=127.0.0.1/emailAddress=abcd@edfh.com"

help:
				@echo "make		- Compile the source code"
				@echo "make clean	- Remove binary file and vim swp files"
				@echo "make gotool	- Run go tool 'fmt' and 'vet'"
				@echo "make ca		- Generate ca files"

.PHONY: clean gotool ca help