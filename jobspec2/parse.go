package jobspec2

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	hcljson "github.com/hashicorp/hcl/v2/json"
)

func Parse(path string, r io.Reader) (hcl.Diagnostics, error) {
	if path == "" {
		if f, ok := r.(*os.File); ok {
			path = f.Name()
		}
	}

	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		return nil, err
	}

	return ParseWithConfig(&ParseConfig{
		Path:    path,
		Body:    buf.Bytes(),
		AllowFS: false,
		Strict:  true,
	})
}

func ParseWithConfig(args *ParseConfig) (hcl.Diagnostics, error) {
	args.normalize()

	c := newJobConfig(args)
	return Decode(c)
}

type ParseConfig struct {
	Path    string
	BaseDir string

	// Body is the HCL body
	Body []byte

	// AllowFS enables HCL functions that require file system accecss
	AllowFS bool

	// ArgVars is the CLI -var arguments
	ArgVars []string

	// VarFiles is the paths of variable data files
	VarFiles []string

	// Envs represent process environment variable
	Envs []string

	Strict bool

	// parsedVarFiles represent parsed HCL AST of the passed EnvVars
	parsedVarFiles []*hcl.File
}

func (c *ParseConfig) normalize() {
	if c.BaseDir == "" {
		c.BaseDir = filepath.Dir(c.Path)
	}
}

func Decode(c *jobConfig) (hcl.Diagnostics, error) {
	config := c.ParseConfig

	file, diags := parseHCLOrJSON(config.Body, config.Path)

	for _, varFile := range config.VarFiles {
		parsedVarFile, ds := parseFile(varFile)
		if parsedVarFile == nil || ds.HasErrors() {
			return nil, fmt.Errorf("unable to parse var file: %v", ds.Error())
		}

		config.parsedVarFiles = append(config.parsedVarFiles, parsedVarFile)
		diags = append(diags, ds...)
	}

	diags = append(diags, c.decodeBody(file.Body)...)

	// if diags.HasErrors() {
	// 	var str strings.Builder
	// 	for i, diag := range diags {
	// 		if i != 0 {
	// 			str.WriteByte('\n')
	// 		}
	// 		str.WriteString(diag.Error())
	// 	}
	// 	return nil, errors.New(str.String())
	// }

	diags = append(diags, decodeMapInterfaceType(&c.Job, c.EvalContext())...)
	diags = append(diags, decodeMapInterfaceType(&c.Tasks, c.EvalContext())...)
	diags = append(diags, decodeMapInterfaceType(&c.Vault, c.EvalContext())...)

	return diags, nil
	// if diags.HasErrors() {
	// 	return diags, nil
	// }
	//
	// return nil, nil
}

func parseFile(path string) (*hcl.File, hcl.Diagnostics) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, hcl.Diagnostics{
			&hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  "Failed to read file",
				Detail:   fmt.Sprintf("failed to read %q: %v", path, err),
			},
		}
	}

	return parseHCLOrJSON(body, path)
}

func parseHCLOrJSON(src []byte, filename string) (*hcl.File, hcl.Diagnostics) {
	if isJSON(src) {
		return hcljson.Parse(src, filename)
	}

	return hclsyntax.ParseConfig(src, filename, hcl.Pos{Line: 1, Column: 1})
}

func isJSON(src []byte) bool {
	for _, c := range src {
		if c == ' ' {
			continue
		}

		return c == '{'
	}
	return false
}
