dep:
	go mod tidy

build-nix: dep
	go build  -o go-aws-dynamic-dns cmd/go-aws-dynamic-dns/main.go

build-win: dep
	go build -o go-aws-dynamic-dns.exe cmd/go-aws-dynamic-dns/main.go 