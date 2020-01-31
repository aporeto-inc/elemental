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

	info := bindataFileInfo{name: "templates/README.md", size: 77, mode: os.FileMode(420), modTime: time.Unix(1578956520, 0)}
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

	info := bindataFileInfo{name: "templates/identities_registry.gotpl", size: 5323, mode: os.FileMode(420), modTime: time.Unix(1578956520, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesModelGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3d\x69\x73\xdc\x36\xb2\x9f\x3d\xbf\x02\x99\x52\x6c\x8e\x77\x44\xe5\xf3\x64\xb5\x55\x8e\x8f\x3c\xd5\xf3\xa1\xb2\xbc\xd9\xaa\xe7\x72\xad\x20\x12\x1c\x21\xe6\x00\x0c\x00\x4a\xd6\x9b\xe2\x7f\x7f\x85\x83\x24\xc0\x13\x1c\xcd\x38\xca\x5b\xf9\x43\x22\xe1\x68\xf4\x85\x46\x77\x03\x6c\x65\x30\xfa\x0a\xd7\x08\x6c\xb7\x20\xbc\x40\x22\x7c\x49\x49\x82\xd7\x39\x83\x02\x53\x12\xbe\x87\x1b\x04\x8a\x62\x36\xc3\x9b\x8c\x32\x01\x82\x19\x00\x00\xcc\x93\x8d\x98\xeb\x9f\xd6\x34\x84\x19\x65\x48\xd0\x10\xd3\x13\x94\xa2\x0d\x22\x02\xa6\x65\x2f\x16\xd7\xf9\x55\x18\xd1\xcd\xc9\x3a\xa5\x57\x30\xe5\x78\x4d\x4e\x36\x6b\x7a\x72\xc5\x29\x69\x0f\xda\x60\x11\x5d\xa3\x34\xbd\x3e\x89\x68\x76\xc7\x05\xcb\x23\x91\x33\xa4\x07\x6e\xb7\xc7\x80\x41\xb2\x46\x20\xbc\xc8\x50\x14\x7e\xba\xcb\xd0\x39\xa3\x37\x38\x46\x8c\x4b\x24\x15\x34\x49\x07\x28\x8a\x7a\x0a\x22\x31\x38\x36\xbd\x4d\x10\xbf\xc1\x14\xc7\x8a\x52\x4f\x40\x45\x31\x5b\xcc\x66\xdb\x2d\x38\x4a\xa1\x40\x5c\xfc\x86\x18\xc7\x94\x80\xd5\xa9\x81\xf8\x56\x35\xbf\x10\x82\xe1\xab\x5c\x20\x5e\x0e\x90\x3c\xdc\x6e\xcd\xe2\x47\x78\x09\x8e\x10\xc9\x37\x72\xde\x55\x8e\xd3\xf8\x35\xc9\x37\x5c\x83\x68\x82\x2e\x8a\xd9\xc9\x89\x14\x8f\x9a\xa1\xa8\x06\x45\x01\x18\xca\x18\xe2\x88\x08\x0e\xc4\x35\x02\x19\xe5\x1c\x5f\xa5\x08\xdc\xc0\x34\x47\x1c\x24\x94\x01\x58\x62\xa1\x88\xd1\xd3\x2b\xcc\x8c\x64\xe7\xe1\x4c\x48\x88\x2d\xf8\x5c\x30\x4c\xd6\xb3\x59\x44\x09\x2f\xe5\x5e\xb3\xef\x88\xc0\x0d\x5a\x82\x23\xb5\x9a\xa4\x42\x4f\xfe\x4d\x2f\x6e\x58\x68\xd0\x26\x7a\xa5\x26\xc6\x7a\xaa\x1c\xa0\x7f\x2a\x8a\xd0\x2c\x52\x4f\x69\x61\x75\xaa\x49\x29\x67\x94\xc2\xa9\x65\x53\xff\x3c\x93\xd8\xe2\x04\x10\x2a\x8c\x6c\xde\xd1\x18\xa5\xe1\x2b\x24\x60\x74\x8d\xe2\x9a\xb1\x76\xef\x6b\x22\xb0\xb8\x33\xcc\x39\x8b\x91\xfa\xb5\x89\x7a\xd5\x4e\x13\xf5\x3b\xbd\xfa\x1d\x45\x22\x9c\xdd\x40\xe6\x07\xef\x14\x54\x3b\x25\xac\x1a\xb7\x8a\x18\x39\x74\x05\x2a\x0d\xb4\x40\x7d\x44\x5c\xc8\xde\xa2\x98\x2f\xd5\xd0\x97\x50\xa0\x35\x65\x77\xab\xae\xa1\x34\x67\x51\x25\x64\x3d\xfe\x5c\x6f\xf5\x55\x1b\xb4\xe9\xa9\x47\x32\x7c\x03\x85\x1c\xd9\x1c\xa8\x3b\x8a\x62\x39\x2b\x26\xf2\x7a\xbb\xed\x1a\x71\xc6\x3f\x52\x2a\xc6\x64\x71\x9e\xe6\x0c\xa6\xa0\x28\xde\x62\x2e\x6c\x69\x40\x90\xca\x16\x9a\x78\xcc\xad\x14\xdd\x67\x8d\xcf\x5f\x9e\xf7\x8e\x94\x04\x9f\x9c\x00\x4b\x3b\x44\xce\x88\x56\x0d\xdc\xa9\x1a\x1c\x60\xa2\x7e\x95\xd8\x86\xb3\x24\x27\x11\x08\xa8\x27\x2e\x8b\x6a\xa5\x60\xd1\xad\x37\x4a\x66\x1a\x8b\x7e\x98\xb5\xfa\xcd\x34\xfe\x2f\x69\x56\xe3\x0e\x41\x46\x31\x11\x88\x01\x41\x01\x04\xd2\xfc\x2a\x84\xfd\x50\x9c\x4e\x92\x5c\xbc\x83\x9c\x04\xc3\xab\x14\xf1\x92\x26\x85\xc6\xea\x14\xc0\x2c\x43\x24\x0e\xfc\x80\x6f\x8b\x25\xa0\x61\x18\x2e\x6c\xb6\x3c\x95\xa0\x0c\xe1\x2f\x14\x34\x03\x94\x3b\x62\x12\x54\xfd\x0a\x01\x41\xb7\x7a\x75\x23\xc7\x43\xf1\x41\xe3\x12\x94\xeb\x87\x61\xd8\xcd\x92\x51\x56\xd1\x5c\xdc\x93\x53\xf2\xc8\xf8\xf7\x52\xb2\x42\x02\xd2\x76\xbe\xc4\x4b\xdb\xa6\x72\x9d\x6a\x19\x9a\x0b\x35\x21\x0c\x86\x76\xcb\x42\xc3\x2f\x1c\x3d\xa5\xb9\x30\xe2\x50\xfb\x2d\xa2\xe4\x06\x31\x61\x4b\x43\x69\x22\xe9\xa3\x7b\x37\x76\xcb\xff\xf6\xab\x9d\xc2\xc4\xe5\xe7\x06\x7e\x45\xc1\xc0\xf0\x25\x48\x11\x09\xe8\xa2\x66\x21\x96\xd3\x7e\xfa\x19\x60\xf0\x77\xd3\xf7\x33\xc0\x7f\xfb\x9b\xcb\xc2\xcf\xf8\x0b\x38\x05\xf4\x33\xfe\x32\xc8\x9a\x57\x28\x81\x79\x2a\x3e\xb0\x18\x31\xc7\xcc\xc4\xba\x03\x50\xd9\x83\xc9\x1a\x24\x18\xa5\x31\x2f\xb5\x35\xa2\x44\x20\xb2\x03\x7f\xec\x05\x83\x05\xf8\xfc\x45\xbb\x01\x0d\x1b\x53\x36\xd7\x24\x55\xae\x8d\x5e\xc5\xc1\xbb\x74\xbe\x1c\xaf\x6a\x69\x4f\x35\xa7\x88\xe6\x84\xa6\xfc\x13\xbd\xc8\x20\xe3\xc8\xa1\xda\xd3\x76\x1b\x5d\x42\xb1\xd4\x20\x0d\xc6\x77\xfb\x9e\x9c\x80\x0f\x6d\x8b\x0d\x6e\x71\x9a\x02\x4a\xd2\x3b\xc5\x59\x68\xba\xd6\xf8\x06\x11\xc3\xf9\x10\xbc\xa7\xfa\x47\xb0\x41\x90\x70\x20\xf5\x84\x21\xd3\xc4\xd1\x0e\xb2\x28\x59\x10\x18\xd9\x86\x61\xa8\xd9\xee\x6b\x0b\x94\xee\x4e\xa1\xff\xde\xca\x1c\x36\x70\x96\xb6\x25\x0c\x9e\x8f\xe0\x00\x8a\x62\xd8\x42\x94\xae\xb0\xad\x0b\x37\xa6\xed\xbe\x1a\x6f\x60\x07\x0b\x80\x89\xe8\x3a\x4b\x91\x08\x5f\x9c\x9f\x9d\x91\x84\x86\x96\x4b\xae\xdd\x79\xa3\xb8\x56\x74\xa0\x7e\x3e\xc2\xfc\x35\x89\xd8\x5d\x26\xa4\x58\x24\x0b\x13\x98\x72\xe4\xe1\x71\x36\x3d\xcd\x8d\x1c\x22\x69\x84\xcd\x69\xa5\x37\x38\xee\xd8\x18\x6f\x3e\x8f\x24\x79\xcd\x08\xa8\x8e\x54\x9a\x71\xc7\x71\x51\x18\xaf\x2e\x34\xc4\x28\x3f\xae\x83\xbe\x53\x20\x98\x72\xc7\x2d\x3e\x6c\xb7\x2a\xfc\xf8\x44\xdf\xa8\x0d\x70\x24\xf9\x68\xb8\x10\x36\x59\x26\x39\xae\x90\x2e\x97\x96\x92\xb8\xfc\x9d\x53\xb2\x9a\x1f\xcf\xc1\x86\xaf\x65\x90\xaa\x7e\xbe\x52\x8d\xff\x56\x6c\x31\x1a\x30\xbf\x34\x5a\xf2\x1e\xdd\x8e\xb0\xb6\x74\x75\xe4\xe1\xde\x7f\x60\x49\x9c\x94\x16\x8d\x00\x0c\x16\xc3\x40\x1a\xca\xf4\x74\x68\x6c\xbd\x9f\x6c\x46\xac\x06\x34\x70\x39\xb3\x2c\xe8\xb1\x1d\x58\x4a\xbe\x4b\x9d\xe3\x94\x59\x81\x28\x08\x46\x04\xbe\x70\x0c\xb5\x11\x3d\xbf\xa6\x79\x1a\xff\x8b\x61\x81\xce\x08\x16\x18\xa6\xf8\x7f\x11\x93\xe2\x54\x91\xaa\x5c\x4a\xe7\x08\x1a\xca\x73\x14\x9e\xe7\x57\x29\x8e\x24\x35\xa0\x01\xf6\x08\x13\xac\xec\xd3\x6d\x07\x58\x24\x1c\xe0\xcd\xb9\x38\x31\xd3\x9d\xf6\xae\xb6\x63\xfb\x54\xf1\x6b\x32\xbb\xda\x27\x6a\xec\xf4\xfd\xfb\xc2\xc2\xd2\x24\x0d\x6a\xcb\x9e\xbc\x7c\xe0\xb8\xf9\x4d\xc3\xe4\x4b\x58\x82\x1b\xfe\x86\x0e\xb8\x1d\xba\x9e\x71\x90\x13\xfc\x47\x5e\xc6\x3c\x72\xce\x44\x5a\xe5\x94\x60\x01\x5c\x1f\x43\xc7\x89\x7a\xae\x85\x4d\xa9\x9c\xe5\xe1\x10\x56\x0b\xd4\x83\xc2\x97\xe5\xc9\x5f\xee\xe3\x4a\xca\xd2\xf0\x34\x40\xcc\x5b\xa9\x9d\x5d\x18\x76\x81\x84\x85\x25\x47\xe2\x30\x0c\x73\x96\x09\x70\x0c\x4a\x57\xc0\x8f\x6b\x7e\xec\x02\xa7\x00\xc7\xc3\x4c\x39\x39\x01\xbf\x22\xf1\xcb\xc5\x87\xf7\x00\x6f\x32\xad\xa6\x9a\x62\x69\x9a\xc1\x06\x32\x7e\x0d\x53\x29\x4e\x15\x4d\x26\x30\x42\xca\xab\xfa\x74\x8d\x39\xc0\x1c\xe4\x5c\xbb\x65\x82\x41\xc2\x33\xc8\x10\x11\xda\xab\x92\x88\x80\xb3\x57\xb2\xef\x1d\x25\x6b\xfa\xea\x97\xb3\x57\x00\x72\xf0\xe1\x0a\x45\xe2\xec\x95\x37\xa3\x0c\x76\xc1\x02\x04\x15\x06\x32\xce\x41\x8c\x51\x56\xb1\x0b\x27\x80\x82\xd3\x53\x40\x70\x6a\xf9\x32\x46\x31\x08\x4e\x97\xf2\x3f\xb6\x4f\xc2\xa5\xc1\x7a\xba\x91\x98\xd5\x16\x74\xd0\xa2\x97\xca\xe7\x7f\xdc\xce\x2c\x2b\x17\x5e\x08\xca\x50\x0c\x1a\xad\x96\x68\x4d\x8f\xa4\x44\x09\xb7\x25\xcc\x1f\x4e\xc1\x7c\x6e\x51\xc7\xbb\x87\x9d\x2a\xc9\x85\xda\xed\x3d\x8b\xff\x0b\x7d\x0b\xba\x01\x96\x3e\x9a\xb3\xa7\x0c\x16\xbd\xb0\xbb\x41\x35\x75\x6c\xf8\x57\x7b\xd3\x72\x2d\x19\xad\x89\x17\x0f\x5a\x13\x0d\x76\x01\x83\xb7\x9a\xc5\x1f\xe1\xed\x42\xeb\xa1\xaf\x1a\xee\x43\x03\x71\x22\xd7\xd4\x21\xfd\x6d\xf8\x4f\x62\x18\x13\xf0\xc5\xcf\xaa\xe3\x87\x9e\xe5\x11\x63\x8e\xc0\x0f\xac\xc7\x3d\x4a\x7c\xda\xa3\x5a\xa1\xd4\xd3\x45\xa7\x2e\x4e\x84\xb4\xbb\x2e\x1a\x45\xf4\x3b\x27\xba\x62\x98\x6b\xc8\xe2\x88\xc6\x28\x6e\x46\x33\xca\xbf\xf5\x56\xb4\x9d\x43\x98\x86\x61\xff\x25\x45\x37\x48\x65\xda\x9b\x1b\x4a\x76\x84\x2f\x53\xc8\xb9\x96\xd9\x59\xbd\xa3\x3c\x71\xac\x60\xb7\xce\xfb\xf2\x34\xee\x8f\x6f\xe6\xfe\x5c\xee\x4d\x98\x94\x69\xe2\x9e\xc4\x89\x17\x1d\x8a\x90\x07\x91\x21\x99\xe8\xa6\xd8\xce\x41\xd9\xcb\x23\x86\x33\x51\x5f\x30\xbd\xa2\x91\x9b\x61\xa2\x51\xae\x7c\x50\x35\x26\xa1\xcc\xf2\x64\x7c\x85\xfe\x8a\x46\x7d\xe2\xbe\x94\x44\xf1\xe8\x17\x18\x7d\x15\x38\xfa\xca\x07\xb0\xbb\x74\xae\x1a\x26\x92\xee\x6b\xab\x15\x8e\x7d\xc8\x26\x1b\x11\x5e\x64\x0c\x13\x91\x04\xf3\xbf\xff\xc8\x57\x3f\xf2\x7f\xcc\x97\x80\x86\xb5\xcb\xae\xa2\xa0\xba\x49\x7b\xb6\x8b\x96\xa8\x3c\x8d\x68\x25\x33\x1d\x7f\xfd\x8a\x08\x62\x50\xa0\x5f\x91\x10\x88\x81\xb0\x15\x5e\x69\xaf\xac\xd3\xec\x35\xf3\x67\xad\x01\xc6\xe4\x30\x14\x21\x7c\xd3\xf4\x48\x8f\x86\x3d\xad\x2e\x80\xc1\xc2\x5d\xa7\xbc\xba\x73\x59\xda\xe3\x17\x34\xd2\x2a\x6d\x16\x5c\x0c\xb0\xe0\xa2\x87\x05\x95\x53\x9e\x31\x9a\x21\x26\x83\x29\x0f\x46\x80\x9c\x4b\x4d\xa8\x13\x7d\xca\xa5\xf7\x67\x4f\x0f\x36\x2a\x35\xff\xbe\xbe\xdf\x6c\x31\xaa\xf2\x51\x7b\xcf\x31\x0b\x42\x63\x6b\x34\x77\x06\x24\x31\x08\xfa\xb6\xc7\xa2\xdd\xa5\x6f\xe3\x16\x86\x9f\x9d\x39\x58\xae\x9b\xba\x0f\x2c\xe5\x5e\x95\xe3\x51\x5c\x26\xf2\xf7\x9a\x3e\x1d\xd9\xc9\x5e\x59\x53\x3d\xc4\xce\x9d\x5a\x1e\x59\x8a\x88\x99\xbc\x90\xbe\xd9\x4f\x96\x6b\x74\x72\x02\x08\x4d\x31\x11\x2b\xb0\xa6\xfa\x49\x04\x6f\xfa\x4d\x4f\xc7\xb3\x9d\x16\x44\xd0\xf1\x2a\x61\xd0\x2e\x34\x27\xe2\x04\x5c\x43\x7e\xce\x50\x82\xbf\x81\x40\xe7\xdc\x94\x26\x39\x29\xb7\x05\x98\x3f\x9f\xb7\xa7\xb7\xd5\x6b\xd5\xa3\x76\xcb\xd6\xc2\xb6\xcb\x35\x0c\xf1\xa9\x37\x48\x37\x3d\xd3\xd3\x5c\x38\x5e\x71\xa6\xdc\x62\x1f\x9e\x17\xf6\x2d\x57\x52\xdf\x71\x19\x45\xb1\x02\xa5\x5b\x2c\xa2\x6b\x90\xec\x4b\x4c\x11\xe4\xfa\x09\x46\xb9\x6b\xe7\x2b\xa7\x7f\x0f\xa2\xd4\xac\x68\xb3\x79\x2c\x06\xf3\x11\xea\x20\xec\xa7\x83\xc1\xe2\x1e\x04\x5c\xc6\x7d\x99\xe5\x04\x36\xf2\xcf\xda\x5a\x99\x16\x4b\x2a\x48\xb7\x68\xbb\x05\xeb\xf6\x0d\x64\x5f\x51\x2c\x43\xba\x4b\x54\x66\xb6\x2f\x5b\xe6\xbe\xec\xf2\xcf\xd1\xb4\x30\x08\x2a\x18\x96\xed\xa9\xba\xcb\xac\x3a\x5b\x80\x40\x06\x62\x55\x86\x02\x4c\x89\xb7\x1a\x81\x95\x9d\xaa\x1f\xc8\x0f\x14\x3a\x23\x02\x4e\x2d\x32\xcd\x54\xe3\x0a\x75\x4e\x1a\x09\x19\xa5\x9f\xf4\x5a\x52\x91\x04\xf3\x9c\x28\xd9\x08\x5a\xae\x60\x3d\x47\x7a\xd6\x01\xfa\x99\xda\x99\xcf\x06\x0f\xd5\x67\x20\xf8\x91\x2f\x56\xe0\x47\x3e\x6f\xba\x5a\x8a\x9c\x56\x86\xc2\x37\x84\x53\x91\x43\x53\x7d\x62\x74\x1f\xf5\x31\xb3\x27\xa8\x4f\x0b\x83\xbf\x96\xfa\x18\xf4\xf7\xae\x3e\x86\x91\x0f\x59\x7d\xec\x5e\xa9\x4b\xe7\x50\x9e\x1f\x30\xcb\x52\xfd\x88\x86\x50\x35\xb4\x4e\x0a\x43\xe0\x71\x27\x5a\x3e\x46\x99\x78\x8d\xa0\x16\x0f\x8c\x9b\x36\xe4\xf2\x58\xb9\xe3\x63\x7d\xdb\xc2\xeb\x77\x8c\xbe\x1a\x63\xe6\xd5\xda\xf2\x83\x5e\xd9\x0e\x8d\xce\xf8\xeb\x3f\x72\x98\x06\x76\xbc\xb4\xb0\xc4\x9f\x41\x82\xa3\x60\x1e\x41\x22\xfd\xd1\x4c\x31\x2f\x61\x74\x03\x20\xd0\x54\xdc\x62\x71\x0d\x62\x9c\x24\x88\x21\x22\xaa\x37\x56\x73\xe7\xd6\x98\x53\x75\xe9\xa5\x57\xf7\xbb\x73\x9e\xb9\xc7\x7a\x8b\x16\xde\x9f\x59\x75\x15\xf8\x1e\xa7\x77\x7f\xb6\x6a\xe4\xd8\xee\x3a\xae\x7b\x81\x3d\xf7\x82\x66\x27\x19\x5a\x6d\x43\x57\x02\xaf\x10\xca\x1a\xcf\xc9\x62\x84\x32\xfd\x82\x0a\x8f\xbc\xa0\x52\x2f\x3f\xbd\x6d\xa4\x5e\xc8\xf7\xee\xd5\x4d\xb0\x3e\x79\x62\xed\xdb\x27\xc5\x6c\xf6\xc4\xbc\x94\x18\xbe\x9b\x2d\x66\x4f\x68\x58\xae\x7c\x46\x04\x0d\x68\x2e\x16\xb3\xd9\x93\x8e\xf7\x3a\xf5\x20\x49\x3c\x46\xdc\x8d\x29\x31\x31\x9b\x5a\x1f\x12\x83\x34\x4c\x66\x4a\x89\xda\xd8\xf8\xed\x6c\xf6\x44\x40\xb6\x46\x62\x59\xa6\x86\x9d\xe7\xd6\xa1\xe2\x30\x5d\xcc\x9e\x98\xdc\xf1\x0f\x35\x03\xf5\x5e\x75\x12\x22\xff\xb4\x4c\x35\xca\x94\xc8\x87\xd6\x37\xf6\x57\xda\xdb\x85\x16\xc2\x73\xfd\xa6\xec\xb9\xc6\x69\xe8\x2d\x99\xda\xb5\xe6\x4d\x88\x7e\xba\xad\x6e\xda\x70\x6c\xf8\x1c\xe5\x4c\x5b\x08\x92\x50\xb6\xd1\xa9\x2b\xae\x13\xd0\x15\xe7\x6b\x32\xbd\xf3\xab\x66\xa9\xa0\x91\xbd\x57\xbf\x28\x9b\x59\x9b\x59\x75\x7e\xf1\x6d\x79\xd1\xf8\x47\x8e\x19\x8a\x5f\x0f\x0d\xdc\xf1\xbc\xae\x12\x5f\x9f\x18\x24\x1c\x4b\xaa\x9d\xbe\xf0\xf5\xb7\x8c\x72\x54\x1f\x59\xa6\xf9\xa3\xc1\xc9\x1d\x8d\xfe\x00\xfa\x8d\xf5\x5c\x07\xcb\x73\xcb\x0a\x1a\x15\xa9\x51\x2f\xf9\x51\x82\x32\x47\xbe\x13\xe1\x2c\x7b\x6c\x51\xbf\x0b\xe0\xb0\xea\xb4\xd1\x10\x9a\x77\x92\xad\x53\xda\x39\x95\x35\x2d\x94\x81\xa0\xa6\x07\x91\x7c\x33\x5f\xdc\x9f\x1c\xde\xeb\xd7\x48\xaa\xbe\x03\x59\x35\x49\x02\x6f\xd0\x24\x01\x7d\xc2\x1b\xf4\x50\xc5\xf3\x4d\x20\x46\x60\x3a\x5f\xd8\xad\x29\xe6\x62\xbe\x98\x40\xe1\x6b\x03\xe6\xc1\x50\x59\xd3\x82\x89\x40\x6b\xc4\x26\x09\xec\x8c\x88\x07\x48\x49\x92\x52\x28\x26\xd1\xf1\x46\xce\x78\x18\x94\x0c\x11\xc6\x50\x32\xf7\xba\x4f\x77\x51\xab\xe9\xfe\x88\x38\x12\xe6\x4a\xe7\x0d\x65\xff\x83\x18\xd5\x9f\xc2\x8c\x66\x47\x6a\x2e\x76\x8f\x0c\xeb\xc3\xa7\x87\x41\xd6\x49\x74\x6a\x7e\x68\x31\x05\x58\x59\x15\xcf\x8d\xc9\x50\xf2\x56\xed\xc2\x46\xe3\x3b\x98\xd5\xe6\xd4\x24\xd3\x78\x7e\x65\x3d\x19\xef\xe6\xde\xd6\x26\x59\x4e\x68\x5d\x7b\x03\xf5\xe0\x9f\x08\x4c\x72\xd4\xc0\xda\x97\xdb\x3c\xbf\xea\x62\x2d\xcf\xaf\xbe\x23\x1f\xc3\x17\x69\x4a\x6f\x51\xfc\xf2\x9a\xe2\x08\x71\x9f\xfd\xa2\x8f\x9c\x33\xa2\x9e\xa7\x4f\x3a\x78\x96\xf5\x55\xa3\x9c\xf7\x3b\xc5\xa4\x85\xc0\xe5\x7c\x09\xe6\x97\x12\x5a\xb1\x54\xae\xd9\x8b\x5c\xd0\xb5\xb9\x50\x89\x07\xf6\xde\x18\x3b\xbc\x98\x00\x99\x17\x0b\xce\xa1\x90\x26\xdc\xcf\x58\x2c\xd5\xfd\x61\x73\x8d\xcb\x8e\xe6\x77\x88\x73\xb8\x46\xba\x57\x76\x5a\xfe\xcf\x01\xc8\x5e\x0b\x10\xbe\x83\xdf\xde\x22\xb2\x16\xd7\xe0\x27\x1f\xc2\xdf\xc1\x6f\x78\x93\x6f\xf4\x14\x5f\xf2\x65\x6b\xbd\x8e\x6c\x51\xf1\xe5\xa1\x28\xc2\x64\x12\x45\x98\xec\x48\x51\xb5\xce\xc1\x29\x82\xdf\x94\xc9\x00\x3f\x85\x3f\xf5\x79\xc2\xfe\xc7\x9d\x11\xe1\x84\xd3\xae\x92\xe0\x6f\xe6\x4b\xc6\xbd\x91\x6b\x32\x02\xbe\x38\x7b\x7b\x1a\x4b\x19\x41\x05\x0d\xac\x17\x7b\x96\xd2\x98\x16\xee\x53\x66\x5a\x49\xa7\xcb\xac\xc4\x62\xff\x32\xf3\x44\x79\x17\x91\xd5\x48\x7f\x1f\x91\x55\x8f\xd0\x43\xd0\xfa\xf4\x5a\x3d\x52\x57\x9f\x3e\x0f\x7d\x7f\x5d\x33\x42\x82\xdb\x94\xc4\x2a\xca\xad\x77\xe7\x35\xf9\xba\xd1\xd7\xad\xf4\x25\xf3\xb8\xc7\x8f\xf4\x60\x82\x43\xef\x4d\x45\xa9\x42\xac\xca\xb3\xea\x84\x43\xcd\x07\xfb\xe3\xe9\x97\x39\x17\x74\x53\x5e\xa2\xd7\x10\xc2\x3a\x6b\xbb\x81\x59\x86\xc9\x5a\x7d\x81\xad\xde\x79\xd5\x90\xde\xe9\xae\xd0\xfc\x1f\xcc\xeb\x8f\xf3\x5b\xe8\x34\x72\xba\x25\xd4\x6e\x51\x18\xb8\xa5\x40\xe8\xde\x58\xdc\xc5\x71\x73\x1f\xef\x3a\xfd\x0b\xf0\x0f\xeb\x5a\xde\x64\xe1\xdc\x21\x66\x05\x1b\x06\xea\x98\x6b\xbf\x76\x6c\xcc\xda\xe5\x91\x9f\xec\xc1\x09\x8e\x14\x67\xdf\x50\x56\xe5\x71\x9c\x27\x14\x55\xab\x33\xbc\x7a\x63\xa5\x53\x83\xf5\x75\x87\xfa\x18\xfe\x2b\xba\x2b\xf3\x55\x63\x4f\x99\xfa\x70\x08\x14\xa0\xf6\x63\x88\x1e\x74\xea\x0c\xea\xcd\x12\xd0\xaf\x46\xfc\xbd\x0b\xd7\x29\xab\x77\x30\xfb\x2c\x97\xfa\xf2\xb3\x9c\xd6\xe2\xf4\x8d\xcd\xe4\x93\x13\xf0\x2f\x04\x22\x9a\xa7\xb1\xe2\x6d\x82\x49\x0c\xb0\x58\x02\x4e\x41\x8a\xc4\x33\x0e\xa2\x6b\x14\x7d\x05\xd4\x7c\x8c\x47\x6f\x11\xd3\xf7\xe9\x98\xc4\xe8\x1b\x8a\x01\xcf\x50\x04\x36\x30\xb3\x65\x36\x84\xe7\x5b\x09\xe2\x25\xe4\xa8\x03\xe1\xf2\xfb\xe0\x4e\x86\x70\x47\x86\x49\x9e\xa6\x96\x8c\xb8\x3b\x72\x03\x33\x4f\x69\xf5\xac\x15\x2c\x24\x8c\xcf\x5a\x58\x5f\x7c\x65\xe5\x41\xbe\x43\x75\x9d\x4a\xcd\x51\xaf\xb6\xea\x4b\xab\x1e\xe5\x74\x5e\x54\x43\x70\x83\xd8\x1d\x80\xf1\x0d\x24\x11\x8a\x81\x64\x80\x42\x4f\x5c\x43\x01\xee\x68\x6e\xde\x72\x29\x49\x13\x84\x62\x70\x95\x0b\x80\x09\xe0\x74\x83\x24\x20\x35\xbd\x64\x25\xc8\x39\x52\xa2\xf6\x7b\x9c\x69\x12\xb5\x2e\x21\xae\xca\x5b\xdf\x03\x94\x1c\x33\x4f\x3d\xd4\xb0\x6d\xc3\x6e\x7b\xa6\x62\x87\x5e\x77\x0c\x3f\x76\xeb\x30\x7f\x7d\xb7\xd3\xde\x22\x6d\x7d\x40\x08\x33\x75\xe1\x58\x89\x56\x0a\x72\xf8\xd6\x61\xac\x8a\x85\xbb\xde\xe9\x14\x45\xdd\x36\xcf\xc6\x49\xe9\x6e\xeb\x63\xb4\x6a\x86\x44\xa1\xf1\x18\xf0\xd8\xae\xe0\xd2\x64\xfa\x7c\x35\x6a\xf9\x5a\xd7\x7a\x9d\xb1\xac\xfc\x67\xb7\xaf\xda\xb1\xa7\x8c\x39\x1d\x58\x8d\x77\x2e\x6e\x98\xbe\xea\x49\x1f\x1c\x17\xc5\xa4\x10\xbf\x76\x28\xab\x69\x45\xe5\x9b\x2c\xdb\xb4\x35\x72\x01\x35\x76\x76\xc7\xaa\x33\x6f\x30\x48\x5d\xb9\x40\xef\x95\xa3\xd3\xb1\xea\x11\x97\xdf\x12\x0c\x29\xc9\x7d\x20\xe9\x9d\xb3\x82\xd5\xae\x29\x68\x8c\xf4\x82\x6e\x32\x4d\xa5\x03\x5d\xf5\xdb\xed\x0a\xba\xfa\x9e\xd1\x19\xed\x7c\xd0\x18\xfa\x2f\x98\x31\x14\x35\xe5\x51\xb7\x6a\x52\x9c\x51\x9e\x70\x9d\x47\xdf\x35\xe0\xaa\x79\xd5\xf1\x2e\xbb\xf1\x18\xdb\x6b\xa5\xd6\x33\x12\xf9\xaf\x6a\xd4\xf8\xdb\x63\xfc\x80\xd6\x17\x5a\x15\x48\xdd\x64\x00\x56\xfd\x5e\xe0\xde\xe0\x54\x20\x56\x3e\x20\x2b\x7b\xeb\x56\x0d\xd4\x19\xe5\x07\x97\x32\x84\xd7\xe4\xbf\x91\xa3\x8a\x75\xab\x81\x6b\x8f\xf2\x82\x6b\x1e\x7f\x5b\x3d\xba\x45\xc3\xab\x7a\xbd\x60\xb5\x3f\xbf\x91\xff\xea\x56\x0d\xd3\x19\xe5\x05\xd7\x4e\x49\x55\x9d\x55\xe3\xaa\x9d\xb6\xf2\x04\xda\xda\x7b\x65\xdb\xaa\x95\x47\xf1\x82\x68\xe5\x99\x6a\x90\x65\xe3\xaa\x9d\x8b\xf2\x04\xda\x46\xd3\xb4\xad\x5a\xa9\x03\x1f\x88\x4d\x83\x69\xd9\xc9\x49\xe6\x51\x7d\x64\xd2\x54\xf4\xaa\x51\xe3\x66\x8f\xf1\x02\x7a\xce\xf0\x06\xb2\xbb\x86\x9a\xd7\xad\x1a\xac\x33\xca\x0b\xee\x47\x04\xe3\xa6\x1d\x2f\xdb\x56\x26\x83\x5b\x8d\xf0\x84\xe8\xde\x78\x6b\x88\xba\x6d\xd5\xcc\x09\x7b\x41\xbc\x40\x11\x43\xce\xf7\xe6\xba\xa5\xfc\x5e\xdf\xf4\x7a\xc2\x6a\x6e\xeb\x0b\x6b\x5b\x5f\x4c\xda\xd6\xee\xb7\x76\x0a\x96\x6a\x31\xb0\xca\x5e\x3f\x58\xf9\x95\xf9\x20\xa1\x06\xa6\x9b\xca\x72\x63\xd5\x00\x3f\x2d\x6c\xbd\x53\x90\xff\xaa\x46\x8d\xa2\x3d\xc6\x0f\x68\x03\x45\x0b\xbf\x51\xe4\xcc\x0a\xfd\x09\x88\x71\xd7\xbb\x3b\x98\xfc\x0e\x3e\x78\xcf\xc2\x0f\xdb\x19\xd7\xd1\x7b\xf8\xe8\x89\x3f\x7a\xe2\x8f\x9e\xf8\xa3\x27\xfe\xe8\x89\x3f\x7a\xe2\x8f\x9e\xf8\xa3\x27\xfe\xe8\x89\x3f\x7a\xe2\xa3\x9e\xf8\xb6\xfd\xd9\xff\x3d\x3e\x6e\xd6\x77\x75\xfe\x85\x11\xbb\x4b\x0e\xfb\x42\xd0\xf5\xf9\x26\xad\xf7\xf9\xcb\xd8\x07\x2c\x7b\x2b\x42\x3c\x05\xaf\x3f\xb5\x14\xf1\xb4\x4a\x9e\xbb\x91\xb7\x53\x59\xe2\x29\x4b\x1c\xa4\x38\xf1\xf7\xe0\xcc\x81\x0a\x15\xef\xce\xbb\xfb\x95\x2b\x1e\xdd\x5d\xdf\xa1\x68\xf1\x34\x01\xfc\xa7\x96\x2e\x9e\xc6\xa5\x07\x51\x9e\x47\x17\xcf\x38\x4f\x21\x76\xeb\x3d\x4d\x3a\x02\x9c\x2a\xc6\x87\xdd\xdb\x06\xd7\x87\xa5\x5d\x61\x85\xd5\xa0\x9e\xed\x52\x1b\x78\x1a\x77\x76\xaf\x10\x3c\xee\x65\x74\x94\xfc\x6d\x17\x58\x19\xaa\xfd\x1b\x7a\x39\x17\x4e\x09\xe0\x0e\x75\xf7\x29\xe7\xd6\x51\xcf\x57\x95\xfd\x0d\x3b\xfc\x35\xf5\xfb\x7e\x4a\xfa\x7a\x31\xd0\x2e\xec\xeb\xc1\x8b\xb0\xaa\xef\x3b\x3e\x36\x58\x78\x7d\xab\xbd\x75\xce\xf3\xf1\x09\xdb\x52\x41\xbc\x2a\xc9\x1a\x9d\x68\x7d\x09\xee\xf1\x91\xf3\x21\xca\xca\xfa\x16\x8a\x75\xd0\x1e\x29\x7f\xea\x4f\xca\x2e\x55\x63\xab\xaf\x7a\xc6\x4b\xa0\xf6\x14\x27\x30\xd5\x62\x0b\x9b\x5d\xcf\x0f\x54\x85\xd6\xb7\xae\xec\xbe\xf9\x7b\xcf\x22\xb3\x38\x01\x38\x6e\x95\x1d\xf5\x2e\x3d\xfb\xd4\xd4\x9e\x2d\x34\x9f\x76\x80\x50\x57\xcc\xec\x64\xe9\x9f\x5d\xb8\xd6\x47\x04\xdf\xb3\x7c\xad\x97\x95\xda\xf1\xb4\xd8\x63\x11\x5b\x97\x48\xff\x2a\xb6\xcf\xa7\x97\xb1\xed\xab\x76\xe2\xa2\xb1\x4b\xb1\xdb\x96\x4a\x0e\xff\x6a\x5b\x88\x07\x55\xf2\xd6\xd3\x8c\x1c\xbc\xf0\xad\xb7\xee\xfe\x45\xca\xdf\xe2\x58\x7d\xa8\x38\x56\xeb\xb6\xb7\xc0\xc7\x53\xbb\x70\xb7\xab\xd1\x1d\x40\x9b\x1a\xbd\x4b\xc9\xdc\x7b\x68\x74\xad\xce\xf7\xaa\x8a\xeb\xa3\x8b\xf7\x72\xde\xbb\x62\xb8\x4c\xb5\x34\x30\x33\x07\xf1\x2e\x08\x76\xc5\x5e\xaa\xa5\xab\x3a\xa1\x89\xbd\x46\xff\x30\x45\xc5\xfc\x49\x17\xfb\x63\xd6\x6f\x4f\x05\x6f\x72\xb1\x83\xfd\xec\xd2\xed\x11\x68\x3d\xf6\xbf\x37\x09\x3d\xa4\xb1\x26\xde\xdc\xfe\xc9\xf5\xe7\x7c\x14\xea\xa1\x56\xa1\xeb\x12\xc7\x78\x19\xba\xee\x59\x87\xaf\x43\x57\x71\xfa\xff\x6f\x35\x3a\x1f\x65\x7a\xa8\x35\xe9\x7c\x95\xc9\x2d\x4a\xb7\x47\x65\x9a\x54\x95\xee\x4f\x54\xa6\xed\x5f\xb7\xfc\xb4\x07\xd7\x86\x8a\x50\xab\x4a\x58\xdb\xa1\x02\xcb\xfd\x51\x47\x8f\x67\xda\x91\x88\xec\x3b\x64\x1e\x66\x21\x6b\x18\xc7\x0c\xf1\x2a\xd7\xde\x5d\xd7\xda\x8b\xef\x87\xab\x6e\xfd\x74\x3b\x5e\xde\xda\xb3\xe4\x9c\x77\x26\xd0\xdf\x1c\xd6\xe5\xe7\xbc\x92\x82\x1d\xc1\x4e\x7f\x11\x3a\xaf\x98\xe6\x60\xa5\xe8\x0e\xc6\xac\xba\x2c\x9d\xcf\xac\xc3\x17\xa7\x1b\xc7\xc2\xa3\x44\x9d\x4f\x79\x49\x47\x65\x55\x9a\x7e\xc2\x5f\x6f\x69\xe6\xeb\x27\xfc\xd9\xbe\xc1\x00\xd4\xa4\xf1\x55\x98\xdf\xf3\xb7\xf9\x86\xcf\x98\xbd\x3c\xc0\xe8\xe2\xc6\xf4\x3b\x8c\x43\xf3\xa4\xf7\x7e\xa3\xc5\x12\xeb\x97\xff\x0b\x00\x00\xff\xff\x16\x5f\xae\x5b\x51\x7f\x00\x00")

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

	info := bindataFileInfo{name: "templates/model.gotpl", size: 32593, mode: os.FileMode(420), modTime: time.Unix(1578956520, 0)}
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

	info := bindataFileInfo{name: "templates/relationships_registry.gotpl", size: 18028, mode: os.FileMode(420), modTime: time.Unix(1578956520, 0)}
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

