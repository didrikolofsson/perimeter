package main

import (
	"os"
	"perimeter/internal/cli"
	"perimeter/internal/index"
	"perimeter/internal/logx"
)

func main() {
	// Scan project root
	files, err := index.ScanDirRecursive(cli.Path)
	if err != nil {
		logx.Logger.Error("Failed to scan project root", "error", err)
		os.Exit(1)
	}
	sourceFiles, err := index.GetSourceFiles(files)
	if err != nil {
		logx.Logger.Error("Failed to get source files", "error", err)
		os.Exit(1)
	}
	content, err := index.GetFileContentFull(sourceFiles[0])
	if err != nil {
		logx.Logger.Error("Failed to get file content", "error", err)
		os.Exit(1)
	}
	logx.Logger.Info(string(content))
}
