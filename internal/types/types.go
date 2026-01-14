package types

import "os"

type File struct {
	Path string
	Info os.FileInfo
}
