// Code generated by go-bindata.
// sources:
// templates/README.md
// templates/identities_registry.gotpl
// templates/model.gotpl
// templates/relationships_registry.gotpl
// DO NOT EDIT!

package static

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

var _templatesReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xcb\x11\x02\x21\x10\x45\xd1\x7d\x47\xf1\x2c\x53\x22\x81\xc6\xbe\x0a\xc5\x47\x6a\x60\x33\xd9\xcf\x79\x2b\x31\x56\xf7\xc3\x36\x4b\x85\x8d\xbe\xb5\xa3\x09\xa1\xf3\x57\x46\x8c\x4c\x04\xa1\x3a\x75\x0a\xca\x75\xfa\x75\xbf\xcc\x24\x69\x78\x43\xcb\x3f\xcd\x7f\xd8\x13\x00\x00\xff\xff\xaa\x97\xff\x85\x4d\x00\x00\x00")

func templatesReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_templatesReadmeMd,
		"templates/README.md",
	)
}

func templatesReadmeMd() (*asset, error) {
	bytes, err := templatesReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/README.md", size: 77, mode: os.FileMode(420), modTime: time.Unix(1528494836, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesIdentities_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x57\x4b\x6b\xe4\x38\x10\xbe\xeb\x57\xd4\x86\xb0\xd8\xd0\x71\x2e\xcb\x1e\xb2\xe4\x10\x86\x19\x68\xd8\x0c\x21\x81\xbd\x34\x39\x68\xdc\x65\x47\xac\x2c\x19\x49\xce\x4c\x63\xf4\xdf\x17\xc9\xf2\x43\x8e\xd3\x8f\xd9\x24\x0b\x9b\x4b\xda\x56\x3d\xbf\x2a\x7d\x55\xae\x69\xfe\x37\x2d\x11\xda\x16\xb2\x07\x34\xd9\x27\x29\x0a\x56\x36\x8a\x1a\x26\x45\xf6\x95\x56\x08\xd6\x12\xc2\xaa\x5a\x2a\x03\x67\xa5\xcc\x68\x2d\x15\x1a\x99\x31\x79\x89\x1c\x2b\x14\x86\xf2\x33\x42\x9e\xa9\x82\x84\x00\x00\xb0\x2d\x0a\xc3\xcc\xce\x29\xeb\x5b\x5a\xc3\x35\x54\xb4\xde\x68\xa3\x98\x28\x1f\x07\x9d\x6c\x1d\xe4\xa0\xf5\x6a\xee\xaf\x6d\x2f\x40\x51\x51\x62\x17\xcc\x43\x8d\x39\x2b\x58\xee\x83\xd1\x2e\x90\x51\x10\x58\x01\xfa\x49\x36\x7c\x7b\x8f\x25\xd3\x06\x55\x24\x0d\x19\x9c\x67\x77\xcd\x37\xce\xf2\x5b\xb9\xc5\x58\xf7\x02\xce\xc7\x10\xe1\xea\x1a\x32\x27\xc3\xb3\xcf\xe3\xcb\x8b\x89\xc2\x59\xdb\x06\x81\x7b\xd4\xc6\x1d\x5b\x7b\x76\xe5\x62\x98\x9a\xb1\xb6\x4f\x68\x15\xb9\x42\xb1\x9d\x7b\x9f\xbc\xb2\x24\xc2\x2c\xa7\x06\x4b\xa9\xd8\xff\x10\x38\xd9\xa8\x1c\xdf\x05\x3c\xca\x19\xd5\xef\x88\xd8\xc5\x07\x42\xd6\xb6\x21\xaa\x73\xb6\x82\x73\x9f\xd9\x44\xe9\xa6\xcb\x14\x62\x8c\x7b\xb9\xb7\x03\x76\x4f\xa3\x8a\x2d\xfe\x58\xc0\x7a\xf3\xb8\x79\xec\x7e\x7e\x38\xc6\x0e\x81\xd9\xfd\x1c\xa0\x60\x05\x70\x14\xfd\xf1\xba\x0b\x1e\xac\x5d\x0c\x37\x0e\xd9\x17\x20\x97\x55\x2d\x1b\xb1\x9d\xd4\x60\x34\x12\x29\xf6\xf6\xda\xa9\x89\x41\xdd\x5a\x1f\xa5\xfb\xbf\x1a\x81\x05\xbb\xb7\x28\x76\xd5\xb6\x80\x5c\xbb\x74\x04\xe3\xab\x9f\xac\x5b\x4a\xc8\xe5\x25\xf8\xe0\xff\x42\xa5\x1d\x98\x0a\x4d\xa3\x84\x06\xf3\x84\x90\x37\x4a\xa1\x30\xf0\x1c\xce\x64\xe1\x5f\x57\x3e\x59\x52\x34\x22\x8f\x74\x93\x14\x0a\x2e\xa9\xf9\xfd\x37\x68\x83\x9d\x61\x78\xdc\xdc\xad\xd7\xa2\x90\x59\xef\xc6\x65\x48\x88\xd9\xd5\xc1\xdc\x2d\x15\xb4\x44\x05\xda\xa8\x26\x37\xad\x25\xde\x7c\x52\x44\xa7\x29\xf4\x1d\xfb\x45\xc9\xca\x55\x33\x11\xae\xa4\x1d\xbc\x29\x2c\xde\x6a\x9f\x6a\x88\x66\x3e\x7f\x36\x4e\xfd\x91\x1c\xe3\xed\x53\x47\xbf\xbb\x24\xf0\xf0\xee\x74\xaf\x11\x83\x6f\x7a\x3b\xc7\xb9\xf7\xd7\x3b\xe9\x2e\xf3\xd1\x8e\x47\xf6\xdb\xf8\x9f\x47\xba\x12\xbb\x84\x8a\x31\xbf\x84\x2d\x78\x4a\x7b\x57\xac\x00\x06\xd7\x50\x64\x2f\x4a\x43\xc5\x2e\xfd\x03\x7e\x61\xd9\x5a\x7f\xae\x6a\xb3\x4b\xd2\xc9\x85\xea\xa1\x89\x08\x64\xc9\xd4\x80\xfb\xc9\xe6\xc2\xbb\xd8\x5c\xc0\x51\xec\xd2\x03\x58\x14\x8c\x7e\xe3\x98\xf4\xb5\x5b\x84\x60\xfe\xae\xd3\xe9\x91\xd1\xdf\x99\xc9\x9f\x86\xea\x87\x68\x07\x16\xdf\xc3\x7b\x3f\xcd\x79\x39\xd5\xdd\xbe\xf6\x62\x90\x8c\x64\x7f\x35\x07\xed\x2b\x7e\x7f\x45\x25\x49\xc9\x02\x6d\xcc\x1e\xb7\x58\xd0\x86\x9b\x17\x66\x05\xe3\xa1\x1a\xaf\x02\xdd\xd1\xe5\x7e\x8c\x23\x3a\x8e\x6e\xd5\x30\x6e\x36\xbd\x01\xbf\x94\x1e\x6a\xf2\xae\x48\xae\x17\x1e\xbc\xd9\xa8\xd7\xf7\x57\x74\xd6\x51\xa1\x45\x66\x0d\xd6\xdd\x9e\xf4\xb8\xfe\x3a\x90\xfc\x72\x38\xfa\xcd\x3a\x4c\x48\x33\xcc\x2e\x7d\x2f\xa5\xf9\x6f\x3a\xf0\xd7\x25\x85\x3b\xde\x28\xca\xc1\xda\x3f\x99\x76\x03\xe1\x70\x27\xbe\x5d\x63\x4e\xc1\x3e\xa9\x53\xf4\xde\x56\xd1\x27\xf7\xca\x3d\xf2\xae\x70\x4f\xac\xd6\xc9\xd4\x6b\x74\xd2\x55\x48\xcd\xf9\x5f\x2d\xc9\x38\x5f\xee\x8b\xac\x0a\x13\xf7\x3a\x72\xe9\x26\xaf\x5b\x09\xc2\xe1\x74\x1b\xf0\x62\x93\x08\x6e\x27\x6a\xfd\x36\xd0\x3d\x45\x81\x4e\xc5\xc6\xbd\xa0\xf7\x6e\x89\xf7\x77\xc3\x79\x00\x86\xa1\x1e\xbc\x52\xce\x01\x7f\x30\x6d\xdc\xed\x67\xc3\x79\x70\x16\xe9\x24\x8e\x29\x0e\x0e\xc4\x25\x91\x8f\xdf\xfb\xf7\x5e\x8e\xd3\x3e\x72\x6c\x40\xcf\xcf\xf9\x2f\x52\x0d\x79\x4f\x21\x74\xc5\x0b\xab\x00\x14\x52\xf9\xe7\x92\x3d\xe3\xb8\x99\x0c\x88\xce\xed\x1c\xe2\xe6\x98\x99\x5f\xe3\xa3\x23\x60\xfd\x58\xc2\x99\x6e\xe4\xff\xea\xa3\x6a\xff\x92\x7e\x04\x49\xc5\xbb\x8a\xa3\x26\x4b\xfe\x09\x00\x00\xff\xff\xec\x98\x22\xd7\x79\x11\x00\x00")

func templatesIdentities_registryGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesIdentities_registryGotpl,
		"templates/identities_registry.gotpl",
	)
}

func templatesIdentities_registryGotpl() (*asset, error) {
	bytes, err := templatesIdentities_registryGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/identities_registry.gotpl", size: 4473, mode: os.FileMode(420), modTime: time.Unix(1537418981, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesModelGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5b\x6f\x1b\x37\x16\x7e\xd7\xaf\xe0\x0a\x6e\x2a\x15\xf2\x38\xcf\xea\xba\x40\xd6\x49\x0a\x63\xe3\xd6\x88\xb3\xe9\x43\x10\xd4\xf4\xcc\x19\x99\xcd\x88\x54\x49\x8e\x63\xaf\x30\xff\x7d\xc1\xdb\x0c\x39\x17\x89\xe3\xd8\x81\x17\x70\x1e\x5a\xeb\x90\xfc\xce\x95\x87\x87\x97\xd9\xe0\xf4\x0b\x5e\x01\xda\x6e\x51\x72\x01\x32\x39\x61\x34\x27\xab\x92\x63\x49\x18\x4d\x7e\xc3\x6b\x40\x55\x35\x99\x90\xf5\x86\x71\x89\x66\x13\x84\x10\x9a\xe6\x6b\x39\x35\x7f\x89\x3b\x9a\x4e\x27\xe6\xef\x15\x4b\xf0\x86\x71\x90\x2c\x21\xec\x08\x0a\x58\x03\x95\xb8\x30\x3d\xb7\x5b\xc4\x31\x5d\x01\x4a\x2e\x36\x90\x26\x1f\xee\x36\x70\xce\xd9\x0d\xc9\x80\x0b\x74\x58\x55\x06\x42\x49\x81\xaa\xaa\x1e\x02\x34\xd3\x8d\x7d\x10\x1f\x71\x41\x32\x2d\xe6\x08\xa0\xf9\x64\xb2\xdd\xa2\x83\x02\x4b\x10\xf2\x23\x70\x41\x18\x45\xcb\x63\x0b\xf9\x4e\x93\x5f\x49\xc9\xc9\x55\x29\x41\xb8\x0e\x4a\x82\x9a\xfb\x01\x59\xa0\x03\xa0\xe5\x5a\x8d\xbb\x2a\x49\x91\xbd\xa1\xe5\x5a\x18\x88\x36\x74\x55\x4d\x8e\x8e\x94\x00\x7a\x84\x56\x1b\x55\x15\xe2\xb0\xe1\x20\x80\x4a\x81\xe4\x35\xa0\x0d\x13\x82\x5c\x15\x80\x6e\x70\x51\x82\x40\x39\xe3\x08\x3b\x29\xb4\x32\x66\x78\x2d\x99\xf5\xcb\x34\x99\x48\x85\xd8\xc1\x17\x92\x13\xba\x9a\x4c\x52\x46\x85\xf3\xda\x76\x7b\xe8\x34\xa0\x78\x0d\x0b\x74\xa0\xb9\x29\x2d\xcc\xe0\x8f\x86\xb9\x35\xa1\x15\x9b\x1a\x4e\x6d\x89\xcd\x50\xd5\xc1\xfc\x55\x55\x89\x33\x75\x3d\xa4\x23\xd5\xb1\x51\xc5\x8d\x08\x9c\xa3\x7d\xd3\xfc\x3d\x51\xd2\x92\x1c\x51\x26\xad\x6f\xce\x58\x06\x45\xf2\x1a\x24\x4e\xaf\x21\x6b\x0c\xeb\xb7\xbe\xa1\x92\xc8\x3b\x6b\x9c\xd3\x0c\xf4\xcf\xb6\xe8\x35\x9d\xe5\xfa\x37\xbb\xfa\x0b\x52\x99\x4c\x6e\x30\x8f\xc3\x3b\x46\x75\x6c\x27\x35\x71\xab\x95\x51\x5d\x97\xa8\x8e\x40\x0f\xea\x3d\x08\xa9\x5a\xab\x6a\xba\xd0\x5d\x4f\xb0\x84\x15\xe3\x77\xcb\xbe\xae\xac\xe4\x69\xed\x64\xd3\xff\xdc\x4c\xd4\x65\x17\xda\xb6\x34\x3d\x39\xb9\xc1\x52\xf5\x6c\x77\x34\x0d\x55\xb5\x98\x54\x23\x6d\xbd\xdd\xf6\xf5\x38\x15\xef\x19\x93\xfb\x7c\x71\x5e\x94\x1c\x17\xa8\xaa\xde\x11\x21\x7d\x6f\x60\x54\x28\x0a\xcb\x23\xc6\xd6\x81\x1e\xc3\xe3\xd3\xe7\x9f\x06\x7b\x2a\x85\x8f\x8e\x90\x17\x1d\xb2\xe4\xd4\x84\x06\xe9\x0d\x0d\x81\x08\xd5\x3f\x95\xb4\xc9\x24\x2f\x69\x8a\x66\x2c\x52\x96\x79\xcd\x69\x36\xef\x8f\x1b\xed\x33\x23\xc5\x30\x66\x13\x7e\x13\x23\xff\x09\xdb\x34\xb2\x63\xb4\x61\x84\x4a\xe0\x48\x32\x84\x51\xaa\xda\x94\xc0\x71\x22\x8e\x57\x49\x31\xef\x51\x27\x27\xf8\xaa\x00\xe1\x74\xd2\x62\x2c\x8f\x11\xde\x6c\x80\x66\xb3\x38\xf0\x6d\xb5\x40\x2c\x49\x92\xb9\x6f\x96\x17\x0a\xca\x2a\xfe\x4a\xa3\x59\x50\x11\xb8\x49\x32\xfd\x13\x23\x0a\x5f\x0d\x77\xeb\xc7\xc7\xb2\x83\x91\x65\xe6\xf8\x27\x49\xd2\x6f\x92\xbd\xa6\x62\xa5\xfc\x46\x4b\xa9\x25\xe3\xcf\x85\x32\x85\x02\x32\x79\xde\xc9\x65\x72\x93\xe3\x53\xb3\x61\xa5\xd4\x03\x92\xd9\xae\xd9\x32\x37\xf8\x55\x10\xa7\xac\x94\xd6\x1d\x7a\xbe\xa5\x8c\xde\x00\x97\xbe\x37\x74\x24\xd2\x21\xbd\xef\x67\x6e\xf5\xdf\xe1\xb0\xd3\x92\x84\xf6\xdc\xd1\x73\x5b\xf9\x66\x23\x12\xd6\x9e\xdd\x76\x5a\x4c\xf5\xdd\x6d\x93\xd7\x90\xe3\xb2\x90\xbf\xf3\x0c\x78\x90\x5f\x32\xd3\x80\x98\x6a\x21\x74\x85\x72\x02\x45\x26\x5c\x98\xa6\x8c\x4a\xa0\xf7\x30\x8c\xcf\x70\x36\x47\x9f\x3e\x9b\xf5\xbf\x95\x5c\x1c\xb9\x51\xae\x55\x51\xfd\x6e\xc5\x6a\x0a\xa0\x76\x39\xe3\xea\xab\x7a\x85\x0b\xd7\xa8\x70\x3d\x37\x26\x32\x26\x71\x00\xbe\x35\x6e\x2c\xed\x5b\xb5\xb7\xd8\xb3\x39\x22\x54\xf6\x25\x54\x90\xc9\xab\xf3\xd3\x53\x9a\xb3\xc4\xab\xcb\x4c\x4d\x67\x65\x55\x4b\xa0\x5b\x0e\xf7\xd4\x15\xed\x7a\x62\xad\xba\x28\x25\x70\x7b\x98\x5b\xf3\xf7\x2f\x5f\xb6\x66\x2b\x53\x25\x7f\xcb\x2b\xbb\xbd\xb1\xdd\xea\x4a\xf1\x03\x7b\xab\x42\x09\x1d\x28\x6d\x51\x8e\x0b\x01\xba\x00\x0e\x14\x53\x76\xd1\x9c\x1d\x80\xb2\xd7\xe5\x5f\x82\xd1\xe5\xf4\x70\x8a\xae\xf4\x1f\x7f\x6a\x7d\xac\x6f\xa6\x97\x66\x94\x2a\xf3\x93\xb3\x52\xc2\x6d\xa7\xff\xe1\xf4\xd2\xfa\xf8\x37\xf8\xba\xc7\x6e\x6e\xb5\x52\xf9\x79\x38\xe7\x28\x59\x75\x0c\xec\x01\x9c\xcd\x77\x83\xb4\x42\xe1\xc5\xae\xbe\xcd\x9c\xf0\x0d\xb4\xdc\x11\x3f\x8b\x89\x17\xf2\x87\xfe\xde\x40\xf9\x43\x25\x13\xc1\xb8\xb7\x97\x40\xb3\x3d\xee\x9c\x07\xb3\xcb\xd6\x64\xe2\x9a\x95\x45\xf6\x07\x27\x12\x4e\x29\x91\x04\x17\xe4\xbf\xc0\x95\x9b\xf5\x66\x43\xb1\x32\x73\xb0\x15\x1a\x07\xc9\x79\x79\x55\x90\x54\x69\x83\x5a\xb0\x07\x84\x12\x9d\x1e\xbf\xf6\xc0\x82\x0c\xc0\xdb\x63\x49\x6e\x87\x07\xf4\x3e\xda\xa1\x9f\x06\xe2\x48\x76\x4e\xc6\x14\xfe\xbd\xe5\xdb\x50\x65\xef\x12\xca\xce\x68\x79\xa0\x42\x0d\x05\x95\x9a\x3f\xfb\xc6\x28\x96\x93\xd6\xca\x61\xf6\x4c\x81\x5e\x3f\x0a\x54\x52\xf2\x77\xe9\xca\x56\x35\x66\xa4\xae\x6a\xc8\x6c\x8e\xc2\xd5\xc2\x94\xfa\x66\xac\x27\x8d\x0b\x4e\xb7\xda\x25\x35\x83\xa6\x53\x72\x62\x0a\x01\xc8\xdc\x3c\xae\xbd\xac\x12\x52\x0b\x62\x3a\x9d\xb4\x82\xe0\x3e\x06\xbb\x00\xe9\x49\x29\x40\x3e\x8e\xc1\x02\x36\x33\x92\x59\xa3\xcd\x23\xad\x16\x67\x2e\x74\x8c\x48\xf6\x10\x46\xe9\x5b\x6e\xaf\x31\xcf\x52\x96\x41\xd6\x5e\x78\x75\xc2\x8f\x36\xc4\xbd\x57\xdb\xf1\x5a\x0c\xd6\x51\x6e\xdb\x38\x50\x4f\x45\xe9\xa2\x95\x79\xda\x85\xd3\x48\x8b\xf9\x21\xe8\x5a\x45\xca\xc9\x46\x36\x27\x51\xaf\x59\x1a\x56\xa4\x2c\x2d\x75\xa6\xd3\x7d\x54\x35\xdc\xcc\x97\xd8\x88\x78\xcd\xd2\x9e\x24\xa2\xc5\x85\xbf\x07\xe5\x99\x7e\xa2\x2c\x63\xe9\xe7\x69\x3b\x2d\x68\xf2\x85\x39\xc0\x0a\xd2\x47\xd8\xed\x52\xd9\x4b\xa4\xff\xc2\xe9\x17\x49\xd2\x2f\x62\x87\xe2\x97\x7d\x53\xaa\xaf\xf2\x8b\x30\x72\x6c\xba\xd0\xf2\x77\xcc\x62\x65\xcf\xd7\x32\xb9\xd8\x70\x42\x65\x3e\x9b\xfe\xf3\x07\xb1\xfc\x41\xfc\x32\x55\xbb\xb9\x66\x09\xd2\x01\xd2\x90\x4c\xa6\x9e\x77\x82\x22\xb2\x52\xac\xa3\xc3\xd4\x13\xbf\x02\x05\x8e\x25\xfc\x0a\x52\x02\x47\x49\xa7\x5c\x38\x3a\x42\xbf\x82\x54\x2a\x76\x52\x94\x1f\x3d\xbd\x1d\x6c\x56\xe1\x90\x02\xb9\x69\x67\xd8\x83\x1d\x36\x1b\xe0\x38\x9b\x87\x7c\xdc\x69\x62\x68\x52\x93\x60\x3b\xeb\x4f\xab\xc8\xef\x9a\xe0\x62\x87\x09\x2e\x06\x4c\x50\x2f\x32\x2b\x72\x03\xf4\x81\xad\x30\xc0\x74\xe6\x25\x8d\x5e\x7b\xd4\x2b\x51\xbf\x25\xd0\x31\xf2\x10\x5a\x33\x60\xec\x92\x7b\x8d\x33\xf6\x35\x88\x04\xa1\x49\xdf\xba\xb6\x18\xe0\x7d\x75\xbd\x65\x3f\xa2\xba\xaf\x47\x74\x2b\xf6\x88\x99\xe3\x25\xe8\x8e\x59\x97\xe8\x85\xb2\x77\x48\x57\x9b\x83\x56\x89\x3b\x90\xd5\xd5\x52\x6d\xae\x31\x74\xc5\x42\x32\x30\xe6\x4c\x4b\xce\x81\x4a\x44\x68\xce\xf8\xda\x24\x67\x21\x19\x87\x4c\xad\xba\xe6\x98\xcb\xec\x17\x4b\x0e\xf1\x4b\xb7\x65\xa5\x0a\x5c\xce\x19\x77\x26\xd4\x3f\x44\x78\x60\xf2\x46\xd3\xb6\x2e\xe5\xfe\x5d\x12\x0e\xd9\x9b\x5d\x1d\xfb\xee\x68\xf6\xda\xd5\x3b\x57\xfe\xc0\x31\x15\x44\x69\x1d\xb4\x25\x6f\x6e\x37\x4c\x40\xb3\x83\xb5\xe4\xf7\x56\xa6\xb0\xb7\x5a\x70\xf4\x84\x98\x9a\xc4\x3b\x75\xcd\xaa\x8d\xf3\x50\x74\x67\x0f\x07\x65\x93\x76\xb8\x3e\x0f\xcc\xa7\xf9\xcf\x1a\xef\x1f\xc7\x88\x92\xc2\x8b\xac\x96\xa9\xea\xa3\xa3\x90\xbe\x50\x83\xdd\x21\x52\xb8\xf8\xf7\x2a\x23\xc9\x1a\x46\xa9\xf2\x81\xac\xe1\x29\x2a\x02\xb7\x12\x38\xc5\xc5\x28\x65\xde\xd8\x41\x4f\x51\x21\x42\x25\xac\x80\x8f\xd2\xe7\x94\xca\xa7\xa8\x4a\x5e\x30\x2c\x47\x29\xf2\x56\x8d\x78\x22\xaa\xec\xd2\x8c\x43\xde\xa3\x57\xbf\xa0\x49\x93\x25\x07\x44\x86\x96\xa8\x30\xc2\xda\x8c\xa3\x59\x20\xd7\x3b\x22\xe4\x74\xde\x22\x9e\xe1\xcd\x74\xee\x04\xb6\x47\xd4\xa2\xbc\xf2\x4e\xa8\xfb\xd7\xf8\x46\xc6\x46\x4b\x51\x5e\xed\x57\x29\x4a\xad\x46\xb5\x1d\x0a\x26\xaf\x8a\x42\x15\x02\x27\xd7\x8c\xa4\xcd\x6d\xf2\xae\x68\x32\x69\xf7\x94\xea\xe3\xfd\x56\x30\x99\x2c\x3e\x1b\x88\xa9\x45\xb3\x35\x53\xe3\xfe\x62\x84\x76\x04\xb8\x9c\x2e\xd0\xf4\x52\xa1\x55\x0b\xbd\x84\xbf\x2a\x25\x5b\xd9\xea\x2f\xdb\x11\x99\xdf\xe2\xe6\x46\x08\xcc\xa3\x6c\x70\x8e\x55\x19\x4a\xe3\xe6\xd2\x42\x6f\x7e\xda\x3c\x2e\x8d\x7a\xde\xba\xf8\x18\x9a\xad\x24\x4a\xce\xf0\xed\x3b\xa0\x2b\x79\x8d\x5e\xc6\xe8\x76\x86\x6f\xc9\xba\x5c\x9b\x21\xb1\x1a\x2a\x6a\xc3\x47\x51\xf4\xa1\xf6\xa3\xa9\x44\xe8\x28\x95\x08\xbd\xa7\x4a\x35\x9f\xc7\x57\x09\xdf\xea\x17\x1d\xe8\x65\xf2\x72\xa8\x48\x8a\x4f\xf9\xd6\x89\x23\x32\x7e\xed\xc3\x8f\xf6\xbd\xc7\xc3\xe9\x6b\x4f\x03\x62\x85\x8e\x5e\x6f\x17\xaa\xba\x9e\xb5\xc4\x9e\x3f\xb4\x9f\xf6\x05\xe2\x43\x7a\xcd\xc4\xe9\x78\xaf\x39\x29\x1e\xc1\x6b\x91\x32\xdf\xc7\x69\x8d\xd4\xdf\xc9\x69\xf5\x5d\x4f\xe2\x1f\x76\x7b\x77\x41\xfa\x56\xb9\xf7\xa9\x5a\xc7\x12\x0a\x6e\xed\xb4\xd5\xaa\x7b\xd7\x3b\x8d\xfe\x86\x18\x5b\x5e\x45\xeb\x79\x38\x50\x4f\x45\x58\x21\x50\xf8\xa6\x56\x55\x4b\x56\x3f\xad\x33\x1b\xd2\xc6\x10\xfe\x3b\xb3\x93\x52\x48\xb6\x76\x0f\x07\x1a\x84\xa4\x66\x70\xb0\xc6\x9b\x0d\xa1\x2b\xfd\x58\x4d\x1f\x31\x37\x48\x67\xa6\x29\xb1\xff\x47\xd3\xe6\xe5\x61\x47\x1c\xcf\x49\x24\x6f\x50\xfb\x7d\x61\x71\x9d\x47\xd8\xc3\xd9\xb8\xcf\xe4\x24\x47\x05\xd0\x56\xdd\x3b\x47\xbf\xa0\x97\x35\x2b\x7b\xd4\x11\x76\xb1\x1c\x7c\x0c\xe8\x19\x5b\x8f\x86\xce\x28\x77\xee\x4a\x8a\xf8\xfb\x37\xd5\x42\x72\x92\x6a\xd3\xbe\x65\xbc\xde\xe8\x07\xa7\x42\x35\x35\xe8\x5e\x1f\x33\x9b\xe3\xb3\xe6\xa1\xa3\x7e\x38\xf8\x05\xee\xdc\x81\xc6\xbe\xd3\xa2\x21\x19\x66\x1a\xc8\x5d\xd0\x34\x19\x66\x40\x9c\x6d\x6d\xbc\x9b\x05\x62\x5f\xac\xff\x07\x19\x37\x67\x1a\x67\x78\xf3\x49\xb1\xfa\xfc\xb3\x1a\xd6\xb1\xf4\x8d\x6f\xe4\xa3\x23\xf4\x07\xa0\x94\x95\x45\xa6\x6d\x9b\x13\x9a\x21\x22\x17\x48\x30\x54\x80\xfc\x51\xa0\xf4\x1a\xd2\x2f\x88\xd9\xa7\x66\xec\x2b\x70\x94\x62\x01\x88\xd0\x0c\x6e\x21\x43\x62\x03\x29\x5a\xe3\x8d\xef\xb3\x5d\x72\xbe\x53\x10\x27\x58\x40\x8f\xc0\xee\x2d\x55\xaf\x41\x44\xe0\xc3\xbc\x2c\x0a\xcf\x47\x22\xec\xb9\xc6\x9b\x48\x6f\x0d\xf0\x9a\xcd\x15\xc6\x27\xe3\xac\xcf\xb1\xbe\x8a\x50\x3f\xd0\x7a\xb2\xff\x59\x47\xd0\xbf\xf3\xc6\x03\x6f\xf4\x0b\x8f\xda\x0c\x2a\x84\x77\xa1\xed\x7f\x4e\x1a\xf2\x3b\x1e\x63\x85\x6d\x3b\xf3\x8e\x3a\x6c\xf3\x9e\x14\xd4\x23\x94\x08\xad\x23\xf0\xe0\x29\x75\x7b\xb9\x99\x2e\xf7\x4e\xab\xf0\xb0\x75\x68\x4f\xa4\xfe\xf9\xf4\x65\x77\x6b\xa3\x76\x36\xed\xe3\xd4\xde\xe1\x7a\xbf\xb7\x1c\xd8\x87\x1e\x56\xd5\xa8\xbd\x62\x53\xb0\xd4\xc3\xaa\x7a\xe5\x5b\x74\x75\x6b\x6d\x2a\x1b\xe9\xfc\x86\x65\xef\x06\x74\xa7\x76\x8e\x41\xef\xcd\xba\xfa\x17\x34\x2c\x07\xdc\x15\xc7\x82\x83\xf6\xdc\xef\xb4\xb8\x0b\x38\x78\x74\xa3\x41\xab\x67\x14\x7a\x70\x99\xeb\xb5\xfb\x74\x83\xde\xea\x39\x06\xdd\x95\x7f\x6d\x74\x4d\xd7\xe8\xfa\xcd\x4b\xd0\x3b\x78\xf4\x92\xc4\x33\xdc\x70\x48\xdb\xde\x6e\xa8\x4e\x15\xaf\x57\x24\x6e\x70\x65\xdb\x00\xd7\xe4\x65\xcf\xd5\x67\xeb\xbe\x33\x8a\x93\x77\xae\xee\x9a\x2c\xc9\xc8\xde\xb4\x47\xc1\xbd\x25\x85\x04\x8e\xaf\x8a\xc0\xfe\x0d\xd5\x80\x06\xbd\xe2\x70\x19\x07\xb2\xa2\xff\x86\x20\x26\x1b\xaa\xc5\xf5\x7b\x45\xe1\xda\xbb\x4f\xaf\xc5\x50\x0c\x5e\xdd\x1a\x85\xe5\x3d\xf8\xf0\x5a\x1b\xaa\xc1\x0c\x7a\x45\xe1\xfa\x07\x20\x75\x63\x4d\x5c\x76\x0f\x49\x22\x41\x3b\xd3\xc4\xd1\x96\x9d\x3d\x7b\x14\xa2\x77\xa8\xd1\x40\x3a\xe2\xb2\x7b\xf0\x11\x09\xda\x15\xd3\xd2\x96\x9d\x4d\x6a\x0c\x62\x3b\x73\x7a\x09\x73\x54\x9e\xd4\x89\xa9\x1d\xe8\x35\xd1\xc8\xe6\xf7\x89\x02\x3d\xe7\x64\x8d\xf9\x5d\x2b\xcc\x1b\xaa\x81\x0d\x7a\x45\xe1\xbe\x07\x9c\xb5\x13\xba\xa3\x2d\xed\x81\x61\xdd\x23\x12\x31\xbc\x78\x33\x88\x86\xb6\x6c\x1f\x41\x46\x21\x5e\x74\xa6\xe2\x85\x37\x15\x2f\x46\x4d\xc5\x0b\x73\x4b\xea\x63\x69\x8a\xc5\x72\xad\x71\x58\xe5\x95\xbd\x5c\x6f\xc0\x0c\xc9\x7d\xb4\x53\x77\x88\x8b\x9c\xce\x15\xa7\xfa\x57\x13\x8d\x88\x7e\x9f\x38\xd0\x96\x88\x9e\x7c\x7b\x85\xb3\x1c\x86\xb7\xa6\xfb\xeb\xe6\xfe\x6d\xc6\x77\x28\xa0\x07\x18\x3f\xed\x4a\xda\xec\xeb\x92\xe7\x32\xfa\xb9\x8c\x7e\x2e\xa3\x9f\xcb\xe8\xe7\x32\x1a\x3d\x97\xd1\xcf\x65\xf4\x73\x19\xed\x10\x9f\xcb\xe8\x08\xc0\xff\xc7\x32\x7a\x3b\xfa\xeb\x87\x88\x07\xad\xad\x22\xbb\x7e\x84\xfb\x40\x9f\x24\x5a\x2e\x0f\xf8\x61\xa2\xe4\xe5\xf7\xff\x2e\xf1\xe1\xbf\x15\x33\x86\x79\x94\x2f\xc6\xbe\xdb\x37\x60\xa1\x0e\xdf\xf0\x25\xd8\x4f\x8f\xf4\x29\xd8\x77\xfb\xb8\xcb\x59\xe2\x7b\x7d\xe2\xf5\xa2\xef\x1b\x2f\xa5\xed\x7f\xa8\x08\x5f\xd6\x63\x44\x19\x1d\x7c\x5b\x1f\x4e\xf7\x91\xca\x3a\x5e\xb1\x1f\xd1\xda\x2f\xea\xf7\x7e\x92\x7b\xbf\xfd\x34\x2b\xe5\xd0\x57\x0b\x3f\xf5\xbd\xaf\xf7\xed\x57\xf5\x7d\x86\xef\xd9\xf6\x7f\x01\x00\x00\xff\xff\x44\x4f\x82\xc2\xe2\x48\x00\x00")

func templatesModelGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesModelGotpl,
		"templates/model.gotpl",
	)
}

func templatesModelGotpl() (*asset, error) {
	bytes, err := templatesModelGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/model.gotpl", size: 18658, mode: os.FileMode(420), modTime: time.Unix(1539628015, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRelationships_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\xdf\x6b\xe3\x38\x10\xc7\xdf\xfd\x57\x0c\x21\x1c\xed\x91\x73\xaf\xdd\xb7\x40\x1f\x8e\xf6\x58\xfa\xd0\xa5\x74\xef\xee\x25\xe4\x41\x67\x4f\x12\xdd\x2a\x92\x57\x91\xb3\x14\xe3\xff\xfd\x90\x1d\x3b\xb6\x23\xd9\x6e\xd8\xc0\x2e\xcc\x14\x0a\xb6\x66\xe6\x3b\x1a\xfd\xc8\x07\x27\x2c\xfa\xc2\xd6\x08\x59\x06\xe1\x67\x34\xe1\x83\x92\x2b\xbe\x4e\x35\x33\x5c\xc9\xf0\x13\xdb\x22\xe4\x79\x10\xf0\x6d\xa2\xb4\x81\xc9\x5a\x85\x2c\x51\x1a\x8d\x0a\xb9\xba\x41\x81\x5b\x94\x86\x89\x49\x10\x44\x4a\xee\x0c\x48\x15\xab\xe8\xb3\xd1\x5c\xae\xe1\x1e\x26\x8b\xe2\x79\x39\x81\x9b\x1b\x90\x4a\x70\x69\xe6\xb0\x67\x3a\xda\x60\xf4\x65\x16\x23\x8b\x23\x15\x63\x10\xec\x99\x06\x8d\xa2\xd0\xdc\x6d\x78\xb2\x7b\xc5\x35\xdf\x19\xfd\x06\xb5\x42\xf8\xea\x1a\x0f\x82\x55\x2a\x23\xe0\x92\x9b\xab\x6b\xc8\x82\x00\x00\x3c\x99\xee\x87\x72\x65\x79\x19\x9e\x65\xa0\x99\x5c\x23\x4c\x51\x1a\x6e\xde\x6c\x0f\x66\x30\xad\xb2\xc2\xfc\xbe\xec\x54\x2b\x89\x6d\x52\x19\xfc\x1b\xf0\x15\xec\x36\x2a\x15\x71\x99\x19\x75\xd3\x13\xa6\x36\xb8\x99\x1b\xa6\xe1\x4b\xfa\xaf\xe0\xd1\xb3\x8a\xf1\x90\xc6\x39\x85\x45\x96\xb5\xe2\xf2\xfc\x29\x2e\x1f\x97\x70\x0f\xbf\xb8\xa7\x97\x15\xf9\x1a\xa5\xad\x0d\x5c\x09\x94\xc7\x09\x85\x0f\x1a\x99\xc1\x6b\xf8\xbd\x9a\x84\xb5\xf2\xe5\x1c\xb6\x2c\x59\xec\x8a\xf5\x5c\xfe\xea\x56\x78\x92\x2b\x05\x47\x99\x4a\xea\xd0\xc3\x84\x69\x94\x66\x06\x53\x16\x55\xdd\xeb\x2a\x37\x65\xbd\x3d\x7c\x92\xd2\xdd\xc8\x52\xa0\xd5\xc4\x6e\xc2\x89\x6d\xdc\xc1\x2f\xcf\x27\x30\xf7\x35\xcb\x4e\xa5\x3d\x93\x46\x3d\x87\x19\x84\x8f\x98\x68\x8c\x98\xc1\xb8\xab\x63\xed\x38\x3a\x07\xa3\x53\x9c\x39\xd3\xa1\x74\x06\x77\x94\x5e\x98\x66\x5b\x34\xa8\x1f\x71\x65\xb7\xb8\xed\x9f\x3f\xaa\x5e\x58\x7f\x74\xf8\x8a\x5f\x53\xae\x31\xee\x2c\x76\x65\xd5\x70\x1d\xba\x9b\x37\x4e\xcd\x27\xfc\x76\x1c\x38\xb8\xda\xa1\xab\x93\x3c\xd6\x16\x4b\xfb\x57\x6e\x9d\xd3\x9e\x36\x6b\x3f\xec\x14\x3e\x83\xa9\xb8\x2d\x36\xc8\x88\x19\xb8\xca\xf7\x35\x44\xdc\x7a\xe6\xdb\xac\x75\xa8\x52\x57\xb5\x77\x45\xb5\xe2\xb6\x2f\xb7\xb3\xa0\xbb\x81\x82\xca\xa2\xc6\x94\xe4\x2a\xeb\x43\x59\xd6\xdd\x90\x02\x54\x67\x43\x7c\xb0\xe7\xe2\x74\xab\xba\x84\x50\xc6\x03\x69\xf3\xfe\x44\xe3\x92\x0c\x7b\xf5\xc8\x0c\x07\xf7\x7b\x38\x52\x5f\x9f\x73\x90\xc7\x1c\xc9\x3f\xa5\xd1\x1c\x77\x9e\x0d\xd1\x3c\x89\x8b\xe5\xf1\x2c\x3a\x32\xb9\x2f\xae\xc6\xbe\x48\x86\xce\xd6\xa1\x14\xef\xb6\x79\xa7\x7c\x65\xf6\xc7\x6a\x7e\xb8\x84\x2b\xa2\xe8\xd9\x6c\x7f\xbd\x25\x47\x77\xfb\xd0\xef\x5e\xdd\x99\x49\xf8\x88\x2b\x96\x0a\xf3\x0f\x13\xe9\xc9\x4f\x40\xd3\x9a\x7e\xb5\x50\x27\x78\x40\x70\x78\x73\xf1\x15\xe0\xd7\x7a\x06\x13\x94\xe9\x76\xd2\x57\xd4\x1f\x42\xa8\x6f\x18\x3f\x6c\x14\x8f\xb0\x58\xec\x77\x5e\x48\xff\xcd\x60\xba\x2f\x56\x38\x09\xdb\xc9\x86\xae\x81\xa2\x03\xfb\xe1\x1b\xa0\x67\xbf\x57\x36\x7c\x26\x47\xdc\xdb\xd3\x24\x7c\x4e\x85\xe1\x89\xe8\x5d\xc6\xca\xc7\xf7\x3b\x3b\x52\xd8\x51\x72\x4f\xc4\xfb\xbc\x3d\x43\x9d\x24\x1e\x2f\xc7\xeb\x46\xa0\x63\xd4\x0b\x78\x7f\x27\xf1\x29\xe0\x95\x2f\x2f\x0c\x78\xa5\x08\x01\x9e\x47\x89\x00\x8f\x00\x8f\x00\xef\x0c\x19\x02\xbc\xba\x8a\xef\x08\x78\xc4\x77\x40\x7c\x47\x7c\x37\xde\xfb\x82\x7c\xf7\xc2\x4c\xb4\x21\x3a\x23\x3a\x23\x3a\xf3\x56\x4b\x74\xe6\x32\xa2\x33\xa2\xb3\x86\x11\x9d\x11\x9d\x11\x9d\xf9\x93\x5c\xfc\xeb\xdb\x23\x0a\x3c\xf9\xfa\x56\xbe\xbc\x30\xdf\x95\x22\xc4\x77\x1e\x25\xe2\x3b\xe2\x3b\xe2\xbb\x33\x64\x88\xef\xea\x2a\x88\xef\xba\x46\x7c\xd7\x36\xe2\xbb\xf1\x11\x3f\x25\xdf\x7d\x44\xd3\xb9\x5d\x5e\xd1\x1e\xf4\xfd\xa5\xf1\xee\x23\x1a\x62\x3b\x8f\x12\xb1\x1d\xb1\x1d\xb1\xdd\x19\x32\xc4\x76\x75\x15\xc4\x76\x5d\x23\xb6\x6b\x1b\xb1\xdd\xf8\x88\x9f\x95\xed\x9e\x99\x7c\xf3\xf0\x9d\x1d\xba\x3c\xe3\x59\x15\xe2\x3c\x8f\x12\x71\x1e\x71\x1e\x71\xde\x19\x32\xc4\x79\x75\x15\xc4\x79\x5d\x23\xce\x6b\x1b\x71\xde\xf8\x88\x1f\x86\xf3\x2c\xa9\x10\x9d\x01\xd1\x19\xd1\x99\xb7\x5a\xa2\x33\x97\x11\x9d\x11\x9d\x35\x8c\xe8\x8c\xe8\x8c\xe8\xcc\x9f\xe4\x7b\x7c\x85\x2b\xff\x77\x5e\x66\x59\xf5\x94\x07\xff\x07\x00\x00\xff\xff\xbd\x9e\xc8\xb9\xa8\x46\x00\x00")

func templatesRelationships_registryGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesRelationships_registryGotpl,
		"templates/relationships_registry.gotpl",
	)
}

func templatesRelationships_registryGotpl() (*asset, error) {
	bytes, err := templatesRelationships_registryGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/relationships_registry.gotpl", size: 18088, mode: os.FileMode(420), modTime: time.Unix(1533220355, 0)}
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
	"templates/README.md": templatesReadmeMd,
	"templates/identities_registry.gotpl": templatesIdentities_registryGotpl,
	"templates/model.gotpl": templatesModelGotpl,
	"templates/relationships_registry.gotpl": templatesRelationships_registryGotpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"README.md": &bintree{templatesReadmeMd, map[string]*bintree{}},
		"identities_registry.gotpl": &bintree{templatesIdentities_registryGotpl, map[string]*bintree{}},
		"model.gotpl": &bintree{templatesModelGotpl, map[string]*bintree{}},
		"relationships_registry.gotpl": &bintree{templatesRelationships_registryGotpl, map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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

