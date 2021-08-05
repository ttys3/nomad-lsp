package nomadstructs

import (
	"log"
	"os"

	"github.com/juliosueiras/nomad-lsp/jobspec2"
	"github.com/sourcegraph/go-lsp"
)

func GetDiagnostics(fileName string, originalFile string) []lsp.Diagnostic {
	result := make([]lsp.Diagnostic, 0)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		log.Println(err)
		return result
	}

	if _, err := os.Stat(originalFile); os.IsNotExist(err) {
		log.Println(err)
		originalFile = fileName
	}

	r, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return result
	}
	hclDiags, err := jobspec2.Parse(originalFile, r)
	if err != nil {
		log.Println(err)
		return result
	}

	for _, diag := range hclDiags {
		result = append(result, lsp.Diagnostic{
			Severity: lsp.DiagnosticSeverity(diag.Severity),
			Message:  diag.Detail,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      diag.Subject.Start.Line - 1,
					Character: diag.Subject.Start.Column - 1,
				},
				End: lsp.Position{
					Line:      diag.Subject.End.Line - 1,
					Character: diag.Subject.End.Column - 1,
				},
			},
			Source: "Nomad",
		})
	}

	return result
}
