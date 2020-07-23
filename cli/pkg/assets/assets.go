// Code generated for package assets by go-bindata DO NOT EDIT. (@generated)
// sources:
// raw-assets/Dockerfile
// raw-assets/baseimages-base.Dockerfile
// raw-assets/baseimages-packages.Dockerfile
package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dockerfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x4f\x6f\xe2\x30\x14\xc4\xef\xfe\x14\x4f\x6c\x8e\xe0\xa0\x3d\x22\x71\x08\x4b\x80\x68\x45\x82\xc2\xb2\x7f\x4e\xc8\x9b\x3c\x82\xa5\xc4\x76\xed\x07\x2d\xb5\xfc\xdd\xab\x42\xa3\x22\xb5\xbd\x8d\x7f\x33\xd6\x8c\x9d\x94\x4b\x98\x25\xdb\x74\x9f\xad\x93\x65\xca\x16\x65\xb1\x86\xe8\x0e\xb0\xd7\xc0\x2a\xd9\xee\x97\x9b\x1d\x4b\xf3\xdf\xbd\x9e\x46\x3d\xbc\xd2\x79\x3a\xcb\x92\x7c\xbf\x28\x8b\xfc\x57\x9a\xcf\xa7\x4a\x2b\xa9\x08\xad\xa8\x48\x9e\x91\x79\x6f\x85\x6a\x10\x22\x39\x84\x48\x2a\x47\xa2\x6d\x61\x32\x05\x9e\xdd\x74\x08\xac\xdc\xe5\xe0\x7d\x6f\x86\xe0\x3d\xaa\x3a\x04\xe6\xbd\x3c\x00\x5f\x09\xb7\xb9\xd0\x51\xab\x12\x1f\x4e\xd2\x62\x87\x8a\x5c\x08\xec\x47\xb1\xf9\x07\x03\xef\xf9\x67\xee\x00\x62\xea\x4c\x6c\xef\x20\xa7\x27\xba\x56\x19\x69\xa0\x1f\x32\xb2\x5f\x04\xfb\x0d\xdf\x60\x91\xfd\x5d\xa7\x13\x20\xec\x8c\xb6\xc2\x5e\x86\x70\x52\x24\x5b\xa0\xa3\x74\x20\x1d\x68\x05\xe6\x62\x24\x68\x0b\x8f\x08\x07\xa9\x6a\x10\xf0\x1f\x89\xd0\xbe\x5f\x02\x67\xf4\xc7\xf6\x23\x91\x71\x93\x38\x76\xa4\xad\x68\x90\x37\x5a\x37\x2d\x0a\x23\x1d\xaf\x74\x17\x5b\x34\xad\xac\x04\xe1\xc8\x5c\x9f\x38\xaa\xf1\x7c\x07\xc7\x7c\xcc\xbf\x73\x12\x96\x37\xcf\xec\xf6\x1d\x1c\xe2\x4a\xd7\xc8\xfe\x14\xe5\xcf\x79\x56\xbe\x9d\x5e\x02\x00\x00\xff\xff\x71\xbb\x60\x22\xe9\x01\x00\x00")

func dockerfileBytes() ([]byte, error) {
	return bindataRead(
		_dockerfile,
		"Dockerfile",
	)
}

func dockerfile() (*asset, error) {
	bytes, err := dockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Dockerfile", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _baseimagesBaseDockerfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x54\x41\x6f\xf2\x46\x10\xbd\xfb\x57\x8c\xf8\x3e\x91\xa4\xca\xda\x6a\xab\x5e\x90\x72\x48\x44\xa2\xe6\x50\x88\xa2\x34\xa7\x4a\xd5\xd8\x1e\xdb\x53\xd6\xbb\x9b\xdd\x31\x04\x28\xff\xbd\x5a\x43\xf9\x0c\x84\x13\xfb\xde\xbc\x37\x6f\x46\xbb\x7e\x7a\x9d\xff\x01\xdb\x6d\xfa\x80\x81\x9e\x5b\xac\x69\xb7\x4b\x92\xc7\xd9\x3b\x4c\x1f\x1f\x9e\xef\x67\x7f\x3f\xbd\xce\x67\x6f\x8f\xb3\xe9\x9d\xb1\x86\x8d\x90\xc7\x42\x78\x49\x49\xf2\x0d\x9e\x2b\x58\xdb\x0e\xb0\x2c\x21\xd8\x96\xa4\x61\x53\x83\x58\x90\x86\x03\x68\x0e\x72\xdb\x73\x08\x85\x6d\x5b\x32\x12\x39\xfa\x74\x1a\xd9\xc0\xaa\x41\x01\x96\xab\x00\x95\xf5\x69\x34\x8b\xff\x6d\x25\x64\xc0\x58\x81\x42\x13\xfa\x7d\x95\xb6\xa6\xee\xed\x22\x0f\x80\x4e\xc0\x61\xb1\xc0\x9a\x02\xb0\x81\xa9\x2d\x16\xe4\x2b\xd6\x14\x00\x3d\x41\x17\xa8\x8c\xa6\xb7\xc9\x37\x40\x53\xee\xd3\xac\x58\x6b\x68\x48\x3b\xe8\x02\x2c\x88\x1c\xb0\x80\x70\xb9\x4e\x93\xd7\x3f\x67\xd1\x53\xd5\x24\xd0\xb9\x12\x85\x40\x7d\xc0\x78\x7c\x04\xd9\x04\x41\xad\x41\x7d\xac\x41\x29\x63\xd5\x01\x50\x9e\xf6\x83\x95\x01\xfe\x4a\xe0\xf0\x6b\x71\x41\x83\x63\xde\xb1\x2e\x15\x85\x40\x46\x18\xf5\x80\xd1\x9c\x87\xa0\x55\x49\xcb\x01\xb8\xd1\x9c\xff\x5c\x9f\x81\x9a\xf3\x7c\xf3\xcb\x25\xe8\x09\x4b\xcd\x86\x2e\x99\xf0\xa1\x59\xe8\xd7\x33\x62\x15\xa7\xf9\x71\x2c\x3a\x7f\x92\x47\x2f\xdb\x53\x17\x53\x74\x3e\x50\xf8\xed\xd2\xff\xc0\xac\xce\xa9\xcf\x8d\xea\x84\xf5\x70\x1f\xb2\xb8\x94\x57\x15\x5f\x82\x7a\xd3\xe2\x19\xea\xd6\xd2\x58\xa3\xac\x23\x13\xc2\x30\x6b\xcd\x27\x83\xa0\x2a\xc8\x0b\x57\x5c\xa0\xd0\xb0\x79\x29\xfe\x73\x70\x1c\x8f\xc1\xb7\xa0\x7c\x05\xd9\x12\x7d\xa6\x39\xcf\xd0\x49\xd6\xdf\xad\xec\xa7\xa4\xbf\x0a\xfd\x56\x1a\x11\x17\x26\x59\xe6\xd6\x64\x96\xa9\xef\x0c\xfc\x0b\x39\x86\xa6\x7f\x18\x2f\xf7\x6f\xbf\xdf\x8d\x32\x6f\xad\x64\x69\x5f\x91\x85\x86\xdb\x30\x39\x81\x72\x36\x93\xef\xb1\x74\xd4\xdb\xc6\xc4\x85\xb6\x86\x8e\xde\x35\x4b\xd3\xe5\x69\x61\xdb\xac\xb5\xad\x55\x1a\xf3\x7d\xbf\xe3\xfd\xd2\x71\x1a\x49\xa3\x74\xf4\xfd\xba\xe7\x20\xb6\xb8\x19\x65\x4e\x77\x35\x9b\xf0\xa5\x20\x8e\x39\xdc\x61\x94\x9d\x55\x6c\xb7\xe9\x4b\xbf\xdb\x77\xf2\x81\xad\xd9\xed\xfa\x90\xfb\xda\x5a\xdb\x1c\x35\xfc\xdf\xf1\x4c\xaa\x94\xf3\x6c\xbe\xb2\xb8\x89\xdf\x84\xb7\xf9\x74\x3e\x01\xc7\x06\x96\x7b\x1c\xae\xff\xe9\xdc\x5a\xc8\x03\x07\x40\x58\x11\xfb\x12\x5a\x12\x54\x87\x47\x0c\xab\x86\x8b\x06\x0a\x34\x57\x02\x39\x45\xad\xa1\xf2\x66\x1f\x88\xdd\xf1\xf9\x1d\x6c\x92\xff\x02\x00\x00\xff\xff\x87\x40\x6d\x1c\xb1\x04\x00\x00")

func baseimagesBaseDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_baseimagesBaseDockerfile,
		"baseimages-base.Dockerfile",
	)
}

func baseimagesBaseDockerfile() (*asset, error) {
	bytes, err := baseimagesBaseDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "baseimages-base.Dockerfile", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _baseimagesPackagesDockerfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0b\xf2\xf7\x55\xa8\xae\xd6\x73\x4a\x2c\x4e\xf5\xcc\x4d\x4c\x4f\xad\xad\xe5\xe2\x0a\x0a\xf5\x53\x28\xc8\x2c\x50\xc8\xcc\x2b\x2e\x49\xcc\xc9\x01\xc9\x07\x54\x96\x64\xe4\xe7\x05\x24\x26\x67\x27\xa6\xa7\x16\xd7\xd6\x72\x01\x02\x00\x00\xff\xff\xb5\xd4\xe0\x91\x39\x00\x00\x00")

func baseimagesPackagesDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_baseimagesPackagesDockerfile,
		"baseimages-packages.Dockerfile",
	)
}

func baseimagesPackagesDockerfile() (*asset, error) {
	bytes, err := baseimagesPackagesDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "baseimages-packages.Dockerfile", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"Dockerfile":                     dockerfile,
	"baseimages-base.Dockerfile":     baseimagesBaseDockerfile,
	"baseimages-packages.Dockerfile": baseimagesPackagesDockerfile,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"Dockerfile":                     &bintree{dockerfile, map[string]*bintree{}},
	"baseimages-base.Dockerfile":     &bintree{baseimagesBaseDockerfile, map[string]*bintree{}},
	"baseimages-packages.Dockerfile": &bintree{baseimagesPackagesDockerfile, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
