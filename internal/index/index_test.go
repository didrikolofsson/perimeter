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
