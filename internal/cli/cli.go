package cli

import (
	"flag"
)

var (
	Path string
)

func init() {
	flag.StringVar(&Path, "path", ".", "Path to project root")
	flag.Parse()
}
