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
	for _, file := range sourceFiles {
		isJestTest, err := index.GetJestTestSignature(file)
		if err != nil {
			logx.Logger.Error("Failed to get jest test signature", "error", err)
			continue
		}
		if isJestTest {
			continue
		}
		signatureHits, err := index.ScanSourceFile(file)
		if err != nil {
			logx.Logger.Error("Failed to scan source file", "error", err)
			continue
		}
		for _, signatureHit := range signatureHits {
			logx.Logger.Info("Signature hit", "path", signatureHit.Path, "line", signatureHit.LineNumber, "signature type", signatureHit.SignatureType)
		}
	}
}
