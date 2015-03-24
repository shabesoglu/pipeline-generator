package pipeline

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _templates_jenkins_normal_job_xml_swp = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9a\x5f\x8c\xdc\x46\x1d\xc7\x9d\xa4\x85\x52\x08\x94\x52\x10\x0f\x45\x5a\x96\x4a\x27\x45\xac\x7d\xb9\x82\x92\x4a\xbe\x4d\x72\x7b\x47\xb8\x34\xf7\x47\xb7\xbb\x09\xf4\x25\x9a\xb5\x67\xbd\x93\x1b\x7b\x5c\xcf\xb8\xc9\x11\xa5\xad\x8a\x78\x82\x07\x90\x40\x02\x21\x90\x40\x6a\x1f\x90\x90\x40\xb4\x6f\x20\xa1\x22\x78\xe1\x81\x27\x28\x88\x3e\xf0\xc0\xdf\x07\x5e\x90\x90\x40\x45\x2d\x5f\xdb\x63\xaf\xd7\x7b\x97\x4b\x8a\x68\x05\x9d\xef\xe9\x23\xdb\xbf\xf9\xcd\xf8\x37\xe3\xf1\x78\x66\x6e\x47\x8b\x97\xd6\x37\x5a\xa7\xec\x87\x2d\xe8\x3e\xcb\xfa\x34\x7d\xef\x70\xf8\xbb\x9d\xa3\xff\xba\xfb\x88\xb5\x4b\xa5\xa4\x7c\x24\x92\xc8\xba\xa5\x0a\xbf\x95\xad\xad\x47\x6d\x2e\x3c\xc2\x0f\xf2\x7b\x72\x5a\xa0\x73\x4d\x24\xbb\x4e\x20\x3a\x31\xf1\x76\x49\x40\xa5\x23\x13\xcf\x09\x98\x9a\xa4\x23\xdb\x13\xa1\x23\x45\x1a\xf9\x1e\x17\xa9\xef\xc4\x2c\xa6\x9c\x45\xb4\x13\xd0\x88\x26\x44\x89\xc4\x51\x34\x8c\x39\x51\xc8\x76\x95\x46\xbb\x2c\x92\x4e\x24\x92\x90\xf0\xce\x55\x31\xb2\xaf\x87\x07\x46\x60\x64\xf4\x16\x55\xaa\xc6\x9d\xd3\xc7\xad\x87\x97\x4e\x2e\x66\x97\x1f\x69\x7f\xb8\xf5\xbe\xfb\x87\x6f\x76\x54\x46\x46\x46\x46\x46\x46\x46\x46\x46\x46\x46\x6f\xa0\x54\x7c\xcc\x7a\x0a\xc7\xa3\xfa\xda\xd3\xc7\x23\x8d\xe3\x6b\x5a\x23\x7d\xed\x37\xd2\x8f\xe9\xe3\x7b\xf4\xf1\x67\x8d\x74\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\xa3\x37\x4f\x04\x0b\xf9\xe3\x38\xbe\x90\x2d\xd4\xef\x9b\xae\xff\x7f\x83\x85\xfc\x2f\xc0\x8b\xe0\xfb\xe0\x19\x30\x02\xab\xe0\x01\xf0\xa7\x77\x5b\xd6\x4b\xe0\x3b\xe0\xeb\xe0\x1c\xf8\x00\xf8\x33\x0a\x7b\x01\x3c\x0b\xbe\x00\xb6\xc0\x09\x70\x3f\xf8\xe3\xbb\x2c\xeb\xa7\xe0\xab\xe0\x29\x10\x83\x0b\xe0\x41\xf0\xf7\x77\x5a\xd6\x1f\xc0\xf3\xe0\x69\xe0\x80\x57\xef\xb5\xac\x5f\x82\xe7\xc0\x4d\xf0\x18\x58\x00\xaf\xbc\xc3\xb2\x7e\x00\xbe\x0c\x24\x78\x0c\xac\x83\x25\xf0\x21\x70\x0f\xf8\xdb\x3d\x96\xf5\x2b\xf0\x34\x78\x1c\xb4\xc0\x6b\x6f\xb7\xac\xbf\x82\xdf\x83\x97\xc0\x0f\xc1\x97\xc0\x67\xc1\x2a\x58\x02\xef\x07\xff\x78\x9b\x65\xbd\x0c\xbe\x07\xbe\x01\x1e\x07\x1e\x78\x04\x7c\x10\xbc\x72\xb7\x65\xfd\x18\x7c\x13\x7c\x11\x3c\x03\x62\xd0\x06\xff\xbc\xcb\xb2\x7e\x02\xae\x80\xe3\xe0\xb7\xc7\xe0\x07\x3e\x0f\x9e\x04\xe7\xc0\x83\xe0\x2e\xf0\xeb\xa3\x68\x5f\xf0\x1c\xf8\x1c\xf8\x14\x38\x03\x16\xc1\xbd\xe0\xe5\x23\xc5\x73\x79\xbe\xb9\x89\xd2\xca\xe5\x3a\x93\xd4\x97\x22\xb2\x63\x9e\x06\x2c\x92\x76\xc0\x94\xbd\x92\x90\xc8\x9b\xf4\x63\xea\x75\xb5\x1b\x1c\x23\x12\xd2\xee\x09\xe7\xc6\x8d\x56\x2d\x9d\x8d\x19\x4d\x5a\x37\x6f\xba\x4e\x9e\x5c\x16\x7a\x1b\x65\xba\xa3\xdc\x40\xa5\xbe\x74\x52\x49\x93\x1d\x1a\x0a\x45\x7b\x22\x1a\xb3\x40\x76\x6f\x15\xe2\xb0\xe1\x5d\x0b\x34\x4d\x78\x37\x8b\xf2\x3c\x53\xc3\x9d\x8b\x79\x70\x99\x69\xea\x90\xd0\xb1\xcc\xe2\x70\x9d\xf2\xac\x51\xcb\xc3\x6b\xb3\xff\xed\xdd\x03\xea\xe0\x7a\xf9\xd5\x25\x9a\x48\x26\xa2\xee\x92\xeb\xcc\x1a\xe0\xe4\x4a\x2f\x6c\x79\x9c\x48\xb9\xdc\xde\xe7\x7e\xa8\x4b\xbf\xb7\xd1\x6e\x15\xb6\xe5\x36\x6c\x67\x97\xec\x25\xfb\x91\x76\x96\x19\xb5\xa5\x5c\x52\x54\x35\x2b\xc9\x19\xa5\x8c\xfb\x97\x13\x12\xc7\x28\xbf\x6c\x5e\x49\x6d\x9f\x91\x91\xe0\xb6\xfe\xb5\x8b\x5d\xfe\x24\xc6\xde\xd6\x27\x3a\x1c\x84\xae\x12\x36\x4a\x95\x48\xaa\x26\x48\x63\x9f\x28\xba\xca\x64\xcc\xc9\xde\x66\xd6\x38\x2a\x49\x29\x1a\x76\xce\x5e\xe6\x78\xa2\x28\x6c\xa0\x7f\x64\xd3\x7d\xe8\xc6\xca\x70\xfd\xe2\xea\x95\xcd\xe1\xc6\xca\xda\xce\xcd\xce\x43\x37\xce\xaf\x0f\xae\xec\xac\x5d\x5a\xef\xaf\x6f\x6d\x7e\x94\xd3\x28\x50\x93\xe5\x53\x78\x58\xcd\x9c\x45\x05\x5e\x57\xfc\x55\x7b\xf9\x70\x40\xb9\x7b\x9d\xea\x77\x40\x45\xca\xd9\x45\xfb\xb4\x7d\xaa\x5d\x36\x52\xa3\xe5\x49\x24\x99\x27\xb8\x48\xec\x73\x38\xeb\x65\x67\x2b\xb5\xb6\xad\xea\x9a\xfb\x6c\x90\x38\x6f\x80\xeb\x8a\x26\x61\xf6\x84\x6b\x36\x6b\xbf\x7e\x74\x48\xe9\x55\xec\x95\x1f\x82\xfd\x98\xbd\x98\x07\xeb\xce\x3d\x63\xd7\x41\xa5\x83\xa0\x7a\xe2\x59\x9f\x88\xfc\xa2\x4b\xd4\xaa\x56\x3a\xd9\xe8\x4e\x83\xe2\xbc\xaa\x06\x0b\x22\x91\xd0\x6d\x21\x55\x4f\x84\x21\x53\x9f\x14\x62\x57\x76\xc7\x04\x5d\xcb\x75\xf6\x4f\x2c\xb3\xe6\x6f\xd1\x89\x96\xfe\x43\x6f\x9b\xbe\xe7\x87\xdd\x18\x81\xb2\x71\xcb\xd6\x36\xea\x6f\x90\x28\x25\x9c\xef\x21\xf4\x99\x7e\x3d\x5f\xa5\xb2\x23\x54\x45\xef\xd0\xac\xef\xd0\xbc\x15\x9b\x95\x73\xd4\x24\xa1\x72\x22\xb8\x5f\x7b\xdb\x3d\x81\x3e\x46\x55\x91\x43\x77\xe9\x59\x5b\xdd\x17\x8f\xa0\xbb\x72\x71\xb8\xa6\x1f\x6e\x2d\x4d\x24\x3e\x8b\x08\xef\x2e\xba\x4e\x79\xda\x18\x52\xfa\xc3\x5e\x6f\xad\xdf\x6f\x8c\x2c\x73\x31\xe1\x45\x93\x2a\xa1\x24\xdc\x4e\xc4\x55\xea\x29\x99\x8f\x64\x43\x6d\xbc\x20\x46\xb2\x18\xcf\x9a\x5e\xf5\xe7\xe0\xe8\x86\xbf\xfd\xe6\xd1\x8f\xa0\x71\x9b\xac\x53\xd5\xfb\x14\x9e\xaa\x17\xea\xb2\x29\x7a\x79\x94\xbd\x6c\xb2\xbc\x9b\x4c\x47\xa1\xf0\x53\x4e\x7b\xe3\xa0\x1c\xc8\x38\x93\xaa\x5d\x3a\xf8\xe2\x7c\xf1\xd3\x3b\xda\xaf\x5c\xf3\x51\x30\x85\x2d\x2b\xa9\xec\x69\x87\x3b\xea\x3b\x7a\x61\xf6\x72\x95\xe5\xb3\xc8\xe3\xa9\x4f\xfd\x1d\x1a\xe4\x4e\xba\x52\x97\x45\x82\x56\x08\x56\x59\x52\x74\xa8\x39\x43\xd1\xa5\xd0\xc1\x1b\xf9\xf5\x4d\x76\x59\x3c\x20\x81\xee\x1b\xe5\x55\x91\x86\x41\xb8\x08\x6c\x2d\x24\x8c\x3b\x4d\x6b\x3d\x38\x7a\xbd\x28\x3c\xfb\x6e\xc8\xa6\x51\xdf\xb1\x34\xe3\xa3\x84\xf7\x20\xf2\xe8\xd4\x80\x71\x10\xc3\xd7\x80\x24\x01\x55\x08\xfc\xb6\xdb\x1c\xb1\x0c\x84\xe0\xdd\x55\x3a\x26\x29\x57\xae\x53\x1a\x8a\xe4\x7c\x10\xe9\x4d\x84\x40\x54\xb7\xf8\xf8\xa4\x8a\x71\x5b\x17\xb1\x52\xcb\x52\xdd\x06\x9f\xbd\xfe\x04\xef\xac\xb8\xd6\xe3\x22\xaa\xbe\x0d\x0d\xab\x55\x1b\x64\x36\x85\x62\xe3\xbd\x62\x24\x99\x1d\x61\x66\x52\xca\xfa\x67\x5f\xd4\x6d\xc1\x79\xe9\x5a\xb3\x14\x2e\x71\x92\x46\x74\xa5\x9c\x51\x68\xaf\x59\x63\xe1\x78\x0d\xe3\xff\x56\xaa\xb2\x4e\x20\x63\xe2\xd1\xd2\x77\xce\x5e\xb8\x7b\x9c\x92\x28\x7f\x07\x7b\xd9\x59\x95\x9c\xf7\x98\x22\xb1\x70\x24\xa9\x9a\x88\x64\x2b\x29\x02\xc7\x27\xa0\x2c\x78\x3e\xe1\xbf\xf4\x3a\x24\xd4\x4b\xf1\xfd\x7b\x62\xea\x26\xf5\x83\xd8\x2f\x45\x07\xc1\x24\x19\xf1\xba\xbd\xbc\xe9\x5c\x42\x91\xc1\x99\x9d\xb5\xbd\x05\xe7\x8f\xff\x13\xb3\xc1\x6c\xe0\x5b\x97\xeb\x11\x53\x8c\x70\x0c\xe7\xe5\xcc\x30\x4e\x04\xe6\x0b\x8a\xd1\x3b\x99\x16\x6e\x17\x99\xf6\xa6\x1f\x19\x45\x02\x9a\xcf\x6b\xb2\x66\xea\x97\x57\x79\x4b\x4d\xd3\x4a\x6f\x45\xe4\x6e\xe5\x3c\xd0\x17\xb9\x6f\x95\xa2\x47\xb3\xdb\x0f\xe5\x4e\xe6\x75\xee\x6c\xa5\xdd\x5d\x4a\xe3\x55\x1a\x63\xd8\xc7\x18\xcb\xa6\x5d\x7e\xce\x9e\xb7\x18\x17\xc1\x8e\x50\xa4\x9a\x0b\xbb\x04\x25\x8d\x89\xa7\x36\xd3\x70\x20\x1e\x45\x9e\xee\x49\xbc\xe4\x73\xc6\x59\xe7\x55\xb2\x27\x75\x42\xa7\xe6\x5e\x33\x17\xfe\x51\x95\x3f\x73\x8b\x1a\xa5\xf9\xb3\xa5\xf8\x33\xb9\xdd\x69\xa4\x8d\x9e\x93\x35\xb3\xb4\x2f\x56\xc9\x45\xab\xf8\x54\x7a\x09\x8b\xb3\x11\x24\x1f\xc6\x5d\xc4\xa3\xbf\x42\x59\x8b\x65\xd3\x0a\x9c\x9d\xb9\x1e\xf2\x96\x9e\x90\x2f\x2f\x9c\xb4\x17\x17\xf0\xbd\xf4\x04\xa6\x38\xc1\xf2\xc2\x70\xf0\x89\xce\xe9\x85\x33\xdd\x7c\xfd\x7f\x01\x6b\xf3\x6f\x1f\x2f\xd6\xff\xe5\xff\xef\x9f\xc5\xc9\x57\xc0\x00\x9c\x05\xaf\x62\x6d\xff\x17\xf0\x22\xf8\x1a\xb8\x04\x1e\x00\xc7\xc0\xcf\x91\xf7\x47\xe0\xbb\xba\x9c\x6f\x1d\x7f\xa3\x76\x2f\x8c\x8c\x8c\x8c\x8c\x8c\x8c\x8c\x8c\x8c\x8c\x8c\xfe\x9f\x94\x2f\xfd\x8b\x35\x5d\xbe\x0f\x90\x8e\x38\x93\x93\xc3\xfe\x59\x50\xee\x39\xc4\x24\xc1\x32\x5d\xd1\x84\x7d\x86\xfa\x7a\x4b\xd8\xde\x7f\x97\xdd\x9b\xdd\x7b\x79\x3d\x85\x35\x37\x64\xa6\xdb\xd0\x97\x99\x9a\x6c\x8a\xed\xb2\x80\x6a\xe9\x7e\x50\x72\xbd\x08\x04\xe6\xb3\x6c\x85\x3b\xdd\x8a\x9f\x9a\xea\x8e\x71\x7d\xe7\x3d\xdf\xa7\x88\x1b\xbb\xec\x07\xd4\x74\xbf\xdd\xa1\x7d\x2b\xbb\x29\x7c\x3a\x8d\xd2\x69\x86\xd9\x68\xbf\xff\xac\xf9\x66\x0b\xbc\xe3\xc2\xaa\x8d\x96\x19\xcf\x8e\x76\xcd\x36\x9d\x3e\xde\xee\x5a\xff\x0e\x00\x00\xff\xff\x1f\xbd\x47\xf7\x00\x40\x00\x00")

func templates_jenkins_normal_job_xml_swp_bytes() ([]byte, error) {
	return bindata_read(
		_templates_jenkins_normal_job_xml_swp,
		"templates/jenkins/.normal-job.xml.swp",
	)
}

func templates_jenkins_normal_job_xml_swp() (*asset, error) {
	bytes, err := templates_jenkins_normal_job_xml_swp_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/jenkins/.normal-job.xml.swp", size: 16384, mode: os.FileMode(420), modTime: time.Unix(1427203435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_jenkins_multi_job_xml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\x4d\x73\xdb\x36\x10\xbd\xfb\x57\x70\x74\xf1\x49\x90\x9d\x36\x9d\x1c\x68\xe5\x43\x89\x67\xd2\x89\x15\x8d\x14\x8d\xcf\x10\xb9\x96\x60\x81\x00\x07\x00\x5b\xa9\x99\xfc\xf7\x2e\x40\x00\x04\x4d\xa5\xb5\xc6\x39\x99\xd8\x2f\x2c\xde\xbe\x7d\x72\xfe\xf6\x50\xf1\xec\x2f\x50\x9a\x49\x71\x73\x79\x4d\xae\x2e\x33\x10\x85\x2c\x99\xd8\xde\x5c\xae\xbf\xdd\x8e\xdf\x5c\xbe\x9d\x5e\xe4\x85\xac\x88\x61\x7b\xca\xc9\x23\x88\x3d\x13\x9a\xd4\xbc\xd9\xda\xbf\x55\xc3\x0d\x7b\x94\x1b\x72\x67\x3f\xfe\x94\x9b\x85\x92\x8f\x50\x98\xac\x0d\xb8\x19\xf9\x84\x71\x08\x1c\xb7\x8e\x77\xd7\xe4\xfa\xf5\x68\x7a\x91\x65\x39\x2d\x0c\xde\xae\x27\xee\x50\x82\x2e\x14\xab\xad\x65\x9a\x4f\xd2\x93\xf5\xee\x01\xea\x8f\x50\x83\x28\xb1\x4b\x06\x7a\xfa\x40\xb9\x86\x7c\x32\xb0\xdb\xe0\x5a\xc9\x1a\x94\xf1\x47\x34\x68\x20\x25\xa3\x1b\x99\xbc\x82\xd5\xc0\x99\x00\xb2\xf0\x1f\x8b\x36\xe7\x18\xdb\x2f\xd1\x8c\xf8\x1c\xc7\x21\x34\xf4\x7f\x45\xde\x90\x3f\x46\x6d\x65\xac\x6d\xa8\xde\xcf\x69\x05\xd3\x9a\x2a\xca\x39\xf0\x0c\x0e\x50\x34\xb6\xf3\x7c\x12\x9d\x21\x5a\x1b\xba\x05\x67\xf9\xfe\x3d\x23\xab\x70\xca\x7e\xfc\xc8\x27\x9d\xaf\xed\x7a\x72\x46\xdb\xee\xdd\x93\xfe\xc3\x73\x5d\x54\x59\xc1\xa9\xd6\x37\xa3\x5d\x53\x6a\x29\x08\x5a\xc8\xbc\xe1\x7c\x35\xbb\x1b\xb5\xb0\x17\x54\x2c\x25\xad\xa6\x46\x35\x08\x67\x38\xb9\x81\x30\x4d\x37\x1c\xca\x00\x75\x3c\x5b\xe7\x86\xcb\x62\xff\xa1\x61\xbc\xbc\xdf\x81\xf8\x28\xff\x16\xda\x28\xa0\x95\x33\x21\x87\x42\xd2\xff\xc6\x0d\x8b\xad\xeb\xe7\x94\x1a\x44\xd9\x42\x46\xb1\xed\x16\x39\xed\x9f\x26\x45\xd1\x28\x05\xc2\xb8\xa0\x50\xe7\xa9\xd9\x75\x60\xbf\x30\xd1\x23\x7f\x06\xeb\x3f\xb4\x99\x71\xc0\xf5\x8e\xea\x76\x88\xee\xeb\x0a\x87\x12\x2d\xbd\x18\xcc\xd5\xc1\x92\x65\xc8\x06\x45\xc5\x16\x90\x13\xcd\xc6\xba\x90\x11\xd1\xf9\x9c\x7e\x16\xa1\xe6\x4c\x8a\x07\xb6\xed\x2a\x63\x3a\xfa\x23\xe5\x1c\xd3\x82\x21\x0d\xb2\x98\x2c\x90\xc2\x95\x0e\x5c\xe8\x0c\x69\x1c\x1c\x6a\xa9\xa1\x44\x06\x05\x40\x13\x4b\x1a\xe8\xf9\x82\x3d\x3d\x61\x90\xb5\xa4\x81\xb8\x38\x1a\x27\xb8\x6c\x38\xe8\x05\x35\x3b\x5c\xff\x81\x29\x8d\xaf\xe8\x61\x09\x38\x6a\x64\x39\xc2\x9b\x9c\x7a\x6d\x0a\x7b\x95\xf5\x1c\x57\x46\x51\x03\xdb\x63\xec\xf7\x84\x6b\x98\x8a\x38\x96\xcc\xe9\x4f\x2f\xad\x33\xa7\x29\xb8\xa4\xca\xbc\xe7\xdc\x3e\xad\x45\x2f\xb5\xf4\x60\x8e\xf9\x8e\x89\x27\x6a\xed\x19\xe7\x6e\x9a\x5f\x05\x26\x2f\x41\xe3\x88\xbb\x5b\x6f\xdf\x7f\xfe\xb2\x5e\x7e\x42\xe5\xfb\xcf\xb0\xb4\xa0\x23\xf7\x57\xc1\x8f\x9f\x1f\x70\x46\xb3\x9d\xe5\x59\x14\xd0\xd3\xce\x8e\x7a\x93\x97\x70\x0f\x09\x87\xca\xdc\x51\xd9\x2f\x43\x4a\x7d\x0b\x88\x61\xa2\xa1\xb6\xed\xae\xff\xd5\x7a\x36\xfb\xb4\x5a\xdd\xae\xbf\x38\x98\x4e\x44\x5c\x3c\xbb\xbd\x13\xab\xea\xdf\xed\x37\x3e\xaf\x9b\x0d\x67\x7a\x17\x05\x00\xfb\x66\x0f\x19\x99\xc3\xc1\xdc\x51\xbc\x99\xa7\x1b\x99\xd3\x86\xd8\x5b\x0b\x54\x10\xd5\x54\xfa\xa8\x0d\x54\x9a\x78\x85\x6d\x6f\x27\xae\x7c\x54\x6b\xaf\x4b\xc4\x75\x10\xa4\xfb\x5b\x6b\x8c\x3f\x38\x2e\x65\xf0\x6b\x73\x4d\x7e\x27\xbf\x8d\x52\xb4\x10\xe1\x56\xe0\xda\x2d\x8b\x92\xea\x7f\x7f\xed\x62\x6b\xb7\xea\x83\xf6\x71\x03\x4f\x47\x7b\x2c\x7f\xed\xc3\x22\x94\x09\x05\x3a\x99\xb3\xcd\xf5\x50\xed\x5d\x83\x53\xb4\xca\x03\x06\x14\xfb\x07\xca\xde\x35\x4f\x71\xeb\x45\x8e\x7d\xe8\xbb\x57\xe4\xd5\xeb\x01\x6c\x09\xaf\xcf\xbd\xee\x84\xac\x0e\x8a\x3e\xbb\xf0\x5c\x96\xb0\x08\x8e\x6e\x98\x81\xd1\x83\xb2\xf6\x3f\x1a\x3b\x2d\xdd\x49\x78\xb4\x9c\x16\x17\xbf\x3f\x3f\xd3\x18\xdf\xc8\x3d\x33\xbb\xb9\xec\x3a\x09\xa2\xf0\x33\x77\x22\x0b\x2f\xc3\xef\xc9\x2b\xcf\x2f\x37\x24\x17\x42\xd2\xdb\xe2\x56\xf6\xee\x15\xad\xeb\x16\xe3\xb3\xb4\xc2\xaf\xc7\xf4\xe2\xdf\x00\x00\x00\xff\xff\x09\x99\x6b\x44\x29\x0b\x00\x00")

func templates_jenkins_multi_job_xml_bytes() ([]byte, error) {
	return bindata_read(
		_templates_jenkins_multi_job_xml,
		"templates/jenkins/multi-job.xml",
	)
}

func templates_jenkins_multi_job_xml() (*asset, error) {
	bytes, err := templates_jenkins_multi_job_xml_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/jenkins/multi-job.xml", size: 2857, mode: os.FileMode(420), modTime: time.Unix(1423172303, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_jenkins_normal_job_xml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x58\xdd\x6f\x1b\xb9\x11\x7f\xd7\x5f\xb1\x70\x03\xb8\x3d\x54\xab\xd8\xd7\xe2\xae\x80\xac\x4b\x2c\x2b\x39\xb5\xb6\x62\x58\x52\xf2\x78\xa0\x76\x47\x12\xcf\xdc\xe5\x82\xe4\x3a\x56\xdd\xfc\xef\x1d\x2e\x3f\x96\xfb\x61\xc7\xc6\xdd\x4b\x80\x40\x2f\xda\x99\xe1\x70\x38\x9f\x3f\x72\xfc\xcb\x7d\xc6\xa2\x3b\x10\x92\xf2\xfc\xec\xf8\x24\x7e\x7d\x1c\x41\x9e\xf0\x94\xe6\xbb\xb3\xe3\xf5\xea\xdd\xf0\xe7\xe3\x5f\x26\x83\x71\x21\xf8\xef\x90\xa8\xc9\x20\x8a\xc6\x24\x51\x28\x2c\x47\xd5\x47\x0a\x32\x11\xb4\xd0\x14\x43\x60\x7c\x77\xc3\x15\x51\x5c\x44\x09\x23\x52\x9e\x1d\xed\xcb\x54\xf2\x3c\x56\x44\xde\xca\xf8\xd2\xb3\x8f\xb4\xb8\xd6\x40\x0e\x72\xc5\xff\x03\x50\x4c\x86\x27\xe3\x51\xf0\x69\xf8\x79\x99\x05\xec\xfa\xcb\x70\x89\x50\x74\x8b\x16\x5d\x34\xb5\xf4\x90\x9b\xf2\x0b\xaf\x27\x90\x5e\x84\xca\xc7\xa3\xfa\x28\xd5\xf7\x2d\x32\x2e\xa0\x80\x3c\x45\x0f\x51\x90\x93\x2d\x61\x12\xc6\xa3\x0e\x5d\x0b\xa3\xc3\x0a\x40\xb5\xe6\x13\x09\x12\xe2\x94\x92\x0d\x67\xf1\xef\x90\xdf\xd2\x5c\xc6\x05\x2d\x80\xd1\x1c\xe2\x6b\xfb\xe7\xda\xac\x39\x44\x05\x2b\x77\x34\x3f\x3b\x4a\x91\x8c\xb1\x39\x0c\x9d\xe8\xd0\x70\xde\xbc\x8e\x7f\x8e\x7f\xb2\x0e\x44\xdd\xda\xb5\x0b\x92\xc1\xe4\xe1\x21\x8a\x57\xf6\x23\xfa\xf2\x65\x3c\xf2\x1c\x27\x2a\x15\xd9\x81\x97\x5d\xba\xaf\x4a\xb8\xe6\x19\x93\x47\x2f\xb0\xd9\x78\xac\x79\x6a\xdc\x81\x6e\xa3\x78\x2e\xe7\x39\x55\x94\xb0\x7f\xf3\x0d\xee\xa3\x05\x65\x92\xb5\xb2\xc3\x1c\x4c\xc6\x3b\xaa\xe2\xf7\x54\x2d\xa7\x57\x47\xde\x0d\x48\x7b\x73\x1a\x9f\xc6\xff\x72\x29\x93\xf0\x7c\x4b\x77\x1f\x4d\xd6\x4e\x4e\xc7\xa3\x26\xc1\x08\x95\x12\xc4\x0d\x64\x5c\xc1\xb4\xe2\x4a\xef\x83\x9e\x2d\xd7\x2d\x61\x27\xab\xf5\x08\x56\xf9\x0a\xad\x5a\xdf\x5c\x56\x8e\xd2\x24\xa7\x6c\xf4\x6c\x6d\xb8\xae\xdf\xa4\xf1\x46\x90\x3c\xd9\xc3\x93\x16\x9e\x57\x32\xcb\x02\x92\xc0\xb6\x5c\x07\xeb\x87\x91\x36\xaf\xe6\xd3\x2d\x05\x51\xd9\x99\x87\x91\xef\x33\xb4\xad\x74\x3c\x6a\x9a\x32\x4e\xa9\x24\x1b\x06\xcb\x72\x93\xf1\xb4\x64\x75\xd2\x77\x19\x66\x81\x80\xa4\xc4\x28\xdc\x85\x1c\x25\x4a\x5c\xd1\xc7\xb1\x9b\xf0\xf7\x90\x83\x20\xaa\x66\x19\xff\x94\x48\xd3\xad\xc6\x6f\xfa\x55\x41\x5b\xe4\xa5\xda\x73\xf1\x41\x4c\x79\x96\x51\xa5\x40\x38\x05\x5d\x86\x4d\x28\x06\x24\xaf\xa2\x3c\xd5\xff\x3e\x71\x71\x2b\x0b\x92\x98\xb2\x30\x4c\x23\xf8\x19\xd3\xfe\x43\xa9\xbc\x80\x53\xdc\xa1\x1b\xf1\x42\x94\x39\x9c\x3b\x97\x5a\xd9\x26\xd1\xb9\x4d\x67\xc5\x35\x67\xcc\x49\x05\x14\x23\x42\x77\x39\x17\xb0\xe0\xd8\xaa\x0e\xc6\x7e\x27\xda\xc3\xf1\x25\xb0\xdc\x13\xc6\xf8\xe7\x29\xe3\x39\xd8\x40\xb4\xa9\x36\x07\x4b\xca\xd2\xe9\x9e\x73\xcc\xd1\x27\x6a\xb3\x54\x94\xc5\x17\xb0\x25\x25\x53\xe7\xc1\x92\xa3\x91\xd5\x83\x42\x2b\xce\xd9\xc4\xca\x8c\x47\x8e\x60\xdb\xa0\x0f\xdc\x76\xe7\xb6\x61\x54\x2a\xbf\x5e\x00\xc3\x58\xde\xc1\x8a\x88\x1d\xa8\x0b\x2a\x6a\xc6\x16\x04\x76\x58\x70\x04\xb8\x4f\x58\x99\x42\x7a\x03\x3b\x3f\x8f\x02\xb2\x2e\x42\x19\x58\x65\x32\x45\xb7\xb7\x0e\x71\x96\x11\xca\x1c\x55\xde\xd2\x62\x45\x76\xd6\x59\xee\xcb\xc6\x20\x6f\x6c\x39\xb1\x1d\x4e\x87\x1d\x07\x26\x1a\x8b\x09\xa3\xb3\xa8\x4d\xc0\xd9\x50\xa5\x52\x7b\xbd\xdd\x31\xc9\x42\xb3\xfe\xf4\x72\xf8\xaa\xcf\xe1\x5e\x41\x2e\xeb\xa1\x3e\x42\x8b\xaa\x3f\x4a\xd0\xdd\x0e\xdd\x68\xe4\xec\x69\xd7\x85\x54\x02\x48\x86\xed\x5c\x9a\x7e\x8e\x92\x6e\x40\xb8\x15\xf1\x0d\x68\x54\x01\x55\x8e\xac\x0c\xb1\x1e\x42\xd8\x70\x46\xfe\xab\xb4\xfa\xae\x0d\xc6\xa8\xbc\xda\xde\x04\xb3\xb6\x2d\xe5\x87\xdf\x5e\x80\xdc\x73\x96\xb6\xfb\xe2\x72\x3d\x9d\xce\x96\xcb\x66\x1b\x44\x26\x17\x88\x6e\x08\x9b\xbc\x1e\x8f\xdc\xdf\x9a\x99\x70\x86\xe3\xfe\xfc\x72\x3d\xd3\x53\x85\x71\xd1\xe0\x65\x05\x03\x65\xce\x64\xf3\xa3\x49\xf3\xbd\xb6\x65\xd4\x78\xf4\x7c\xff\xf8\x74\x09\x7d\x6e\x25\x20\xbd\x22\x79\x89\xb5\x7b\xb0\x79\x85\x79\xe0\x43\xe0\x70\x96\xdb\x01\x67\x68\x9f\xe3\x27\x3f\x44\xf6\x87\x71\x0e\xc6\x89\x6d\x30\xd7\x5c\x2a\xd3\x44\x7e\xe5\xfc\x56\x36\x7b\x4c\x9b\x39\x68\x4c\x96\x47\x77\x0e\x8f\x84\xbe\x09\x92\xca\xb4\x9d\x4f\x82\x14\x85\x4f\xb3\xf6\xf4\x23\x98\x9a\x55\x28\xe2\xb7\xf8\x6f\xaa\xff\x9d\x07\xab\x3c\x4c\xf0\x72\x08\x8f\xfe\x11\xbf\xae\xe1\x51\x45\xbc\x22\x45\x05\x6d\x30\xd5\x45\x66\x63\xeb\x68\x83\xde\x01\xf9\x95\x7d\x9f\x8f\xeb\x2c\x2e\xc1\xda\xc4\xa3\x6f\x4a\x0d\x8d\x5f\x8a\xf0\x2c\x42\x5f\x01\x66\x1b\xd6\xfb\xe4\xd5\xc3\xf9\x7a\x7e\x79\xf1\xdb\x62\x7d\x75\x3e\xbb\xf9\x32\x7c\xf5\xf0\x7e\xbe\xfa\xed\x66\xf6\x71\xbe\x9c\x7f\x58\xfc\x9d\x41\xbe\x53\xfb\xb3\x9f\xb0\x6e\xda\x2b\xeb\xaa\x4b\xf1\xf3\x82\x4a\x24\x1f\x2a\x2f\xd8\xe9\xd0\xa1\x3f\x1f\x0c\x76\x0f\x6a\xda\x49\x27\xc8\x8d\xd4\xfd\xa6\xd0\x60\xd5\x4b\x3a\x2d\x05\x67\x53\x55\x4a\x7a\x6c\x6f\x65\x13\xa3\x7d\xc7\x8f\xdf\xf1\xe3\x77\xfc\xf8\x1d\x3f\x46\x5d\xfc\x38\xea\x01\x80\x35\xf8\xfa\xb6\x47\xe3\x23\x5d\xdf\xc0\x80\xfa\x71\x60\xc9\xc8\x1d\x5c\x92\x0d\x30\x3b\x0c\x30\xa2\x98\x8f\x90\x2e\x78\x6a\xdf\x28\x42\x09\x2c\xc1\x90\x3f\x68\x61\x8b\x84\xe4\x37\x9c\x64\x2e\xb1\xdd\xe7\xa0\xee\x44\x69\xab\x01\xa5\xc6\xd5\x8c\x27\xb7\xe6\x04\x7b\xc8\x2f\xf8\xe7\xdc\x80\xcd\x8a\x84\x40\xde\x2d\xfa\xaa\x5c\x57\x99\x03\xb2\x4f\xab\xea\x48\x0d\xcc\x60\xc3\x96\x87\x39\x6b\x2a\xc5\x9f\xaa\x45\x1e\x3c\x81\xcc\x5f\x80\xcb\x43\x54\xfe\x87\x30\x79\x17\x91\x3f\x81\xc7\x9f\x40\xe3\x8f\x63\xf1\xe7\x22\xf1\x16\x0e\x7f\x3e\x0a\x6f\x24\x55\x95\xc7\xe1\x25\x08\x5b\xef\x0e\xa2\xf8\xad\x7b\x63\x84\xa2\x8d\xc1\x5d\x9d\x24\xbc\x38\xb8\xc7\xc5\x78\x8a\x1f\x6e\x8d\xaf\xcc\x50\xe2\xcd\x49\xfc\xe3\x69\x7c\x52\x57\xa7\x7b\x75\xd5\xce\xb7\x3e\xd6\x45\x59\x61\xb9\x2a\x00\xc1\xb3\x6c\xb5\x60\x4b\x99\x1e\x4a\x5a\xde\xef\xa4\xe5\x2c\xdd\x5f\x99\xaa\xde\x58\xdf\xc0\x6c\xcb\x93\x35\x45\x02\x43\xbd\x9d\x37\xdc\xde\x73\xf9\xbb\x49\xe5\xc5\xa5\x5d\x79\x14\xc0\x0c\x4c\x5c\xb6\x21\xc9\xed\x8a\x5f\x12\xa9\x96\x65\x92\x80\x94\xdb\x92\xd9\xc8\x3d\xca\x0e\x40\x94\xcd\xb4\x77\xd5\x41\x96\x4a\xcf\xf0\xdd\x61\x82\x3d\xfa\x3d\xe3\x1b\xc2\x96\xa0\x14\x96\x4d\x9d\x92\x2d\x41\x0f\x62\xdc\xc1\x3c\x25\xe5\x38\xf8\xde\xe1\x5a\x10\x85\xa0\xb9\x72\x7e\x0b\x40\xc3\x63\x02\xfd\xbd\xf1\xd1\x98\xf7\x5d\xef\xc6\x8d\xc7\xf1\xe5\x1e\x58\x98\xfe\x59\x46\x72\xcc\xdc\xbf\x44\x0e\x69\x47\x09\x9e\x4d\x87\x25\x23\x3b\x9a\x0c\xe0\xbe\xe0\x42\x45\xd7\xf3\xeb\xd9\xe5\x7c\x31\xb3\xb7\x81\xb3\x57\x7f\x85\x64\xcf\xa3\xa3\x57\x0f\x9e\xf3\x71\x76\xa3\x2f\x06\x5f\x8e\xa2\xff\x45\x49\xa9\xa2\xe1\xf6\x24\x1a\xa6\xc7\xc3\xe3\xbf\x39\x25\xfa\xfe\xb0\xfc\xf5\xed\xf3\x16\x9f\xda\xc5\xfd\x2f\xbc\x01\xba\xc7\x11\x19\x61\x05\x02\x2e\x1a\xee\x89\x48\x23\x73\x51\xc1\x8d\xea\x67\x91\x41\x85\x9b\xcc\x61\x9d\x5f\xb4\x5f\xfd\xf9\x1b\x6e\x6e\x79\xca\x8e\x19\x77\x9d\x2c\xca\x0d\xa2\x81\x7d\xfb\xc9\xc2\x60\x9b\x7e\xa7\x5f\xe1\xec\x0e\x46\x65\x56\x7d\x62\x29\x9e\x04\x63\x12\xb1\x27\x2d\x28\x76\xdc\xa0\x44\x52\xbc\xe5\x18\xbd\x33\x7d\x83\x5b\xe3\x28\xd0\x03\x25\x6c\x49\x4f\x8a\xd4\xa5\x96\xa7\x2b\x3e\xcf\x53\x7a\x47\x53\xbc\xdb\x3b\xbc\xdb\xa5\xf7\x39\xc2\x58\xff\xc4\xcb\x41\xd0\x07\x7a\x0e\xef\xb8\x6f\x45\xb2\xd7\x17\x51\x6f\x15\xf1\x59\xde\x69\x26\xa4\x59\x00\xba\x87\xe7\xec\x30\xdf\x76\xea\xba\x43\x76\xf2\xfa\x22\x2a\xd5\x07\x64\x5b\xc9\x80\xe0\x0d\xd0\x00\x73\x96\x15\xea\x60\x6d\xf3\x18\xbc\xc3\xe8\xf3\x4b\xff\xc1\x7a\x3d\xb4\x80\x7b\x65\x9e\x55\x1a\x2f\x5a\xa4\xc4\x52\xce\xe2\x04\xa3\x2e\xca\x4c\x1e\xa4\x82\x4c\xc6\x8d\x62\x8f\xab\xdc\xf3\xd7\x60\x3b\x55\xe2\x2a\xbe\xee\x4e\x6c\xbb\xa3\xcf\xaf\x6a\x49\xe7\xba\x7f\x82\xb0\xec\xc7\x10\x96\x55\xd7\xbb\x30\xd9\x1c\xd4\x08\xe6\x80\x89\x4d\xc7\x7c\x9d\x77\xbd\xd2\xd6\x4f\x7f\xee\xc1\x1e\x71\xac\x9d\x91\xda\xb8\x86\x57\x5b\xcd\xb2\x20\x02\x4d\xc3\x4e\x4d\xff\x0b\x69\x63\x9b\xb6\xdf\x1a\x92\x43\x2b\xaa\x1f\x02\xfe\xd9\x71\x5b\x30\x37\x5e\xba\x5d\xfb\xbe\xdf\xa7\xf4\xd9\x8a\x35\x44\xbd\x76\x8c\x3a\x98\xae\xbb\x75\xd4\xba\x79\x6f\xe2\x1a\x0e\x78\xd9\xb6\x28\xa5\xfa\xbe\x5a\x43\xaa\x9a\x14\x0a\x5a\x43\x3e\x51\xb5\x5f\xf0\xda\x12\x57\x49\x8f\xb1\x6b\xf7\xb5\x67\xdb\x0b\xfd\xd7\x3a\xe5\xcb\xd5\xf5\x3d\x1f\x86\x2d\x3e\x40\x40\xff\x0f\x00\x00\xff\xff\x29\x21\x12\x07\xca\x1e\x00\x00")

func templates_jenkins_normal_job_xml_bytes() ([]byte, error) {
	return bindata_read(
		_templates_jenkins_normal_job_xml,
		"templates/jenkins/normal-job.xml",
	)
}

func templates_jenkins_normal_job_xml() (*asset, error) {
	bytes, err := templates_jenkins_normal_job_xml_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/jenkins/normal-job.xml", size: 7882, mode: os.FileMode(420), modTime: time.Unix(1427203417, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_jenkins_pipeline_xml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x53\xcb\x6e\xdb\x30\x10\xbc\xfb\x2b\x08\xa1\x57\x53\x4e\x0f\x45\x0e\x34\xd3\xc2\x6d\x80\x16\x7d\xa4\x68\xda\x6b\xc1\x48\x6b\x9a\x2d\xb5\x24\xf8\xb0\x1d\x04\xf9\xf7\xae\x1e\x4c\xa2\xf4\x01\x04\x3d\x49\x9c\x99\x9d\x9d\xd5\x52\xe2\xec\xd8\x59\xb6\x87\x10\x8d\xc3\x75\x75\xc2\x57\x15\x03\x6c\x5c\x6b\x50\xaf\xab\xaf\x97\xe7\xcb\xd3\xea\x4c\x2e\x44\x04\xde\x1a\x75\xe5\x2c\xff\x01\xf8\xd3\x60\xe4\xde\x78\xb0\x06\x81\xbf\xa6\x07\x19\x5c\x5f\x4c\xc0\x37\x03\x07\xe6\x6d\xd6\x86\x0c\xdb\x89\x5c\x16\xf9\x72\x64\x5e\xae\xf8\x29\x7f\x51\xc9\x05\x63\x02\x55\x07\xf2\xe6\x86\xf1\x8f\xf4\xc2\x6e\x6f\x45\x3d\x20\x3d\xb5\x35\x36\x41\x78\x73\x84\x26\x27\x17\xa2\xdc\x2a\x1b\x41\xd4\x8f\xe1\x7b\xe9\xe7\x0c\x19\xe6\xb2\x11\xea\x25\x3e\x38\x0f\x21\x19\x88\xac\xb1\x2a\xc6\x75\xb5\xcb\x6d\x74\xc8\x3b\x47\x39\x79\x1f\xfc\xd9\xc5\xa8\xb9\x7e\x6f\x62\xaa\xea\xa1\xac\x71\x9d\x77\x08\x98\xbe\x78\x68\x86\x66\x04\x3e\xf1\x83\x7c\x5f\x6e\x1e\xba\x8c\x26\xff\x9c\x7d\xa0\xb7\x26\xc4\xf4\xce\x5d\xcd\x25\x77\xe8\x18\xa5\xfe\xff\x2c\xa2\xfe\x7d\x48\x81\xee\xd3\xb6\xd4\x45\x79\xb2\xa2\x6c\x33\xa4\xd7\xc4\x9d\x3b\xbc\xd2\x3a\x80\x56\x09\xda\xc2\x95\x05\xfc\x85\x2d\xe6\x1b\x67\x73\x87\x64\x3d\x3a\x97\xe3\xe0\xeb\x68\x4f\xa8\x25\x52\x26\xf2\x99\x4e\x77\x1d\xf7\x2a\xa9\xfb\xeb\xf0\x10\xea\x25\xd9\xb7\xd4\xef\x2d\xd2\xf2\xf7\xca\xca\xe7\xa2\x7e\x84\x14\x9f\xcd\x4e\xa1\x86\x99\x4f\x81\x7a\x89\xb2\xd6\x1d\x3e\x28\xcc\xca\x5e\x06\xa3\x35\xfd\x25\x32\x85\x4c\xca\x3f\x31\x7d\x05\x4d\x0a\x47\x7f\x3e\xed\x27\xd2\xfd\x79\xf2\x76\xe4\xe2\x57\x00\x00\x00\xff\xff\x5f\x40\xf2\x3f\x94\x03\x00\x00")

func templates_jenkins_pipeline_xml_bytes() ([]byte, error) {
	return bindata_read(
		_templates_jenkins_pipeline_xml,
		"templates/jenkins/pipeline.xml",
	)
}

func templates_jenkins_pipeline_xml() (*asset, error) {
	bytes, err := templates_jenkins_pipeline_xml_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/jenkins/pipeline.xml", size: 916, mode: os.FileMode(420), modTime: time.Unix(1423648470, 0)}
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
	"templates/jenkins/.normal-job.xml.swp": templates_jenkins_normal_job_xml_swp,
	"templates/jenkins/multi-job.xml":       templates_jenkins_multi_job_xml,
	"templates/jenkins/normal-job.xml":      templates_jenkins_normal_job_xml,
	"templates/jenkins/pipeline.xml":        templates_jenkins_pipeline_xml,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"templates": &_bintree_t{nil, map[string]*_bintree_t{
		"jenkins": &_bintree_t{nil, map[string]*_bintree_t{
			".normal-job.xml.swp": &_bintree_t{templates_jenkins_normal_job_xml_swp, map[string]*_bintree_t{}},
			"multi-job.xml":       &_bintree_t{templates_jenkins_multi_job_xml, map[string]*_bintree_t{}},
			"normal-job.xml":      &_bintree_t{templates_jenkins_normal_job_xml, map[string]*_bintree_t{}},
			"pipeline.xml":        &_bintree_t{templates_jenkins_pipeline_xml, map[string]*_bintree_t{}},
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
