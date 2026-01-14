package index

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"perimeter/internal/logx"
	"perimeter/internal/types"
)

func validateRoot(root string) error {
	info, err := os.Stat(root)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return errors.New("root is not a directory")
	}
	// Check that package.json exists in the root directory
	if _, err := os.Stat(filepath.Join(root, "package.json")); err != nil {
		return errors.New("package.json not found in root directory")
	}
	return nil
}

func ScanDirRecursive(root string) ([]types.File, error) {
	if err := validateRoot(root); err != nil {
		return nil, err
	}

	files := []types.File{}
	filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			logx.Logger.Debug("Failed to inspect path", "path", p)
			return err
		}
		if info.IsDir() {
			return nil
		}
		files = append(files, types.File{Path: p, Info: info})
		return nil
	})
	return files, nil
}

func IsSourceFile(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".js", ".jsx", ".ts", ".tsx":
		return true
	default:
		return false
	}
}

func GetSourceFiles(files []types.File) ([]types.File, error) {
	sourceFiles := []types.File{}
	for _, file := range files {
		if !IsSourceFile(file.Path) {
			continue
		}
		sourceFiles = append(sourceFiles, file)
	}
	return sourceFiles, nil
}

func GetFileContent(file types.File, maxBytes int64) ([]byte, error) {
	f, err := os.Open(file.Path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var buf []byte
	if maxBytes != 0 {
		r := io.LimitReader(f, maxBytes)
		b, _ := io.ReadAll(r)
		buf = b
	}
	if maxBytes == 0 {
		b, _ := io.ReadAll(f)
		buf = b
	}
	return buf, nil
}

func GetFileContentFull(file types.File) ([]byte, error) {
	return GetFileContent(file, 0)
}
