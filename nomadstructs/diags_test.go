package nomadstructs

import "testing"

func TestGetDiagnostics(t *testing.T) {
	diags := GetDiagnostics("/tmp/nomad-lsp-569356428", "/home/ttys3/repo/docker/nomad/drone-runner-podman.nomad")
	t.Logf("diags=%+v", diags)
}
