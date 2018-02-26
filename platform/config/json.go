// Code generated by go-bindata.
// sources:
// config/local.json
// config/production.json
// config/staging.json
// DO NOT EDIT!

package config

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

var _configLocalJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x90\xcd\x6a\xeb\x30\x10\x85\xd7\xf2\x53\x98\x59\xdf\xd8\x72\x14\x63\x5d\x51\x4a\x69\x29\x74\x11\x5a\x68\xc9\x03\xc8\x8a\x5c\x8b\xca\x1e\x23\x8d\x09\x34\xe4\xdd\x8b\x9c\xba\xd9\xe9\xe7\xfb\xe6\xe8\xe8\x9c\x31\x18\xf5\x60\x41\xe5\xe0\xd1\x68\x0f\xff\x32\x06\x01\x91\x20\x57\xb0\x6c\x5a\x1d\xed\x21\xf8\x44\xf4\x44\x93\x2a\xcb\x05\xec\x31\x92\x92\x5c\xf2\x05\x9a\xe6\xd6\x3b\xf3\x78\x43\xaf\xa7\x18\x08\x54\x9e\xa8\x65\xac\x8d\x38\x07\x93\xc2\xce\x19\x63\xd0\x39\xbf\xae\x19\x4c\x9a\xfa\x24\x96\x7a\x72\x65\xba\x89\x25\x64\x8c\x5d\x32\x76\x49\xae\x1d\xb4\xf3\x7f\x62\xc0\x21\xb1\x6f\x22\xdf\xeb\x36\xe6\x77\x7a\x72\xd1\xd1\x03\x8a\x62\xb4\x74\xc2\xf0\x75\x0f\xbf\x9e\x3e\xc5\xd5\x0a\xf6\xd3\xe1\x98\xbc\x39\x6e\x4e\x36\xd2\x66\xbb\x52\xa3\x45\x50\x0b\x64\xf0\xfb\x79\x3c\x4e\xe8\x46\x02\xb5\x16\xae\xb6\x4d\xc1\x0b\x5e\x54\xaa\xe6\x7c\x29\xcc\x20\x0e\x3a\xd0\x13\x8e\x14\xb4\xa1\x0f\x13\xdc\x44\x2f\x3a\xf6\xa0\xa0\x6d\x4c\xd5\xc9\x9a\x6b\xbe\xad\xad\xd8\xed\xea\xda\x36\xf6\xbf\x34\xb5\x94\xa6\x91\x3b\x21\x6b\xde\x34\x5d\x5b\x5d\xe7\x74\xb3\xf7\xaf\x78\xb4\x87\xf7\xfd\x2d\xf0\xf6\xc3\x82\x0b\x21\xd2\x33\xb3\xcb\x4f\x00\x00\x00\xff\xff\x34\x79\x81\x13\xac\x01\x00\x00")

func configLocalJsonBytes() ([]byte, error) {
	return bindataRead(
		_configLocalJson,
		"config/local.json",
	)
}

func configLocalJson() (*asset, error) {
	bytes, err := configLocalJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/local.json", size: 428, mode: os.FileMode(420), modTime: time.Unix(1519610869, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configProductionJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8f\xdf\xaa\xd4\x30\x10\xc6\xaf\xd3\xa7\x28\x73\xed\xb6\x3d\xdb\x96\xc6\x20\x22\x8a\xe0\xc5\x41\x41\x39\x0f\x90\xa6\xd3\xd3\x60\x9a\x09\xc9\x94\xa2\xcb\xbe\xbb\xa4\xeb\x9e\xbd\xcb\x9f\xdf\x6f\xbe\xf9\x2e\x85\x00\xaf\x57\x04\x55\x42\x88\x34\x6d\x86\x2d\x79\x78\x57\x08\x88\x44\x0c\xa5\x82\xe3\x32\xea\x84\x2f\xd1\x65\x6c\x61\x0e\x49\xd5\x75\x70\xfa\x4f\x45\x6d\xe5\x91\x77\x8a\xbf\x0f\x2c\x6c\xa3\xb3\xe6\xf3\x03\xbe\xbd\x52\x64\x50\xa5\x6c\x64\x73\x0c\xc6\x44\x5b\x34\x39\xf3\x52\x08\x01\xb3\x75\xf7\xb3\x80\xa0\x79\xc9\x62\xad\x83\xad\xf3\x4f\xaa\xa1\x10\xe2\x5a\x88\x6b\x76\x71\xd5\xd6\xbd\x89\x91\xd6\xcc\xfe\x68\xcb\x67\x3d\xa6\xf2\xc3\x82\xce\xd1\xa7\xc7\x4e\x1f\xe1\xbf\xa6\xf7\x74\x97\x22\xbe\xe6\x86\xaa\x84\x2d\x9d\x76\x4c\x7c\x3a\xdf\x29\x8f\x04\xea\x80\x0c\xfd\xfd\xea\xa7\x40\xd6\x33\xa8\xa3\xb1\xaa\x6b\xc6\xc4\x1e\xf9\xa4\x83\xad\x3c\x92\xdf\xb5\x73\xc8\x95\xa1\x35\xb7\x14\x90\x56\x1d\xf9\x0b\x79\x8e\xda\xf0\x2f\x13\x6d\xe0\x6f\x3a\x2d\xa0\x60\x1c\xcc\xd3\x2c\xfb\x46\x37\xe7\x1e\xdb\xae\xeb\x7b\x1c\xf0\xbd\x34\xbd\x94\x66\x90\x5d\x2b\xfb\x66\x18\xe6\xf1\xe9\x36\x67\xde\x9c\xfb\x4e\x13\xbe\xfc\x7c\x7e\x84\x27\xc4\xa9\xcb\xb1\x15\xc5\x57\x75\x6e\xda\xf6\x58\xbb\xb8\xfe\x0b\x00\x00\xff\xff\x57\x3f\x82\x13\xc2\x01\x00\x00")

func configProductionJsonBytes() ([]byte, error) {
	return bindataRead(
		_configProductionJson,
		"config/production.json",
	)
}

func configProductionJson() (*asset, error) {
	bytes, err := configProductionJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/production.json", size: 450, mode: os.FileMode(420), modTime: time.Unix(1519611139, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configStagingJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8f\xdd\x8a\xdb\x30\x10\x85\xaf\xe5\xa7\x30\x73\xdd\xd8\xde\x38\xc6\xaa\x28\xa5\xb4\x14\x7a\xb1\xb4\xd0\xb2\x0f\x30\x56\x26\x89\xa8\xac\x11\xd2\x18\xd3\x2e\x79\xf7\x22\xa7\xd9\xdc\xe9\xe7\xfb\xe6\xcc\x79\xad\x14\x04\x9c\x09\x4c\x0d\x59\xf0\xec\xc2\x19\xde\x55\x0a\x12\xb3\x40\x6d\x60\xbb\x4c\x98\xe9\x25\xf9\xc2\x5c\x44\x62\x36\x6d\x1b\x3d\xfe\x69\xb8\x6f\x02\xc9\xca\xe9\xf7\x86\xc5\x65\xf2\xce\x7e\x7e\xc0\xb7\x57\x4e\x02\xa6\xd6\x9d\xee\xb6\xc1\x94\x79\x49\xb6\x04\xbe\x56\x4a\xc1\xc9\xf9\xfb\x59\x41\x44\xb9\x14\xb1\xc5\xe8\xda\xf2\x93\x5b\xa8\x94\xba\x56\xea\x5a\x5c\x9a\xd1\xf9\x37\x31\xf1\x5c\xd8\x1f\x7d\xfd\x8c\x53\xae\x3f\x60\x74\xd9\xc9\xa7\xc7\x52\x1f\xe1\xbf\x87\x6b\xbe\x5b\x89\xce\x8e\x43\xf1\x96\xbc\x5b\x29\xcb\x6e\x7f\xa7\x02\x31\x98\x0d\xb2\xfc\xf7\x6b\x38\x46\x76\x41\xc0\x6c\x95\x4d\xdb\x0a\x65\x09\x24\x3b\x8c\xae\x09\xc4\x61\x45\xef\x49\x1a\xcb\x73\xa9\xa9\x20\xcf\x98\xe4\x0b\x07\x49\x68\xe5\x97\x4d\x2e\xca\x37\xcc\x17\x30\x30\x8d\xf6\xe9\xa4\x87\x0e\xbb\xfd\x40\xfd\xe1\x30\x0c\x34\xd2\x7b\x6d\x07\xad\xed\xa8\x0f\xbd\x1e\xba\x71\x3c\x4d\x4f\xb7\x39\xa7\xc5\xfb\xef\x7c\xa4\x97\x9f\xcf\x8f\xf0\x4c\x74\x3c\x94\xd8\x86\xd3\xd9\xec\xbb\xbe\xdf\xd6\xae\xae\xff\x02\x00\x00\xff\xff\x48\xb3\xac\x41\xc0\x01\x00\x00")

func configStagingJsonBytes() ([]byte, error) {
	return bindataRead(
		_configStagingJson,
		"config/staging.json",
	)
}

func configStagingJson() (*asset, error) {
	bytes, err := configStagingJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/staging.json", size: 448, mode: os.FileMode(420), modTime: time.Unix(1519610890, 0)}
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
	"config/local.json": configLocalJson,
	"config/production.json": configProductionJson,
	"config/staging.json": configStagingJson,
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
	"config": &bintree{nil, map[string]*bintree{
		"local.json": &bintree{configLocalJson, map[string]*bintree{}},
		"production.json": &bintree{configProductionJson, map[string]*bintree{}},
		"staging.json": &bintree{configStagingJson, map[string]*bintree{}},
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
