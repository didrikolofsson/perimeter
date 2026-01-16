package main

import (
	"os"
	"perimeter/internal/cli"
	"perimeter/internal/index"
	"perimeter/internal/logx"
	"perimeter/internal/types"
)

func main() {
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

	signatureHits := []types.SignatureHit{}
	for _, file := range sourceFiles {
		isJestTest, err := index.GetJestTestSignature(file)
		if err != nil {
			logx.Logger.Error("Failed to get jest test signature", "error", err)
			continue
		}
		if isJestTest {
			continue
		}
		sigs, err := index.ScanSourceFile(file)
		if err != nil {
			logx.Logger.Error("Failed to scan source file", "error", err)
			continue
		}
		signatureHits = append(signatureHits, sigs...)
	}

	signatureSpans := []types.SignatureSpan{}
	for _, signatureHit := range signatureHits {
		signatureSpan, err := index.ExpandSignatureHitSpan(signatureHit)
		if err != nil {
			logx.Logger.Error("Failed to expand signature hit span", "error", err, "signature hit", signatureHit)
			continue
		}
		signatureSpans = append(signatureSpans, signatureSpan)
	}

	for _, sp := range signatureSpans {
		logx.Logger.Info("Signature span", "path", sp.Path, "start line", sp.StartLine, "end line", sp.EndLine, "content", sp.Content)
	}
}
