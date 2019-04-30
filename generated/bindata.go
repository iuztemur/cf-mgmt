// Code generated by go-bindata.
// sources:
// files/cf-mgmt.sh
// files/cf-mgmt.yml
// files/pipeline.yml
// files/security-group.json
// files/vars.yml
// DO NOT EDIT!

package generated

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _filesCfMgmtSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xd4\x2f\x2d\x2e\xd2\x4f\xca\xcc\xd3\x4f\xcd\x2b\x53\x48\x4a\x2c\xce\xe0\x2a\x4e\x2d\x51\xd0\x4d\xe5\x4a\x4e\x51\x48\xce\xcf\x4b\xcb\x4c\xd7\x2d\x4a\x2d\xc8\xe7\x4a\x4e\xd3\xcd\x4d\xcf\x2d\x51\x50\x71\x76\x8b\xf7\x75\xf7\x0d\x89\x77\xf6\xf7\xf5\x75\xf4\x73\xe1\x02\x04\x00\x00\xff\xff\x4b\xdb\x02\xbb\x43\x00\x00\x00")

func filesCfMgmtShBytes() ([]byte, error) {
	return bindataRead(
		_filesCfMgmtSh,
		"files/cf-mgmt.sh",
	)
}

func filesCfMgmtSh() (*asset, error) {
	bytes, err := filesCfMgmtShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/cf-mgmt.sh", size: 67, mode: os.FileMode(493), modTime: time.Unix(1518199076, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesCfMgmtYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcc\x4b\x6a\xc4\x30\x0c\x00\xd0\xbd\x4e\x21\xb2\xae\x9b\xbd\x2f\x53\x84\xab\x38\x22\xfe\x21\xc9\xa1\xa1\xf4\xee\x25\xc3\x30\x30\xeb\x07\x2f\x84\x00\xa3\x90\x6f\x5d\x6b\xc4\x22\x6d\xfe\x00\x48\xa5\xcc\x5f\xca\xd6\xa7\x26\x8e\x80\xe8\xd7\xe0\x88\xdf\x3d\x1d\xac\xe1\xc1\x80\xf8\x64\xfc\x55\x1e\xdd\xc4\xbb\x5e\x11\x87\x9c\xdd\xa9\x18\xeb\x29\x89\x6d\x4d\x5b\xa8\xb9\xfa\x07\x3a\xe5\x88\x4b\x21\x67\xf3\xe5\x0f\x01\xa4\x8d\xe9\x76\xef\x01\x1b\x55\x8e\x98\x7a\xdb\x24\x87\xbb\x03\xd0\xd9\x6e\x1b\xe4\xfb\x9b\xac\x49\x56\x27\x3b\x5e\xf5\xa7\xed\xf0\x1f\x00\x00\xff\xff\x06\xd9\x38\x16\xc7\x00\x00\x00")

func filesCfMgmtYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesCfMgmtYml,
		"files/cf-mgmt.yml",
	)
}

func filesCfMgmtYml() (*asset, error) {
	bytes, err := filesCfMgmtYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/cf-mgmt.yml", size: 199, mode: os.FileMode(420), modTime: time.Unix(1554829763, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesPipelineYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\xc9\x6e\xdb\x30\x10\x86\xef\x7e\x8a\x39\xc6\x40\xd4\xdc\x75\x33\x62\x25\x30\xe0\x25\xb0\xdc\x14\x45\x11\x10\x8c\x34\x56\xd8\x52\x4b\x49\xca\x85\xdf\xbe\xa0\xe4\xc8\x16\x45\xc5\x6b\x91\x02\xd6\x2d\x30\xc7\x33\xe3\xff\xe3\x90\x7f\x28\x50\xa6\xb9\x08\x50\xba\x3d\x07\x12\x1a\xa3\x0b\x41\x9a\x2c\x59\xe4\x08\xcc\xd2\x1e\x80\x5a\x67\xe8\x42\xc4\x54\x0f\xa0\x0c\x75\x7b\x00\x00\xb9\x60\x2e\xdc\xdc\x44\x4c\x11\x1d\x49\x72\xc1\xfa\xfd\x62\xe5\x55\xd0\x24\x78\xab\x2d\x96\x1f\xf5\xfb\x55\x0d\xc5\x62\x74\x94\x60\x51\x84\xa2\x2a\xa2\x3f\x34\xaa\xb0\x44\xa1\x58\x51\xae\xb3\xed\x7e\xa7\xdf\xef\xf5\x7e\xa6\xaf\xbb\x5d\x0b\xa4\x0a\x9d\x54\x44\xb2\x07\x90\x71\x9a\xe8\x0c\x0e\x44\xa8\xcc\x9f\x04\xb0\xc9\xe2\x82\x12\x39\x16\x61\x8a\xca\x5f\x66\x12\x80\x25\xe3\x75\x41\xee\x02\x76\xa7\x43\xe5\x5d\xb0\x74\xe2\x28\x56\x5f\xd6\x31\x2f\x42\x33\x2a\x68\x2c\xcb\xae\x01\xfc\xef\xfe\xc2\x9b\x90\xe1\x6c\x32\x18\x4d\x75\xf3\x72\x2d\x15\xc6\x24\x4c\x63\xca\x92\x8d\x50\x00\x5f\x7d\x6f\x4e\x46\x43\x1d\x90\x4b\x14\x84\x85\xd5\xd2\xd3\xc0\xf7\xbf\xcd\xe6\xc5\x5a\x46\xa5\xfc\x93\x8a\xed\xe2\xfd\x6c\xfa\x30\x7a\x24\xc3\xd1\x5c\x2f\x97\xfd\x91\x90\x89\x6d\xc0\x78\xe4\x4d\x17\xc4\xf7\xee\xe7\xde\xa2\x88\xe1\x0c\x13\x45\x24\x06\x02\x55\x15\x36\x9e\x3d\x92\xb1\xf7\xec\x8d\x75\x08\x4f\x23\xc2\x71\x85\x7c\x9b\xe5\x81\x4c\x1e\x27\x0b\x72\x3f\x9b\x4c\x06\xd3\x61\x5d\x9f\x4a\x79\xf9\x46\x05\x86\x4e\xf9\xd3\xce\x10\xbf\x91\xa7\xd3\xdf\xd4\xdf\x90\xc8\xdc\xfc\x12\x83\x5c\x30\xb5\x76\x22\x91\xe6\xd9\xf9\x73\xd0\xcc\xd7\x21\x69\x19\x09\x53\xaa\x0a\x0d\x95\x92\x45\x89\x13\xe2\x92\xe6\x5c\x9d\x8f\x48\xcb\x2a\x25\x86\x2e\xfc\xb0\x97\x7e\xd9\xe1\xb8\xb7\x78\xc7\xd3\xe4\xb9\x47\xb2\x8a\x6b\x88\x1c\xcf\xb8\x6f\x8a\x30\xe3\x16\x6c\x9f\xc7\x7a\xb1\x0e\x9a\x09\x6d\x57\x1f\x8b\x23\x70\x32\xc1\x56\xfa\xef\x83\xef\x28\x73\xc6\x74\xe6\x97\x03\xad\x83\xa5\x5a\x47\xac\xdd\x49\x34\xe4\xda\x1a\x0b\x14\x2b\x16\xa0\x43\x83\x00\xe5\x3f\x81\xd6\xa8\xd0\x81\x6a\x58\x8e\xba\x44\x75\xd7\x77\xb1\xe1\x32\x93\xec\xc1\xf6\x41\xed\x8e\xa0\xd5\x34\x7e\x3c\x69\xef\x4e\x22\xa3\x01\x9e\x38\x68\xb7\x90\x67\xe1\x3b\x4d\xad\xcb\x61\xe7\x65\x55\xb2\xe3\xd6\xe6\x2c\x4b\x85\x4c\xe3\x71\x0e\xab\x36\x32\x27\x99\x92\x8e\xe0\x3e\x5b\x62\x12\xdc\x0c\xca\xa9\x04\xcb\xef\x5d\x86\xa1\xd9\x4a\xc7\xd0\x64\x58\x57\xc8\xca\xb0\x3c\xee\xfe\x1f\x90\x55\x3f\x57\x40\x73\x38\x78\x22\xbb\x95\x78\x48\x33\xd2\x28\x57\x84\xf9\xde\xfc\xd9\x9b\x57\x41\xda\xd8\xa0\xb8\xc8\xbe\xd8\x08\x6e\xdf\x1c\xbf\xf3\x54\xd1\x8b\xef\x0e\x0b\xf5\xaa\xd0\x15\x60\x3f\x07\xd6\x46\x27\x3b\xad\xa3\xdf\x67\xec\xd8\x6e\xdb\xdf\x66\x0e\xa4\xd9\xbd\xd5\x1c\x89\xb5\xf5\xa5\xc6\xf4\xa5\x9f\xe9\x99\x2c\xbd\x5c\x01\xd6\xcf\x3e\xa4\xb7\x72\xdb\x36\xc5\xa9\x07\xf4\xfe\xe7\x05\x5b\x91\x2b\xc0\x7d\x3a\x24\xf3\x68\x66\x32\xe5\x54\xb1\x34\x71\x24\x46\x31\x26\xea\x22\x0f\xe6\xe5\xbf\xab\x8d\x6b\xf6\xd8\x59\xb6\x36\xf7\xa9\x78\x0f\x04\xb4\x17\xf4\xb1\x04\x2d\x4a\x6c\x9f\x16\x38\xd2\x24\xcf\x4e\x3a\x7e\xcd\xf1\xbd\x6d\x9a\xae\xcb\x1c\xcb\xb6\x26\xbb\x41\x6d\x3c\x48\x34\x54\xfa\x1b\x00\x00\xff\xff\x12\x0a\xda\xc2\x3e\x20\x00\x00")

func filesPipelineYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesPipelineYml,
		"files/pipeline.yml",
	)
}

func filesPipelineYml() (*asset, error) {
	bytes, err := filesPipelineYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/pipeline.yml", size: 8254, mode: os.FileMode(420), modTime: time.Unix(1556638141, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesSecurityGroupJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\x8e\xe5\x02\x04\x00\x00\xff\xff\x44\xd2\x68\x70\x03\x00\x00\x00")

func filesSecurityGroupJsonBytes() ([]byte, error) {
	return bindataRead(
		_filesSecurityGroupJson,
		"files/security-group.json",
	)
}

func filesSecurityGroupJson() (*asset, error) {
	bytes, err := filesSecurityGroupJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/security-group.json", size: 3, mode: os.FileMode(420), modTime: time.Unix(1510432475, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesVarsYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x31\x8f\xd4\x30\x10\x85\x7b\xff\x8a\xa7\xdb\x7a\xef\x44\x41\x93\xee\xc4\x2d\x12\x12\x82\x13\x82\x3a\xf2\xda\x13\xef\x20\xdb\x13\x8d\x27\xbb\xca\xbf\x47\x4e\xf6\x40\x50\x52\x25\x2f\x79\x6f\xe6\x9b\x77\x58\x65\x51\x24\x36\x28\xcd\x82\x45\xd9\x25\xb6\xb1\x8b\x71\x51\x1e\xfe\xa8\xb3\xfa\x1a\x2e\x03\x8a\x6f\x46\xea\xf6\x60\x98\xd0\xd6\x66\x54\x10\xa5\x78\xae\x6e\x57\xe3\xae\x06\x77\xc0\xd2\x48\xe1\x43\x90\xa5\x1a\x6e\x6c\x17\xcc\xa4\x85\x5b\x63\xa9\x30\x41\x50\xf2\x46\x10\x4d\xed\xa9\xcd\x3e\x50\x73\x3d\x32\x72\xec\xe9\x97\xd3\xeb\xb7\xd3\x87\xe7\xef\xa7\x17\x1c\xf1\xa3\x11\x42\x66\xaa\x36\x36\x0a\x4a\x86\x23\x66\xdf\xda\x4d\x34\x42\xa6\xff\x59\xf5\x16\x1f\xf0\xf0\xe0\x0e\xf7\xe9\xb8\x4f\x9f\x44\xb1\x78\x1f\xf6\x97\x1d\xca\xfd\x05\xd0\x11\x7f\x13\x98\xe0\xcc\x75\x7b\xe6\xe8\x67\x1c\x21\x35\xaf\xa8\x44\x91\x22\xb8\xf3\x71\x4d\xf8\xfc\xf2\xfc\xea\xba\x61\xfc\x67\x79\x96\x94\xba\x21\xd3\x95\xf2\xb6\x33\x4c\xc7\x92\x8a\x21\x48\x29\xbe\xc6\x06\xae\xb0\x0b\x61\xe6\x99\x32\x57\x72\x59\xd2\xb8\xd9\x07\x7c\xfa\xf2\xf1\xab\x3b\xc0\xb8\x10\xb8\x1a\xe9\xd5\xe7\x8e\x62\xca\x29\x91\x62\x99\xa3\x37\x7a\x8a\x94\xc9\x08\x3f\xe5\xdc\x20\xd5\x75\xfb\xf1\x6e\x19\xf0\xee\x7d\x71\xbd\x05\xa9\x13\xa7\x45\xbd\xf5\xe2\x22\x2b\x05\x13\x5d\xdd\xfe\x7d\x8c\xac\x03\x1e\x9f\x76\xd5\xfd\x3e\x67\xb9\xa1\xcd\x14\x78\x5a\xb7\x0b\xfa\xf9\x8d\xf4\x4a\xda\x91\xdf\x70\x71\xdd\x2e\xe8\x7f\x1f\xd7\x92\xf7\x12\x76\xdb\x56\xc1\xaf\x00\x00\x00\xff\xff\xcc\xf5\x97\x32\x8c\x02\x00\x00")

func filesVarsYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesVarsYml,
		"files/vars.yml",
	)
}

func filesVarsYml() (*asset, error) {
	bytes, err := filesVarsYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/vars.yml", size: 652, mode: os.FileMode(420), modTime: time.Unix(1556638287, 0)}
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
	"files/cf-mgmt.sh": filesCfMgmtSh,
	"files/cf-mgmt.yml": filesCfMgmtYml,
	"files/pipeline.yml": filesPipelineYml,
	"files/security-group.json": filesSecurityGroupJson,
	"files/vars.yml": filesVarsYml,
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
	"files": &bintree{nil, map[string]*bintree{
		"cf-mgmt.sh": &bintree{filesCfMgmtSh, map[string]*bintree{}},
		"cf-mgmt.yml": &bintree{filesCfMgmtYml, map[string]*bintree{}},
		"pipeline.yml": &bintree{filesPipelineYml, map[string]*bintree{}},
		"security-group.json": &bintree{filesSecurityGroupJson, map[string]*bintree{}},
		"vars.yml": &bintree{filesVarsYml, map[string]*bintree{}},
	}},
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

