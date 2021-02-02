package gitfs

import (
	"fmt"
	"io/fs"
	"reflect"
	"sort"
	"testing"
	"testing/fstest"
)

func TestFS(t *testing.T) {
	fsys, err := New("https://github.com/forensicanalysis/fslib")
	if err != nil {
		t.Fatal(err)
	}

	fs.WalkDir(fsys, ".github", func(path string, d fs.DirEntry, err error) error {
		fmt.Println("path", path)
		return err
	})

	var names []string
	entries, err := fs.ReadDir(fsys, ".github")
	for _, entry := range entries {
		names = append(names, entry.Name())
	}
	fmt.Println(names)

	err = fstest.TestFS(fsys, "LICENSE")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNew(t *testing.T) {
	want := []string{
		".github", ".gitignore",
		"bufferfs", "example_test.go", "fallbackfs", "fat16", "fsio", "fslib.go", "fslib_test.go",
		"fstest", "go.mod", "go.sum", "gpt", "LICENSE", "mbr", "ntfs", "osfs", "Readme.md",
		"registryfs", "replace", "systemfs", "testdata",
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"New", args{"https://github.com/forensicanalysis/fslib"}, want, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsys, err := New(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var names []string
			entries, err := fs.ReadDir(fsys, ".")
			for _, entry := range entries {
				names = append(names, entry.Name())
			}
			sort.Strings(names)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(names, tt.want) {
				t.Errorf("New() got = %v, want %v", names, tt.want)
			}
		})
	}
}
