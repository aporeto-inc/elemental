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

	info := bindataFileInfo{name: "templates/README.md", size: 77, mode: os.FileMode(420), modTime: time.Unix(1515709395, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesIdentities_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x41\x6b\xe3\x3c\x10\xbd\xfb\x57\x0c\xa1\x7c\x38\xd0\x38\x97\x8f\x3d\x14\x7a\x28\x65\x0b\x81\x6d\x29\x2d\xec\xa5\xf4\xa0\x3a\x63\x77\x58\x59\x0a\xd2\xa4\xdd\x22\xfc\xdf\x17\xcb\x96\xe2\xa4\x4e\xd6\xdd\x86\xf6\x54\xc9\x33\x6f\x9e\xe6\xbd\x91\xb2\x12\xf9\x2f\x51\x22\x38\x07\xd9\x3d\x72\x76\xa9\x55\x41\xe5\xda\x08\x26\xad\xb2\x1b\x51\x21\xd4\x75\x92\x50\xb5\xd2\x86\x61\x52\x12\x3f\xaf\x9f\xb2\x5c\x57\x73\xb1\xd2\x06\x59\xcf\x48\xe5\x73\x94\x58\xa1\x62\x21\x27\x49\x52\xac\x55\x0e\xa4\x88\xd3\x29\xb8\x04\x00\x1a\x6c\x23\x54\x89\x6d\x85\xfb\x15\xe6\x54\x50\xee\x2b\xd8\x06\xbd\x89\x89\x08\xd9\x1d\x96\x64\x19\xcd\x62\x89\x8a\x89\xdf\x52\xe7\xb2\x6b\xbd\x44\x99\x7d\xf7\xeb\x86\x53\x5d\x87\xaf\x53\xe7\x00\xd5\xb2\x81\xa9\x93\x64\x3e\x07\x1f\xfa\x13\x8d\x25\xad\xc0\x20\xaf\x8d\xb2\xc0\xcf\x08\xf9\xda\x18\x54\x0c\x2f\xdd\x37\x5d\xf8\xed\xca\x43\xb7\xac\xfb\xb9\xe9\x14\x0a\xa9\x05\x7f\xfb\x1f\x5c\x87\x13\x9b\x74\x71\xbb\x58\xa8\x42\x67\xa1\x4c\x5d\x43\x5b\xbc\x65\x55\x90\x78\x92\x78\xa5\xe3\x11\x22\x0f\x01\x0a\x5f\x81\x94\x65\xa1\x72\x0c\x14\xfa\x59\x50\x68\xe3\x37\x4b\x7a\x41\x05\x14\x10\x94\xa8\xb0\x63\xb9\xa7\x48\x1a\x63\x2d\x1b\x52\xe5\xb4\xd7\xd3\x7e\x8a\x4b\x12\x00\xfb\x4a\x9c\x3f\x6f\xe0\xc7\x0b\x05\x90\x0b\xdb\xfa\x65\x57\x15\xd8\xc8\xe2\x9d\x73\xd6\x25\x40\xe8\xdf\x0d\xbe\xee\xc9\x4b\xa7\x1d\x81\x59\x50\xb3\x59\x2e\xb1\x10\x6b\xc9\x01\xa7\x43\x51\x24\x13\x80\xa0\xf7\x4e\x37\x2e\x05\x63\xa9\xcd\x27\x5a\x9e\x07\x84\xfd\x2d\x0f\x45\xd2\x18\x3b\xbe\xe5\x31\xe5\xe8\x2d\x0f\xac\xbe\xa4\xed\x97\x5a\x31\x2a\xfe\xb8\xe1\xc5\x50\xea\x08\xd7\x1f\x2e\x78\xc8\xfc\x03\x99\x9f\x9e\x01\xe7\x80\x0a\x50\x9a\x43\x5b\x17\xf6\x4e\x6b\x86\xd9\x11\xa6\xe4\xbf\xa1\xac\x5b\xb9\x36\x42\x42\x5d\xff\x20\xcb\x6e\x43\x63\x4b\xb8\xa3\xe9\x38\x62\x8a\xc6\xe8\x38\x34\x4a\x87\x0b\x1e\x9a\xa8\xbf\xe8\xf8\x4f\x83\x75\x24\x1d\xf7\x8e\xde\x97\x6a\x79\x21\x65\xc7\x88\xd0\x6e\xa4\x93\x12\xf0\x37\x59\x26\x55\x06\xb3\x13\xda\x4e\x8e\xad\x9c\x74\x0a\x0f\x8f\xbb\x57\x98\x9f\x8c\x24\x96\x1c\x0a\x70\x91\xf1\xb8\x96\x1f\xea\xe5\xe9\xfb\xd3\xfb\xf3\xbd\x08\x03\x42\x92\xb0\x68\xaf\xc5\x0a\xce\xa1\x12\xab\x87\xd6\x24\xc3\x94\x47\x12\x6a\x82\x4e\x70\xc3\xe4\xec\x7c\x80\xde\x2c\x06\x77\x80\x27\x74\x0a\x27\x9e\x4e\x2f\xe1\xa2\xa5\x17\x90\x27\xce\x85\x98\xba\x9e\x9c\x35\xc9\xfd\x42\x07\x8f\xdc\x5b\xf4\x9f\x39\x7e\xbb\x32\xba\xf2\x75\xb6\x7e\xdb\xc4\x63\x0b\x6b\x75\x4e\x82\x71\x09\xac\x7b\x93\xe8\x69\x6c\x3d\x66\x3d\xa8\xb4\x25\xb9\xef\x0d\xdb\x31\xc0\x46\x84\x07\xff\xef\x63\x74\x9f\xdf\x1f\x7c\x05\xa4\xf4\x5c\xba\xd4\x3d\xb7\x7d\x74\xe4\x2e\xce\xe6\x72\x7f\x4f\xad\x71\x6c\x4b\x1c\x0e\x5e\xea\x23\x9c\x30\x66\xd2\x77\xe6\x2f\xd4\x76\x1f\x73\xc6\x7b\x77\x9c\xc6\xfd\x9d\xb9\x6f\xfe\x06\x2e\x84\xba\x27\x48\x73\x09\xd4\xc9\x9f\x00\x00\x00\xff\xff\x1a\x7d\x3b\xec\xc9\x0b\x00\x00")

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

	info := bindataFileInfo{name: "templates/identities_registry.gotpl", size: 3017, mode: os.FileMode(420), modTime: time.Unix(1521232602, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesModelGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xdd\x6f\xdb\x38\x12\x7f\xf7\x5f\xc1\x13\xba\xbb\xf6\xc1\x51\xfa\xec\xbd\x2c\xd0\x6b\xd3\x45\x70\xed\x36\xa8\x8b\xde\x43\x51\x6c\x18\x69\xe4\xb0\x91\x48\x95\xa2\xdc\xe4\x0c\xfd\xef\x07\x7e\x49\xa4\x3e\x2c\x39\x4d\x76\xfb\xe0\x97\x20\x1a\x92\xbf\xf9\xe4\x70\x86\x74\x8e\xa3\x5b\xbc\x01\xb4\xdb\xa1\x70\x0d\x22\x7c\xc9\x68\x42\x36\x25\xc7\x82\x30\x1a\xfe\x81\x33\x40\x55\x35\x9b\x91\x2c\x67\x5c\xa0\xf9\x0c\x21\x84\x82\x24\x13\x81\xfe\xaf\xb8\xa7\x51\x30\xd3\xff\x6f\x88\xb8\x29\xaf\xc3\x88\x65\xa7\x38\x67\x1c\x04\x3b\x21\x34\x3a\x85\x14\x32\xa0\x02\xa7\x7a\xc9\x6e\x87\x38\xa6\x1b\x40\xe1\x3a\x87\x28\xfc\x70\x9f\xc3\x25\x67\x5b\x12\x03\x2f\xd0\x49\x55\x69\x2c\x29\x0e\xaa\xaa\x7a\x09\xd0\x58\x0d\x2e\x66\xb3\x59\x8d\xf0\x8c\x2c\xd1\x33\xa0\x65\x86\x56\x67\x28\x3c\xa7\x65\x56\x48\x61\x4f\x4f\xe5\x0a\x35\xa0\xe0\x51\x55\x21\x0e\x39\x87\x02\xa8\x28\x90\xb8\x01\x94\xb3\xa2\x20\xd7\x29\xa0\x2d\x4e\x4b\x28\x50\xc2\x38\xc2\x42\x70\x72\x5d\x0a\x50\xdc\xf5\xf2\x17\x96\x66\x0c\x11\x84\x33\x21\x11\x3b\xf8\x85\xe0\x84\x6e\x66\xb3\x88\xd1\xc2\x9a\x69\xb7\x3b\xb1\x82\x52\x9c\xc1\x12\x3d\x53\xdc\xa4\xb0\x7a\xf1\x47\xcd\xdc\xe8\x6c\xc4\xa6\x9a\x53\x5b\x62\xbd\x54\x4e\xd0\xff\x55\x55\x68\x6d\x53\x2f\xe9\x48\x75\xa6\x55\xb1\x2b\x3c\x6b\x2a\x63\x36\xff\x1b\xab\x69\xa7\xbc\x65\x31\xa4\xe1\x39\x15\x44\xdc\x1b\xcd\x2f\x62\x50\x9f\x6d\xb9\x6a\x3a\x4b\xd4\x37\xbb\xfe\x02\x91\x08\x67\x5b\xcc\xa7\xe1\x9d\xa1\x3a\x40\xc2\x9a\xb8\x53\x92\xca\xa9\x2b\x54\xc7\x83\x03\xf5\x1e\x0a\x21\x47\xab\x2a\x58\xaa\xa9\x2f\xb1\x80\x0d\xe3\xf7\xab\xbe\xa9\xac\xe4\x51\xed\x41\x3d\xff\x92\x93\x2d\x16\x12\xbd\x35\xdd\x0c\x54\xd5\x72\x56\xe9\x58\x23\x09\xa2\x4c\x78\x93\x2e\x8a\xf7\x8c\x89\x26\xd6\xfa\xb5\xbc\x4c\x4b\x8e\x53\x54\x55\x6f\x48\x21\x5c\xbb\x61\x94\x4a\x0a\x4b\x26\xac\xad\xe3\x6d\x0a\x8f\x4f\x9f\xff\x39\x38\xd3\xf8\xf8\x25\xa3\x02\xa8\x70\xdc\x29\x4a\x4e\xb5\x2f\x49\xaf\x2f\x0b\x44\xa8\xfa\x94\x42\x87\xb3\xa4\xa4\x11\x9a\xb3\x89\x22\x2d\xda\x0c\xe7\x8b\x7e\x7f\x2b\xaf\x68\x61\x86\xa1\x9b\xb0\x99\x59\x6d\xf2\x46\x05\x8c\x72\x46\xa8\x00\x8e\x04\x43\x18\x45\x72\x4c\xca\x3d\x4d\xd2\x87\x68\x96\xfb\xea\x78\xaa\x26\x04\xcb\x04\x63\x34\x53\xc2\xac\xce\x10\xce\x73\xa0\xf1\x7c\x1a\x8b\x5d\xb5\x44\x2c\x0c\xc3\x85\x6b\x9c\x9f\x25\x94\x51\xff\x85\x42\x33\xa0\x85\xe7\x33\xc1\xd4\x27\x46\x14\xbe\x69\xee\xc6\xa9\x4f\x65\x0d\x2d\xcb\xdc\xf2\x0f\xc3\xb0\xed\x67\x6d\x91\x89\x06\x63\xa5\xf8\x4e\x7b\xc9\x9c\xfe\xe7\x52\x1a\x44\x02\xe9\x44\x6c\xa5\xd3\xf9\xc5\xf2\xa9\xd9\xb0\x52\xa8\x05\xe1\x7c\xdf\x3e\x5a\x68\xfc\xca\x8b\x59\x56\x0a\xe3\x14\xb5\x13\x23\x46\xb7\xc0\x85\xeb\x13\x15\x95\xb4\x13\xfd\x5a\xed\xe2\x61\x46\x97\x7f\x7b\x76\x94\x83\xd9\xb2\xe7\x9e\x99\xbb\xca\x35\x1b\x11\x90\x39\x76\xdb\x6b\x31\x39\x77\xbf\x4d\x5e\x41\x82\xcb\x54\xbc\xe3\x31\x70\x2f\xe5\xc4\x7a\x00\x31\x39\x42\xe8\x06\x25\x04\xd2\xb8\xb0\xc1\x1a\xe9\x00\x39\xdc\x30\x2e\xc3\xf9\x02\x7d\xfa\xac\x0f\xe8\x56\xa2\xb1\xe4\x46\xb9\x56\x75\xf2\xce\x88\x55\xd7\x01\x4d\x89\x52\x1f\x4b\xfe\xc1\xe2\x9f\xb0\xda\x26\xda\x06\x1f\x81\x17\x84\x51\x4f\xfd\xad\xa1\x7d\xaf\xba\x06\x7b\xbe\x40\x84\x1a\x97\xbb\xb9\x14\x44\xf8\xe2\xf2\xe2\x82\x26\x2c\xb4\x52\x54\x4a\x2c\xe7\xfc\x1f\x2b\x00\xda\x07\x7f\x26\xa7\x48\xc1\x71\x7b\x99\x3d\x9c\xc7\x4f\x2f\x53\x39\x95\x91\x94\x79\xb7\x73\x2d\xdf\xb2\xf8\x6e\x87\x32\x7c\x0b\x92\xaa\x0a\xc3\x99\x2c\xae\xac\xe8\xd2\xca\x0a\xda\x2a\x27\x8d\x70\xf5\xa5\x60\x74\x15\x9c\x04\xe8\x5a\xfd\xf3\xa7\x12\xd8\x18\x3c\xb8\xd2\xab\x64\xf9\x1a\xbe\x2d\x05\xdc\x19\x27\xfd\x01\xdf\x46\x8c\x60\xcf\x1a\x99\x57\x87\xb3\x84\x94\x4b\x39\x71\x04\x70\xbe\xd8\x0f\xd2\x0a\xd8\x9f\xf7\xcd\x6d\xa2\xd8\x35\xc6\x6a\x4f\x08\x2c\xbb\x71\x6f\x4b\x55\x55\xc8\x4a\xcf\x30\xae\xaa\x6b\xdf\x29\x17\x94\x08\x82\x53\xf2\x3f\xb7\x68\x6f\x95\xa3\x8a\xb1\x87\xd3\x62\xd8\xb3\x47\x7a\xcb\x92\xa1\x12\xd3\x6e\x92\xbd\x06\x5c\xa0\x47\xa9\x3c\x50\xa7\xf4\xb0\xf9\xb3\x95\xd0\x74\xad\xed\x49\xfa\x4b\x81\x4a\x4a\xbe\x96\xb6\xc0\x92\x6b\x0e\x94\x5e\x2e\x99\x2f\x90\x9f\xc4\x74\x6d\xaa\xd7\x3a\xd2\x58\x77\xd8\x24\x1c\xd6\x0c\x9a\x49\xf2\xe4\x95\xe7\x13\xc4\x36\x58\x6d\xc3\x02\x69\x01\x6d\x88\x20\x68\x86\xb5\xcf\xb4\x09\xd6\x20\x1c\xbe\x05\x88\xa7\x31\x81\xc7\x66\x4e\x62\x63\x86\xc5\x44\x3b\x4c\x33\x00\x3a\x43\x24\xee\x57\xb3\x2f\x71\xdf\x60\x1e\x47\x2c\x86\xb8\x9d\xc2\x55\x96\x99\xac\x5a\x4f\xde\x9e\x98\xb9\xf7\x1d\xaa\xb6\xbb\x18\x38\x5c\x27\x49\xa7\xc4\xfb\xd1\x4e\x51\xd7\xd3\x5a\xe2\x57\x50\x44\x9c\xe4\xc2\x18\x46\x5a\x85\x45\x7e\x85\xc1\xa2\x52\x6d\x7a\x35\x47\x56\x37\x4d\x58\x4e\x75\xd3\x2b\x16\xf5\xec\xbe\x13\x29\x0c\x7c\x1d\x94\x27\xf8\x44\x59\xcc\xa2\xcf\x41\x7b\x3f\x29\xf2\x5a\xdf\x18\x78\xfb\xce\x9f\x76\x25\x2d\x50\x44\xff\xc6\xd1\xad\x20\xd1\x6d\xb1\x47\xf1\xab\x6e\xe4\x3a\x47\xbb\xd2\x72\xca\x46\x53\x22\x75\x34\x35\xe2\x24\x99\x08\xd7\x39\x27\x54\x24\xf3\xe0\x5f\x3f\x15\xab\x9f\x8a\xdf\x02\x59\x70\x37\x09\x56\x79\xb1\x21\xe9\xac\xb5\x30\x9e\x1b\x38\xdd\xf5\xe1\x2e\xdd\xfa\x3b\x08\xd9\xbc\x69\x37\xfe\x0e\x42\x0a\xdb\xd9\xa6\xae\x6b\x7b\x27\x98\x7d\xc8\x21\x02\xb2\x6d\x67\x99\x67\x7b\xb4\x1f\xe0\x38\x5f\xf8\x7c\xec\xdd\x8a\x6f\x1c\x9d\x64\x3a\x59\xd5\x2b\xb1\x6c\xf4\xba\x6a\xae\x07\xd4\xac\x93\xe9\x86\x6c\x81\x3e\xb2\xa6\x03\x4c\xe7\xce\x3e\xec\xd5\xb9\xce\xb8\xfd\xda\xa2\x33\xe4\x20\x78\x21\xe8\xdf\x33\x7d\xc4\x29\x89\xb1\x50\x27\x05\x89\x41\x2b\x1a\x95\x9c\x03\x15\x88\xd0\x84\xf1\x4c\xef\xd6\x42\x30\x0e\xb1\xcc\x8d\xba\x9b\xd5\x95\x44\xc9\x61\x7a\x82\x35\xac\xe4\xe1\xcf\x39\xe3\x56\x03\xf5\x51\xf8\x1d\xd1\xb9\xa2\xed\xec\x1e\xfc\x5a\x12\x0e\xf1\xf9\xbe\x89\x7d\x17\x9a\x7e\x68\x37\xe7\x93\xba\x43\xfa\xc0\x31\x2d\x88\x54\xd3\x1b\x0b\xcf\xef\x72\x56\x40\x53\xcd\x1a\xf2\x7b\x23\x84\x3f\x5b\xa6\x1c\xe5\x91\x40\xef\xd3\xc0\x0e\xcb\x31\xce\x7d\x59\xad\x01\x2c\x94\xd9\xe3\x7e\xce\x1d\x70\xe8\xe2\x57\x85\xf7\x8f\x33\x44\x49\xea\x74\x80\x2d\xdb\xd4\xcd\xa0\x4f\x5f\xca\xc5\xb6\x2d\xf4\x13\x7a\xaf\x32\x82\x64\x70\x90\x2a\x1f\x48\x06\x3f\xa2\x22\x70\x27\x80\x53\x9c\x1e\xa4\xcc\xb9\x59\xf4\x37\x2b\x34\xa8\x5e\xf8\x22\x4d\xd9\x37\x88\x5f\xde\x30\x12\x35\xb1\xbd\x4f\x35\x1d\x6a\x17\x54\x5d\x52\xb4\xd4\xd2\x91\x3b\x1f\xd0\x6e\xd9\xd4\x14\x72\xdd\x17\x46\x68\x47\x80\xab\x60\x89\x82\x2b\x89\x56\x2d\x55\xde\x79\x51\x0a\xb6\x01\x0a\x1c\x0b\xb5\x63\x86\x6c\x04\x2d\xdb\xc0\x01\x4e\x6e\x84\xc0\x7c\x92\x0d\x2e\xb1\x4c\xf6\x74\x9a\x57\x97\xea\xc8\x6f\xf3\xb8\xd2\xea\x39\xb9\xe0\x29\x34\xdb\x08\x14\xbe\xc5\x77\x6f\x80\x6e\xc4\x0d\x7a\x3e\x45\xb7\xb7\xf8\x8e\x64\x65\xa6\x97\x4c\xd5\x50\x52\x1b\x3e\x92\x92\xe0\xb4\x80\x27\x53\x89\xd0\x83\x54\x22\xf4\x81\x2a\xd5\x7c\x9e\x5e\x25\x7c\xa7\x1e\x8e\xd0\xf3\xf0\xf9\xd0\xc1\x90\xa4\x0c\x8b\x49\xf9\xc7\x38\xf1\xb5\x5c\x70\xa0\x0f\x3f\x9a\x67\xa5\xc7\xd3\xd7\xd4\xc0\x53\x85\xbe\xa0\x93\x45\x26\x54\xcc\x5b\x62\x2f\x1e\xdb\x4f\x63\x81\xf8\x98\x5e\xd3\x71\x7a\xb8\xd7\xac\x14\x4f\xe0\xb5\x89\x32\x3f\xc4\x69\x8d\xd4\x7f\x9d\xd3\xfe\xd6\xc2\xe4\x7b\x72\xc4\x0f\x51\x88\x3c\x92\x1f\x7a\x07\x7b\xa6\x92\x04\xa5\x40\x5b\x65\xce\x02\xfd\x86\x9e\xd7\x32\x99\x46\xcd\x9f\xe2\x3e\x5a\x18\x0c\xe8\x59\x5b\xaf\x86\xce\x2a\xdb\xd2\x93\xd4\x5e\x8a\xe5\x10\x91\x84\x44\xaa\x83\x79\xcd\x78\xdd\x0f\x78\xad\x6b\x4d\xf5\xa6\xd7\xd7\x13\xba\xeb\x6b\x7e\x91\xa0\xae\x54\x6f\xe1\xde\xf6\x3d\x63\x8d\xfc\x90\x0c\x73\x05\x64\xef\xcf\x9a\x20\x18\x10\x47\x35\x4b\x24\x41\xdb\x25\x62\xb7\x32\x6c\xf6\xb1\x6d\x1a\x9f\xb7\x38\xff\x24\x19\x7d\xfe\x55\x2e\xdb\xb9\x76\xda\xce\x8c\xe5\x4e\x4f\xd1\x7f\x01\x45\xac\x4c\x63\xd5\x1a\x25\x84\xc6\x88\x88\x25\x2a\x18\x4a\x41\xfc\x52\xa0\xe8\x06\xa2\x5b\xc4\xcc\xdb\x33\xfb\x06\x1c\x45\xb8\x00\x44\x68\x0c\x77\x10\xa3\x22\x87\x08\x65\x38\x9f\x4d\xba\xbd\x7d\x23\x01\x5e\xe2\x02\x7a\xc4\xb4\xaf\xa9\xbd\x46\x28\x3c\xbf\x25\x65\x9a\x3a\x7e\x29\xfc\x99\x19\xce\x27\x7a\x68\x80\xd7\x7c\x21\x31\x3e\x69\x07\x7d\x9e\xe6\x9f\x09\xca\x7b\x3a\xcf\xc6\x7f\xec\xe1\xcd\xef\x3c\xfc\xe0\x5c\x3d\xfb\xd4\x46\x90\x41\xbb\x0f\x6d\xfc\xc7\x20\x3e\xbf\xb3\x43\x6c\xa0\x4c\xd0\xfc\xd2\x46\xb3\x58\x33\x99\xa2\xba\x9d\x78\xd0\x97\xc2\x82\xd5\xe8\x46\xb0\x79\xc0\xdc\xf9\xf5\x76\x01\x08\xb9\xd4\x55\xb7\x94\x97\x95\xbc\x83\xe3\x24\x33\x67\xa9\xea\x6d\x56\x03\x3d\xd7\x49\x55\x1d\xd4\x17\x35\x87\x73\xbd\xac\xaa\x13\xe7\xb2\xad\x53\xab\x7d\xb2\x92\xb9\xe4\x55\x6f\xa3\x35\xa8\x95\x85\xee\x7d\x6c\x40\xc8\x23\xaf\x06\x9c\x33\x0e\xce\x41\xf9\xe8\x1d\x4d\xef\x1d\x6c\x87\xaa\xa5\x6e\xcd\x1b\xc5\xf5\x6e\xd7\xeb\x51\x97\xaa\x71\x5b\xf3\xa6\xe2\xda\x82\xa6\x1e\x1d\x3a\xba\x4f\xda\xbc\xd5\xca\x3a\x44\x5a\x70\x52\x22\xe7\x71\x4e\x53\x4c\x10\x50\x92\xf6\x7a\x5f\x0d\xfb\xdc\x69\x99\x3d\x80\xf3\xf0\x9d\x63\x9f\x6b\x77\x3b\x14\xe1\x9c\x08\x25\x29\x9a\xeb\x0c\xef\xa1\x2e\x1e\x28\xbe\xbd\x16\x3b\x44\x81\x93\xaa\x0a\x1c\x77\x5a\xad\x82\x89\xfc\x0f\x64\xd5\xc3\xc9\x63\x84\xfa\x39\xb5\xe3\x69\x20\xc0\x72\x0e\x91\xbf\x8b\x1b\x9a\x0d\x5a\x67\xce\x84\x90\xf5\x1e\x59\x2c\x64\x4d\x5c\xf5\x3c\x55\xb4\xde\x27\x46\x79\x38\x37\xa0\x7a\xc0\x10\xb4\xbc\xcd\xe8\x28\xd0\x6b\x92\x0a\xe0\xea\xa7\x44\xf5\x58\x43\xd3\x70\xde\x9c\x71\x44\xc6\x81\x6c\xe8\x7f\xc0\xc9\x2f\x0d\xcd\x20\xba\x73\xa6\x20\x66\x58\x78\x68\x19\x16\x26\x05\xd6\x83\xe3\xb9\xaf\x79\x31\xd1\x74\xfd\xad\x25\xaa\xc7\x46\x51\x9c\x37\xd3\x7a\xac\xa1\x69\x34\x6f\xce\x38\xa2\x2c\xd1\x1c\x30\xf9\x69\x70\xcc\xc8\x28\x84\x7b\x47\x64\x86\x6a\xd2\xaa\x7b\x8b\x34\x01\xae\x95\x71\x2d\x65\xd5\xb9\xce\x18\xc5\x72\xee\x7a\x2c\x98\x25\xad\xba\xb7\x41\x13\xe0\xda\xa2\x19\xca\xaa\xd3\xb3\x8f\x61\xf9\x07\xac\x73\xae\x4e\x3e\x4e\xd5\x39\xe6\xef\x9e\x9a\xa4\xe5\x71\x67\x8c\xc2\x5d\x72\x92\x61\x7e\xef\xed\x9d\x86\xa6\x01\xbd\x39\xa3\x88\xef\x01\xc7\xfe\x59\x6f\x29\x2b\x73\x4f\x5a\x8f\x4f\xc0\xf2\xdf\x57\x24\x96\xa6\xac\xda\x77\xae\xa3\x58\xeb\xd6\x5e\x5c\x3b\x7b\x71\x3d\x79\x2f\xae\xf5\x8b\x57\x83\xa2\xbe\x0d\x8a\x1d\x1b\x47\x29\xaf\xcd\x93\x9d\x85\xd1\x04\xfb\xeb\xe8\x7a\x78\x3c\x1e\x3a\x2f\x56\x08\xd5\x24\x2d\x96\x3b\x63\x1c\xce\x13\xcb\x91\x69\x44\x20\x85\xdc\xf9\x65\xc6\xe1\x0d\xdf\x5f\xd0\xcc\x0c\x30\x7e\xda\xae\xc6\x54\x4f\xc7\x96\x06\x1d\x5b\x9a\x63\x4b\x73\x6c\x69\x8e\x2d\xcd\x18\xe2\xb1\xa5\x39\xb6\x34\xc7\x96\xe6\xd8\xd2\x1c\x5b\x9a\x63\x4b\xf3\x83\xb4\x34\xff\x0f\x00\x00\xff\xff\xf4\x2d\xa3\xeb\x8c\x3d\x00\x00")

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

	info := bindataFileInfo{name: "templates/model.gotpl", size: 15756, mode: os.FileMode(420), modTime: time.Unix(1521846993, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRelationships_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x94\x41\x6b\xdb\x30\x14\xc7\xef\xfa\x14\x7f\x4c\x19\x09\xa4\xf6\xce\x81\x1c\x46\x07\xa5\x87\x8d\x91\xb2\x53\xc8\x41\x95\x5f\x1c\x51\xfb\xc9\xc8\xcf\x19\x41\xe8\xbb\x8f\xd8\x49\x9b\x30\x97\x30\x5c\x4a\x8f\x7a\xf2\xfb\xbd\xff\x4f\x12\xae\xb5\x79\xd6\x05\x21\x04\xa4\x8f\x24\xe9\x9d\xe3\x8d\x2d\x5a\xaf\xc5\x3a\x4e\x7f\xea\x8a\x10\xa3\x52\xb6\xaa\x9d\x17\x24\x85\x95\x6d\xfb\x94\x1a\x57\x65\xba\x76\x9e\xc4\xdd\x5a\x36\x19\x95\x54\x11\x8b\x2e\x13\xa5\x8c\xe3\x46\xc0\x2e\x77\xe6\x51\xbc\xe5\x02\x0b\x24\xab\x6e\xbd\x4e\x90\x65\x60\x57\x5a\x96\x39\x76\xda\x9b\x2d\x99\xe7\x59\x4e\x3a\x37\x2e\x27\xa5\xd4\x4e\x7b\x78\x2a\xbb\xe9\xcd\xd6\xd6\xcd\x92\x0a\xdb\x88\xdf\xe3\x65\x44\xba\x1c\xda\x57\x2a\xcb\x70\xb1\x03\x4f\xd2\x7a\x6e\x20\x5b\x42\xe5\x72\x2a\x2f\xc9\xa9\xda\xb4\x6c\x2e\x7b\x26\xd3\x6b\x73\x10\x94\xc2\x11\x3d\x9c\x54\x45\xd5\x93\x2d\x5b\x99\x4c\x4f\x0d\x43\xac\xc5\xb5\x69\x21\x1e\x9a\x43\x80\xd7\x5c\x10\x6e\x88\xc5\xca\xfe\x70\x2b\x33\xdc\x9c\x98\x98\x2f\xfa\xbb\xbb\xd4\x8f\xf1\xad\xb9\xab\x10\xce\x48\x31\x3e\xe4\xfd\x6a\x8d\x05\xbe\x0c\x27\x0a\x0a\x38\x04\xb9\x85\xdd\xa0\x10\x4c\x4a\xe2\xd7\x04\xe9\xb7\xb2\x74\x7f\x9a\x3b\x4f\x5a\x68\x8a\xaf\xfd\x6c\xe0\xbc\x3c\x47\xa5\xeb\x55\xd3\xbd\x88\xf5\x93\x73\x25\x7a\x66\x4f\x3d\xfa\xd5\xda\x13\xcb\x41\xe8\x95\x7d\x4f\x82\xc4\x74\x8c\xe4\x04\x06\x92\x10\x5e\x3e\x8f\x31\xc1\x1c\xe2\x5b\x9a\x9d\x21\x89\xf3\xd3\xe7\x71\xa6\x06\xaa\x57\x6c\x7e\xd7\xf9\x90\x4d\x5f\x1e\x67\xd3\x76\x8c\x91\x36\x7d\x9c\x5f\x5a\xcc\xf6\x13\xa4\xf9\xbf\xb3\xfd\x4e\x25\x0d\x9c\x6d\x5f\x1e\x67\x93\x77\x8c\x0f\xb5\xb9\x27\xf9\x47\x65\x49\xe2\x2d\xed\x46\xca\x14\x24\x1f\x6d\xf2\x43\xf3\xfe\x4d\x9b\xc3\xe6\x68\xa3\x4a\xf3\xfe\x5d\xde\xfe\x03\x6f\xdc\x67\x48\x73\x51\x8d\xfd\x0f\xfb\xb8\x8e\xea\x6f\x00\x00\x00\xff\xff\xf1\xb5\xf1\x24\x61\x07\x00\x00")

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

	info := bindataFileInfo{name: "templates/relationships_registry.gotpl", size: 1889, mode: os.FileMode(420), modTime: time.Unix(1518473635, 0)}
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

