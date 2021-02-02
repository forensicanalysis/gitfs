package gitfs

import (
	"io/fs"
	"strings"
	"time"
)

type GitEntry struct{ info fs.FileInfo }

func (g *GitEntry) Name() string {
	if g.info.Name() == "" {
		panic("empty name")
	}
	if g.info.Name() == "/" {
		return "."
	}
	return strings.Trim(g.info.Name(), "/")
}
func (g *GitEntry) Mode() fs.FileMode {
	if g.IsDir() {
		return fs.ModeDir
	}
	return 0
}

func (g *GitEntry) Size() int64 { return g.info.Size() }

func (g *GitEntry) ModTime() time.Time { return time.Time{} }

func (g *GitEntry) Sys() interface{} { return nil }

func (g *GitEntry) IsDir() bool { return g.info.IsDir() }

func (g *GitEntry) Type() fs.FileMode { return g.Mode() }

func (g *GitEntry) Info() (fs.FileInfo, error) { return g, nil }
