package butterflyembed

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRuntimeCreation(t *testing.T) {
	bf, err := CreateRuntime(".")
	if err != nil {
		t.Fatal(err)
	}
	if !exists(bf.baseDir) || !exists(bf.generatedDir) {
		t.Fatal("runtime not created")
	}
	bf.AddTargetCode("testModule", "// just for testing")
	if !exists(filepath.Join(bf.generatedDir, "/testModule.go")) {
		t.Fatal("test file not created")
	}
	bf.Clear()
	if exists(bf.baseDir) {
		t.Fatal("runtime not cleared")
	}
}

func exists(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}
	return true
}
