package gitfs

import (
	"io"
	"io/fs"

	"github.com/go-git/go-billy/v5"
)

type GitFile struct {
	file      billy.File
	path      string
	fs        *GitFS
	dirOffset int
}

func (g *GitFile) Read(bytes []byte) (int, error) {
	return g.file.Read(bytes)
}

func (g *GitFile) Close() error {
	return g.file.Close()
}

func (g *GitFile) Stat() (fs.FileInfo, error) {
	return g.fs.Stat(g.path)
}

func (g *GitFile) ReadDir(n int) ([]fs.DirEntry, error) {
	entries, err := g.fs.ReadDir(g.path)

	// directory already exhausted
	if n <= 0 && g.dirOffset >= len(entries) {
		return nil, nil
	}

	// read till end
	if n > 0 && g.dirOffset+n > len(entries) {
		err = io.EOF
	}

	if n > 0 && g.dirOffset+n <= len(entries) {
		entries = entries[g.dirOffset : g.dirOffset+n]
		g.dirOffset += n
	} else {
		entries = entries[g.dirOffset:]
		g.dirOffset += len(entries)
	}

	return entries, err
}
