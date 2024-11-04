build: clean
	go build -o build/savescum cmd/main.go

clean:
	rm -f build/savescum

build-linux: clean
	GOOS=linux GOARCH=amd64 go build -o build/savescum-linux cmd/main.go

build-bsd: clean
	GOOS=freebsd GOARCH=amd64 go build -o build/savescum-freebsd cmd/main.go

clean-bsd:
	rm -f build/savescum-freebsd

deploy: build-bsd build-linux
	scp build/savescum-freebsd 192.168.50.12:/mnt/big-6-disk-pool/others/bin
	scp build/savescum-linux 192.168.50.13:/home/ashwin/bin/savescum
