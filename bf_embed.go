//go:generate 7z a -r .\internal\bf_embed.zip .\internal\bf_embed\*
package butterflyembed

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"io"
	"os"
	"path/filepath"
)

//go:embed internal/bf_embed.zip
var bfRuntime []byte

func cacheRuntime(bf *BFRuntime) error {
	var bReader = bytes.NewReader(bfRuntime)
	var wkdir = bf.baseDir

	r, err := zip.NewReader(bReader, bReader.Size())
	if err != nil {
		return err
	}

	for _, file := range r.File {
		destPath := filepath.Join(wkdir, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(destPath, os.ModePerm)
			continue
		}
		if err := writeZipFile(destPath, file); err != nil {
			return err
		}
	}
	return nil
}

func writeZipFile(destPath string, file *zip.File) error {
	archiveFile, err := file.Open()
	defer func() { archiveFile.Close() }()
	if err != nil {
		return err
	}

	if os.MkdirAll(filepath.Dir(destPath), os.ModePerm) != nil {
		return err
	}
	destinationFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer func() { destinationFile.Close() }()

	_, err = io.Copy(destinationFile, archiveFile)
	return err
}
