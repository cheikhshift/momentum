// Code generated by go-bindata.
// sources:
// web/test.tmpl
// web/your-404-page.tmpl
// web/your-500-page.tmpl
// tmpl/ang.tmpl
// tmpl/jquery.tmpl
// tmpl/server.tmpl
// DO NOT EDIT!

package momentum

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

var _webTestTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x56\x48\xcc\x4b\x57\xa8\xad\xe5\xaa\xae\x56\x28\x4e\x2d\x2a\x4b\x2d\x52\xa8\xad\x05\x04\x00\x00\xff\xff\xdf\xac\x12\x18\x16\x00\x00\x00")

func webTestTmplBytes() ([]byte, error) {
	return bindataRead(
		_webTestTmpl,
		"web/test.tmpl",
	)
}

func webTestTmpl() (*asset, error) {
	bytes, err := webTestTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/test.tmpl", size: 22, mode: os.FileMode(420), modTime: time.Unix(1510775959, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webYour404PageTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x92\x41\x6f\x13\x31\x10\x85\xcf\xe4\x57\x4c\x7c\x6d\xbc\x26\x0a\xd0\x16\xd9\x2b\xb5\x34\xa8\xaa\x80\x50\x28\x42\x70\x33\xeb\xc9\x7a\x12\xaf\xbd\xb1\x27\x89\xc2\xaf\x47\xd9\x04\xb5\x27\xcf\xcc\x93\x9e\xbe\xf7\x64\x3d\xbe\x5b\x7c\x78\xfa\xf5\x75\x0e\x9e\xbb\x50\x8f\xf4\xf1\x81\x60\x63\x6b\x04\x46\x51\x8f\x00\xb4\x47\xeb\x8e\x03\x80\x1e\x4b\x09\xdf\x70\xb3\xa5\x8c\x0e\x3a\x64\x0b\x6c\xdb\x02\x52\x9e\xf5\xe1\xd4\x78\x9b\x0b\xb2\x11\x5b\x5e\xca\x2b\xf1\x52\x8a\xb6\x43\x23\x76\x84\xfb\x3e\x65\x16\xd0\xa4\xc8\x18\xd9\x88\x3d\x39\xf6\xc6\xe1\x8e\x1a\x94\xc3\x32\x01\x8a\xc4\x64\x83\x2c\x8d\x0d\x68\xa6\x13\x28\x3e\x53\x5c\x4b\x4e\x72\x49\x6c\x62\x3a\x5b\xbf\xd2\x4c\x1c\xb0\xbe\x0d\x36\xae\xa1\xb7\x2d\x6a\x75\xba\x1c\xe9\xd5\x7f\x7c\xfd\x27\xb9\xc3\x19\xc6\x4f\xeb\x7b\x0c\x21\x4d\x60\x9f\x72\x70\x63\xad\xfc\xb4\x1e\x3d\x67\x5c\x3d\x6e\x31\x1f\x60\x49\xb9\xf0\x04\xd8\x63\x84\x27\x64\x8f\xf9\xbc\xdc\xa6\xc4\x85\xb3\xed\xe1\xe1\x7b\xf5\x1c\xbf\x34\x99\x7a\x86\x92\x1b\x23\x3c\x73\x5f\xde\x2b\xd5\x24\x87\xd5\x6a\x73\xf4\xab\x9a\xd4\xa9\xd3\x28\x67\xd5\xb4\x9a\x56\x25\x50\x57\x75\x14\xab\x55\x11\x40\x91\xb1\xcd\xc4\x07\x23\x8a\xb7\xb3\xab\x37\xf2\xe6\xf2\xe3\xef\xd5\xe5\xee\xc2\xa9\xe2\xba\xcf\x9b\x5e\xc5\xc5\xe3\x3e\xd0\xa7\xdd\x8f\xf2\xb0\xbc\xbb\xff\x79\xb1\xbe\x5e\x74\xad\xb2\x6a\xee\xf1\xc6\xb5\xfc\xf7\x4b\x99\xf9\x7e\x69\xdb\x77\x73\x77\xfd\xf6\x75\x14\xd0\xe4\x54\x4a\xca\xd4\x52\x34\xc2\xc6\x14\x0f\x5d\xda\x16\x51\x6b\x75\x62\x1d\xc0\x87\x9a\x4e\xed\x68\x35\xfc\x83\x7f\x01\x00\x00\xff\xff\xb5\x60\xf2\x4b\x17\x02\x00\x00")

func webYour404PageTmplBytes() ([]byte, error) {
	return bindataRead(
		_webYour404PageTmpl,
		"web/your-404-page.tmpl",
	)
}

func webYour404PageTmpl() (*asset, error) {
	bytes, err := webYour404PageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/your-404-page.tmpl", size: 535, mode: os.FileMode(493), modTime: time.Unix(1510764845, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webYour500PageTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x92\x41\x6f\x13\x31\x10\x85\xcf\xe4\x57\x4c\x7c\x6d\xbc\x26\x0a\xd0\x16\xd9\x2b\xb5\x34\xa8\xaa\x80\x50\x28\x42\x70\x33\xeb\xc9\x7a\x12\xaf\xbd\xb1\x27\x89\xc2\xaf\x47\xd9\x04\xb5\x27\xcf\xcc\x93\x9e\xbe\xf7\x64\x3d\xbe\x5b\x7c\x78\xfa\xf5\x75\x0e\x9e\xbb\x50\x8f\xf4\xf1\x81\x60\x63\x6b\x04\x46\x51\x8f\x00\xb4\x47\xeb\x8e\x03\x80\x1e\x4b\x09\xdf\x70\xb3\xa5\x8c\x0e\x3a\x64\x0b\x6c\xdb\x02\x52\x9e\xf5\xe1\xd4\x78\x9b\x0b\xb2\x11\x5b\x5e\xca\x2b\xf1\x52\x8a\xb6\x43\x23\x76\x84\xfb\x3e\x65\x16\xd0\xa4\xc8\x18\xd9\x88\x3d\x39\xf6\xc6\xe1\x8e\x1a\x94\xc3\x32\x01\x8a\xc4\x64\x83\x2c\x8d\x0d\x68\xa6\x13\x28\x3e\x53\x5c\x4b\x4e\x72\x49\x6c\x62\x3a\x5b\xbf\xd2\x4c\x1c\xb0\xbe\x0d\x36\xae\xa1\xb7\x2d\x6a\x75\xba\x1c\xe9\xd5\x7f\x7c\xfd\x27\xb9\xc3\x19\xc6\x4f\xeb\x7b\x0c\x21\x4d\x60\x9f\x72\x70\x63\xad\xfc\xb4\x1e\x3d\x67\x5c\x3d\x6e\x31\x1f\x60\x49\xb9\xf0\x04\xd8\x63\x84\x27\x64\x8f\xf9\xbc\xdc\xa6\xc4\x85\xb3\xed\xe1\xe1\x7b\xf5\x1c\xbf\x34\x99\x7a\x86\x92\x1b\x23\x3c\x73\x5f\xde\x2b\xd5\x24\x87\xd5\x6a\x73\xf4\xab\x9a\xd4\xa9\xd3\x28\x67\xd5\xb4\x9a\x56\x25\x50\x57\x75\x14\xab\x55\x11\x40\x91\xb1\xcd\xc4\x07\x23\x8a\xb7\xb3\xab\x37\xf2\xe6\xf2\xe3\xef\xd5\xe5\xee\xc2\xa9\xe2\xba\xcf\x9b\x5e\xc5\xc5\xe3\x3e\xd0\xa7\xdd\x8f\xf2\xb0\xbc\xbb\xff\x79\xb1\xbe\x5e\x74\xad\xb2\x6a\xee\xf1\xc6\xb5\xfc\xf7\x4b\x99\xf9\x7e\x69\xdb\x77\x73\x77\xfd\xf6\x75\x14\xd0\xe4\x54\x4a\xca\xd4\x52\x34\xc2\xc6\x14\x0f\x5d\xda\x16\x51\x6b\x75\x62\x1d\xc0\x87\x9a\x4e\xed\x68\x35\xfc\x83\x7f\x01\x00\x00\xff\xff\xb5\x60\xf2\x4b\x17\x02\x00\x00")

func webYour500PageTmplBytes() ([]byte, error) {
	return bindataRead(
		_webYour500PageTmpl,
		"web/your-500-page.tmpl",
	)
}

func webYour500PageTmpl() (*asset, error) {
	bytes, err := webYour500PageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/your-500-page.tmpl", size: 535, mode: os.FileMode(493), modTime: time.Unix(1510764845, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplAngTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x55\xc1\x6e\xdb\x38\x10\xbd\xeb\x2b\x5e\x85\xa2\x96\x10\x45\x6e\x0f\xbb\x07\xbb\x4a\x10\x14\x3d\x14\x58\xa0\x41\xda\x9e\x0c\xa3\xa0\xe5\xb1\x45\x9b\x21\x19\x92\x52\xe2\x35\xfc\xef\x0b\x52\x92\x25\x07\xd9\xc5\x1a\x30\x20\x8d\xde\x1b\xce\xbc\x47\x0e\x3f\xdb\xd2\x70\xed\x60\x4d\x59\xc4\x95\x73\xda\xce\xa6\x53\xb6\x63\x2f\xf9\x56\xa9\xad\x20\xa6\xb9\xcd\x4b\xf5\x18\x62\x53\xc1\x57\x76\xca\xe4\xb6\x16\xcc\xec\xec\xf4\x53\xfe\x47\xfe\x67\xff\x9e\x3f\x72\x99\xef\x6c\x7c\xf3\x79\xda\xe6\xbc\x89\xfa\xe4\xee\xa0\xa9\x88\x1d\xbd\xb8\xe9\x8e\x35\xac\x8d\xc6\x37\x11\x30\x9d\xe2\x4e\x13\x66\xb8\xbb\xff\x06\x92\x6b\xad\xb8\x74\x28\x05\x27\xe9\xf2\x08\x68\x98\x09\x80\x02\xe7\x55\xd4\xba\x16\x94\x4c\xee\x34\x4d\x32\x2c\x96\xe9\x3c\xea\x70\x96\x0c\x67\x82\xff\xed\xd1\x9b\x5a\x96\x8e\x2b\x99\xa8\xd5\x2e\x83\x36\xb4\xe1\x2f\x29\x8e\x3d\xd2\x19\x14\x58\x2c\x33\xe8\x79\x04\x6c\x94\x49\x34\xb8\x84\x5a\xed\x5a\x10\xc0\x37\xf0\xdc\xbc\x62\xf6\xfb\xb3\xbc\x37\x4a\x93\x71\x87\x44\xa7\x3d\xa0\xcd\xb4\x47\xd1\x65\xc7\x6d\xff\x70\x85\x78\x11\xe3\x0a\xda\x3f\x2d\x63\xcc\xa0\x33\x34\x28\x7c\xfa\x85\x5e\xce\x3b\xba\x75\x26\xd7\xb5\xad\x92\xa4\xc1\xbb\xa2\x80\xac\x85\xc0\x87\x0f\x41\x2c\xb5\xf1\x84\xa2\x40\xac\x56\x3b\x2a\x5d\x9c\xe2\xb6\xa3\x61\xe8\x33\x69\x32\xec\x53\xcc\xce\x5f\x48\x96\x6a\x4d\xbf\x1e\xbe\x7d\x51\x8f\x5a\x49\x92\x2e\xd9\xa7\xbe\x8a\xc2\xd7\xf3\xc6\xd7\x26\x4d\xdb\x72\x4e\x51\xfb\x37\xe4\x6a\x23\x43\x6d\x3b\xc5\x65\x12\x7f\x88\xd3\x79\x74\x8a\xa2\xe8\x4e\x53\xbe\x61\xa5\x53\xe6\xd0\x8b\x7f\x16\xf9\xbd\xdf\x38\xe9\xb1\x37\xe2\x77\xeb\xd8\xd1\x3a\x66\x1c\x66\x03\x2e\x3d\xe2\x94\x61\xc5\xac\x77\x3c\x8e\xb3\x8a\xd8\x9a\x8c\x9d\xe1\x78\xca\xbc\xfb\x63\xac\x21\xab\x95\xb4\xe4\x39\xc0\x69\x1e\x36\x4b\xa9\xe4\x86\x6f\x2d\x32\xa0\x5f\xac\x36\xe2\x81\x9e\xc6\x96\x37\x64\x56\x59\x6d\x44\x86\x35\x73\x2c\x2b\x99\x10\x2b\x56\xee\xd3\xde\x7d\xc1\xac\xfb\x6a\x8c\xf2\x7b\xe0\x78\x0a\xdb\xc7\xff\x7c\xd5\x79\x28\x39\x49\x7b\x8f\x42\x63\xc9\xf1\x2c\x30\xf0\x48\xae\x52\xeb\x19\xc2\x22\x18\x7d\xa8\x8d\x98\xb5\x39\x42\x7f\x57\x21\x82\x2b\x24\x48\x02\x18\xef\x0a\xc4\xf7\xdf\x7f\xfc\x8c\xbd\xc9\x43\xe4\xd7\xcf\x18\xf0\xa1\xef\xc1\xe9\x7c\x4f\x07\x9b\xf8\xc2\xd3\x5c\x90\xdc\xba\x0a\x37\xf8\x88\x14\xb7\x88\x6f\xbd\x8b\x23\xfb\xdb\xfe\x10\xc7\x69\x90\x13\x69\x36\xaa\xc7\x7f\x9b\xb5\x88\x51\xb4\x53\x1c\x5d\xa9\xdd\xeb\x19\x70\x4a\x73\x57\x91\x4c\x7a\x2d\x61\xeb\xb2\x24\x6b\xbf\x74\x1a\x0e\xae\x60\x2c\x4a\x38\xc7\xae\xe2\x16\xbd\xd8\x78\xe6\x42\x60\x45\x21\x40\x6b\x30\x7b\x90\x65\x65\x94\x54\xb5\x15\x87\x0b\xea\xeb\x3c\xcf\x15\x49\xb8\x8a\xd0\xaf\x05\x6e\xc1\x1a\xc6\x05\x5b\x09\x8a\x2e\xc1\x28\x95\xb4\x4a\x50\x2e\xd4\x76\x28\x6e\xfe\x3a\x6d\x68\x96\xe4\xfa\x0c\xc9\x83\x30\x17\x84\x37\xaa\x29\x5f\xb7\xdd\xd2\x30\xf0\xa2\x11\xfe\x34\x9c\x08\x90\xdf\x5f\xff\x47\xb5\x37\xf5\xf1\xd3\x87\x75\x49\xa0\xca\xb2\x1e\x59\xd4\xf1\x54\x98\x77\x0d\x99\xee\xc8\xda\x41\xae\x67\xee\xaa\x81\x6e\x1d\x73\xb5\xcd\x5f\xeb\xf6\x3e\x89\x73\xa1\xbc\xfb\xd7\xaa\x21\x23\xd8\x21\x4e\xf3\xd2\xda\x64\xb2\xe6\x56\x0b\x76\x98\x64\x98\x48\x25\x69\x72\xd9\xe4\x48\xcb\x0d\x13\x96\x2e\xc4\x78\x53\xbb\xff\x84\x8d\xe5\xeb\x47\x51\x58\xaf\x1b\x44\xad\x5e\xdf\x24\x77\xa3\xd1\x70\xa7\x29\x3d\x8e\x4e\x2d\x0a\x7f\x47\x74\xec\x76\xb7\xff\xa5\x94\xfe\xfa\x34\xe2\x30\x63\x32\xec\xe9\x90\xa1\x61\xa2\x1e\xf8\x1b\x65\x90\xf8\xa9\xc0\xfd\x1d\x63\x4c\x7f\xe8\xae\xf1\x69\x0e\x8e\x9b\x02\x1f\xe7\xe0\xd7\xd7\x63\xeb\xf8\x26\xf1\xd0\x05\x5f\x2e\xf6\x74\x58\xa2\x28\x5e\x25\xf5\xbf\xae\x81\x16\x37\x1c\xb0\xee\xe9\xd4\xab\xd0\xc1\x82\x48\x17\x1d\x3c\xd0\x53\x4d\x76\xdc\x76\x3b\x7a\xfe\x6d\xb2\xa1\x9d\x41\x0f\xf4\x34\x06\x5e\xe2\xfa\x05\x80\x2d\xb9\x1f\x24\x36\x17\xb3\xf9\xb2\xa2\xdf\x83\xa4\xe1\x6e\xf0\xf6\x9c\x6f\xf6\x7f\x02\x00\x00\xff\xff\x2e\xc3\xe0\xef\x36\x08\x00\x00")

func tmplAngTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplAngTmpl,
		"tmpl/ang.tmpl",
	)
}

func tmplAngTmpl() (*asset, error) {
	bytes, err := tmplAngTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/ang.tmpl", size: 2102, mode: os.FileMode(420), modTime: time.Unix(1510772037, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplJqueryTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xce\xcd\x8a\x83\x30\x14\x40\xe1\xbd\x4f\x11\xb2\x37\xc1\x9f\x99\xc5\x60\x84\x19\xd0\xd9\x55\xa4\xa5\xd0\xee\x24\xda\xe4\x4a\x93\xd8\xdc\x88\xe6\xed\x4b\x71\x77\x56\x87\xaf\x42\xe9\x61\x09\x09\x21\xe8\xa5\xa0\x3a\x84\x05\x7f\x38\x97\x6e\x9c\xd8\xfc\x5a\x27\x1f\x99\x74\x86\x1f\x99\x16\x2c\x67\x19\x33\x60\xd9\x8c\x34\x21\x04\x6c\x98\x94\x87\x10\x05\x45\x3d\xe4\x5f\xdf\xa9\xde\x54\xa9\x70\x57\xed\x5d\x77\xd8\x34\x83\x19\xbb\xdb\xff\xdf\x23\x2b\xda\xd8\xaf\x70\xd9\x9e\xbf\xbd\xda\xaf\xe7\x93\x0a\xa5\xf8\x1c\xa4\x77\x88\xce\x83\x02\x2b\xe8\x60\x9d\x8d\xc6\xad\x48\xeb\x8a\x1f\xb0\xfa\x1d\x00\x00\xff\xff\x21\x9c\xac\xf7\xa1\x00\x00\x00")

func tmplJqueryTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplJqueryTmpl,
		"tmpl/jquery.tmpl",
	)
}

func tmplJqueryTmpl() (*asset, error) {
	bytes, err := tmplJqueryTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/jquery.tmpl", size: 161, mode: os.FileMode(420), modTime: time.Unix(1510769033, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplServerTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x4d\x4f\xdb\x40\x10\x3d\xdb\xbf\x62\xb4\x87\x6a\x2d\x1c\x07\x55\x3d\x15\x4c\x0f\x5c\xa0\x6a\x0b\x22\x54\xaa\x84\x38\x6c\xd6\x63\x62\xd7\xd9\xdd\xce\x8e\x93\xb8\x28\xff\xbd\x5a\x7f\x09\x04\xa7\x9e\x12\x8d\xdf\xbc\xf7\x66\xe7\xed\x9e\x7b\x4d\x95\x63\xe0\xce\x61\x2e\x18\x0f\xbc\xac\xd5\x4e\x0d\x55\x71\x11\x47\x31\x40\xd9\x1a\xcd\x95\x35\x50\x7b\xc2\x3f\x2d\x7a\xde\xda\x2d\x1a\x6e\xb7\xb2\xa5\x26\x75\xaa\x6b\xac\x2a\xd2\x40\x91\x6a\xd5\x34\x6b\xa5\x7f\x27\xcf\x31\x00\xec\x14\xc1\x61\xc3\xec\x20\x07\x83\x7b\xf8\xf5\xfd\xdb\x15\xb3\xbb\x1b\x58\x64\x72\x16\x03\x44\x3d\x20\xb3\x86\x50\x15\x9d\x67\xc5\xa8\x37\xca\x3c\x21\xe4\xb3\xb2\x4c\x20\xf0\x45\x51\x55\xca\x01\xde\x83\x57\x01\x0c\x79\x0e\x9f\x06\xb9\x28\x0a\x82\xbe\xd5\x1a\xbd\x87\x1c\xe4\x20\x9e\x05\xd2\xd6\x07\xe0\xc7\xd3\xd3\x24\x20\x21\xaa\x4a\x90\xc1\x71\xa8\x8a\xdb\x9b\xd5\xfd\xd7\xd5\xcd\x0f\x31\xf0\x40\x14\x31\x75\x30\xfd\x9f\x66\x92\x01\x92\x39\x45\x1e\x67\x17\xde\x59\xe3\xf1\x1e\x0f\x9c\xa4\x93\x72\x3f\x56\x68\x3c\x82\x56\xac\x37\x20\x31\x99\xc9\x22\x6d\x8d\xb7\x0d\x66\x8d\x7d\x92\xe2\xda\xec\x54\x53\x15\x30\x88\x9f\xcd\x98\x49\xf1\x19\x89\x2c\xc1\x67\x78\xab\xd7\x1b\x17\xf0\x05\xc4\x0a\x69\x87\x04\x7b\xb2\x8c\x60\x2c\x4c\x28\xf1\x6e\xdf\x31\x85\x52\x35\x1e\x21\x99\x5c\x0e\xbf\x47\xc0\x50\x9d\xa5\xdf\x76\xce\x03\xc2\xe8\x34\x74\x1e\xcf\xe2\x78\x58\xb4\x47\xaa\x54\x53\xfd\x7d\xb5\x38\xbb\xae\x87\xd9\x7b\x04\x13\xe4\xf0\xf0\x18\xba\x4b\x4b\x32\xd4\x1c\x54\x06\x02\xaa\x67\x0c\x5b\xb1\xeb\x3a\xdb\x28\x7f\xb3\x37\xb7\x64\x1d\x12\x77\xd2\x25\xd3\x01\x42\x20\xc9\x5c\xeb\x37\x12\x8d\xb6\x05\xfe\xbc\xbb\xbe\xb4\x5b\x67\x0d\x1a\x96\x2e\x81\x13\x10\xb9\x80\x13\x78\xe7\xab\x5d\xd7\x0f\xee\x31\x79\x61\x9e\x90\x5b\x32\x3d\x65\x6d\x2b\x23\xc5\x87\x61\x09\xe1\xdb\x18\x4b\x87\xa6\xcf\x49\x0a\x21\xeb\xc0\xd4\x62\xd2\x8f\x5c\x95\xaf\xf3\x33\x65\x67\x8c\x1c\xf2\x98\xf2\x2b\x54\x05\x92\x14\x97\xd6\x30\x1a\x5e\x84\x26\x91\x82\x50\xce\x35\x95\x56\xe1\x94\x96\x87\xc5\x7e\xbf\x5f\x94\x96\xb6\x8b\x96\x9a\xc1\x7a\x31\xe5\x61\x22\x34\x85\x9c\x8f\x58\x8e\x97\x2e\x0c\x33\x2e\xee\xad\xa3\x97\x89\xfe\x1f\x57\xb5\xb7\xe6\x1d\x13\xfd\x2d\xf0\x4c\x95\x79\xaa\xca\xee\xb5\x93\xd1\xca\x0b\x74\x28\xc7\xe7\xcb\xe1\x41\xb9\x88\xa7\xf7\xc6\x93\xce\xc5\x32\xc4\xa4\x54\x9a\x2d\x75\x59\xed\xc5\xc5\x8c\xfb\x17\x00\x00\xff\xff\xea\xbb\x4f\x7f\x97\x04\x00\x00")

func tmplServerTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplServerTmpl,
		"tmpl/server.tmpl",
	)
}

func tmplServerTmpl() (*asset, error) {
	bytes, err := tmplServerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/server.tmpl", size: 1175, mode: os.FileMode(420), modTime: time.Unix(1510840022, 0)}
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
	"web/test.tmpl":          webTestTmpl,
	"web/your-404-page.tmpl": webYour404PageTmpl,
	"web/your-500-page.tmpl": webYour500PageTmpl,
	"tmpl/ang.tmpl":          tmplAngTmpl,
	"tmpl/jquery.tmpl":       tmplJqueryTmpl,
	"tmpl/server.tmpl":       tmplServerTmpl,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"ang.tmpl":    &bintree{tmplAngTmpl, map[string]*bintree{}},
		"jquery.tmpl": &bintree{tmplJqueryTmpl, map[string]*bintree{}},
		"server.tmpl": &bintree{tmplServerTmpl, map[string]*bintree{}},
	}},
	"web": &bintree{nil, map[string]*bintree{
		"test.tmpl":          &bintree{webTestTmpl, map[string]*bintree{}},
		"your-404-page.tmpl": &bintree{webYour404PageTmpl, map[string]*bintree{}},
		"your-500-page.tmpl": &bintree{webYour500PageTmpl, map[string]*bintree{}},
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
