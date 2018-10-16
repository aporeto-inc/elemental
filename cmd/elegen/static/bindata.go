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

var _templatesIdentities_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\x4b\x6b\xe4\x38\x17\xdd\xfb\x57\xdc\x2f\x84\x0f\x1b\x2a\xce\x66\x98\x45\x86\x2c\x42\xd3\x0d\x05\x93\x26\x24\x30\x9b\x22\x0b\xb5\xeb\xda\x11\x23\x4b\x46\x92\xd3\x5d\x18\xfd\xf7\x41\xb2\xfc\x90\xcb\xf5\x4a\xd2\xd5\xd0\xd9\xa4\xca\xd2\x7d\x9d\xa3\x7b\x74\x5d\x15\xc9\xfe\x25\x05\x42\xd3\x40\xfa\x84\x3a\xfd\x24\x78\x4e\x8b\x5a\x12\x4d\x05\x4f\xbf\x92\x12\xc1\x98\x28\xa2\x65\x25\xa4\x86\x8b\x42\xa4\xa4\x12\x12\xb5\x48\xa9\xb8\x46\x86\x25\x72\x4d\xd8\x45\x14\xbd\x12\x09\x71\x04\x00\x40\xd7\xc8\x35\xd5\x1b\x6b\xac\xee\x49\x05\xb7\x50\x92\x6a\xa5\xb4\xa4\xbc\x78\xee\x6d\xd2\xa5\xdf\x07\x8d\x33\xb3\x7f\x4d\x73\x05\x92\xf0\x02\xdb\x64\x9e\x2a\xcc\x68\x4e\x33\x97\x8c\xb2\x89\x0c\x1b\x81\xe6\xa0\x5e\x44\xcd\xd6\x8f\x58\x50\xa5\x51\x06\xbb\x21\x85\xcb\xf4\xa1\xfe\xc6\x68\x76\x2f\xd6\x18\xda\x5e\xc1\xe5\x90\x22\xdc\xdc\x42\x6a\xf7\xb0\xf4\xf3\xf0\xf0\x6a\x64\x70\xd1\x34\x7e\xc3\x23\x2a\x6d\x97\x8d\xb9\xb8\xb1\x39\x8c\xdd\x18\xd3\x15\xb4\x08\x42\x21\x5f\x4f\xa3\x8f\x1e\x99\x28\xc0\x2c\x23\x1a\x0b\x21\xe9\x6f\x08\x9c\xa8\x65\x86\x3f\x05\x3c\xc2\x28\x51\x3f\x11\xb1\xab\x33\x42\xd6\x34\x3e\xab\x4b\xba\x80\x4b\x57\xd9\xc8\xe8\xae\xad\x14\x42\x8c\xbb\x7d\x1f\x07\xec\x9e\x83\xca\xd7\xf8\x63\x06\xeb\xd5\xf3\xea\xb9\xfd\x78\x76\x8c\x2d\x02\x93\xfe\xec\xa1\xa0\x39\x30\xe4\xdd\xf2\xb2\x4d\x1e\x8c\x99\x4d\x37\x4c\xd9\x11\x90\x89\xb2\x12\x35\x5f\x8f\x38\x18\x9c\x04\x86\x9d\xbf\x66\xec\xa2\x37\x37\xc6\x65\x69\xff\x2f\x06\x60\xc1\xec\x25\xc5\x2c\x9a\x06\x90\x29\x5b\x0e\xa7\x6c\xf1\x46\xde\x92\x28\xba\xbe\x06\x97\xfc\x3f\x28\x95\x05\x53\xa2\xae\x25\x57\xa0\x5f\x10\xb2\x5a\x4a\xe4\x1a\x5e\xfd\x9a\xc8\xdd\xe3\xd2\x15\x1b\xe5\x35\xcf\x02\xdb\x38\x81\x9c\x09\xa2\xff\xfc\x03\x1a\xef\xa7\xbf\x3c\xee\x1e\x96\x4b\x9e\x8b\xb4\x0b\x63\x2b\x8c\x22\xbd\xa9\xbc\xbb\x7b\xc2\x49\x81\x12\x94\x96\x75\xa6\x1b\x13\x39\xf7\x71\x1e\xac\x26\xd0\x9d\xd8\x2f\x52\x94\x96\xcd\x98\x5b\x4a\x5b\x78\x13\x98\xed\x6a\x57\xaa\xcf\x66\x7a\xff\xac\xac\xf9\x73\x74\x4c\xb4\x4f\xad\xfc\x6e\x62\xaf\xc3\x9b\xd3\xa3\x06\x0a\xbe\xea\xfc\x1c\x17\xde\xb5\x77\xdc\x36\xf3\xd1\x81\x07\xf5\x5b\xb9\x8f\x47\x86\xe2\x9b\x98\xf0\xa1\xbe\x98\xce\x44\x4a\xba\x50\x34\x07\x0a\xb7\x90\xa7\x5b\xd4\x10\xbe\x49\xfe\x82\xff\xd1\x74\xa9\x3e\x97\x95\xde\xc4\xc9\xa8\xa1\x3a\x68\x02\x01\x99\x73\xd5\xe3\x7e\xb2\x3b\xff\x2c\x74\xe7\x71\xe4\x9b\xe4\x00\x16\x39\x25\xdf\x18\xc6\x1d\x77\xb3\x10\x4c\x9f\xb5\x36\x1d\x32\xea\x3b\xd5\xd9\x4b\xcf\xbe\xcf\xb6\x57\xf1\x3d\xba\xf7\x66\xcd\xcb\x88\x6a\xe7\xb5\xad\x8b\x64\x10\xfb\x9b\x29\x68\x5f\xf1\xfb\x0e\x93\x38\x89\x66\x64\x63\xf2\x75\x8d\x39\xa9\x99\xde\x72\xcb\x29\xf3\x6c\xec\x02\xfa\xa9\x22\x52\xe1\x5b\xe0\xde\xb6\xfc\x85\xa0\x7b\x43\x2e\x74\x7f\x0f\xa8\x47\x21\xf4\x7b\x49\x69\x8b\x7c\x0f\x35\x1f\xc6\x94\xbf\xd8\xf6\xd3\x13\x5c\x9c\x81\xfe\xf5\x83\xc1\xaa\x73\xe0\x5e\x1f\x0e\xc9\x51\xcb\xac\xed\xda\x27\xe7\x36\x50\xa5\xfd\xbd\x37\xe9\x7d\x7f\xba\x26\x52\xd0\xea\x5c\x72\x9c\x12\x1c\x28\x7e\x3e\x1d\xf5\x61\xc7\x72\xe7\xe9\x3a\xaf\x56\xfc\x7f\xce\xe0\x81\xd5\x92\x30\x30\xe6\x6f\xaa\xec\xd5\x7d\xc6\x83\xb9\x2d\x04\x47\xf3\x34\x63\xfa\xdb\xb1\xb5\x5b\x42\x7e\x21\x67\x01\xe4\x27\x75\xb7\xda\xdb\xde\xea\xe4\xfe\x7e\x44\xd6\xd2\xf7\x42\x2b\x15\x8f\xa3\x06\x2b\x2d\x4f\x72\x3a\x5d\xc9\xb9\x3d\x36\xd6\x2b\x91\x50\xfa\x79\xf6\x36\x08\x69\xe7\x5a\x3b\x70\xfb\xc5\xf1\xac\xed\xb6\x8d\x32\xb8\x1f\x99\x75\xb3\x76\xfb\x2d\x48\x74\xbc\x6d\x98\xba\xbb\xe8\x26\x72\xf1\xee\x18\xf3\xc0\x50\x54\x7d\x54\xc2\x18\xe0\x0f\xaa\xb4\x55\x6c\xda\xaf\xfb\x60\x81\x4d\x6c\xd5\xfd\xe0\xb8\x39\xb7\xe5\xfc\x6f\xd5\x7b\x5b\xe4\xb4\x9f\x10\x8c\x47\xcf\x4d\xd1\x5f\x84\xec\xeb\x1e\x43\x68\xc9\xf3\x83\x36\xe4\x42\xba\xef\x05\x7d\xc5\x61\xee\xef\x11\x9d\xfa\x39\x74\x9f\x86\xb7\xe9\x2e\x55\x3a\x02\xd6\xf3\xca\xce\xf8\x7d\xf7\x5d\x3f\x59\xec\x7f\x05\x3e\x42\xa4\xc2\x37\x01\x2b\x4d\x26\xfa\x2f\x00\x00\xff\xff\x51\x27\x8c\xfd\xd7\x14\x00\x00")

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

	info := bindataFileInfo{name: "templates/identities_registry.gotpl", size: 5335, mode: os.FileMode(420), modTime: time.Unix(1539649739, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesModelGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\x59\x6f\xdc\x36\xb7\xef\xf3\x2b\xd8\x81\x9b\x4a\xe9\x58\xce\xf3\xa4\x2e\x90\x9b\xa5\x30\x6e\xd2\x1a\x71\x90\x3e\x18\x46\x4d\x4b\xd4\x98\x8d\x86\x54\x48\xca\xcb\x1d\xe8\xbf\x5f\x70\x93\x48\x2d\x33\x94\x97\x5c\x5f\x7c\xce\x43\xeb\x39\x24\xcf\xce\xc3\x73\x48\x8a\x25\x4c\xbf\xc1\x15\x02\x9b\x0d\x48\x4e\x90\x48\xde\x52\x92\xe3\x55\xc5\xa0\xc0\x94\x24\x7f\xc2\x35\x02\x75\x3d\x9b\xe1\x75\x49\x99\x00\xd1\x0c\x00\x00\xe6\xf9\x5a\xcc\xf5\x5f\xfc\x96\xa4\xf3\x99\xfe\x7b\x45\x13\x58\x52\x86\x04\x4d\x30\x3d\x40\x05\x5a\x23\x22\x60\xa1\x7b\x6e\x36\x80\x41\xb2\x42\x20\x39\x29\x51\x9a\x7c\xb9\x2d\xd1\x31\xa3\x57\x38\x43\x8c\x83\xfd\xba\xd6\x28\x24\x17\xa0\xae\x9b\x21\x88\x64\xaa\x71\x08\xc5\x57\x58\xe0\x4c\xb1\x39\x01\x51\x3c\x9b\x6d\x36\x60\xaf\x80\x02\x71\xf1\x15\x31\x8e\x29\x01\xcb\x43\x83\xf2\xa3\x02\xbf\x11\x82\xe1\x8b\x4a\x20\x6e\x3b\x48\x0e\x1a\xea\x7b\x78\x01\xf6\x10\xa9\xd6\x72\xdc\x45\x85\x8b\xec\x3d\xa9\xd6\x5c\xa3\xe8\xa2\xae\xeb\xd9\xc1\x81\x64\x40\x8d\x50\x62\x83\xba\x06\x0c\x95\x0c\x71\x44\x04\x07\xe2\x12\x81\x92\x72\x8e\x2f\x0a\x04\xae\x60\x51\x21\x0e\x72\xca\x00\xb4\x5c\x28\x61\xf4\xf0\x86\x33\x63\x97\x79\x32\x13\x12\x63\x0f\x3f\x17\x0c\x93\xd5\x6c\x96\x52\xc2\xad\xd5\x36\x9b\x7d\x2b\x01\x81\x6b\xb4\x00\x7b\x8a\x9a\x94\x42\x0f\xfe\xaa\x89\x1b\x15\x1a\xb6\x89\xa6\xd4\xe5\x58\x0f\x95\x1d\xf4\x5f\x75\x9d\x58\x55\x37\x43\x7a\x5c\x1d\x6a\x51\xec\x08\xcf\x38\xca\x36\xed\xdf\x33\xc9\x2d\xce\x01\xa1\xc2\xd8\xe6\x13\xcd\x50\x91\xbc\x43\x02\xa6\x97\x28\x6b\x15\xeb\xb6\xbe\x27\x02\x8b\x5b\xa3\x9c\xa3\x0c\xa9\x9f\x5d\xd6\x1b\x38\xcd\xd5\x6f\x7a\xf1\x2f\x4a\x45\x32\xbb\x82\x2c\x0c\xdf\x21\x68\x7c\x3b\x69\x80\x1b\x25\x8c\xec\xba\x04\x8d\x07\x3a\xa8\x3e\x23\x2e\x64\x6b\x5d\xcf\x17\xaa\xeb\x5b\x28\xd0\x8a\xb2\xdb\xe5\x50\x57\x5a\xb1\xb4\x31\xb2\xee\x7f\xac\x27\xea\xb2\x8f\xda\xb4\xb4\x3d\x19\xbe\x82\x42\xf6\xec\x76\xd4\x0d\x75\xbd\x98\xd5\x13\x75\xbd\xd9\x0c\xf5\x38\xe2\x9f\x29\x15\xbb\x6c\x71\x5c\x54\x0c\x16\xa0\xae\x3f\x62\x2e\x5c\x6b\x40\x50\x48\x08\xcd\x03\xc6\x36\x8e\x1e\x42\xe3\xf4\xec\xe5\x68\x4f\x29\xf0\xc1\x01\x70\xbc\x43\x54\x8c\x68\xd7\xc0\x83\xae\xc1\x01\x26\xea\xa7\xe4\x36\x99\xe5\x15\x49\x41\x44\x03\x79\x89\x1b\x4a\x51\x3c\xec\x37\xca\x66\x9a\x8b\x71\x9c\xad\xfb\xcd\x34\xff\x6f\x69\xd9\xf2\x0e\x41\x49\x31\x11\x88\x01\x41\x01\x04\xa9\x6c\x93\x0c\x87\xb1\x38\x5d\x24\x49\x7c\x40\x9c\x1c\xc3\x8b\x02\x71\x2b\x93\x62\x63\x79\x08\x60\x59\x22\x92\x45\x61\xc8\x37\xf5\x02\xd0\x24\x49\x62\x57\x2d\x2f\x24\x2a\x23\xf8\x1b\x85\xcd\x20\xe5\x9e\x99\x04\x55\x3f\x21\x20\xe8\x5a\x53\x37\x76\x7c\x2c\x3d\x68\x5e\x22\x4b\x3f\x49\x92\x61\x95\xec\x54\x15\xad\xc4\x3d\x35\x25\x97\x8c\x7f\x16\x52\x15\x12\x91\x8e\xf3\x96\x2f\x1d\x9b\x2c\x9d\x86\x0c\xad\x84\x1a\x90\x44\xdb\x66\x4b\xac\xf1\xd7\x9e\x9f\xd2\x4a\x18\x73\xa8\xf9\x96\x52\x72\x85\x98\x70\xad\xa1\x3c\x91\x8c\xc9\x7d\x37\x75\xcb\xff\x8e\xbb\x9d\xe2\xc4\xd7\xe7\x1a\x7e\x43\xd1\x96\xee\x0b\x50\x20\x12\xd1\xb8\x55\x21\x96\xc3\x5e\xbd\x06\x18\xfc\x66\xda\x5e\x03\xfc\xeb\xaf\xbe\x0a\x4f\xf1\x19\x38\x04\xf4\x14\x9f\x6d\x55\xcd\x3b\x94\xc3\xaa\x10\x7f\xb1\x0c\x31\x2f\xcc\x64\xba\x01\x50\xd9\x82\xc9\x0a\xe4\x18\x15\x19\xb7\xde\x9a\x52\x22\x10\xb9\x83\x7e\x5c\x82\x51\x0c\x4e\xcf\x74\x1a\xd0\x89\x31\x16\xdc\x8a\xd4\x49\xac\xfe\x32\x6c\xb5\x79\x50\x37\xab\xb1\x69\x56\xb3\xd0\xf9\x4b\x95\xbf\xac\x6b\x15\x69\x95\x7c\xa1\x1f\xaa\xa2\xf0\x94\x11\x18\xd2\x8d\x8b\xa1\x4c\x3a\xd6\x49\x09\x19\x47\x8f\x35\xab\xbf\x50\x8d\x3f\x32\x66\x49\x92\x44\x6b\xec\x49\xb9\x5e\xd2\x61\xb3\x89\x04\x63\xee\x68\x8d\xe7\x2a\xff\xca\xc0\xee\xeb\x79\x06\x77\x14\x03\x4c\xc4\xd0\x9a\x86\x44\xf2\xe6\xf8\xe8\x88\xe4\x34\x71\x52\x63\x9d\x56\x1b\x3f\x91\x59\x88\xcd\x48\x76\xa4\x76\xdd\x94\x6e\x2d\xbb\x48\x21\x60\x77\x98\x4d\xbb\x76\x67\x10\x26\x6d\xae\x52\xc9\x7f\x67\x46\x6c\x9f\x09\x9b\x8d\x4a\xd6\xbf\xd0\x0f\xd2\x10\x60\x4f\x4a\x0b\x72\x58\x70\xa4\x6a\x10\x4f\x30\xa9\x17\x45\xd9\x22\x90\xfa\x3a\xff\x97\x53\xb2\x9c\xef\xcf\xc1\x85\xfa\xe3\x1f\x25\x8f\xb1\xcd\xfc\x5c\x8f\x92\x95\x56\xf2\xa9\x12\xe8\xa6\xd7\x7f\x7f\x7e\x6e\x6c\xfc\x27\xba\xde\xa1\x37\x9b\x30\xc8\x25\x72\x3c\xec\x4b\x5e\x95\x0f\xec\x40\x18\xc5\xdb\x91\x74\x5c\xe1\xc5\xb6\xbe\xad\x9f\xbb\x0a\x5a\x6e\xf1\x9f\xc5\xcc\x09\x37\xfb\x6e\x79\x26\xed\x21\xa7\x13\xa7\xcc\x29\xe7\x40\xb4\xc3\x9c\xb1\x17\xd9\x4c\x5a\xcc\x2f\x69\x55\x64\x7f\x33\x2c\xd0\x11\xc1\x02\xc3\x02\xff\x0f\x62\xd2\xcc\xaa\xde\x93\xa4\x74\xfc\xeb\xb8\xc6\x5e\x72\x5c\x5d\x14\x38\x95\xd2\x80\x0e\xda\x3d\x4c\xb0\x0a\x13\xd7\x03\x68\x91\xf0\x90\x77\xc7\xe2\xdc\x0c\xf7\xe0\x43\xb0\x7d\x37\x04\x87\x81\xcc\x9c\x0c\xa9\xbd\x06\x33\xe8\xb1\xe2\xca\x06\x94\xad\xde\xf2\x40\xb9\x32\xf0\x92\x65\x77\xf6\x4d\x11\x2c\xc7\x9d\x55\x5b\x97\xad\x9e\x5c\xbf\x70\x50\x11\xfc\xbd\xb2\x95\x83\x1c\x33\x51\x56\x39\x24\x8a\x81\xbf\x52\xeb\x6a\x4b\x8f\x75\xb8\xb1\xce\x69\x43\x7b\xd2\x10\x68\x3b\x25\x6f\xed\x42\x69\xe7\x71\x63\x65\x19\x90\x3a\x28\xe6\xf3\x59\xc7\x09\xee\xa2\xb0\x13\x24\x1c\x2e\x39\x12\x8f\xa3\x30\x8f\x4c\x84\x33\x60\x57\xe5\x30\xad\x85\xa9\x0b\x1c\x02\x9c\x3d\x84\x52\x86\x96\xdb\x4b\xc8\xb2\x94\x66\x28\xeb\x2e\xbc\x2a\xe0\x07\x2b\xe2\xce\xab\xed\x74\x29\x46\x73\x58\x5b\xb9\x8f\xe4\xb2\x41\xb2\x28\x61\x9e\x76\xd2\x3a\x51\x63\xae\x0b\xda\x56\x9e\x32\x5c\x8a\x76\x33\xf0\x1d\x4d\xfd\x6a\x80\xa6\x95\x8a\x74\xaa\x8f\xcc\x04\xdb\xf9\x12\xea\x11\xef\x68\x3a\x10\x44\x14\xbb\xe8\xfb\x28\x3f\xf3\x53\x42\x33\x9a\x9e\xcd\xbb\x61\x41\x81\x4f\xf4\x1e\xa2\x17\x3e\xfc\x6e\xe7\x52\x5f\x3c\xfd\x2f\x98\x7e\x13\x38\xfd\xc6\xb7\x08\x7e\x3e\x34\xa5\x86\x32\xbf\x00\x25\x87\x86\x0b\xc5\x7f\x4f\x2d\x86\xf7\x7c\x2d\x92\x93\x92\x61\x22\xf2\x68\xfe\xdb\xcf\x7c\xf9\x33\xff\x7d\x2e\x0b\xea\x76\x09\x52\x0e\xd2\x82\x74\xa4\x8e\x7b\x4e\x11\x98\x29\x36\xde\xa1\xf3\x89\x3f\x10\x41\x0c\x0a\xf4\x07\x12\x02\x31\x90\xf4\xd2\x85\x83\x03\xf0\x07\x12\x52\xc4\x5e\x88\xea\x96\x4f\xbd\x0e\x26\xaa\x30\x94\x22\x7c\xd5\x8d\xb0\x7b\x5b\x74\x36\x42\x31\x8a\x7d\x3a\x76\x43\xd7\x57\xa9\x0e\xb0\xbd\xf5\xa7\x93\xe4\xf7\x55\x70\xb2\x45\x05\x27\x23\x2a\x68\x16\x99\x15\xbe\x42\xe4\x81\xb5\x30\x42\x34\x72\x82\xc6\xa0\x3e\x9a\x95\x68\x58\x13\xe0\x10\x38\x18\x3a\x33\xa0\x3b\x01\x20\xc9\x40\x34\x36\x0b\xe2\x7e\x93\xde\x8a\x8d\x8d\xda\x6c\x6d\xe8\xb9\x0a\xd7\xa0\x7b\x2d\x3d\x41\xb5\xb1\xee\xe2\x96\xbc\x56\x31\x38\x57\xa5\xad\x1e\x1c\x83\xc3\x43\xf0\xca\xa9\x6f\x6d\x95\xb0\xa3\xbc\xd7\x9e\x07\x9c\x7f\x6d\xf6\x1f\x30\x0b\xfd\x81\x7d\x33\x2d\xc1\x0b\x69\x3f\x1f\x2e\x8b\x8d\x2e\x45\x44\x32\x07\x5d\xed\x16\xdf\xbc\x94\xc9\x7d\x90\x20\xb5\xbb\x7d\x97\xb7\x9b\x77\x46\xc1\xad\xa0\xfc\x1a\x8b\xf4\x12\xe4\x0f\x25\x7b\x0a\x39\xf2\x17\xc2\xa5\xd7\xae\xe5\x18\xd0\x04\x38\x04\x2f\xa2\x21\x15\xc5\x53\x54\x64\x8c\xcd\x4b\x53\xba\x1e\x43\x29\x1d\x2c\xcb\x42\xef\x5d\x13\x4a\x00\xc1\x45\x9b\x45\x42\xf0\x32\x40\x9f\x66\x0f\x78\x62\xdd\xa1\x88\x47\x66\x82\x6c\x73\xe4\xd8\x5d\x5d\x55\x79\xc6\xdb\xe3\xc3\x9d\xaa\xb7\xe5\x9b\x1e\x67\x80\x38\x07\x3f\x69\xca\xee\xda\x73\xc4\xdf\x7f\xaf\x60\x11\xb9\x0b\x52\xec\x98\xbe\x84\x04\xa7\xd1\x3c\x85\x44\x46\x82\x52\x29\x2f\x67\x74\x0d\x20\xd0\x52\x5c\x63\x71\x09\x32\x9c\xe7\x88\x21\x22\x9a\xa3\x8d\xb9\xb7\x49\xc4\xa9\xaa\x92\x35\xf5\x28\x40\xbf\xf1\xac\x7b\x96\xd8\x95\x85\x8f\x04\xbf\x9f\x0e\x95\x39\x9d\xdd\xac\xb1\x20\xf9\x92\x0f\x39\xd7\xac\xf1\x20\xcf\xb1\xda\x1f\xbd\x1d\x24\x73\x40\xac\x0a\x11\x9c\x21\x1d\x04\xd3\x8a\x69\x7d\x90\x9c\xb2\xb5\xce\xb9\xb8\xa0\x0c\x65\x32\x99\xd6\xce\xa3\xb7\x81\x2a\x86\xc2\x33\x72\x43\x4a\xd6\xad\x8c\x51\x66\xbd\x44\xfd\x50\x1e\xd2\x3a\xd5\x7b\x05\xdb\xd8\x4c\xea\x7b\x85\x19\xca\xde\x6f\xeb\x38\x74\xfa\x1d\xe0\x6b\x4d\x1e\xf5\x85\x41\xc2\xb1\x94\xda\x6b\x4b\xde\xdf\x94\x94\xa3\x76\x63\xca\x80\x3f\x1b\x9e\xfc\xde\x32\x8f\x54\xeb\xdc\x5c\x07\xfc\xb9\x63\x73\xc4\x98\xcf\xba\xd5\x87\x45\x65\x72\x31\x3f\xed\x1e\xf1\x80\xf8\xb5\xc2\xd7\xf3\x97\x8e\xaa\x9a\x63\x0c\x1f\xbe\x90\x83\xe3\xd6\x59\x9c\x9c\x7e\x50\x18\x81\xd7\x68\x92\x28\x5f\xf0\x1a\x3d\x45\x41\xd0\x8d\x40\x8c\xc0\x62\x92\x30\xef\xcd\xa0\xa7\x28\x10\x26\x02\xad\x10\x9b\x24\xcf\x11\x11\x4f\x51\x94\xbc\xa0\x50\x4c\x12\xe4\x83\x1c\xf1\x44\x44\xd9\x26\x19\x43\xf9\x80\x5c\xc3\x8c\x26\x6d\x94\x1c\x61\x19\x75\x58\x45\x13\xb4\x4d\x19\x88\x3c\xbe\x3e\x62\x2e\xe6\x71\x07\xf8\x09\x96\xf3\xd8\x32\x6c\x92\x2e\x5e\x5d\x38\x67\xa6\xc3\xab\x52\xcb\x63\x2b\x25\xaf\x2e\x76\x8b\x14\x24\x96\x9b\x1c\x8d\x0a\x98\xbc\x29\x0a\x7a\x8d\xb2\xb7\x97\x14\xa7\x88\x87\x78\x93\x0e\xbb\x47\x44\x1d\x9c\x76\x9c\x49\x47\xf1\x68\xc4\xa7\x16\xed\x8e\x8b\x1c\xf7\x2f\xc5\xa4\xc7\xc0\xf9\x7c\x01\xe6\xe7\x12\x5b\xbd\x50\x99\xf4\x9b\x4a\xd0\x95\x29\xea\xb2\x2d\x9e\x79\x1f\x33\xb7\x4c\x40\x16\xa4\x83\x63\x28\xab\x4b\x12\x36\x97\x16\x6a\x4f\xa3\x4b\xe3\x5c\x8b\xe7\xac\x8b\x8f\x21\xd9\x4a\x80\xe4\x13\xbc\xf9\x88\xc8\x4a\x5c\x82\x57\x21\xb2\x7d\x82\x37\x78\x5d\xad\xf5\x90\x50\x09\x25\xb4\xa5\x23\x21\xea\xac\xea\xd1\x44\xc2\x64\x92\x48\x98\xdc\x51\xa4\x86\xce\xe3\x8b\x04\x6f\xd4\x5d\x39\xf0\x2a\x79\x35\x96\x24\x85\x87\x7c\x63\xc4\x09\x11\xbf\xb1\xe1\x57\x73\x93\xee\xe1\xe4\x35\x9b\x7c\xa1\x4c\x07\xaf\xb7\x0b\x99\x5d\x47\x1d\xb6\xe3\x87\xb6\xd3\x2e\x47\x7c\x48\xab\x69\x3f\x9d\x6e\x35\xcb\xc5\x23\x58\x2d\x90\xe7\xbb\x18\xad\xe5\xfa\x07\x19\xad\x39\xc2\x4d\xdc\xe2\xd9\x39\xe2\x55\xd7\x6f\x07\x2f\x01\xf7\x34\x21\xd1\xad\xad\xb4\x4a\x74\xe7\xd4\xb6\x95\x5f\x03\x43\xd3\xab\x60\x39\xf7\x47\xf2\xa9\x00\x2d\x78\x02\x5f\x35\xa2\x2a\xce\x9a\x5d\x07\x5d\x90\xb6\x8a\x70\x6f\xf0\xbe\xad\xb8\xa0\x6b\x7b\x25\xab\xc5\x90\xb4\x7b\x18\x6b\x58\x96\x98\xac\xd4\x35\x60\x75\x72\xd4\x62\xfa\xa4\x9b\x12\xf3\x7f\x30\x6f\xef\x74\xf7\xd8\xe9\xec\x70\x58\xac\xc3\xb6\x30\x78\xad\x45\xe8\xc3\xe9\x78\x48\xe5\x66\xd3\xd1\xcf\x7b\x63\xf0\xbb\xb3\xf7\x68\x36\xa3\xfc\x2e\xee\x46\x89\xc1\x81\x06\xc6\x36\xa3\x51\x6f\x94\x3d\x4e\xc1\x45\xf8\xb1\xba\x6c\xc1\x39\x4e\x95\x6a\x3f\x50\xd6\x14\xfa\xde\x5e\x6e\x03\xf5\xba\x37\xa7\x47\x7a\x57\xbc\xbd\x42\xae\xae\x64\x7f\x43\xb7\x76\x43\x63\xd7\xd1\xc9\x18\x0f\x91\x42\xd4\xdf\xf1\x1d\x61\xa7\xdd\xf5\xbd\x5a\x00\xfa\xcd\xd8\x7f\x94\x70\xbb\xa7\xf1\x09\x96\xa7\x92\xd4\xd9\x6b\x39\xac\xa7\xe9\x2b\x57\xc9\x07\x07\xe0\x6f\x04\x52\x5a\x15\x99\xd2\x6d\x8e\x49\x06\xb0\x58\x00\x4e\x41\x81\xc4\x2f\x1c\xa4\x97\x28\xfd\x06\xa8\xb9\xc4\x4b\xaf\x11\xd3\x9b\x9f\x98\x64\xe8\x06\x65\x80\x97\x28\x05\x6b\x58\xba\x36\xdb\xc6\xe7\x47\x89\xe2\x2d\xe4\x68\x80\x61\x7b\x4b\x75\x50\x21\xdc\xb3\x61\x5e\x15\x85\x63\x23\xee\xf7\x5c\xc3\x32\xd0\x5a\x23\xb4\xa2\x58\xe2\x38\xd5\xc6\x3a\x0b\xb5\x55\x80\xf8\x9e\xd4\xb3\xdd\xb7\xb5\xbc\xfe\xbd\xab\x5b\xb0\x54\x9b\xbb\x8d\x1a\xa4\x0b\x6f\xc3\xb6\xfb\xa2\xbe\x4f\xef\x70\x8a\x16\x36\xdd\xc8\x3b\x69\xb3\xcd\xb9\x29\xd4\x8c\x90\x2c\x74\x4e\xb6\xbc\x8f\x54\xba\xcb\xcd\x7c\xb9\x73\x5a\xb9\x67\xe0\xfb\xa3\x35\x91\xfc\xe7\xc2\x97\xfd\xd2\x46\x56\x36\x1e\xae\xce\xdd\x20\xbf\xde\x5b\x8e\xd4\xa1\xfb\x75\x3d\xa9\x56\x6c\x13\x96\x66\x58\xdd\xac\x7c\x8b\xbe\x6c\x9d\xa2\xb2\xe5\xce\x6d\x58\x0e\x16\xa0\x5b\xa5\xb3\x04\x06\x2f\xcc\xc8\x7f\x5e\xc3\x72\xc4\x5c\x61\x24\x18\x52\x96\xfb\x8b\x14\xb7\x1e\x05\x07\xae\x25\xe8\xf4\x0c\xc2\xee\xdd\xd1\x70\xda\x5d\xb8\xc6\xde\xe9\x39\x05\xbb\x4d\xff\xba\xd8\x15\x5c\x61\x57\x57\xd9\xbc\xde\xde\x5d\xb6\x24\x9c\x60\xc9\x50\xda\xb5\x76\x0b\xb5\xa2\x38\xbd\x02\xf1\x7a\x37\x31\x5a\xc4\x0d\x78\x39\x70\xa3\xa1\x73\x8d\x21\x88\x92\xb3\xaf\x6e\x9b\x0c\x48\xf3\xde\xb6\x07\xa1\xfb\x80\x0b\x81\x98\x3a\x46\x75\x5a\x5b\xa8\x46\xea\xf5\x0a\xc3\x4b\x19\xc2\x2b\xf2\xdf\xc8\xf3\xc9\x16\x6a\xf0\xba\xbd\x82\xf0\x9a\x2b\x0d\x4e\x8b\x86\x68\x7c\x4d\x6b\x10\x2e\xe7\x1e\x97\xd3\xda\x42\x35\x4e\xaf\x57\x10\x5e\x77\x03\xa4\x69\x6c\x80\xcb\xfe\x26\x49\x20\xd2\xde\x34\xb1\xb0\x65\xaf\x66\x0f\xc2\xe8\x6c\x6a\xb4\x28\x2d\x70\xd9\xdf\xf8\x08\x44\xda\x67\xd3\xc0\x96\xbd\x22\x35\x04\x63\x37\x72\x3a\x01\x73\x52\x9c\x54\x81\xa9\xeb\xe8\x0d\x50\xf3\xe6\xf6\x09\x42\x7a\xcc\xf0\x1a\xb2\xdb\x8e\x9b\xb7\x50\x8d\xd6\xeb\x15\x84\xf7\x33\x82\x59\x37\xa0\x5b\xd8\xd2\x6c\x18\x36\x3d\x02\x31\xfa\x07\x6f\x1a\xa3\x86\x2d\xbb\x5b\x90\x41\x18\x4f\x7a\x53\xf1\xc4\x99\x8a\x27\x93\xa6\xe2\x89\x3e\x25\x75\x71\x29\x88\xc1\x65\x5b\xc3\x70\x55\x17\xe6\xce\x4c\x8b\x4c\x83\xec\xe7\x90\x4d\x87\x30\xcf\xe9\x1d\x71\xca\x7f\x0d\x50\xb3\xe8\xf6\x09\x43\xda\x61\xd1\xe1\x6f\x27\x73\x86\xc2\x78\x69\xba\x3b\x6f\x1e\x2e\x33\x7e\x40\x02\x3d\x42\xf8\x69\x67\xd2\xba\xae\x4b\x9e\xd3\xe8\xe7\x34\xfa\x39\x8d\x7e\x4e\xa3\x9f\xd3\x68\xf0\x9c\x46\x3f\xa7\xd1\xcf\x69\xb4\xc5\xf8\x9c\x46\x07\x20\xfc\xff\x98\x46\x6f\xfa\x5f\x8f\xdc\xe3\xf2\xbc\x3e\x82\x09\xff\xbc\x7c\xf8\x3d\x93\x50\x0c\xfa\x9b\xe4\x49\xf4\x4e\xcf\x76\x5d\xd3\x7d\xb0\x17\x4e\xa6\xf0\xf5\x7f\xfa\xce\xc9\xdd\xde\x03\x98\x26\xde\x9d\xde\x3c\x99\x42\xe2\x51\x5e\x3e\xf9\x11\x9a\x79\xa4\x57\x50\xee\xae\xbb\xfb\xbd\x85\xb2\x73\x76\xfd\x80\x17\x51\xa6\x19\xe0\x3f\xf5\x5d\x94\x69\x5a\x7a\xda\x1f\x9a\x8e\xbc\x8e\x32\x69\x69\xf0\xde\x48\x79\xdc\x39\xaf\x59\x7d\x5a\x4e\x97\x58\xa6\xb6\x7a\xdf\x5d\x9e\x41\x99\xa6\x9b\xbb\x3f\x86\xb2\x3b\xf7\x18\x78\xfc\xa4\xff\x59\xdf\xb6\x57\x50\x92\xa0\x94\xc3\x7b\x0c\x65\x60\x12\xec\x76\xfe\xc1\x47\x51\x04\xab\x90\x7f\xa1\xeb\xc7\xbd\x8b\x12\xa4\x58\xf7\x75\x94\x00\x1d\x25\xcd\x23\x29\xbb\xfb\x46\x71\xd0\xf7\x6b\x1b\x6f\xf5\xdf\x3d\x60\x63\x1d\x27\xe8\x39\x0e\xe3\x2b\xbd\xaf\xe3\x02\x3e\xfc\x7a\x8c\xb7\x39\x42\x5f\xdb\xf0\xd8\xde\xf1\x86\x44\xb8\x28\x77\x79\x7a\x03\xe7\xc1\xef\x48\xf4\xbf\xc7\x70\x9f\xdc\xa8\x5d\x75\xbd\x7c\xa4\xa7\x3c\x42\x1f\xe7\x78\x68\xfd\xfe\xa8\x97\x3a\x5e\x0c\x3d\xd5\x71\xef\xc7\x37\x42\x24\xbc\x57\x8c\xef\xac\xf3\x50\xdf\x04\xeb\x70\x65\x6c\x72\x17\xe6\x06\x16\x67\x09\x18\xfa\x20\xdb\xac\xcd\x3b\x9f\x79\x6a\xd4\x3c\xe9\x30\xc7\xcc\x96\xde\x07\xc4\xfd\x6f\x41\x2b\x31\xfa\x35\xe8\xf6\x8f\x41\x9d\x0f\x40\xfb\xeb\x7d\x9b\x66\xfd\x6f\x00\x00\x00\xff\xff\xbb\x10\x7c\xfc\xe7\x58\x00\x00")

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

	info := bindataFileInfo{name: "templates/model.gotpl", size: 22759, mode: os.FileMode(420), modTime: time.Unix(1539653963, 0)}
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

