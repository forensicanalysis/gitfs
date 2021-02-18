<h1 align="center">gitfs</h1>

<p  align="center">
 <a href="https://codecov.io/gh/forensicanalysis/gitfs"><img src="https://codecov.io/gh/forensicanalysis/gitfs/branch/master/graph/badge.svg" alt="coverage" /></a>
 <a href="https://godocs.io/github.com/forensicanalysis/gitfs"><img src="https://godocs.io/github.com/forensicanalysis/gitfs?status.svg" alt="doc" /></a>
</p>

Read a remote git repository as [io/fs.FS](https://golang.org/pkg/io/fs/#FS).

## Example

``` go
func main() {
	// init file system
	fsys, _ := gitfs.New("https://github.com/boltdb/bolt")

	// read root directory
	data, _ := fs.ReadFile(fsys, "README.md")

	// print files
	fmt.Println(string(data)[:4])
	// Output: Bolt
}
```
