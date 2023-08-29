module github.com/juliosueiras/nomad-lsp

go 1.19

// Pinned dependencies are noted in github.com/hashicorp/nomad/issues/11826
// https://github.com/hashicorp/nomad/blob/29893023f7da52d76479399825480f8536eb337b/go.mod#L9
replace github.com/hashicorp/hcl => github.com/hashicorp/hcl v1.0.1-0.20201016140508-a07e7d50bbee

replace github.com/hashicorp/hcl/v2 => github.com/hashicorp/hcl/v2 v2.9.2-0.20220525143345-ab3cae0737bc

require (
	github.com/creachadair/jrpc2 v0.43.0
	github.com/davecgh/go-spew v1.1.1
	github.com/hashicorp/go-cty-funcs v0.0.0-20230112231712-1aba97313821
	github.com/hashicorp/hcl/v2 v2.16.0
	github.com/hashicorp/nomad v1.4.6
	github.com/hashicorp/nomad/api v0.0.0-20230210211211-29893023f7da
	github.com/mitchellh/reflectwalk v1.0.2
	github.com/shoenig/test v0.6.1
	github.com/sourcegraph/go-lsp v0.0.0-20200429204803-219e11d77f5d
	github.com/stretchr/testify v1.8.1
	github.com/zclconf/go-cty v1.12.1
	github.com/zclconf/go-cty-yaml v1.0.3
)

require (
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-cidr v1.1.0 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/bmatcuk/doublestar v1.3.4 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/cronexpr v1.1.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/hcl v1.0.1-vault-3 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/exp v0.0.0-20230210204819-062eb4c674ab // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
