package gitfs

import (
	"io/fs"
	"sort"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

type GitFS struct{ fsys billy.Filesystem }

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
	file, err := g.fsys.Open(name)
	return &GitFile{path: name, fsys: g.fsys, File: file}, err
}

func (g *GitFS) Stat(name string) (fs.FileInfo, error) { return g.fsys.Stat(name) }

func (g *GitFS) ReadDir(name string) (entries []fs.DirEntry, err error) {
	infos, err := g.fsys.ReadDir(name)
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		entries = append(entries, &GitEntry{info})
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })
	return entries, err
}

type GitFile struct {
	billy.File
	path string
	fsys billy.Filesystem
}

func (g *GitFile) Stat() (fs.FileInfo, error) {
	return g.fsys.Stat(g.path)
}

type GitEntry struct{ info fs.FileInfo }

func (g *GitEntry) Name() string { return g.info.Name() }

func (g *GitEntry) IsDir() bool { return g.info.IsDir() }

func (g *GitEntry) Type() fs.FileMode { return g.info.Mode() }

func (g *GitEntry) Info() (fs.FileInfo, error) { return g.info, nil }
