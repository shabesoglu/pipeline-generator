// Code generated by go-bindata.
// sources:
// templates/jenkins/multi-job.xml
// templates/jenkins/normal-job.xml
// templates/jenkins/pipeline.xml
// DO NOT EDIT!

package pipeline

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

var _templatesJenkinsMultiJobXml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x55\x4b\x6f\xdb\x38\x10\xbe\xe7\x57\x08\xbe\xe4\x14\x2a\xc1\x3e\xd0\x83\xe2\x6e\xe3\x26\x40\x17\x49\x6a\xd8\x35\x7a\xa6\xa5\x89\xc4\x98\x22\x05\x92\xda\xb5\x51\xf4\xbf\xef\xf0\x21\x8a\x8a\x8d\x4d\x82\xdd\x93\xc4\x6f\x1e\x1c\x7e\xfc\x66\x58\x7c\xdc\xb7\x3c\xfb\x0b\x94\x66\x52\x5c\x9f\x5f\x91\xcb\xf3\x0c\x44\x29\x2b\x26\xea\xeb\xf3\xcd\xb7\xbb\x8b\x0f\xe7\x1f\xe7\x67\x45\x29\x5b\x62\xd8\x8e\x72\xf2\x0c\x62\xc7\x84\x26\x1d\xef\x6b\xfb\x6d\x7b\x6e\xd8\xb3\xdc\x92\x07\xfb\xf3\xa7\xdc\x2e\x95\x7c\x86\xd2\x64\xde\xe1\x7a\x16\x02\x2e\x06\xc7\x0b\x6f\xf8\xe3\x8a\x5c\xfd\x36\x9b\x9f\x65\x59\x41\x4b\x83\xbb\xeb\xdc\x2d\x2a\xd0\xa5\x62\x9d\x45\xe6\x45\x9e\xae\xac\x75\x07\xd0\x7d\x86\x0e\x44\x85\x55\x32\xd0\xf3\x27\xca\x35\x14\xf9\x11\x6e\x9d\x3b\x25\x3b\x50\x26\x2c\x11\xd0\x40\x2a\x46\xb7\x32\x39\x05\xeb\x80\x33\x01\x64\x19\x7e\x96\x3e\xe6\x10\xcb\xaf\x10\x46\x7e\x0e\x17\x83\xeb\x50\xff\x25\xf9\x40\x7e\x9f\xf9\xcc\x98\xdb\x50\xbd\x7b\xa4\x2d\xcc\x3b\xaa\xf6\x94\x73\xe0\x19\xec\xa1\xec\x6d\xe9\x45\x1e\xad\x83\xbb\x36\xb4\x06\x87\xfc\xf8\x91\x91\xf5\xb0\xca\x7e\xfe\x2c\xf2\xd1\xe6\xcb\xce\xdf\x51\xb7\x3b\x78\x3e\x3d\x79\xa1\xcb\x36\x2b\x39\xd5\xfa\x7a\xd6\xf4\x95\x96\x82\x20\x42\x1e\x7b\xce\xd7\x8b\x87\x99\xe7\xbd\xa4\x62\x25\x69\x3b\x37\xaa\x47\x3e\x87\x95\xbb\x11\xa6\xe9\x96\x43\x35\x70\x1d\xd7\xd6\xb8\xe5\xb2\xdc\xdd\xf4\x8c\x57\xdf\x1b\x10\x9f\xe5\xdf\x42\x1b\x05\xb4\x75\x10\x8a\x68\x08\x7a\xd5\xef\x38\xd9\xa6\x7b\x4b\xaa\x23\x2f\x9b\xc8\x28\x56\xd7\x28\xea\x70\x34\x29\xca\x5e\x29\x10\xc6\x39\x0d\x79\x5e\xc2\xae\x02\xfb\x87\x81\x81\xf9\x77\xc8\xfe\xc6\x47\xc6\x0b\xee\x1a\xaa\xfd\x25\xba\xbf\x4b\xbc\x94\x88\x4c\x7c\x30\x56\x0f\x48\x96\xa1\x1a\x14\x15\x35\xa0\x26\xfa\xad\x35\xa1\x22\xa2\xf1\x2d\xf5\x2c\x87\x9c\x0b\x29\x9e\x58\x3d\x66\xc6\x70\xb4\x47\xc9\x39\xa5\x0d\x40\xea\x64\x39\x59\x52\x45\x5b\x3d\x68\x61\x04\x52\x3f\xd8\x77\x52\x43\x85\x0a\x1a\x08\x4d\x90\xd4\x31\xe8\x05\x6b\x7a\xa1\x20\x8b\xa4\x8e\xd8\x39\x1a\x6f\x70\xd5\x73\xd0\x4b\x6a\x1a\xec\xff\x23\x28\xf5\x6f\xe9\x7e\x05\x78\xd5\xa8\x72\xa4\x37\x59\x4d\xca\x14\x76\x2b\x6b\x39\xac\x8d\xa2\x06\xea\x43\xac\xf7\x84\xe9\x38\x14\x79\xac\x98\x1b\x40\x93\xb0\x11\x4e\x43\xb0\x49\x95\xf9\xc4\xb9\x3d\x9a\x67\x2f\x45\x26\x34\xc7\x78\xa7\xc4\x13\xb9\x76\x8c\x73\x77\x9b\x5f\x05\x06\xaf\x40\xe3\x15\x8f\xbb\xde\x7d\xfa\x72\xbf\x59\xdd\xe2\xe8\xfb\x57\xb7\x34\xa1\x13\xf7\x57\xc1\x0f\x5f\x9e\xf0\x8e\x16\x8d\xd5\x59\x9c\xa0\xa7\x8d\xa3\xf4\xf2\xff\xa2\x3d\x14\x1c\x8e\xe6\x51\xca\xa1\x19\x52\xe9\x5b\x42\x0c\x13\x3d\xb5\x65\x8f\xf5\xaf\x37\x8b\xc5\xed\x7a\x7d\xb7\xb9\x77\x34\x9d\xf0\x38\x7b\x73\x79\x27\x5a\x35\x9c\x3b\x74\x7c\xd1\xf5\x5b\xce\x74\x13\x07\x00\xd6\xcd\x9e\x32\xf2\x08\x7b\xf3\x40\x71\x67\x9e\x76\x64\x41\x7b\x62\x77\x2d\x71\x82\xa8\xbe\xd5\x07\x6d\xa0\xd5\x24\x4c\x58\xbf\x3b\x71\xe9\xe3\xb4\x0e\x73\x89\xb8\x0a\x86\xd1\xfd\xcd\x83\xf1\xc5\x71\x21\x47\xcf\xcd\x15\xf9\x95\xfc\x32\x4b\xd9\x42\x86\xfd\x80\xf3\x5d\x16\x47\x6a\x78\x80\x6d\x63\x6b\xd7\xea\x47\xe5\x63\x07\x9e\xf6\x0e\x5c\xfe\xbf\x07\x8b\x54\x26\x12\x48\x88\x9d\x50\x1a\xf6\xb0\xcf\xa5\xf6\xc9\x26\x49\xec\xc1\x1b\xbb\x83\x2f\x7a\x3c\x5e\x3c\xd8\xd4\x1c\x1f\xe7\x46\x81\x6e\xa4\x1f\xf2\x01\x13\x76\xf0\x05\x79\x15\xb9\x98\x8c\xc1\x42\x2a\x7c\x4b\x28\xb7\x73\x65\xf8\x4d\xa7\x30\x97\x6a\x7e\x73\xbf\xb9\xb5\xb2\xb3\xff\x93\x09\xdd\x71\x30\xe0\x9f\x94\x30\x40\x27\x58\x6c\x81\x17\x45\x15\xf9\x2b\x87\x4f\x19\xc4\x06\x9a\x48\xd5\xf7\xf6\x77\x45\xbb\xce\x3f\x7b\xef\x6a\x88\xc0\xd7\xfc\xec\x9f\x00\x00\x00\xff\xff\x4d\x74\xcd\x5a\x0f\x0a\x00\x00")

func templatesJenkinsMultiJobXmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesJenkinsMultiJobXml,
		"templates/jenkins/multi-job.xml",
	)
}

func templatesJenkinsMultiJobXml() (*asset, error) {
	bytes, err := templatesJenkinsMultiJobXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/jenkins/multi-job.xml", size: 2575, mode: os.FileMode(420), modTime: time.Unix(1438004167, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesJenkinsNormalJobXml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x59\xeb\x6f\xdb\x38\x12\xff\xee\xbf\x42\xc8\x15\xc8\xdd\xe2\x22\x37\xd9\xc3\xed\x1e\xe0\x78\xdb\x38\x6e\xd7\x7b\x49\x1a\xf8\xd1\xfd\xb8\xa0\xa5\xb1\xcd\x86\x12\x05\x91\x4a\x63\xe4\xfa\xbf\xdf\xf0\x29\xea\x11\x27\xc5\xee\x97\x02\x45\x80\x40\x9c\x19\x3e\xe6\xfd\x23\x3d\xfa\xe5\x21\x63\xd1\x3d\x94\x82\xf2\xfc\xfc\xf8\x34\x7e\x7d\x1c\x41\x9e\xf0\x94\xe6\xdb\xf3\xe3\xd5\xf2\xdd\xc9\xcf\xc7\xbf\x8c\x07\xa3\xa2\xe4\x9f\x20\x91\xe3\x41\x14\x8d\x48\x22\x51\x58\x0c\xf5\x20\x05\x91\x94\xb4\x50\x14\x43\x60\x7c\x3b\xe7\x92\x48\x5e\x46\x09\x23\x42\x9c\x1f\xed\xaa\x54\xf0\x3c\x96\x44\xdc\x89\xf8\xca\xb3\x8f\x94\xb8\x5a\x81\xec\xc5\x92\xff\x17\xa0\x18\x9f\x9c\x8e\x86\xc1\xd0\xf0\xf3\x2a\x0b\xd8\xf5\xc8\x70\x49\x29\xe9\x06\x4f\x74\xd9\x5c\xa5\x87\xdc\x94\xbf\xf1\xeb\x04\xd2\x37\xe1\xe2\xa3\x61\xad\x8a\x1e\xdf\x21\xe3\x12\x0a\xc8\x53\xb4\x10\x05\x31\xde\x10\x26\x60\x34\xec\xd0\x95\x30\x1a\xac\x00\x5c\xd6\x0c\x91\x20\x20\x4e\x29\x59\x73\x16\x7f\x82\xfc\x8e\xe6\x22\x2e\x68\x01\x8c\xe6\x10\xdf\xda\x8f\x5b\x33\x67\x1f\x15\xac\xda\xd2\xfc\xfc\x28\x45\x32\xfa\x66\x7f\xe2\x44\x4f\x0c\xe7\xcd\xeb\xf8\xe7\xf8\x27\x6b\x40\x5c\x5b\x99\xf6\x86\x64\x30\x7e\x7c\x8c\xe2\xa5\x1d\x44\x5f\xbe\x8c\x86\x9e\xe3\x44\x85\x24\x5b\xf0\xb2\x0b\x37\xd2\xc2\x35\xcf\x1c\x79\xf8\x15\x67\x36\x16\x6b\x6a\x8d\x3b\xd0\x4d\x14\xcf\xc4\x2c\xa7\x92\x12\xf6\x1b\x5f\xe3\x3e\x4a\x50\x24\x59\x2b\x3a\x8c\x62\x22\xde\x52\x19\xbf\xa7\x72\x31\xb9\x3e\xf2\x66\x40\xda\x9b\xb3\xf8\x2c\xfe\x8f\x0b\x99\x84\xe7\x1b\xba\xfd\x68\xa2\x76\x7c\x36\x1a\x36\x09\x46\xa8\x12\x50\xce\x21\xe3\x12\x26\x9a\x2b\xbc\x0d\x7a\xb6\x5c\xb5\x84\x9d\xac\x5a\xa7\x64\xda\x56\x78\xaa\xd5\xfc\x4a\x1b\x4a\x91\xdc\x62\xc3\x17\xaf\x86\xf3\xfa\x8f\x34\x5a\x97\x24\x4f\x76\x70\xf0\x84\x17\x5a\x66\x51\x40\x12\x9c\x2d\x57\xce\xfa\x61\xa8\x8e\x57\xf3\xe9\x86\x42\xa9\xcf\x99\x87\x9e\xef\x3b\x68\x7b\xd1\xd1\xb0\x79\x94\x51\x4a\x05\x59\x33\x58\x54\xeb\x8c\xa7\x15\xab\x83\xbe\xcb\x30\x13\x4a\x48\x2a\xf4\xc2\x7d\xc8\x91\x65\x85\x33\xfa\x38\x76\x13\xfe\x1e\x72\x28\x89\xac\x59\xc6\x3e\x15\xd2\x54\xa9\xf1\x9b\x3e\x2b\x68\x93\xbc\x92\x3b\x5e\x7e\x28\x27\x3c\xcb\xa8\x94\x50\xba\x05\xba\x0c\x1b\x50\x0c\x48\xae\xbd\x3c\x51\x5f\xbf\xf3\xf2\x4e\x14\x24\x31\x69\x61\x98\x46\xf0\x33\x86\xfd\x87\x4a\x7a\x01\xb7\x70\x87\x6e\xc4\x8b\xb2\xca\xe1\xc2\x99\xd4\xca\x36\x89\xce\x6c\x2a\x2a\x6e\x39\x63\x4e\x2a\xa0\x18\x11\xba\xcd\x79\x09\x37\x1c\x4b\xd5\xde\x9c\xdf\x89\xf6\x70\x7c\x0a\x2c\x76\x84\x31\xfe\x79\xc2\x78\xee\x0f\xdb\x26\xdb\x20\xac\x28\x4b\x27\x3b\xce\x31\x48\x0f\x24\x67\x25\x29\x8b\x2f\x61\x43\x2a\x26\x2f\x82\x29\x47\x43\xbb\x0e\x0a\x2d\x39\x67\x63\x2b\x33\x1a\x3a\x82\xad\x83\xde\x73\x9b\xad\xdb\x86\x51\x21\xfd\xfc\x12\x18\x3a\xf3\x1e\x96\xa4\xdc\x82\xbc\xa4\x65\xcd\xd8\x40\x89\x25\x16\x1c\x01\x1e\x12\x56\xa5\x90\xce\x61\xeb\x1b\x52\x40\x56\x59\x28\x82\x53\x99\x50\x51\xf5\xad\x43\x9c\x66\x84\x32\x47\x15\x77\xb4\x58\x92\xad\x0d\x5b\x37\xb2\x4e\xc8\x1b\x5b\x8e\x6d\x89\x53\x7e\xc7\x8e\x89\x87\xc5\x88\x51\x61\xd4\x26\x60\x73\xd0\xb1\xd4\x9e\x6f\x77\x4c\xb2\xf0\x58\x7f\x79\x3e\x3c\x6b\x73\x78\x90\x90\x8b\xba\xab\x0f\xf1\x44\xfa\x43\x96\x74\xbb\x45\x33\x1a\x39\xab\xed\xaa\x10\xb2\x04\x92\x61\x3d\x17\xa6\xa0\xa3\xa4\xeb\x10\x6e\x46\x3c\x07\x05\x2b\x40\xc7\xc8\xd2\x10\xeb\x2e\x84\x15\x67\xe8\x47\x95\x5d\xef\xd6\x80\x0c\x6d\xd5\xf6\x26\x18\xb5\x6d\x29\xdf\xfd\x76\x25\x88\x1d\x67\x69\xbb\x30\x2e\x56\x93\xc9\x74\xb1\x68\xd6\x41\x64\xf2\x12\xe1\x0d\x61\xe3\xd7\xa3\xa1\xfb\xac\x99\x09\x67\xd8\xef\x2f\xae\x56\x53\xd5\x56\x18\x2f\x1b\xbc\xac\x60\x20\x8d\x4e\x36\x3e\x9a\x34\x5f\x6c\x5b\x87\x1a\x0d\x5f\x6e\x1f\x1f\x2e\xa1\xcd\xad\x04\xa4\xd7\x24\xaf\x30\x77\xf7\x36\xae\x30\x0e\xbc\x0b\x1c\xd0\x72\x3b\x60\x13\xed\x33\xfc\xf8\x87\xc8\xfe\xa1\x9f\x83\x7e\x62\x2b\xcc\x2d\x17\xd2\x54\x91\x5f\x39\xbf\x13\xcd\x22\xd3\x66\x0e\x1a\xad\xe5\xc9\x9d\x43\x95\xd0\x36\x41\x50\x99\xb2\xf3\x7b\x49\x8a\xc2\x87\x59\xbb\xfd\x11\x0c\x4d\xed\x8a\xf8\x2d\x7e\x4d\xd4\xd7\x45\x30\xcb\xe3\x04\x2f\x87\xf8\xe8\x5f\xf1\xeb\x1a\x1f\x69\xe2\x35\x29\x34\xb6\xc1\x50\x2f\x33\xeb\x5b\x47\x1b\xf4\x76\xc8\x67\xf6\x7d\x39\xb0\xb3\xc0\x04\x73\x13\x55\x5f\x57\x0a\x1b\x7f\x2d\xc4\xb3\x10\x7d\x09\x18\x6d\x98\xef\xe3\x57\x8f\x17\xab\xd9\xd5\xe5\x1f\x37\xab\xeb\x8b\xe9\xfc\xcb\xc9\xab\xc7\xf7\xb3\xe5\x1f\xf3\xe9\xc7\xd9\x62\xf6\xe1\xe6\x9f\x0c\xf2\xad\xdc\x9d\xff\x84\x79\xd3\x9e\x59\x67\x5d\x8a\xc3\x4b\x2a\x90\xbc\xd7\x56\x30\x01\xdd\xa5\xbf\x1c\x0d\x76\x15\xed\xf7\xa8\xa4\x19\x20\xda\xcc\xd0\x8a\xf1\xb2\xfe\xee\xf5\x6a\x20\xfb\xe6\x34\xfe\x77\x7c\x76\x64\x6b\x54\x27\x72\x1a\xf9\xf0\x4d\x61\x4c\x5d\xa0\x3a\x75\x0a\x1b\x9e\xce\x4f\x05\x06\x36\xa2\x89\xfc\xbe\xa3\xd2\xef\xa8\xf4\x3b\x2a\xfd\x8e\x4a\x8d\x13\xf2\xde\x2d\x43\x58\x59\x43\xba\x6f\xbb\xe1\x3e\x51\xf6\x0d\xb8\xa8\xdf\x1c\x16\x8c\xdc\xc3\x15\x59\x03\xb3\xdd\x00\x3d\x8a\x01\x09\xe9\x0d\x4f\xed\xd3\x47\x28\x81\x39\x18\xf2\x07\x2d\xc4\x92\x90\x7c\xce\x49\xe6\x22\xd5\x0d\x07\x75\x29\x4a\x5b\x15\x28\x35\xa6\x66\x3c\xb9\x33\x1a\xec\x20\xbf\xe4\x9f\x73\x03\x61\x35\x09\xaf\x07\x6e\xd2\xb3\x72\xdd\xc5\x1c\x3c\x3e\xbc\x54\x47\x6a\x60\x3a\x1b\xd6\x3c\x8c\x59\x93\x29\x5e\xab\x16\x79\x70\x00\xef\x7f\x05\xda\x0f\xb1\xfe\x9f\x42\xfa\x5d\x9c\x7f\x00\xe5\x1f\xc0\xf8\x4f\x23\xfc\x97\xe2\xfb\x16\xba\x7f\x39\xb6\x6f\x04\x95\x8e\xe3\xf0\x6a\x85\xb5\x77\x0b\x51\xfc\xd6\x3d\x5d\x42\xd1\x46\xf6\x2e\x4f\x12\x5e\xec\xdd\x9b\x65\x3c\xc1\x81\x9b\xe3\x33\x33\x94\x40\xd4\xf4\xe3\x59\x7c\x5a\x67\xa7\x7b\xcc\x55\xc6\xb7\x36\x56\x49\xa9\x11\xa2\x76\x40\xf0\xda\xab\x27\x6c\x28\x53\x5d\x49\xc9\xfb\x9d\x94\x9c\xa5\xfb\x8b\x98\xae\x8d\xf5\xbd\xce\x96\x3c\x51\x53\x04\x30\x5c\xb7\xf3\x34\xdc\xab\x97\xbf\xf1\x68\x2b\x2e\xec\xcc\xa3\x00\x67\x60\xe0\xb2\x35\x49\xee\x96\xfc\x8a\x08\xb9\xa8\x92\x04\x84\xd8\x54\xcc\x7a\xee\x49\x76\x80\xa2\x6c\xa4\xbd\xd3\x8a\x2c\xa4\x6a\xe2\xdb\xfd\x18\x6b\xf4\x7b\xc6\xd7\x84\x2d\x40\x4a\x4c\x9b\x3a\x24\x5b\x82\x1e\xc5\x38\xc5\x3c\x25\xe5\xd8\xf9\xde\xe1\x5c\x28\x8b\x92\xe6\xd2\xd9\x2d\x40\x0d\x4f\x09\xf4\xd7\xc6\x27\x7d\xde\x77\x69\x1c\x35\xde\xdc\x17\x3b\x60\x61\xf8\x67\x19\xc9\x31\x72\xff\x16\x39\xfc\x1e\x25\xa8\x9b\x72\x4b\x46\xb6\x34\x19\xc0\x43\xc1\x4b\x19\xdd\xce\x6e\xa7\x57\xb3\x9b\xa9\xbd\x63\x9c\xbf\xfa\x3b\x24\x3b\x1e\x1d\xbd\x7a\xf4\x9c\x8f\xd3\xb9\xba\x6e\x7c\x39\x8a\xfe\x17\x25\x95\x8c\x4e\x36\xa7\xd1\x49\x7a\x7c\x72\xfc\x0f\xb7\x88\xba\x95\x2c\x7e\x7d\xfb\xb2\xc9\x67\x76\x72\xff\xc3\x71\x00\xef\xb1\x45\x46\x98\x81\x80\x93\x4e\x76\xa4\x4c\x23\x73\xfd\xc1\x8d\xea\xc7\x96\x81\x06\x4e\x46\x59\x67\x17\x65\x57\xaf\x7f\xc3\xcc\x2d\x4b\xd9\x36\xe3\x2e\xa9\x45\xb5\x46\x34\xb0\x6b\x3f\x84\x2c\xf1\x62\x32\x07\xa5\xa7\xe8\xb7\xfc\xa7\x0a\x15\x88\x7f\x5b\xe1\xff\x39\x08\xc4\x1d\x6f\xcb\x64\xa7\xee\x7b\x3e\x51\xb5\x04\x66\x68\xd0\x3c\xa5\x5e\x55\x49\x9b\xfa\xd8\xdc\x05\x8b\x4f\xc0\x77\x73\xd4\x6f\x10\x57\x3c\xdf\x2e\x64\x4a\x79\xf8\xc3\x44\x4d\x0c\x97\xbf\x24\x92\xdc\x7a\x9d\xea\xfc\xdc\x01\x61\x72\xb7\x48\x08\x83\x77\x44\xc7\x33\x9e\x0c\x6d\xd4\x21\xf7\x19\xef\x49\x65\x0f\x3c\x6b\x18\x78\xd8\x6f\xbc\x6b\x44\x3f\x81\xa5\x32\x3d\x44\x53\x9d\x06\xb6\x42\xf8\x4e\x0b\x8a\x3d\x2b\x50\x22\xc5\xdb\xa7\x59\x77\xaa\x6e\xd6\x2b\x6c\xa6\xaa\x25\x87\x45\xfd\xa0\x48\x5d\xac\xf2\x74\xc9\x67\x79\x4a\xef\x69\x5a\xa1\x49\x1d\xf6\xea\xd0\xfb\xac\x61\x4e\x7f\x40\xf5\xa0\x92\xf6\x28\xef\xb8\x4d\x1b\x46\xf5\x0f\x58\xa2\x5b\x8e\x49\xb3\x84\xa8\x2e\x98\xb3\xfd\x6c\xd3\xa9\x8c\x1d\xb2\x93\x57\x0f\x04\x42\x7e\x40\xb6\x95\x0c\x08\xfe\x00\x0a\xa2\x4f\xb3\x42\xee\xed\xd9\xfc\x35\xa6\xc3\xe8\xb3\x4b\xbf\x62\xfd\xc1\x01\x0f\xd2\x3c\x77\x35\x5e\x1a\x49\x85\xc5\x30\x8b\x13\xf4\x7a\x59\x65\x62\x2f\x24\x64\x22\x6e\x94\xcb\x58\x67\xaf\x7f\x9e\xb0\x7d\x39\xd6\xfe\x75\x6f\x15\xb6\xbf\xf8\xf8\xd2\x53\x3a\xcf\x30\xa7\x08\x6c\x7f\x0c\x81\xad\xbe\x21\x87\xc1\xe6\xc0\x5a\xd0\x49\x8d\x6f\x3a\xc7\x57\x71\xd7\x2b\x6d\xed\xf4\xd7\x2a\xf6\x8c\x61\x1b\x26\x6d\x78\xa8\xf7\xc5\x16\x9d\x85\x3b\x84\xb0\x2d\x58\x04\x8b\x6a\x83\xed\x4b\xcd\x37\xf1\x32\xfb\x8c\xf2\xcd\xb7\xcb\xb0\x13\x04\x40\xe9\xff\x01\x00\x00\xff\xff\x4d\xf2\xd5\x75\x48\x1f\x00\x00")

func templatesJenkinsNormalJobXmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesJenkinsNormalJobXml,
		"templates/jenkins/normal-job.xml",
	)
}

func templatesJenkinsNormalJobXml() (*asset, error) {
	bytes, err := templatesJenkinsNormalJobXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/jenkins/normal-job.xml", size: 8008, mode: os.FileMode(420), modTime: time.Unix(1438004167, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesJenkinsPipelineXml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x53\xcb\x6e\xdb\x30\x10\xbc\xfb\x2b\x08\xa1\x57\x53\x4e\x0f\x45\x0e\x34\xd3\xc2\x6d\x80\x16\x7d\xa4\x68\xda\x6b\xc1\x48\x6b\x9a\x2d\xb5\x24\xf8\xb0\x1d\x04\xf9\xf7\xae\x1e\x4c\xa2\xf4\x01\x04\x3d\x49\x9c\x99\x9d\x9d\xd5\x52\xe2\xec\xd8\x59\xb6\x87\x10\x8d\xc3\x75\x75\xc2\x57\x15\x03\x6c\x5c\x6b\x50\xaf\xab\xaf\x97\xe7\xcb\xd3\xea\x4c\x2e\x44\x04\xde\x1a\x75\xe5\x2c\xff\x01\xf8\xd3\x60\xe4\xde\x78\xb0\x06\x81\xbf\xa6\x07\x19\x5c\x5f\x4c\xc0\x37\x03\x07\xe6\x6d\xd6\x86\x0c\xdb\x89\x5c\x16\xf9\x72\x64\x5e\xae\xf8\x29\x7f\x51\xc9\x05\x63\x02\x55\x07\xf2\xe6\x86\xf1\x8f\xf4\xc2\x6e\x6f\x45\x3d\x20\x3d\xb5\x35\x36\x41\x78\x73\x84\x26\x27\x17\xa2\xdc\x2a\x1b\x41\xd4\x8f\xe1\x7b\xe9\xe7\x0c\x19\xe6\xb2\x11\xea\x25\x3e\x38\x0f\x21\x19\x88\xac\xb1\x2a\xc6\x75\xb5\xcb\x6d\x74\xc8\x3b\x47\x39\x79\x1f\xfc\xd9\xc5\xa8\xb9\x7e\x6f\x62\xaa\xea\xa1\xac\x71\x9d\x77\x08\x98\xbe\x78\x68\x86\x66\x04\x3e\xf1\x83\x7c\x5f\x6e\x1e\xba\x8c\x26\xff\x9c\x7d\xa0\xb7\x26\xc4\xf4\xce\x5d\xcd\x25\x77\xe8\x18\xa5\xfe\xff\x2c\xa2\xfe\x7d\x48\x81\xee\xd3\xb6\xd4\x45\x79\xb2\xa2\x6c\x33\xa4\xd7\xc4\x9d\x3b\xbc\xd2\x3a\x80\x56\x09\xda\xc2\x95\x05\xfc\x85\x2d\xe6\x1b\x67\x73\x87\x64\x3d\x3a\x97\xe3\xe0\xeb\x68\x4f\xa8\x25\x52\x26\xf2\x99\x4e\x77\x1d\xf7\x2a\xa9\xfb\xeb\xf0\x10\xea\x25\xd9\xb7\xd4\xef\x2d\xd2\xf2\xf7\xca\xca\xe7\xa2\x7e\x84\x14\x9f\xcd\x4e\xa1\x86\x99\x4f\x81\x7a\x89\xb2\xd6\x1d\x3e\x28\xcc\xca\x5e\x06\xa3\x35\xfd\x25\x32\x85\x4c\xca\x3f\x31\x7d\x05\x4d\x0a\x47\x7f\x3e\xed\x27\xd2\xfd\x79\xf2\x76\xe4\xe2\x57\x00\x00\x00\xff\xff\x5f\x40\xf2\x3f\x94\x03\x00\x00")

func templatesJenkinsPipelineXmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesJenkinsPipelineXml,
		"templates/jenkins/pipeline.xml",
	)
}

func templatesJenkinsPipelineXml() (*asset, error) {
	bytes, err := templatesJenkinsPipelineXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/jenkins/pipeline.xml", size: 916, mode: os.FileMode(420), modTime: time.Unix(1432204627, 0)}
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
	"templates/jenkins/multi-job.xml":  templatesJenkinsMultiJobXml,
	"templates/jenkins/normal-job.xml": templatesJenkinsNormalJobXml,
	"templates/jenkins/pipeline.xml":   templatesJenkinsPipelineXml,
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
		"jenkins": &bintree{nil, map[string]*bintree{
			"multi-job.xml":  &bintree{templatesJenkinsMultiJobXml, map[string]*bintree{}},
			"normal-job.xml": &bintree{templatesJenkinsNormalJobXml, map[string]*bintree{}},
			"pipeline.xml":   &bintree{templatesJenkinsPipelineXml, map[string]*bintree{}},
		}},
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
