default:
	env CGO_ENABLED=0 GOAMD64=v3 go build -ldflags "-s -w" -o nomad-lsp.amdv3
	env CGO_ENABLED=0 go build -ldflags "-s -w"
	cp nomad-lsp.amdv3 ~/go/bin/nomad-lsp
