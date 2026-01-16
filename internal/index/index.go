package index

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"perimeter/internal/logx"
	"perimeter/internal/types"
	"strings"
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

type ExpressEndpointPattern struct {
	Pattern string
	Type    types.ExpressRouteType
}

var expressEndpointPatterns = []ExpressEndpointPattern{
	{Pattern: ".get(", Type: types.ExpressEndpointGet},
	{Pattern: ".post(", Type: types.ExpressEndpointPost},
	{Pattern: ".put(", Type: types.ExpressEndpointPut},
	{Pattern: ".delete(", Type: types.ExpressEndpointDelete},
}

func IsExpressRoute(line string) bool {
	for _, endpoint := range expressEndpointPatterns {
		if strings.Contains(line, endpoint.Pattern) {
			return true
		}
	}
	return false
}

func GetExpressEndpointType(line string) (types.ExpressRouteType, error) {
	for _, endpoint := range expressEndpointPatterns {
		if strings.Contains(line, endpoint.Pattern) {
			return endpoint.Type, nil
		}
	}
	return "", errors.New("no matching express endpoint type")
}

var jestTestPatterns = []string{
	"it(",
	"describe(",
}

func IsJestTest(line string) bool {
	for _, pattern := range jestTestPatterns {
		if strings.Contains(line, pattern) {
			return true
		}
	}
	return false
}

func GetJestTestSignature(file types.File) (bool, error) {
	f, err := os.Open(file.Path)
	if err != nil {
		logx.Logger.Info("Failed to open file", "error", err)
		return false, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if IsJestTest(line) {
			return true, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return false, err
	}
	return false, nil
}

func ScanSourceFile(file types.File) ([]types.SignatureHit, error) {
	f, err := os.Open(file.Path)
	if err != nil {
		logx.Logger.Info("Failed to open file", "error", err)
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	signatureHits := []types.SignatureHit{}
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		if IsExpressRoute(line) {
			signatureHits = append(signatureHits, types.SignatureHit{
				Path:          file.Path,
				LineNumber:    lineNumber,
				SignatureType: types.ExpressRoute,
			})
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return signatureHits, nil
}
