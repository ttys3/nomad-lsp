module github.com/juliosueiras/nomad-lsp

go 1.16

replace github.com/hashicorp/hcl => github.com/hashicorp/hcl v1.0.1-0.20201016140508-a07e7d50bbee

require (
	github.com/creachadair/jrpc2 v0.18.0
	github.com/davecgh/go-spew v1.1.1
	github.com/hashicorp/go-cty-funcs v0.0.0-20200930094925-2721b1e36840
	github.com/hashicorp/hcl/v2 v2.9.2-0.20210407182552-eb14f8319bdc
	github.com/hashicorp/nomad v1.1.3
	github.com/hashicorp/nomad/api v0.0.0-20210804153318-c67b69bd0cc6
	github.com/mitchellh/reflectwalk v1.0.2
	github.com/sourcegraph/go-lsp v0.0.0-20200429204803-219e11d77f5d
	github.com/stretchr/testify v1.7.0
	github.com/zclconf/go-cty v1.8.4
	github.com/zclconf/go-cty-yaml v1.0.2
)
