package butterflyembed

import (
	"errors"
	"os"
	"path/filepath"
)

const cacheFolder = "/butterfly_embed"
const generateFolder = "/generated_code"

var (
	ErrDuplicatedModule = errors.New("Duplicated Module")
	ErrCantCleanRuntime = errors.New("Couldn't clear previous runtime for new code")
	ErrUnzipRuntime     = errors.New("Error unzipping runtime")
)

type BFRuntime struct {
	root         string
	module       string
	baseDir      string
	generatedDir string
}

func CreateRuntime(root string) (*BFRuntime, error) {
	var runtime = &BFRuntime{
		root:         root,
		module:       "butterfly_embed",
		baseDir:      filepath.Join(root, cacheFolder),
		generatedDir: filepath.Join(root, cacheFolder, generateFolder),
	}
	if err := runtime.Clear(); err != nil {
		return nil, errors.Join(ErrCantCleanRuntime, err)
	}
	if err := cacheRuntime(runtime); err != nil {
		return nil, errors.Join(ErrUnzipRuntime, err)
	}
	return runtime, nil
}

func CreateTemp(root string, callback func(*BFRuntime) error) error {
	var bf, err = CreateRuntime(root)
	if err != nil {
		return err
	}
	defer func() { _ = bf.Clear() }()
	return callback(bf)
}

func (bf *BFRuntime) AddTargetCode(moduleName, content string) error {
	var path = filepath.Join(bf.generatedDir, moduleName+".go")
	if _, err := os.Stat(path); err == nil {
		return errors.Join(ErrDuplicatedModule, errors.New("Attempted to create "+moduleName+" multiple times"))
	}
	generatedFile, err := os.Create(path)
	defer func() { generatedFile.Close() }()
	if err != nil {
		return err
	}
	_, err = generatedFile.WriteString(content)
	return err
}

func (bf *BFRuntime) GoModule() string {
	return bf.module
}

func (bf *BFRuntime) CompileDir() string {
	return bf.baseDir
}

func (bf *BFRuntime) GenerateDir() string {
	return bf.generatedDir
}

func (bf *BFRuntime) Clear() error {
	return os.RemoveAll(bf.baseDir)
}
