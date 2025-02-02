default:
	env CGO_ENABLED=0 GOAMD64=v3 go build -ldflags "-s -w" -o nomad-lsp.linux.amdv3
	env CGO_ENABLED=0 go build -ldflags "-s -w" -o nomad-lsp.linux.amd64
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o nomad-lsp.darwin.amd64
	env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o nomad-lsp.darwin.arm64
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o nomad-lsp.amd64.exe
	env CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags "-s -w" -o nomad-lsp.arm64.exe
	cp nomad-lsp.linux.amdv3 ~/go/bin/nomad-lsp

clean:
	rm nomad-lsp.*
