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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

var _templatesIdentities_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\x4d\x6b\xe4\x38\x13\xbe\xfb\x57\xd4\x1b\xc2\x8b\x0d\x1d\xe7\xb2\xec\x21\x4b\x0e\x61\x98\x81\x86\xcd\x10\x12\xd8\x4b\x93\x83\xc6\x5d\x76\xc4\xca\x92\x91\xe4\xcc\x34\x46\xff\x7d\x91\x2c\x7f\xc8\xed\xfe\x4a\x32\x3d\x30\xb9\xa4\xdb\xd2\x53\x1f\x4f\xa9\x1e\x95\xbb\x22\xd9\xbf\xa4\x40\x68\x1a\x48\x9f\x50\xa7\x9f\x04\xcf\x69\x51\x4b\xa2\xa9\xe0\xe9\x57\x52\x22\x18\x13\x45\xb4\xac\x84\xd4\x70\x51\x88\x94\x54\x42\xa2\x16\x29\x15\xd7\xc8\xb0\x44\xae\x09\xbb\x88\xa2\x57\x22\x21\x8e\x00\x00\xe8\x1a\xb9\xa6\x7a\x63\xc1\xea\x9e\x54\x70\x0b\x25\xa9\x56\x4a\x4b\xca\x8b\xe7\x1e\x93\x2e\xfd\x3e\x68\x1c\xcc\xfe\x35\xcd\x15\x48\xc2\x0b\x6c\x83\x79\xaa\x30\xa3\x39\xcd\x5c\x30\xca\x06\x32\x6c\x04\x9a\x83\x7a\x11\x35\x5b\x3f\x62\x41\x95\x46\x19\xec\x86\x14\x2e\xd3\x87\xfa\x1b\xa3\xd9\xbd\x58\x63\x88\xbd\x82\xcb\x21\x44\xb8\xb9\x85\xd4\xee\x61\xe9\xe7\xe1\xe1\xd5\x08\x70\xd1\x34\x7e\xc3\x23\x2a\x6d\x97\x8d\xb9\xb8\xb1\x31\x8c\xcd\x18\xd3\x25\xb4\x08\x5c\x21\x5f\x4f\xbd\x8f\x1e\x99\x28\xe0\x2c\x23\x1a\x0b\x21\xe9\x6f\x48\x9c\xa8\x65\x86\x3f\x85\x3c\xc2\x28\x51\x3f\x91\xb1\xab\x33\x52\xd6\x34\x3e\xaa\x4b\xba\x80\x4b\x97\xd9\x08\x74\xd7\x66\x0a\x21\xc7\xdd\xbe\x8f\x23\x76\xcf\x41\xe5\x6b\xfc\x31\xc3\xf5\xea\x79\xf5\xdc\x7e\x3c\x3b\xc7\x96\x81\x49\x7f\xf6\x54\xd0\x1c\x18\x72\x48\x97\x6d\xd8\x60\xcc\x6c\xa0\x61\xb0\x8e\xfa\x4c\x94\x95\xa8\xf9\xda\xb1\x3f\xc0\x03\x48\x67\xa9\x19\x83\x7b\xa0\x31\x2e\x32\xfb\x7f\x31\x90\x09\x66\x6f\x21\xcc\xa2\x69\x00\x99\xb2\x29\x70\xca\x16\x6f\xac\x55\x12\x45\xd7\xd7\xe0\x48\xf9\x07\xa5\xb2\x04\x4a\xd4\xb5\xe4\x0a\xf4\x0b\x42\x56\x4b\x89\x5c\xc3\xab\x5f\x13\xb9\x7b\x5c\x3a\x12\xa3\xbc\xe6\x59\x80\x8d\x13\xc8\x99\x20\xfa\xcf\x3f\xa0\xf1\x76\xfa\x0b\xe3\xee\x61\xb9\xe4\xb9\x48\x3b\x37\x36\xc3\x28\xd2\x9b\xca\x9b\xbb\x27\x9c\x14\x28\x41\x69\x59\x67\xba\x31\x91\x33\x1f\xe7\xc1\x6a\x02\xdd\x29\xfd\x22\x45\x69\x2b\x18\x73\x5b\xc6\x96\xde\x04\x66\x3b\xd9\xa5\xea\xa3\x99\xde\x39\x2b\x0b\x7f\x8e\x8e\xf1\xf6\xa9\x95\xdc\x4d\xec\xb5\x77\x73\xba\xd7\x40\xb5\x57\x9d\x9d\xe3\xdc\xbb\x96\x8e\xdb\x06\x3e\xda\xf1\xa0\x78\x2b\xf7\xf1\x48\x57\x7c\x13\x13\x3e\xe4\x17\xd3\x19\x4f\x49\xe7\x8a\xe6\x40\xe1\x16\xf2\x74\xab\x34\x84\x6f\x92\xbf\xe0\x7f\x34\x5d\xaa\xcf\x65\xa5\x37\x71\x32\x6a\xa5\x8e\x9a\x40\x34\xe6\x4c\xf5\xbc\x9f\x6c\xce\x3f\x0b\xcd\x79\x1e\xf9\x26\x39\xc0\x45\x4e\xc9\x37\x86\x71\x57\xbb\x59\x0a\xa6\xcf\x5a\x4c\xc7\x8c\xfa\x4e\x75\xf6\xd2\x57\xdf\x47\xdb\x2b\xf7\x1e\xad\x7b\xb3\xce\x65\x44\xb5\x33\xda\xd6\xe5\x31\x08\xfc\xcd\x94\xb4\xaf\xf8\x7d\x07\x24\x4e\xa2\x19\xd9\x98\x7c\x5d\x63\x4e\x6a\xa6\xb7\xcc\x72\xca\x7c\x35\x76\x11\xfd\x54\x11\xa9\xf0\x2d\x74\x6f\x23\x7f\x21\xe9\x1e\xc8\x85\xee\x48\x5c\xaa\x47\x21\xf4\x7b\x8b\xd2\x26\xf9\x9e\xd2\x7c\x58\xa5\xfc\xc5\xb6\xbf\x3c\xc1\x95\x19\xe8\x5f\x3f\x0c\xac\x3a\x03\xee\x95\xe1\x90\x1c\xb5\x95\xb5\x5d\xfb\xe4\xcc\x06\xaa\xb4\xbf\xf7\x26\xbd\xef\x4f\xd7\x44\x0a\x5a\x9d\x4b\x8e\x53\x82\x03\xc9\xcf\x87\xa3\x3e\xec\x58\xee\x3c\x5d\xe7\xd5\x8a\xff\xcf\x01\x1e\x58\x2d\x09\x03\x63\xfe\xa6\xca\x5e\xdd\x67\x3c\x98\xdb\x42\x70\x74\x9d\x66\xa0\xbf\x5d\xb5\x76\x4b\xc8\x2f\xac\x59\x40\xf9\x49\xdd\xad\xf6\xb6\xb7\x3a\xb9\xbf\x1f\x91\xb5\xe5\x7b\xa1\x95\x8a\xc7\x5e\x83\x95\xb6\x4e\x72\x3a\x5d\xc9\xb9\x3d\xd6\xd7\x2b\x91\x50\xfa\x79\xf6\x36\x70\x69\xe7\x5a\x3b\x70\xfb\xc5\xf1\xac\xed\xb6\x8d\x22\xb8\x1f\xc1\xba\x59\xbb\xfd\x16\x04\x3a\xde\x36\x4c\xdd\x9d\x77\x13\x39\x7f\x77\x8c\x79\x62\x28\xaa\xde\x2b\x61\x0c\xf0\x07\x55\xda\x2a\x36\xed\xd7\xbd\xb3\x00\x13\x5b\x75\x3f\x38\x6e\xce\x6d\x39\xff\x9b\xf4\xde\x16\x39\xed\x67\x03\xe3\xd9\x73\x53\xf4\x17\x21\xfb\xbc\xc7\x14\xda\xe2\xf9\x41\x1b\x72\x21\xdd\xf7\x82\xbe\xe2\x30\xf7\xf7\x8c\x4e\xed\x1c\xba\x4f\xc3\xdb\x74\x97\x2a\x1d\x41\xeb\x79\x65\x67\xfc\xbe\xfb\xae\x9f\x29\xf6\xbf\x02\x1f\x21\x52\xe1\x9b\x80\x95\x26\x13\xfd\x17\x00\x00\xff\xff\x11\x03\x2d\x1e\xcb\x14\x00\x00")

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

	info := bindataFileInfo{name: "templates/identities_registry.gotpl", size: 5323, mode: os.FileMode(420), modTime: time.Unix(1554955843, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesModelGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\x6b\x73\xdc\x36\x92\x9f\x3d\xbf\x02\x99\x52\xbc\x33\xde\x11\xe5\xcf\x93\xd5\x56\xf9\x1c\x27\xa5\x3a\x3b\x51\x45\xbe\xec\x07\x97\x6b\x0d\x91\xcd\x11\x62\x0e\x40\x03\xa0\x6c\xdd\xd4\xfc\xf7\x2b\x3c\x48\x00\x7c\x0c\x41\x3d\x7c\xba\x3a\xf9\x43\xa2\xc1\xa3\xd1\x2f\x34\xba\x1b\x60\x97\x38\xfd\x8c\x37\x80\x76\x3b\x94\x5c\x80\x4c\x5e\x33\x9a\x93\x4d\xc5\xb1\x24\x8c\x26\xbf\xe1\x2d\xa0\xfd\x7e\x36\x23\xdb\x92\x71\x89\x16\x33\x84\x10\x9a\xe7\x5b\x39\x37\x7f\x6d\x58\x82\x4b\xc6\x41\xb2\x84\xb0\x13\x28\x60\x0b\x54\xe2\xa2\xee\x25\xf2\xaa\xba\x4c\x52\xb6\x3d\xd9\x12\x99\x5e\x41\x51\x5c\x9d\xa4\xac\xbc\x11\x92\x57\xa9\xac\x38\x98\x81\xbb\xdd\x31\xe2\x98\x6e\x00\x25\x17\x25\xa4\xc9\xfb\x9b\x12\xce\x39\xbb\x26\x19\x70\xa1\xd6\xd7\xd0\x14\x8a\x68\xbf\x77\x53\x80\x66\xe8\xd8\xf6\xb6\x41\xfc\x89\x0b\x92\x69\x22\x22\x01\xed\xf7\xb3\xe5\x6c\xb6\xdb\xa1\xa3\x02\x4b\x10\xf2\x4f\xe0\x82\x30\x8a\xd6\xa7\x16\xe2\x5b\xdd\xfc\x4a\x4a\x4e\x2e\x2b\x09\xa2\x1e\xa0\xd8\xb3\xdb\xd9\xc5\x8f\xc8\x0a\x1d\x01\xad\xb6\x6a\xde\x65\x45\x8a\xec\x0d\xad\xb6\xc2\x80\x68\x83\xde\xef\x67\x27\x27\x8a\xf3\x7a\x86\xa6\x1a\xed\xf7\x88\x43\xc9\x41\x00\x95\x02\xc9\x2b\x40\x25\x13\x82\x5c\x16\x80\xae\x71\x51\x81\x40\x39\xe3\x08\xd7\x58\x68\x62\xcc\xf4\x06\x33\x2b\xb4\x79\x32\x93\x0a\x62\x07\xbe\x90\x9c\xd0\xcd\x6c\x96\x32\x2a\x6a\x91\x3a\xf6\x1d\x51\xbc\x85\x15\x3a\xd2\xab\x29\x2a\xcc\xe4\x3f\xcd\xe2\x96\x85\x16\x6d\x6a\x56\x6a\x63\x6c\xa6\xaa\x01\xe6\xaf\xfd\x3e\xb1\x8b\xb8\x29\x1d\xac\x4e\x0d\x29\xf5\x8c\x5a\x38\x4e\x36\xee\xef\x99\xc2\x96\xe4\x88\x32\x69\x65\xf3\x8e\x65\x50\x24\x3f\x83\xc4\xe9\x15\x64\x8e\xb1\x7e\xef\x1b\x2a\x89\xbc\xb1\xcc\x39\xcb\x40\xff\x6c\xa3\xde\xb4\xb3\x5c\xff\x66\x97\x7f\x41\x2a\x93\xd9\x35\xe6\x71\xf0\x4e\x51\xb3\x09\x92\xa6\x71\xa7\x89\x51\x43\xd7\xa8\xd1\x40\x0f\xd4\x1f\x20\xa4\xea\xdd\xef\xe7\x2b\x3d\xf4\x35\x96\xb0\x61\xfc\x66\xdd\x37\x94\x55\x3c\x6d\x84\x6c\xc6\x9f\x9b\x5d\xbc\xee\x82\xb6\x3d\x6e\x24\x27\xd7\x58\xaa\x91\xed\x81\xa6\x63\xbf\x5f\xcd\xf6\x13\x79\xbd\xdb\xf5\x8d\x38\x13\x7f\x30\x26\xc7\x64\x71\x5e\x54\x1c\x17\x68\xbf\x7f\x4b\x84\xf4\xa5\x81\x51\xa1\x5a\x58\x1e\x31\xb7\x51\xf4\x98\x35\x3e\x7c\x7c\x31\x38\x52\x11\x7c\x72\x82\x3c\xed\x90\x15\xa7\x46\x35\x48\xaf\x6a\x08\x44\xa8\xfe\xa9\xb0\x4d\x66\x79\x45\x53\xb4\x60\x91\xb8\x2c\x9b\x95\x16\xcb\x7e\xbd\xd1\x32\x33\x58\x0c\xc3\x74\xea\x37\x33\xf8\xbf\x66\xa5\xc3\x1d\xa3\x92\x11\x2a\x81\x23\xc9\x10\x46\xca\xfc\x6a\x84\xe3\x50\x9c\x4e\x92\x5a\xbc\x87\x9c\x9c\xe0\xcb\x02\x44\x4d\x93\x46\x63\x7d\x8a\x70\x59\x02\xcd\x16\x71\xc0\x77\xfb\x15\x62\x49\x92\x2c\x7d\xb6\x3c\x57\xa0\x2c\xe1\xaf\x34\x34\x0b\x54\x04\x62\x92\x4c\xff\xc4\x88\xc2\x57\xb3\xba\x95\xe3\x43\xf1\xc1\xe0\xb2\xa8\xd7\x4f\x92\xa4\x9f\x25\xa3\xac\x62\x95\xbc\x23\xa7\xd4\x91\xf1\xef\x95\x62\x85\x02\x64\xec\x7c\x8d\x97\xb1\x4d\xf5\x3a\xcd\x32\xac\x92\x7a\x42\xb2\x38\xb4\x5b\x96\x06\xfe\x3e\xd0\x53\x56\x49\x2b\x0e\xbd\xdf\x52\x46\xaf\x81\x4b\x5f\x1a\x5a\x13\xe9\x10\xdd\xb7\x63\xb7\xfa\xef\xb0\xda\x69\x4c\x42\x7e\x6e\xf1\x67\x58\x1c\x18\xbe\x42\x05\xd0\x05\x5b\x3a\x16\x12\x35\xed\xe5\x4f\x88\xa0\x7f\xd8\xbe\x9f\x10\xf9\xfb\xdf\x43\x16\x7e\x20\x1f\xd1\x29\x62\x1f\xc8\xc7\x83\xac\xf9\x19\x72\x5c\x15\xf2\x77\x9e\x01\x0f\xcc\x4c\x66\x3a\x10\x53\x3d\x84\x6e\x50\x4e\xa0\xc8\x44\xad\xad\x29\xa3\x12\xe8\x2d\xf8\xe3\x2f\xb8\x58\xa2\x0f\x1f\x8d\x1b\xd0\xb2\x31\x75\xb3\x23\xa9\x71\x6d\xcc\x2a\xbf\x5b\xb4\x9c\x1f\xd4\xf6\x6a\x6a\x8f\xac\x39\xe8\xc2\xa3\x2a\x3c\xd6\x0d\x8b\x0c\x4b\xde\xb3\x8b\x12\x73\x01\x01\x3b\x22\x8d\xba\x55\x32\xc8\x94\x6a\x19\x30\xb1\xfb\xfa\xe4\x04\xfd\xde\x35\xe5\xe8\x2b\x29\x0a\xc4\x68\x71\xa3\x59\x8e\x6d\xd7\x86\x5c\x03\xb5\x22\x49\xd0\x6f\xcc\xfc\x89\xb6\x80\xa9\x40\x4a\x81\x38\xd8\x26\x01\xb7\x10\x52\xcd\x82\x85\x15\x7a\x92\x24\x46\x1e\xb1\x46\x42\x2b\xf5\x14\xfa\xef\xac\xe5\x49\x0b\x67\x65\x74\x92\xc5\x8b\x11\x1c\xd0\x7e\x7f\xd8\x74\xd4\xda\xe4\xeb\xc2\xb5\x6d\xbb\xeb\x56\xb0\xb0\x17\x4b\x44\xa8\xec\x3b\x64\x41\x26\xaf\xce\xcf\xce\x68\xce\x12\xcf\x57\x37\x7e\xbe\x55\x5c\x2f\x6c\x18\xf3\x35\xdb\x3e\xe6\x56\x0d\x51\x44\xe0\xf6\xb4\xda\x0f\x1c\x77\x69\xac\x1f\x5f\xa5\x0a\xff\xd6\x16\x3d\xbc\x35\x77\x3b\x1d\x3d\xbc\x67\xbf\x68\x35\x3d\x52\xd4\xa2\x1c\x17\x02\x74\x50\x14\x10\xa6\xf8\xa2\x57\xae\x01\x28\x7e\x7d\xfa\x4b\x30\xba\x9e\x1f\xcf\xd1\xa5\xfe\xe3\xdf\x9a\x1e\x2b\x9b\xf9\x27\x2b\xbf\xdf\xe0\xeb\x08\x4f\x6a\xef\x44\x9d\xc7\xc3\x67\x8c\xc2\x43\xcb\x77\x04\xe0\x62\x79\x18\x48\x4b\xcc\xcf\x0f\x8d\x75\x9a\xee\x13\xbf\x3e\xa0\x1b\xab\x99\x67\xdb\x8e\xfd\x58\x50\xf1\x5a\x6d\x28\xc1\xb8\x17\x3b\xa2\xc5\x88\xa8\x96\x81\x19\xb5\x3e\xb8\xb8\x62\x55\x91\xfd\x8b\x13\x09\x67\x94\x48\x82\x0b\xf2\xdf\xc0\x95\x08\x75\x70\xa9\x96\x32\xc6\xb6\x25\xf6\xa3\xe4\xbc\xba\x2c\x48\xaa\xa8\x41\x2d\xb0\x47\x84\x12\x6d\x39\xbe\xf6\x80\x05\x19\x00\x6f\xcf\x25\xb9\x9d\x1e\xb4\xf7\xb5\x1d\xfb\xf6\x3e\xae\xc9\xee\xb7\x98\x40\xaf\xd7\x5d\x1f\x8a\xe4\x6a\x63\x71\x50\x5b\xee\xc9\x31\x47\x81\x67\xee\xef\xac\x29\x84\xe5\xa4\xe5\x22\x98\x18\x39\xa0\xeb\x6f\x02\x55\x94\x7c\xa9\xea\x30\x45\xcd\x99\x48\xab\x9a\xb2\x58\xa2\xd0\x2d\x30\xa1\x9d\x99\xeb\x61\x53\x2b\x67\x6d\xb6\x93\x66\x01\x37\x28\x79\x5d\x9f\xc9\xf5\x3e\x6e\xa4\xac\x8c\x4d\x0b\xc4\xbc\x93\x8d\xb9\x0d\xc3\x2e\x40\x7a\x58\x0a\x90\x0f\xc3\xb0\x60\x99\x05\xc9\x50\x7d\x48\xc7\x71\x2d\x8e\x5d\xe8\x14\x91\xec\x3e\x98\xd2\x77\x94\x5e\x61\x9e\xa5\x2c\x83\xac\x7d\xa8\x6a\x63\x1e\xcd\x88\x5b\x9f\xa4\xd3\xa9\x18\x74\x98\xeb\x34\xc1\x80\xe3\x1c\x45\x8b\x26\xe6\x71\x7b\xc8\x13\x39\xe6\xab\x60\xdd\x2b\x52\x4e\x4a\xe9\x32\x8f\x3f\xb3\x34\x0c\x3d\x58\x5a\x69\x4b\xa7\xc7\x28\x5f\xd0\xed\x97\x58\x8d\xf8\x99\xa5\x1d\x23\x62\x39\xf7\x49\x11\x25\xd2\xff\xc0\xe9\x67\x49\xd2\xcf\xe2\x00\x76\x9f\x82\x1c\xd4\x44\xd2\x63\x37\xb1\xc6\x71\x08\xd9\x7c\x2b\x93\x8b\x92\x13\x2a\xf3\xc5\xfc\x1f\x3f\x8a\xf5\x8f\xe2\x9f\x73\x15\x53\xbb\x83\x41\x8b\xcd\x35\x19\xfb\xb9\xec\x88\x2a\xd2\x37\x6b\x64\x66\x4e\xf9\x5f\x81\x02\xc7\x12\x7e\x05\x29\x81\xa3\xa4\x73\x88\x9f\x9c\xa0\x5f\x41\x2a\x12\x3b\x86\xa3\x1d\x3f\x75\x06\xd8\xbd\xce\x21\x05\x72\xdd\xb6\x7b\x47\x07\x78\x36\xb0\xe2\x62\x19\xae\x53\xe7\x74\x43\x96\x1a\xb3\xd7\x39\x15\x5a\x6e\x75\x97\x05\x17\x07\x58\x70\x31\xc0\x82\xc6\xf4\x97\x9c\x95\xc0\xd5\x91\x1d\xc1\x08\x54\x09\xa5\x09\x2e\xd0\xd3\x07\x47\x3c\x7b\x06\xb0\x59\x78\x7b\xbc\x97\x51\xcd\xc1\xd1\xcf\x22\x74\x8a\x3c\x08\xad\xad\xd1\xde\x19\x98\x66\x68\x31\xb4\x3d\x96\xdd\x2e\x93\xa6\x5d\x5a\x7e\xf6\xc6\xe0\xc2\x34\xf5\x9f\x14\x6a\xd2\x55\x3d\x1e\xb2\x3a\xc3\x73\xaf\xe1\xf3\xc8\x4e\x8e\x8a\x9a\xcd\x10\x3f\x76\xae\x99\x4e\x72\x1d\xe7\x9a\xc9\x4b\x74\x7a\x8a\x5e\x7a\xc1\xee\xc9\x09\xa2\xac\x20\x54\xae\xd1\x86\x99\x6b\x30\xd1\x74\xd6\xd1\xc4\x78\xb4\xeb\x41\x44\x3d\xd7\x55\x07\xed\x42\x38\xb1\xab\x1f\x6b\xf4\xbc\x5f\x71\x56\x9d\x35\x5b\x3e\xf7\xde\x8f\xc1\x45\xa9\xe2\x80\x28\x5a\xf6\x7e\x5a\x31\x77\x49\x45\x2b\x00\x47\xab\xf8\x4a\x64\x7a\x85\xf2\xfb\x22\x3f\xc5\x02\xc2\x33\x73\x1d\xf4\x1b\x3a\xba\xcc\xd0\x9b\xe8\xf9\xa2\x9f\x4d\xcb\x49\x6c\xb2\x32\x17\xa5\x8d\x74\xcf\xb1\xa2\x10\x97\x65\x61\xf2\xea\x94\x51\x44\x49\xe1\x9c\x4e\x8c\x22\xb2\x21\x75\x7e\x7a\x62\x98\xa2\x17\x5f\xd8\x0d\x7a\x48\xd9\x3d\xdf\xf4\xd8\x44\x73\xc2\x5d\x6d\x8e\xb2\xbf\x8e\xf6\xcc\x3c\xdb\x48\x72\xf4\x83\x59\xd9\x3f\x14\xcf\xc4\x9b\x2f\x15\x2e\x16\xfe\x49\xb9\xf4\xc4\x5f\x62\x4a\xd2\xc5\x3c\xc5\x54\x59\xa2\x52\x33\x2f\xe7\x6c\x8b\x30\x32\x54\x7c\x25\xf2\x0a\x65\x24\xcf\x81\x03\x95\xcd\xb5\xcb\x3c\xc8\x17\x09\xa6\x83\x6a\xb3\x7a\x5c\xb6\xa9\x7d\xcf\xd9\xa6\x45\x0c\x18\xdf\x1f\x4e\xb5\x38\xbd\xf4\xd7\x90\x91\x7e\x31\x00\x62\xd6\xe8\x50\xa8\x5b\x83\x7e\xbd\x76\x77\xa1\x6c\x5d\xe3\x64\x00\xa5\xb9\xb9\x20\x23\x37\x17\xfa\xc6\x35\xd6\x61\xb3\x0b\xc5\x26\x50\x48\x8e\x98\xb2\x91\x86\x27\xcf\x9e\xd9\xed\x40\x49\x31\x7b\xb6\x9f\xcd\x9e\xd9\x44\xe4\xe1\x04\xcb\x7e\xf6\x8c\x25\xf5\xca\x67\x54\xb2\x05\xab\xe4\x72\x36\x7b\xd6\x93\x27\x77\x83\x14\xf1\x04\x44\x78\x64\x13\x6a\x77\x8e\x39\x5b\x0e\xd2\x30\x99\x29\x35\x6a\x63\xe3\x77\xb3\xd9\x33\x89\xf9\x06\xe4\x0a\x01\xd7\x09\x9f\xe0\x99\x43\xa2\x39\xcc\x96\xb3\x67\x24\xd7\x03\x7e\x70\x0c\x34\x1b\x22\xf0\x37\xff\x8b\xea\xf3\x49\x32\x2d\x73\x2d\xf2\x43\xeb\xaf\xd1\x8f\x62\xae\x17\x5e\x2e\x8d\x10\x5e\x98\xbb\x9c\x17\x06\xa7\x43\x77\x38\x7a\x6b\xd8\x94\xab\x79\x32\xa1\xc3\x65\x92\x59\x3e\xa7\x15\x37\xdb\x90\xe6\x8c\x6f\x4d\x64\x20\x24\xe3\x90\x39\xce\x3b\x32\xa3\xe3\x46\xbb\xd4\x62\xa9\xb0\x66\xbc\x56\x2e\xfd\x43\x1b\x26\x67\xcb\xde\xe8\xb6\x5d\x9d\x2d\xf8\x52\x11\x0e\xd9\x9b\x43\x03\xeb\x18\x7c\xd2\x09\xe3\x5d\x62\xbf\xe7\x98\x0a\xa2\xa8\x0e\xfa\x92\x37\xdf\x4a\x26\xc0\xa5\x46\x6d\xf3\x1f\x16\xa7\x70\x34\x7c\x41\xe6\x6d\xc3\xdc\xf8\x22\x73\xcf\xd4\x58\x15\x71\xa8\xd7\xfc\xa8\x41\xd9\xd8\x24\x0c\x0e\x07\x0c\xcf\xf2\xa7\x50\xa3\x9c\x6f\x12\xb0\xea\xb4\xd5\x90\xd8\xfb\x49\xa5\x35\xce\x40\x79\x81\xa7\x47\x0b\xe3\x68\xe1\xe8\x01\x5a\x6d\xe7\xcb\xbb\x93\x63\xf8\x32\x74\x26\x7f\x07\xb2\x1c\x49\x92\x6c\x61\x92\x80\xde\x93\x2d\x3c\x56\xf1\x7c\x93\xc0\x29\x2e\xe6\x4b\xbf\xb5\x20\x42\xce\x97\x13\x28\x7c\x63\xc1\x3c\x1a\x2a\x1d\x2d\x84\x4a\xd8\x00\x9f\x24\xb0\x33\x2a\x1f\x21\x25\x79\xc1\xb0\x9c\x44\xc7\x2f\x6a\xc6\xe3\xa0\xe4\x10\x61\x1c\xf2\x1e\xb2\xfa\xf1\x4c\xdc\x69\x30\x80\x31\xd4\x98\xc2\x1d\xf7\x06\x87\xfc\xad\xde\x08\xad\xc6\x77\xb8\x74\x16\xcd\x86\x35\xa2\xba\xf4\x5e\x4b\xf4\xfb\x7c\x0e\x45\x47\xa4\xa8\x2e\xc7\x29\x8a\xa1\xca\x8f\x3c\x06\xe9\x4b\x5e\x15\x05\xfb\x0a\xd9\xeb\x2b\x46\x52\x10\x31\xaa\x64\xac\xf1\x19\xd5\x2f\x26\x26\xd9\xe4\x95\xcb\x7e\xaa\x79\x7f\x31\x42\x3b\x08\x7c\x9a\xaf\xd0\xfc\x93\x82\xb6\x5f\x69\xaf\xe5\x55\x25\xd9\xc6\xa6\x72\xb2\x03\x6a\x79\x07\x21\x3b\x1c\x30\x8f\x62\xc1\x39\x96\xca\xba\xc5\xed\xa3\x95\xce\x5c\xb6\xd7\xf8\xd4\xd3\xfc\x0e\x84\xc0\x1b\x30\xbd\xaa\xd3\x73\x0d\x1e\x80\xec\x8d\x44\xc9\x3b\xfc\xed\x2d\xd0\x8d\xbc\x42\x2f\x63\x08\x7f\x87\xbf\x91\x6d\xb5\x35\x53\x62\xc9\x57\xad\x6e\x1d\xd5\xa2\xef\x8b\x1f\x8a\x22\x42\x27\x51\x44\xe8\x2d\x29\x6a\xd6\x79\x70\x8a\xf0\x37\xfd\x7c\x16\xbd\x4c\x5e\x0e\x39\x89\xf1\x27\x81\x15\xe1\x84\x83\xa0\x91\xe0\x9f\xf6\x71\xed\xbd\x91\x5b\x08\x68\xe8\x8d\xc0\x39\xfa\x10\x5e\xa9\xe0\x62\xd1\xc2\x7a\x79\xcf\x52\x1a\xd3\xc2\xfb\x94\x99\x51\xd2\xe9\x32\xab\xb1\xb8\x7f\x99\x45\xa2\x7c\x1b\x91\x39\xa4\xbf\x8f\xc8\x9a\x47\x16\x09\xea\x7c\x0d\xa0\x1f\x61\xe8\xd7\xf8\x87\x3e\x09\x70\x8c\x50\xe0\xb6\x35\xb1\x9a\x72\xef\x5d\x85\x23\xdf\x34\xc6\x7a\x5c\xb1\x64\x1e\x0f\xb8\x58\x11\x4c\x08\xe8\xbd\x6e\x28\xd5\x88\x35\x79\x3e\x13\x8b\x3b\x3e\xf8\xef\xf9\x5f\x57\x42\xb2\x6d\x9d\xbe\x77\x10\x12\x97\x35\xdc\xe2\xb2\x24\x74\xa3\x3f\x0a\xd0\x57\xbb\x0e\xd2\x3b\xd3\x95\xd8\xff\xa3\xb9\xfb\x14\xa4\x83\x4e\x2b\xa7\x58\x43\xed\x17\x85\x85\x5b\x0b\x84\xdd\x1b\x8b\xfb\x38\x6e\x6f\x02\x42\x7f\x78\x89\xfe\xe9\x5d\x08\xd8\x04\x55\x38\xc4\xcf\x4c\x5a\x18\xd0\x33\xb7\x99\x0d\x9d\x59\x5e\x16\x2d\xfe\xb1\x43\x09\x29\xc9\x49\xaa\x39\xfb\x0b\xe3\x4d\x8a\x23\xb8\xbc\x69\x5a\x83\xe1\xcd\xed\xae\xc9\x9a\xb9\xef\x49\xf4\xf7\x19\x9f\xe1\xa6\x4e\xe5\x8c\x5d\xa2\x0e\xe1\xb0\xd0\x80\xba\xd7\x30\x03\xe8\xb8\xe4\xe2\xf5\x0a\xb1\xcf\x56\xfc\x83\x0b\xbb\x6c\xce\x3b\x5c\x7e\x50\x4b\x7d\xfc\x49\x4d\xeb\x70\xfa\xda\x67\xf2\xc9\x09\xfa\x17\xa0\x94\x55\x45\xa6\x79\x9b\x13\x9a\x21\x22\x57\x48\x30\x54\x80\xfc\x9b\x40\xe9\x15\xa4\x9f\x11\xb3\xcf\x40\xd9\x57\xe0\xe6\xc6\x81\xd0\x0c\xbe\x41\x86\x44\x09\x29\xda\xe2\xd2\x97\xd9\x21\x3c\xdf\x2a\x10\xaf\xb1\x80\x1e\x84\xeb\x27\xeb\xbd\x0c\x11\x81\x0c\xf3\xaa\x28\x3c\x19\x89\x70\xe4\x16\x97\x91\xd2\x1a\x58\x6b\xb1\x54\x30\x3e\x18\x61\x7d\x8c\x95\x55\x04\xf9\x01\xd5\x2e\xcb\x58\xc1\xa0\xb6\x9a\x4b\x93\x01\xe5\xb4\x97\x8d\x44\x20\x22\x10\x46\xd7\xc0\x6f\x10\xce\xae\x31\x4d\x21\x43\x8a\x01\x1a\x3d\x79\x85\x25\xba\x61\x95\xbd\x45\xd6\x92\xa6\x00\x19\xba\xac\x24\x22\x14\x09\xb6\x05\x05\x48\x4f\xaf\x59\x89\x2a\x01\x5a\xd4\x71\xef\x45\x6c\x0e\x33\x24\x24\x54\x79\xfd\xc9\x45\x8e\x53\xd8\x35\xa9\x73\x7b\x19\xa6\x87\xed\x5a\x76\x3b\x32\x4b\x79\xe8\xfe\xeb\xf0\x35\x7b\x8f\xf9\xeb\xb5\x3d\x23\x8f\x59\x03\x91\x76\x5e\xb6\xe2\x52\x5f\x78\x35\xa2\x55\x82\x3c\x9c\x90\x1f\xfb\xb0\x2a\x5c\xef\x74\x8a\xa2\xee\xda\x67\xe3\xa4\x4c\xb0\xf7\xd8\xb2\x99\xa1\x50\x68\x3d\x43\x38\xf6\x3f\x2a\x6c\x33\x7d\xbe\x1e\xb5\x7c\xfe\x33\xa2\xe3\xc1\x58\x56\xfd\xf3\xdb\xd7\xdd\xd8\x53\xc5\x9c\x01\xac\xd6\x1d\x66\x18\xa6\xaf\x07\xd2\x07\xc7\xfb\xfd\xa4\x10\xdf\x39\x94\xcd\xb4\x7d\xe3\x9b\xac\xba\xb4\xb5\x72\x01\x0e\x3b\xbf\x63\xdd\x9b\x37\x38\x48\x5d\xbd\x40\xaf\xda\xab\x7f\x41\xc7\x7a\x40\x5c\x71\x4b\x70\xd0\x92\xfb\x9d\x16\x37\xc1\x0a\x5e\xbb\xa1\xa0\x35\x32\x0a\x7a\xf0\xcc\xcd\xeb\xf7\xdb\x0d\xf4\xd6\xc8\x29\xd0\x6b\xf7\xbc\x0d\x5d\xb7\x6b\xe8\xfa\x35\x70\x30\x3a\x78\x0e\x9c\xc4\x2f\x58\x72\x48\xdb\xd2\x76\xad\x35\x29\xde\xa8\x48\xb8\xc1\x63\x36\x07\xb8\x69\x5e\xf7\xbc\x37\x6b\x3d\x32\x8b\x5a\xc9\xbb\xf4\xa9\xbb\x6c\x93\xc1\xdd\xf5\x47\x81\xfb\x85\x14\x12\xb8\xbe\xde\xf3\x7a\x5d\xab\x01\x1a\x8c\x8a\x83\xcb\x38\x90\x0d\xfd\x4f\x08\x74\xd2\xb5\x5a\xb8\xfe\xa8\x28\xb8\xf6\xfd\x99\xd7\x63\x5a\x0c\xbc\xa6\x37\x0a\x96\xf7\x14\xd6\xeb\x75\xad\x06\x66\x30\x2a\x0a\xae\x9f\x9b\x6a\x3a\x9b\xc6\x75\x37\x7f\x15\x09\xb4\xb3\x4d\xea\xb6\x75\x27\xa1\x12\x05\xd1\x4b\x38\x39\x90\x75\xe3\xba\x9b\x94\x8a\x04\xda\x45\xd3\xb6\xad\x3b\x39\x84\x18\x88\x6d\xcb\xe9\x19\xcc\x49\x76\x52\x1b\xa6\xb6\xa2\x37\x8d\x06\x37\x7f\x4c\x14\xd0\x73\x4e\xb6\x98\xdf\xb4\xd4\xdc\xb5\x1a\xb0\xc1\xa8\x28\xb8\x7f\x00\xce\xda\x06\xbd\x6e\x5b\xdb\x54\x6e\x33\x22\x12\x62\x78\x2b\x6c\x20\x9a\xb6\x75\x3b\x39\x1c\x05\xf1\x02\x52\x0e\xc1\x87\x15\xa6\xa5\xfe\x30\xc5\xf6\x46\xc2\x6a\x6f\xeb\x0b\x6f\x5b\x5f\x4c\xda\xd6\x17\xe6\x39\x80\x0f\x4b\xb7\x58\x58\x75\x6f\x1c\xac\xea\xd2\xbe\x89\x74\xc0\x4c\x53\xfd\x29\x7c\x33\x20\x4e\x0b\x3b\x77\xf9\xea\x5f\xd3\x68\x50\xf4\xc7\xc4\x01\x6d\xa1\xe8\xe1\x37\x8a\x9c\x5d\x61\x38\x13\x31\xee\x83\xf7\x47\x95\xdf\xc1\x19\x1f\x58\xf8\x71\x7b\xe5\x26\x8c\x4f\x9e\x5c\xf2\x27\x97\xfc\xc9\x25\x7f\x72\xc9\x9f\x5c\x72\xf4\xe4\x92\x3f\xb9\xe4\x4f\x2e\x79\x0d\xf1\xc9\x25\x7f\x72\xc9\x47\x5d\xf2\x5d\xf7\x13\xc4\x3b\x7c\x68\x65\x6e\xef\xe2\x8b\x74\xf4\xd7\xc5\x8a\x85\x60\x4a\x49\x4c\x5a\xef\xc3\xc7\xb1\x4f\x2a\xee\xad\x52\xd6\x14\xbc\xfe\x57\xeb\x65\x4d\xab\x2a\x73\x3b\xf2\x6e\x55\x3b\x6b\xca\x12\x0f\x52\x41\xeb\x7b\x70\xe6\x81\xaa\x69\xdd\x9e\x77\x77\xab\xa9\x35\xba\xbb\xbe\x43\x65\xad\x69\x02\xf8\xff\x5a\x5f\x6b\x1a\x97\x1e\x77\x0d\x01\xf3\x85\xef\x79\x81\x49\x58\x0d\x62\xd2\xd9\x10\x94\xda\x7a\xd8\x4d\x6f\x71\x7d\x5c\x6a\x97\x34\x58\x1d\x54\xc0\xdb\x14\xb0\x9a\xc6\x9d\xdb\x97\xb1\x1a\x77\x3f\x7a\xca\x56\x75\xbf\x02\x3f\x54\xbf\x2a\x89\xf2\x3a\x82\x32\x56\x3d\xfb\x60\x5c\xff\x7b\xcb\x59\x49\x5e\x41\xf8\x18\xf0\x7e\x2a\x5a\x45\x31\xcd\xaf\x6b\x15\x41\x7f\xd2\x94\xb7\x1a\x1f\xbb\x58\x46\x7d\x4a\xbc\x0b\x0e\xf7\xf1\x09\xbb\x5a\x29\xa2\x0a\x29\x59\x3d\xe8\x7c\xa8\x1c\xf1\x0d\xee\x43\x54\x55\x8a\xad\x93\x14\xa0\x3d\x52\xfd\x27\x9e\x94\xdb\x14\x4d\x22\x79\x74\x05\xa0\xee\xa7\x3b\x7e\xb1\xa4\xbd\xcf\xae\x17\x0f\x54\x84\x29\xb6\xac\xd2\x7d\xf3\xf7\x7b\xd5\x58\x7a\xde\x57\x64\xe9\xce\x65\x93\x62\x28\xbc\x93\xfd\xee\x3b\xc6\x4b\xdd\xd2\xc2\xcc\xca\xe5\x36\x08\xf6\x1d\xbf\xba\xa5\xaf\x8a\x86\x3d\x7e\x47\xcb\xf4\x35\xcc\x9e\x74\xfb\xd3\xec\x99\x76\x55\x07\xfb\x10\x38\x3c\xac\x07\x3f\xd0\x1f\xfb\x3e\xbf\x73\x56\x84\xc7\xfa\xff\xb1\x62\x3e\x8d\x8c\x1f\xac\xa4\xcf\x10\x43\x1f\x67\x4d\x1f\x9c\x65\x1c\x44\xe3\xea\xf7\x97\xf8\x89\x62\xda\xc3\x15\xfa\x79\xbe\x1b\xaf\xf4\x13\x59\x1e\x22\xda\xf7\x88\xb5\x07\x7e\xa9\x88\x28\x37\xa4\xd9\xba\x31\x05\x23\xa2\xfc\x94\x07\x2b\x1b\xf1\x60\xcc\x72\x25\x24\x62\x66\x3d\x7c\x21\x89\x71\x2c\x22\xca\x49\xc4\xd4\x5b\xf1\x55\xf6\x7f\x02\x00\x00\xff\xff\x76\x87\x35\x72\x66\x64\x00\x00")

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

	info := bindataFileInfo{name: "templates/model.gotpl", size: 25702, mode: os.FileMode(420), modTime: time.Unix(1554963055, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRelationships_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\x4d\x6b\xf3\x46\x10\xc7\xef\xfa\x14\x83\x11\x25\x29\xae\xda\xe4\xb9\x19\x72\x28\x71\x79\xc8\x21\x21\xb8\x2f\x17\xe3\xc3\xd6\x1a\xdb\xdb\xae\x57\xca\x7a\xe5\x60\x84\xbe\x7b\x59\xeb\xc5\x92\xbc\x2b\x29\x26\x86\x06\x66\x02\x01\x6b\x67\xe6\x3f\xb3\x9a\x5d\xff\x70\xcc\x96\xff\xb2\x35\x42\x9a\x42\xf0\x3b\xea\xe0\x31\x92\x2b\xbe\x4e\x14\xd3\x3c\x92\xc1\x0b\xdb\x22\x64\x99\xe7\xf1\x6d\x1c\x29\x0d\xa3\x75\x14\xb0\x38\x52\xa8\xa3\x80\x47\x3f\xa3\xc0\x2d\x4a\xcd\xc4\xc8\xf3\xf6\x4c\x81\x42\x71\x8c\xdb\x6d\x78\xbc\x9b\xe1\x9a\xef\xb4\x3a\x40\xe5\x15\xcc\x6c\xeb\x9e\xb7\x4a\xe4\x12\xb8\xe4\xfa\xe6\x16\x52\xcf\x03\x00\x47\xa6\x87\xbe\x5c\x69\x96\x87\xa7\x29\x28\x26\xd7\x08\x3e\x4a\xcd\xf5\xc1\xf4\x31\x06\xbf\xcc\x0a\x93\x87\xbc\xdb\x46\x12\xd3\x68\x1e\xfc\x13\xf0\x15\xec\x36\x51\x22\xc2\x3c\x33\xaa\xba\x27\xf8\x26\xb8\x9e\x1b\xfc\xe0\x35\xf9\x5b\xf0\xe5\x73\x14\x62\x91\xc6\xda\xc2\x3c\x4d\x1b\x71\x59\xf6\x14\xe6\x1f\x17\xf0\x00\x3f\xd8\xdb\x4b\x8f\xf9\x6a\xa5\xad\x35\xdc\x08\x94\xa7\x86\x82\x47\x85\x4c\xe3\x2d\xfc\x52\x36\x61\x2c\x7f\x38\x81\x2d\x8b\xe7\x3b\xad\xb8\x5c\x2f\x7e\xb4\x2b\x3c\xc9\x55\x04\x27\x99\x52\xaa\xd8\xc3\x98\x29\x94\x7a\x0c\x3e\x5b\x96\xbb\xd7\x56\xae\xcb\x3a\xf7\xf0\x49\x4a\xfb\x46\xe6\x02\x8d\x4d\x6c\x27\x1c\x99\x8d\x2b\xfc\xb2\x6c\x04\x13\xd7\x66\x99\x56\x9a\x9d\xd4\xea\x29\x3a\x08\xa6\x18\x2b\x5c\x32\x8d\x61\x5b\xc7\xd8\x69\x75\x02\x5a\x25\x38\xb6\xa6\x43\x69\x0d\x6e\x29\xbd\x32\xc5\xb6\xa8\x51\x4d\x71\x65\x46\xdc\xec\x9f\x3b\xaa\x7a\xb1\xee\xe8\x60\x86\x6f\x09\x57\x18\xb6\x5e\x76\x69\xe5\x72\x15\xba\x9b\xd4\x4e\xcd\x0b\xbe\x9f\x16\x0a\x57\xb3\x74\x73\x96\xc7\xd8\x7c\x61\xfe\xf2\xd1\x39\xdf\xd3\x7a\xed\xc5\xa4\xf0\x31\xf8\xe2\xee\x38\x20\x03\x3a\xb0\x95\xef\xda\x10\x71\xe7\xe8\xb7\x5e\x6b\x5f\xa5\xb6\x6a\xef\x8f\xd5\x8a\xbb\xae\xdc\xd6\x82\xee\x7b\x0a\xca\x8b\x1a\x52\x92\xad\xac\x6f\x79\x59\xf7\x7d\x0a\x50\x9e\x0d\xf1\xcd\x9c\x8b\xf3\x51\xb5\x09\xa1\x0c\x7b\xd2\x66\xdd\x89\x86\x25\xe9\xf7\xea\x90\xe9\x0f\xee\xf6\xb0\xa4\xbe\xbd\xe4\x20\x0f\x39\x92\xbf\x49\xad\x38\xee\x1c\x03\x51\x3f\x89\xf3\xc5\xe9\x2c\x5a\x32\xd9\x2f\xae\xda\x5c\xc4\x7d\x67\xab\x28\xc5\x39\x36\x1f\x94\x2f\xcd\x7c\x59\x4d\x8a\x4b\xb8\xa4\x82\x8e\x61\xfb\xe3\x10\x9f\xdc\xcd\x87\x6e\xf7\xf2\xce\x8c\x83\x29\xae\x58\x22\xf4\x5f\x4c\x24\x67\x5f\x01\x75\xab\xfb\x55\x42\xad\xe0\x1e\xc1\xfe\xe1\xe2\x2b\xc0\xb7\xaa\x83\x11\xca\x64\x3b\xea\x2a\xea\x57\x21\xa2\x77\x0c\x1f\x37\x11\x5f\xe2\xf1\x65\x7f\xf0\x42\xfa\x67\x0c\xfe\xfe\xf8\x86\xe3\xa0\x99\xac\xef\x1a\x38\xee\xc0\xbe\xff\x06\xe8\x98\xf7\xd2\xfa\xcf\xe4\x80\x7b\xdb\x8f\x83\xe7\x44\x68\x1e\x8b\xce\xd7\x58\xfa\xb8\xbe\x67\x07\x0a\x5b\x4a\xee\x88\xf8\x98\xb7\x63\xa9\x95\xc4\xe1\x65\x79\x5c\x0b\xb4\xac\x3a\x01\xef\xcf\x38\x3c\x07\xbc\xfc\xe1\x95\x01\x2f\x17\x21\xc0\x73\x28\x11\xe0\x11\xe0\x11\xe0\x5d\x20\x43\x80\x57\x55\xf1\x89\x80\x47\x7c\x07\xc4\x77\xc4\x77\xc3\xbd\xaf\xc8\x77\xaf\x4c\x2f\x37\x44\x67\x44\x67\x44\x67\xce\x6a\x89\xce\x6c\x46\x74\x46\x74\x56\x33\xa2\x33\xa2\x33\xa2\x33\x77\x92\xab\xff\xfa\x36\x45\x81\x67\xbf\xbe\xe5\x0f\xaf\xcc\x77\xb9\x08\xf1\x9d\x43\x89\xf8\x8e\xf8\x8e\xf8\xee\x02\x19\xe2\xbb\xaa\x0a\xe2\xbb\xb6\x11\xdf\x35\x8d\xf8\x6e\x78\xc4\x97\xe4\xbb\xef\xa8\x5b\xb7\xcb\x0c\xcd\x41\xdf\x5f\x1b\xef\xbe\xa3\x26\xb6\x73\x28\x11\xdb\x11\xdb\x11\xdb\x5d\x20\x43\x6c\x57\x55\x41\x6c\xd7\x36\x62\xbb\xa6\x11\xdb\x0d\x8f\xf8\xaa\x6c\xf7\xcc\xe4\xc1\xc1\x77\x66\xe9\xfa\x8c\x67\x54\x88\xf3\x1c\x4a\xc4\x79\xc4\x79\xc4\x79\x17\xc8\x10\xe7\x55\x55\x10\xe7\xb5\x8d\x38\xaf\x69\xc4\x79\xc3\x23\xfe\x37\x9c\x67\x48\x85\xe8\x0c\x88\xce\x88\xce\x9c\xd5\x12\x9d\xd9\x8c\xe8\x8c\xe8\xac\x66\x44\x67\x44\x67\x44\x67\xee\x24\x9f\xf1\x2b\x5c\xfe\xbf\xf5\x30\x4d\xcb\x4f\x99\xf7\x5f\x00\x00\x00\xff\xff\xb5\x07\x23\xe4\x6c\x46\x00\x00")

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

	info := bindataFileInfo{name: "templates/relationships_registry.gotpl", size: 18028, mode: os.FileMode(420), modTime: time.Unix(1554956824, 0)}
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

