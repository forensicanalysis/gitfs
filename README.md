<h1 align="center">gitfs</h1>

<p  align="center">
 <a href="https://github.com/forensicanalysis/gitfs/actions"><img src="https://github.com/forensicanalysis/gitfs/workflows/CI/badge.svg" alt="build" /></a>
 <a href="https://goreportcard.com/report/github.com/forensicanalysis/gitfs"><img src="https://goreportcard.com/badge/github.com/forensicanalysis/gitfs" alt="report" /></a>
 <a href="https://godocs.io/github.com/forensicanalysis/gitfs"><img src="https://godocs.io/github.com/forensicanalysis/gitfs?status.svg" alt="doc" /></a>
</p>

> Read a remote git repository using Go io/fs abstraction

⚠️ This Go package requires [io/fs](https://tip.golang.org/pkg/io/fs) which is part of Go 1.16 (Release in February 2021).

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
