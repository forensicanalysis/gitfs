package gitfs_test

import (
	"fmt"
	"io/fs"

	"github.com/forensicanalysis/gitfs"
)

func Example() {
	// init file system
	fsys, _ := gitfs.New("https://github.com/boltdb/bolt")

	// read root directory
	data, _ := fs.ReadFile(fsys, "README.md")

	// print files
	fmt.Println(string(data)[:4])
	// Output: Bolt
}
