package index

import (
	"os"
	"path/filepath"
	"perimeter/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateRoot(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "package.json")
	err := os.WriteFile(p, []byte("{}"), 0644)
	require.NoError(t, err)
	require.NoError(t, validateRoot(dir))
	// Remove package.json
	require.NoError(t, os.Remove(p))
	require.Error(t, validateRoot(dir))
}

func TestIsSourceFile(t *testing.T) {
	assert.True(t, IsSourceFile("test.js"))
	assert.True(t, IsSourceFile("test.jsx"))
	assert.True(t, IsSourceFile("test.ts"))
	assert.True(t, IsSourceFile("test.tsx"))
	assert.False(t, IsSourceFile("test.txt"))
}

func TestGetSourceFiles(t *testing.T) {
	files := []types.File{
		{Path: "test.js", Info: nil},
		{Path: "test.jsx", Info: nil},
		{Path: "test.ts", Info: nil},
		{Path: "test.tsx", Info: nil},
		{Path: "test.txt", Info: nil},
	}
	sourceFiles, err := GetSourceFiles(files)
	require.NoError(t, err)
	assert.Len(t, sourceFiles, 4)
}

func TestScanSourceFile(t *testing.T) {
	tests := scanSourceFileTests

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			// Create temporary file with test content
			dir := t.TempDir()
			filePath := filepath.Join(dir, "test.js")
			err := os.WriteFile(filePath, []byte(tt.Code), 0644)
			require.NoError(t, err)

			// Update expected paths
			expected := make([]types.SignatureHit, len(tt.Expected))
			for i, exp := range tt.Expected {
				expected[i] = exp
				expected[i].Path = filePath
			}

			signatureHits, err := ScanSourceFile(types.File{Path: filePath, Info: nil})
			require.NoError(t, err)
			assert.Equal(t, expected, signatureHits)
		})
	}
}

func TestExpandSignatureHitSpan(t *testing.T) {
	tests := expandSignatureHitSpanTests

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			// Create temporary file with test content
			dir := t.TempDir()
			filePath := filepath.Join(dir, "test.js")
			err := os.WriteFile(filePath, []byte(tt.Code), 0644)
			require.NoError(t, err)

			// Update hit and expected paths
			hit := tt.Hit
			hit.Path = filePath
			expected := tt.Expected
			expected.Path = filePath

			signatureSpan, err := ExpandSignatureHitSpan(hit)
			require.NoError(t, err)
			assert.Equal(t, expected, signatureSpan)
		})
	}
}
