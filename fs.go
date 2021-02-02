package gitfs

import (
	"fmt"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"io/fs"
	"sort"
)

type GitFS struct {
	fs billy.Filesystem
}

func New(url string) (*GitFS, error) {
	repository, err := git.Clone(memory.NewStorage(), memfs.New(), &git.CloneOptions{URL: url})
	if err != nil {
		return nil, err
	}

	worktree, err := repository.Worktree()
	if err != nil {
		return nil, err
	}
	return &GitFS{worktree.Filesystem}, err
}

func (g *GitFS) Open(name string) (fs.File, error) {
	info, err := g.Stat(name)
	if err != nil {
		return nil, err
	}
	if name == "." || info.IsDir() {
		return &PseudoDir{fs: g, path: name}, nil
	}
	file, err := g.fs.Open(name)
	return &GitFile{path: name, fs: g, file: file}, err
}

func (g *GitFS) Stat(name string) (fs.FileInfo, error) {
	if !fs.ValidPath(name) {
		return nil, fmt.Errorf("invalid path: %s", name)
	}
	info, err := g.fs.Lstat(name)
	return &GitEntry{info: info}, err
}

func (g *GitFS) ReadDir(name string) (entries []fs.DirEntry, err error) {
	if !fs.ValidPath(name) {
		return nil, fmt.Errorf("invalid path: %s", name)
	}
	infos, err := g.fs.ReadDir(name)
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		e := &GitEntry{info}
		entries = append(entries, e)
	}

	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })
	return entries, err
}
