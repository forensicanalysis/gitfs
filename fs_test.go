package gitfs

import (
	"io/fs"
	"reflect"
	"sort"
	"strings"
	"testing"
	"testing/fstest"
)

func TestFS(t *testing.T) {
	fsys, err := New("https://github.com/forensicanalysis/fslib")
	if err != nil {
		t.Fatal(err)
	}

	fs.WalkDir(fsys, ".github", func(path string, d fs.DirEntry, err error) error {
		return err
	})

	var names []string
	entries, err := fs.ReadDir(fsys, ".github")
	for _, entry := range entries {
		names = append(names, entry.Name())
	}

	err = fstest.TestFS(fsys, "LICENSE")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNew(t *testing.T) {
	want := strings.Split(".gitignore LICENSE Makefile README.md appveyor.yml " +
		"bolt_386.go bolt_amd64.go bolt_arm.go bolt_arm64.go bolt_linux.go bolt_openbsd.go " +
		"bolt_ppc.go bolt_ppc64.go bolt_ppc64le.go bolt_s390x.go bolt_unix.go bolt_unix_solaris.go " +
		"bolt_windows.go boltsync_unix.go bucket.go bucket_test.go cmd cursor.go cursor_test.go " +
		"db.go db_test.go doc.go errors.go freelist.go freelist_test.go node.go node_test.go " +
		"page.go page_test.go quick_test.go simulation_test.go tx.go tx_test.go", " ")
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"New", args{"https://github.com/boltdb/bolt"}, want, false},
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
