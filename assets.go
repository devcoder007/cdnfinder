// Code generated by go-bindata.
// sources:
// assets/cnamechain.json
// assets/resourcefinder.js
// DO NOT EDIT!

package cdnfinder

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _assetsCnamechainJson = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x96\xc1\x6e\xf3\x36\x0c\xc7\xef\xdf\x53\x18\x39\xec\xf4\xc5\x43\xd1\x43\xce\x69\xba\xad\xc1\xda\x00\x5b\x33\x14\x5b\xd1\x03\x2d\xd3\xb6\x10\x59\x0a\x24\x39\x8e\xf3\xf4\x93\x1d\x5b\x62\x9c\x38\xbd\x85\xfc\xff\x24\x51\x14\xc9\xf8\xf3\x47\x14\x7d\xce\x62\x26\x38\x4a\x6b\x62\x5b\xe9\x44\x25\x8d\x45\x13\x4b\xb4\xb3\x9f\xd1\x6c\xdb\x7a\x9e\x5a\xcf\xec\xeb\xe7\x19\x0e\xd0\x9c\xa5\x32\x66\xaa\x9c\x00\x21\x3b\xb6\x40\xbf\x13\xb1\x3c\xb0\x83\x12\xf8\x00\x2c\x3b\x6b\x24\x62\x9a\xe3\x34\x90\x4a\x73\x77\xb5\x45\x56\x48\x25\x54\xce\xdd\x8d\xfa\x40\x47\x60\x6e\x44\x12\xdb\x84\x01\x2b\xd0\x23\x82\x97\x0e\xf2\x0c\x13\xaa\x4a\x33\xad\xa4\xf5\xa7\x95\x70\x52\x32\x5a\x79\x21\x1c\x2c\x41\xee\xb0\x4f\x4c\x9c\xe8\x8e\xee\x7c\x01\xb1\x76\x9e\x1a\xf0\x7b\x6d\x7f\xd9\x06\xed\xc4\x95\x24\x59\x5b\xb6\xb6\x57\x13\x14\x55\x0e\x24\xeb\x4f\x9d\x63\xf5\xbc\x09\x88\xa8\xb0\x00\xeb\x96\xd7\x4a\xef\x3c\xe7\xbc\xd1\x0b\xd8\x68\x73\xf6\x7b\xdc\x34\xc6\x62\x49\x0e\xfc\xcd\x25\x7c\x05\x26\x5c\xa8\xcb\x4c\x26\x9a\x01\x58\xf5\x76\x00\x52\xb9\x58\x78\xf5\x79\xb3\x58\x8c\x24\xa5\xf3\x6b\x69\x0f\xd2\x16\xa8\xc9\x5d\x9c\x7e\x0e\xce\xd0\xf5\x39\x23\x5b\x8f\xf5\x9c\x91\xc8\x27\x75\x26\xa7\xb7\xe7\x59\x13\x73\xd5\xeb\xce\x08\x1a\xeb\x2a\x63\x88\xad\xe0\x12\xba\x9b\x8f\x81\xe1\xf0\x1b\xc0\xe3\xb9\xa8\xee\x10\xad\xeb\x5b\xe8\x91\xde\xf1\x5a\x17\xc7\xb6\x0b\x68\x9c\xee\x9e\x2b\xd7\xce\xa8\xc3\x2b\xef\x11\x53\x46\xb8\xbf\xfe\x59\x6e\xb6\xeb\xd7\x5f\x27\xf8\x73\xc1\x0b\xd0\xbe\x25\x56\xde\x73\x0b\x1a\xa2\xbb\x86\xda\xfe\x65\xae\x9c\xee\x55\x98\x6b\xe3\xf8\xa6\x50\x03\x9b\x10\xf6\x13\x42\xe6\x8c\x50\xaa\xbf\x77\xd6\x48\x0c\x6f\x36\x92\x73\xa5\x72\x81\xdd\xbe\x7f\x74\x3f\x07\xe5\x2c\x98\x46\xa6\x9c\x81\x75\xfd\x78\x83\x69\x54\x65\xab\xe4\xd6\xea\x7e\xdf\xca\xb8\x62\x77\x83\xc2\x25\x7a\xc8\xe9\x88\x13\x71\xaa\xaa\x44\xa0\x1b\xc6\x6c\x37\xc4\x38\x62\x0a\x9e\xa0\x96\x9c\x8e\x80\x97\xde\x15\x98\x9a\x24\xfb\x85\xe7\x45\xcd\x65\x1a\x8a\x9e\x4b\x06\x7b\x32\x38\xd7\xad\x6d\x2a\x01\x84\x30\x4c\x42\x89\x81\x30\x16\xb4\x6d\xc7\x28\xa3\x90\xac\xbf\x21\x5c\x49\x49\xd8\x93\x68\xd6\xbd\xcb\x33\xbb\x23\xb9\xc9\x9f\xd8\xd0\x49\x26\x0c\xbd\xc8\x2b\x82\xc1\x0f\x4c\x22\x8a\x64\x4a\xd9\xbd\xe6\x61\x2e\xbf\xe2\x01\xc5\x63\xd8\x42\xc8\x3a\xf5\x1a\x2f\x51\xb8\x74\x58\x22\x93\x3c\x5c\xcb\x4e\x48\x25\xd0\x3f\xb9\x37\x38\xd2\xe3\x7b\xc0\x18\x71\x17\x98\x10\x4b\x49\x77\xc6\x94\x83\x54\x87\x8b\x57\xb0\x30\x54\x4c\x1f\xe3\x1b\xd7\x5a\xe9\x68\x5d\x42\x1e\x6a\xa2\xec\x9c\x73\xde\x3a\xef\x82\xee\x9d\xe7\x25\xe7\x77\x19\xdd\x06\xf5\x30\x44\xf5\xb7\x1b\xf4\xc8\x2c\xa6\xd1\xd5\xec\x34\xbc\xdc\x0b\x24\x0f\xf4\xde\x39\x2e\x9e\xc7\xd4\x3c\xb3\x74\xbf\xf7\xd6\x71\x45\xb8\xce\x38\xe0\x05\xf2\xde\x7a\x46\x7f\xd1\xa0\x12\x50\xfe\x6b\xa3\xb3\xe8\x20\x8f\x13\x6e\x73\x0d\x07\x6e\x9b\x00\x59\x88\xdc\xef\xb2\x92\x7d\xe3\x5e\x8c\xfe\xd8\xa2\xc0\x4c\xb5\x9a\x5f\xe1\x3d\x1e\x3c\xa8\xb8\x34\xc8\xa4\x2f\xa3\x0f\xd7\x4d\xaa\x36\xd1\xf2\x54\x91\x39\x07\xcd\x43\x9c\xc4\x0d\x14\xca\xc7\xf8\x6f\x6b\x78\xa0\xe1\x65\x1e\x5f\xbb\x4f\x28\xe9\x27\xce\x7f\x67\x73\xf6\xf5\xe3\xeb\xff\x00\x00\x00\xff\xff\x91\x52\xf6\x3c\x9b\x09\x00\x00"

func assetsCnamechainJsonBytes() ([]byte, error) {
	return bindataRead(
		_assetsCnamechainJson,
		"assets/cnamechain.json",
	)
}

func assetsCnamechainJson() (*asset, error) {
	bytes, err := assetsCnamechainJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/cnamechain.json", size: 2459, mode: os.FileMode(436), modTime: time.Unix(1488991204, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsResourcefinderJs = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x56\x5f\x73\xdb\x36\x0c\x7f\xf7\xa7\x40\xfc\x50\x4b\x67\x5b\x4a\x5f\xe3\x65\xb9\xdd\xd6\x5d\xbb\xcb\xda\xbb\xe4\x31\x73\xef\x68\x09\xb6\x99\x48\xa2\x46\x52\x76\xbd\xcc\xdf\x7d\x00\x49\xfd\xb1\xd7\xb8\x7d\x68\x64\x10\x04\x7e\xf8\xe1\x0f\xb1\x13\x1a\x6a\xb1\x41\xb8\x05\x8d\x7f\x37\x52\x63\x34\xd9\xe3\x8a\x45\x93\x38\xc9\x34\x0a\x8b\x51\x3c\x1b\x01\xd8\x19\x88\x3c\xd7\x68\xcc\x62\xc4\xb7\xe8\x4b\x35\x3a\x43\x43\x57\x5f\x8f\x5e\x66\x0e\xc6\x62\x39\xb4\xe5\x25\x93\x78\x31\x72\x0a\x1b\xb4\x1f\x95\xb1\x95\x28\xd9\xe3\xba\xa9\x32\x2b\x55\x15\x19\xab\x63\x78\x25\x27\x69\x0a\xc6\xaa\xa2\xc0\x0a\xd6\x5a\x95\xb0\xb5\xb6\xbe\x49\xd3\x15\x0a\x9d\x9b\x4c\x0b\x9b\x6d\x51\x9b\x24\x53\x65\xfa\x4c\xde\x2b\x51\xa4\x8d\x91\xd5\x66\xfe\x2c\x76\x82\x14\x64\x6d\xe7\x56\xcd\xc9\xcd\xdc\x6e\x71\xbe\x0d\xbe\xe6\x6a\x3d\x17\xf3\x46\x17\x1c\x87\x3e\xb0\x27\x00\x1f\x04\xc1\xa8\x70\x0f\x0f\xb8\xf9\xf0\xad\x8e\x26\x5f\xa3\xbb\x9b\xf5\xbf\x5b\x1b\xdb\x9a\xbe\x4c\x7c\xf7\x17\xb9\x8f\x9e\xbe\xa6\xcb\x69\x3c\x99\xc1\x44\xba\x58\xf8\xba\x46\x4b\x00\x08\xae\x4e\x4a\xc6\x15\x69\x8c\x9f\xde\x2f\x13\xab\x1e\xad\x26\x48\x91\xd3\x3b\x42\xc6\x87\x10\xa1\x0e\x21\x72\x90\xb5\x56\x2b\xb1\x2a\x0e\x90\x0b\x2b\xa0\xd1\x12\xf6\x5b\x49\x5a\x7b\x84\x5c\x55\x96\xee\x10\x30\xb1\xb2\x41\x3d\x53\x95\x51\x05\x26\x85\xda\x38\xaa\x86\xfe\xab\xa6\x28\x9c\xa3\xd1\x71\x34\x4a\xd3\xcf\x8a\xc2\xc1\x1c\xd6\x4a\xc3\x16\x45\x4e\x6c\xb9\x6f\x14\x64\x5e\x56\xb9\xdc\xc9\xbc\x11\x05\xa8\xd5\x33\x66\x96\x2e\x3c\xca\xb2\x2e\xe4\xfa\x00\xaa\xb1\x75\xc3\x92\x3f\x1d\x60\xff\x93\xef\x52\x74\xa0\xd6\x2e\x15\xc9\x47\x6f\x32\x14\x40\xde\x64\xd8\x3a\x19\x64\x33\x88\x62\x8e\x96\x15\x89\xdf\x5e\xeb\xf5\x18\xa4\x2f\x78\x98\xd1\x47\xd1\xe0\x0c\xe4\x22\x08\xb3\xbc\xe2\x84\x84\x98\x4e\x23\xff\xe3\xf1\xcb\xe7\xc4\x38\x6e\x09\x6f\xe7\x85\xd9\x20\x94\x91\xbc\xbd\x5e\xc8\x9f\x82\x34\xa1\x0a\xda\xd8\xed\x42\x4e\xa7\xb1\x27\x9d\xdc\x91\xe5\x70\xfc\x24\x97\x09\xd7\xc5\x22\x14\x02\x81\x38\x3d\x74\xa2\x05\x91\x01\xfe\xdf\x10\x47\x0f\xdc\x27\xa2\x0f\xef\x89\x8e\x96\x64\xe8\xc9\x9d\x2e\x7d\x5a\xfa\x4c\x75\x7a\x0b\xce\x15\x87\x5b\x8a\x17\xd4\x58\x2b\x6d\x87\xfc\xc9\x8a\x98\xef\xd8\x73\xee\xe8\x3f\x33\x83\x95\x30\xc8\x9d\x99\xab\x52\xc8\x6a\x16\x72\xe4\xe8\xf3\x48\x87\x30\xc7\x64\x9c\xa8\x02\x6f\x7f\xcc\x58\xd9\x0a\x39\xfa\xe2\x72\x9f\xf0\x2f\xef\x2b\xe9\x7a\x39\xfe\x21\xed\x7c\xab\xe5\x1c\x02\xe9\x2c\x7b\x8b\x71\x3e\x23\x46\x3d\xd3\x67\xde\x1c\x5f\x49\x5f\x1b\x27\x15\x75\x0e\xed\x44\x39\xbe\x60\x4f\x1a\xe6\x89\xcc\xf9\xd3\x96\x34\x9e\x05\x70\xeb\x00\x5d\x02\xd3\x4f\xa7\x4e\x31\x4d\x73\x2c\xd0\xe2\x45\xf8\x6d\xb2\xd3\xb4\x52\x39\x3e\x1b\x10\x75\x4d\x01\x89\xdc\xf8\x49\x16\x28\x1d\x9d\x26\xe9\x8c\x5c\x9f\xf9\xd8\xd5\xc7\x88\x6b\xe2\x37\x1a\xbf\x49\xa5\xf6\x3c\x4c\xc2\xf4\x25\xa9\x1f\xaa\x89\xd0\x1b\x43\x13\x87\x46\x2b\xc7\x97\xa8\xea\x21\x20\x7b\xc0\x0c\xe5\x8e\xa6\xc0\xa0\xa8\x78\x22\xa3\xe9\xcb\x8a\x86\xe1\x0c\x8c\xfc\x87\xba\xaf\x8d\x79\xd6\x76\xc0\xa0\x21\x87\xc9\x71\x06\xfa\x78\xbb\xe6\xb8\x14\x51\xeb\xd6\x25\x8c\x7c\x0e\x0c\xd1\x2f\x16\xca\x35\x44\x57\x11\x23\x89\x43\xe1\xf0\x37\xe9\xb5\x77\x93\x95\xca\x0f\x8f\x2c\xbb\x83\x73\xd1\x0d\x5c\xc7\x3d\xf3\x43\x1c\x64\x3d\x5e\x00\x89\x07\x19\x1d\xbc\x3e\xfe\x3c\xb8\x8f\x5a\x9d\x18\xde\xbd\x03\x87\x05\x7e\x26\xcb\x01\x8f\x47\xd8\xa7\xbd\xd5\x5e\xb6\x0a\x00\xdf\x39\x0c\xcf\xe2\x9b\xe7\xf4\x88\x35\x15\xe7\xf8\xfa\x92\xd2\xea\x60\xdd\x0b\x1b\x94\x8e\xa3\xcb\xd6\xa6\xb7\xf0\xbe\x2d\xda\x7a\x2b\x2a\xab\x4a\x2a\xc5\x42\xa2\xb9\x02\xa3\xf8\x7d\x31\x88\x9c\x2f\x8b\x95\x9d\xfb\x96\x0d\x39\x06\x49\x35\xbb\x13\xb2\xa0\x87\x09\x9d\x89\x41\x87\xbf\x39\x56\x3d\x3b\x67\x63\x95\xde\xc0\x7b\xb5\x47\xfd\x2b\xb5\x5e\x14\x73\xd3\x8d\x4f\x5d\x8e\xbb\xdb\x5d\xb6\x6b\xa1\x0d\x7e\xaa\x6c\x74\x3e\x85\xe3\x45\xa7\xba\xa2\x7e\x7a\x69\x7f\x1e\x03\x21\x6f\x32\xe2\xa9\x23\x46\xd8\x43\x4b\x8a\x11\x3b\x04\x5a\x0d\xa0\x10\x34\x0e\xe8\x5a\x4d\x35\x83\x5d\x99\xd7\xc4\x03\x5b\x78\xd3\x68\xdf\x0f\x27\x7d\xef\xde\xde\x47\x49\xd1\x59\x7a\xd5\xe5\xa6\x52\xf4\x80\x13\xf3\xf4\xea\xb7\xcd\xf9\x41\x6b\xe2\x73\xd0\x91\xa5\xd9\xcc\x68\x1b\x11\x19\xf2\x62\x70\x6c\xbb\xb8\xc6\x2a\x0a\xad\x3e\xeb\x94\xa9\x28\xad\xb0\x8d\xf1\x2b\x04\x37\xa6\x9f\xfd\x6d\x05\xfb\x53\xb8\x22\xaa\x27\xa6\xc9\x08\xb3\x99\xb4\xeb\xc6\xb0\x2b\x26\xaf\x63\x64\x1c\xe3\x1b\x18\xff\xfe\xcb\xa7\xfb\xf1\x71\x12\xd6\x14\x2c\x88\x06\x7f\xe1\x74\xf0\xc0\x1c\xec\xb0\xd9\xff\x67\xf2\x5e\x89\x9c\x1f\x1a\x2b\xa9\xc9\x26\x30\xa5\xfb\x53\xfa\x5b\x1a\xcc\xda\x5d\x29\x2c\x13\x7d\x47\x78\xc1\xd9\x6c\x76\x8b\x68\x82\x9c\x75\xde\x3a\xfb\xd8\xdb\x48\x06\x6b\x4f\xae\xb2\xa6\x24\xb6\x09\x02\x2d\x58\xa4\xd5\x4d\xee\xd0\x28\xc1\xf3\xd0\xdb\x70\x63\xed\xbe\xbd\x5a\xff\x10\x47\x5e\x39\x5e\x7c\x67\xf7\x3a\x9b\x6d\x41\xb3\x1b\x40\xa1\xe1\x12\xfc\x26\x2d\x0f\x6c\xc6\xf0\x5f\x00\x00\x00\xff\xff\x93\xab\x08\xd6\x60\x0b\x00\x00"

func assetsResourcefinderJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsResourcefinderJs,
		"assets/resourcefinder.js",
	)
}

func assetsResourcefinderJs() (*asset, error) {
	bytes, err := assetsResourcefinderJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/resourcefinder.js", size: 2912, mode: os.FileMode(436), modTime: time.Unix(1488956466, 0)}
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
	"assets/cnamechain.json":   assetsCnamechainJson,
	"assets/resourcefinder.js": assetsResourcefinderJs,
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
	"assets": &bintree{nil, map[string]*bintree{
		"cnamechain.json":   &bintree{assetsCnamechainJson, map[string]*bintree{}},
		"resourcefinder.js": &bintree{assetsResourcefinderJs, map[string]*bintree{}},
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
