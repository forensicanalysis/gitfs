package gitfs

import (
	"io"
	"io/fs"
)

type PseudoDir struct {
	fs        *GitFS
	dirOffset int
	path      string
}

func (p *PseudoDir) Stat() (fs.FileInfo, error) {
	return p.fs.Stat(p.path)
}

func (p *PseudoDir) Read([]byte) (int, error) { return 0, fs.ErrInvalid }

func (p *PseudoDir) Close() error { return nil }

func (p *PseudoDir) ReadDir(n int) ([]fs.DirEntry, error) {
	entries, err := p.fs.ReadDir(p.path)

	// directory already exhausted
	if n <= 0 && p.dirOffset >= len(entries) {
		return nil, nil
	}

	// read till end
	if n > 0 && p.dirOffset+n > len(entries) {
		err = io.EOF
	}

	if n > 0 && p.dirOffset+n <= len(entries) {
		entries = entries[p.dirOffset : p.dirOffset+n]
		p.dirOffset += n
	} else {
		entries = entries[p.dirOffset:]
		p.dirOffset += len(entries)
	}

	return entries, err
}
