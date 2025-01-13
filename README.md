<h1 align="center">gitfs</h1>

<p  align="center">
  <a href="https://godocs.io/github.com/forensicanalysis/gitfs">
    <img src="https://godocs.io/github.com/forensicanalysis/gitfs?status.svg" alt="doc" />
  </a>
</p>

Read a remote git repository as [io/fs.FS](https://golang.org/pkg/io/fs/#FS).

## Example
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fforensicanalysis%2Fgitfs.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fforensicanalysis%2Fgitfs?ref=badge_shield)


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


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fforensicanalysis%2Fgitfs.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fforensicanalysis%2Fgitfs?ref=badge_large)