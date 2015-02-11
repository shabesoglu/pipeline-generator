package pipeline

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
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

func templates_jenkins_multi_job_xml() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xac, 0x56,
		0x4d, 0x73, 0xdb, 0x36, 0x10, 0xbd, 0xfb, 0x57, 0x70, 0x74, 0xf1, 0x49,
		0x90, 0x9d, 0x36, 0x9d, 0x1c, 0x68, 0xe5, 0x43, 0x89, 0x67, 0xd2, 0x89,
		0x15, 0x8d, 0x14, 0x8d, 0xcf, 0x10, 0xb9, 0x96, 0x60, 0x81, 0x00, 0x07,
		0x00, 0x5b, 0xa9, 0x99, 0xfc, 0xf7, 0x2e, 0x40, 0x00, 0x04, 0x4d, 0xa5,
		0xb5, 0xc6, 0x39, 0x99, 0xd8, 0x2f, 0x2c, 0xde, 0xbe, 0x7d, 0x72, 0xfe,
		0xf6, 0x50, 0xf1, 0xec, 0x2f, 0x50, 0x9a, 0x49, 0x71, 0x73, 0x79, 0x4d,
		0xae, 0x2e, 0x33, 0x10, 0x85, 0x2c, 0x99, 0xd8, 0xde, 0x5c, 0xae, 0xbf,
		0xdd, 0x8e, 0xdf, 0x5c, 0xbe, 0x9d, 0x5e, 0xe4, 0x85, 0xac, 0x88, 0x61,
		0x7b, 0xca, 0xc9, 0x23, 0x88, 0x3d, 0x13, 0x9a, 0xd4, 0xbc, 0xd9, 0xda,
		0xbf, 0x55, 0xc3, 0x0d, 0x7b, 0x94, 0x1b, 0x72, 0x67, 0x3f, 0xfe, 0x94,
		0x9b, 0x85, 0x92, 0x8f, 0x50, 0x98, 0xac, 0x0d, 0xb8, 0x19, 0xf9, 0x84,
		0x71, 0x08, 0x1c, 0xb7, 0x8e, 0x77, 0xd7, 0xe4, 0xfa, 0xf5, 0x68, 0x7a,
		0x91, 0x65, 0x39, 0x2d, 0x0c, 0xde, 0xae, 0x27, 0xee, 0x50, 0x82, 0x2e,
		0x14, 0xab, 0xad, 0x65, 0x9a, 0x4f, 0xd2, 0x93, 0xf5, 0xee, 0x01, 0xea,
		0x8f, 0x50, 0x83, 0x28, 0xb1, 0x4b, 0x06, 0x7a, 0xfa, 0x40, 0xb9, 0x86,
		0x7c, 0x32, 0xb0, 0xdb, 0xe0, 0x5a, 0xc9, 0x1a, 0x94, 0xf1, 0x47, 0x34,
		0x68, 0x20, 0x25, 0xa3, 0x1b, 0x99, 0xbc, 0x82, 0xd5, 0xc0, 0x99, 0x00,
		0xb2, 0xf0, 0x1f, 0x8b, 0x36, 0xe7, 0x18, 0xdb, 0x2f, 0xd1, 0x8c, 0xf8,
		0x1c, 0xc7, 0x21, 0x34, 0xf4, 0x7f, 0x45, 0xde, 0x90, 0x3f, 0x46, 0x6d,
		0x65, 0xac, 0x6d, 0xa8, 0xde, 0xcf, 0x69, 0x05, 0xd3, 0x9a, 0x2a, 0xca,
		0x39, 0xf0, 0x0c, 0x0e, 0x50, 0x34, 0xb6, 0xf3, 0x7c, 0x12, 0x9d, 0x21,
		0x5a, 0x1b, 0xba, 0x05, 0x67, 0xf9, 0xfe, 0x3d, 0x23, 0xab, 0x70, 0xca,
		0x7e, 0xfc, 0xc8, 0x27, 0x9d, 0xaf, 0xed, 0x7a, 0x72, 0x46, 0xdb, 0xee,
		0xdd, 0x93, 0xfe, 0xc3, 0x73, 0x5d, 0x54, 0x59, 0xc1, 0xa9, 0xd6, 0x37,
		0xa3, 0x5d, 0x53, 0x6a, 0x29, 0x08, 0x5a, 0xc8, 0xbc, 0xe1, 0x7c, 0x35,
		0xbb, 0x1b, 0xb5, 0xb0, 0x17, 0x54, 0x2c, 0x25, 0xad, 0xa6, 0x46, 0x35,
		0x08, 0x67, 0x38, 0xb9, 0x81, 0x30, 0x4d, 0x37, 0x1c, 0xca, 0x00, 0x75,
		0x3c, 0x5b, 0xe7, 0x86, 0xcb, 0x62, 0xff, 0xa1, 0x61, 0xbc, 0xbc, 0xdf,
		0x81, 0xf8, 0x28, 0xff, 0x16, 0xda, 0x28, 0xa0, 0x95, 0x33, 0x21, 0x87,
		0x42, 0xd2, 0xff, 0xc6, 0x0d, 0x8b, 0xad, 0xeb, 0xe7, 0x94, 0x1a, 0x44,
		0xd9, 0x42, 0x46, 0xb1, 0xed, 0x16, 0x39, 0xed, 0x9f, 0x26, 0x45, 0xd1,
		0x28, 0x05, 0xc2, 0xb8, 0xa0, 0x50, 0xe7, 0xa9, 0xd9, 0x75, 0x60, 0xbf,
		0x30, 0xd1, 0x23, 0x7f, 0x06, 0xeb, 0x3f, 0xb4, 0x99, 0x71, 0xc0, 0xf5,
		0x8e, 0xea, 0x76, 0x88, 0xee, 0xeb, 0x0a, 0x87, 0x12, 0x2d, 0xbd, 0x18,
		0xcc, 0xd5, 0xc1, 0x92, 0x65, 0xc8, 0x06, 0x45, 0xc5, 0x16, 0x90, 0x13,
		0xcd, 0xc6, 0xba, 0x90, 0x11, 0xd1, 0xf9, 0x9c, 0x7e, 0x16, 0xa1, 0xe6,
		0x4c, 0x8a, 0x07, 0xb6, 0xed, 0x2a, 0x63, 0x3a, 0xfa, 0x23, 0xe5, 0x1c,
		0xd3, 0x82, 0x21, 0x0d, 0xb2, 0x98, 0x2c, 0x90, 0xc2, 0x95, 0x0e, 0x5c,
		0xe8, 0x0c, 0x69, 0x1c, 0x1c, 0x6a, 0xa9, 0xa1, 0x44, 0x06, 0x05, 0x40,
		0x13, 0x4b, 0x1a, 0xe8, 0xf9, 0x82, 0x3d, 0x3d, 0x61, 0x90, 0xb5, 0xa4,
		0x81, 0xb8, 0x38, 0x1a, 0x27, 0xb8, 0x6c, 0x38, 0xe8, 0x05, 0x35, 0x3b,
		0x5c, 0xff, 0x81, 0x29, 0x8d, 0xaf, 0xe8, 0x61, 0x09, 0x38, 0x6a, 0x64,
		0x39, 0xc2, 0x9b, 0x9c, 0x7a, 0x6d, 0x0a, 0x7b, 0x95, 0xf5, 0x1c, 0x57,
		0x46, 0x51, 0x03, 0xdb, 0x63, 0xec, 0xf7, 0x84, 0x6b, 0x98, 0x8a, 0x38,
		0x96, 0xcc, 0xe9, 0x4f, 0x2f, 0xad, 0x33, 0xa7, 0x29, 0xb8, 0xa4, 0xca,
		0xbc, 0xe7, 0xdc, 0x3e, 0xad, 0x45, 0x2f, 0xb5, 0xf4, 0x60, 0x8e, 0xf9,
		0x8e, 0x89, 0x27, 0x6a, 0xed, 0x19, 0xe7, 0x6e, 0x9a, 0x5f, 0x05, 0x26,
		0x2f, 0x41, 0xe3, 0x88, 0xbb, 0x5b, 0x6f, 0xdf, 0x7f, 0xfe, 0xb2, 0x5e,
		0x7e, 0x42, 0xe5, 0xfb, 0xcf, 0xb0, 0xb4, 0xa0, 0x23, 0xf7, 0x57, 0xc1,
		0x8f, 0x9f, 0x1f, 0x70, 0x46, 0xb3, 0x9d, 0xe5, 0x59, 0x14, 0xd0, 0xd3,
		0xce, 0x8e, 0x7a, 0x93, 0x97, 0x70, 0x0f, 0x09, 0x87, 0xca, 0xdc, 0x51,
		0xd9, 0x2f, 0x43, 0x4a, 0x7d, 0x0b, 0x88, 0x61, 0xa2, 0xa1, 0xb6, 0xed,
		0xae, 0xff, 0xd5, 0x7a, 0x36, 0xfb, 0xb4, 0x5a, 0xdd, 0xae, 0xbf, 0x38,
		0x98, 0x4e, 0x44, 0x5c, 0x3c, 0xbb, 0xbd, 0x13, 0xab, 0xea, 0xdf, 0xed,
		0x37, 0x3e, 0xaf, 0x9b, 0x0d, 0x67, 0x7a, 0x17, 0x05, 0x00, 0xfb, 0x66,
		0x0f, 0x19, 0x99, 0xc3, 0xc1, 0xdc, 0x51, 0xbc, 0x99, 0xa7, 0x1b, 0x99,
		0xd3, 0x86, 0xd8, 0x5b, 0x0b, 0x54, 0x10, 0xd5, 0x54, 0xfa, 0xa8, 0x0d,
		0x54, 0x9a, 0x78, 0x85, 0x6d, 0x6f, 0x27, 0xae, 0x7c, 0x54, 0x6b, 0xaf,
		0x4b, 0xc4, 0x75, 0x10, 0xa4, 0xfb, 0x5b, 0x6b, 0x8c, 0x3f, 0x38, 0x2e,
		0x65, 0xf0, 0x6b, 0x73, 0x4d, 0x7e, 0x27, 0xbf, 0x8d, 0x52, 0xb4, 0x10,
		0xe1, 0x56, 0xe0, 0xda, 0x2d, 0x8b, 0x92, 0xea, 0x7f, 0x7f, 0xed, 0x62,
		0x6b, 0xb7, 0xea, 0x83, 0xf6, 0x71, 0x03, 0x4f, 0x47, 0x7b, 0x2c, 0x7f,
		0xed, 0xc3, 0x22, 0x94, 0x09, 0x05, 0x3a, 0x99, 0xb3, 0xcd, 0xf5, 0x50,
		0xed, 0x5d, 0x83, 0x53, 0xb4, 0xca, 0x03, 0x06, 0x14, 0xfb, 0x07, 0xca,
		0xde, 0x35, 0x4f, 0x71, 0xeb, 0x45, 0x8e, 0x7d, 0xe8, 0xbb, 0x57, 0xe4,
		0xd5, 0xeb, 0x01, 0x6c, 0x09, 0xaf, 0xcf, 0xbd, 0xee, 0x84, 0xac, 0x0e,
		0x8a, 0x3e, 0xbb, 0xf0, 0x5c, 0x96, 0xb0, 0x08, 0x8e, 0x6e, 0x98, 0x81,
		0xd1, 0x83, 0xb2, 0xf6, 0x3f, 0x1a, 0x3b, 0x2d, 0xdd, 0x49, 0x78, 0xb4,
		0x9c, 0x16, 0x17, 0xbf, 0x3f, 0x3f, 0xd3, 0x18, 0xdf, 0xc8, 0x3d, 0x33,
		0xbb, 0xb9, 0xec, 0x3a, 0x09, 0xa2, 0xf0, 0x33, 0x77, 0x22, 0x0b, 0x2f,
		0xc3, 0xef, 0xc9, 0x2b, 0xcf, 0x2f, 0x37, 0x24, 0x17, 0x42, 0xd2, 0xdb,
		0xe2, 0x56, 0xf6, 0xee, 0x15, 0xad, 0xeb, 0x16, 0xe3, 0xb3, 0xb4, 0xc2,
		0xaf, 0xc7, 0xf4, 0xe2, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x99,
		0x6b, 0x44, 0x29, 0x0b, 0x00, 0x00,
	},
		"templates/jenkins/multi-job.xml",
	)
}

func templates_jenkins_normal_job_xml() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xec, 0x58,
		0xdd, 0x6f, 0xdb, 0x38, 0x12, 0x7f, 0xf7, 0x5f, 0x21, 0xe4, 0x0a, 0xe4,
		0x6e, 0x71, 0x96, 0x9b, 0xec, 0x1d, 0x76, 0x0f, 0x70, 0xbc, 0x6d, 0x1c,
		0xb7, 0xeb, 0x43, 0x92, 0x06, 0xb1, 0xdd, 0x3e, 0x2e, 0x68, 0x69, 0x6c,
		0x73, 0x4d, 0x89, 0x02, 0x49, 0xa5, 0xf1, 0xe5, 0xfa, 0xbf, 0xef, 0x50,
		0xfc, 0x10, 0xf5, 0x91, 0x34, 0x01, 0xf6, 0x65, 0x81, 0xc0, 0x2f, 0xe6,
		0xcc, 0x90, 0x9c, 0xef, 0xf9, 0x51, 0xe3, 0x5f, 0xee, 0x33, 0x16, 0xdd,
		0x81, 0x90, 0x94, 0xe7, 0x67, 0xc7, 0x27, 0xf1, 0xdb, 0xe3, 0x08, 0xf2,
		0x84, 0xa7, 0x34, 0xdf, 0x9e, 0x1d, 0xaf, 0x96, 0x1f, 0x86, 0x3f, 0x1f,
		0xff, 0x32, 0x19, 0x8c, 0x0b, 0xc1, 0x7f, 0x87, 0x44, 0x4d, 0x06, 0x51,
		0x34, 0x26, 0x89, 0x42, 0x61, 0x39, 0xaa, 0x16, 0x29, 0xc8, 0x44, 0xd0,
		0x42, 0x53, 0x0c, 0x61, 0x0f, 0x50, 0x5c, 0x40, 0x01, 0x79, 0x8a, 0xe7,
		0x50, 0x90, 0x93, 0x0d, 0x61, 0x12, 0xc6, 0xa3, 0x0e, 0x5d, 0x0b, 0xe3,
		0xb1, 0x05, 0x08, 0x65, 0x97, 0x48, 0x90, 0x10, 0xa7, 0x94, 0xac, 0x39,
		0x8b, 0x7f, 0x87, 0x7c, 0x4f, 0x73, 0x19, 0x17, 0xb4, 0x00, 0x46, 0x73,
		0x88, 0x6f, 0xec, 0x9f, 0x1b, 0xb3, 0xe7, 0x10, 0x15, 0xac, 0xdc, 0xd2,
		0xfc, 0xec, 0x28, 0x45, 0x32, 0x5a, 0x70, 0x18, 0x3a, 0xd1, 0xa1, 0xe1,
		0xbc, 0x7b, 0x1b, 0xff, 0x1c, 0xff, 0x74, 0x64, 0x4e, 0xc6, 0xb3, 0x15,
		0x91, 0xfb, 0x6b, 0x92, 0xc1, 0xe4, 0xe1, 0x21, 0x8a, 0x97, 0x76, 0x11,
		0x7d, 0xfb, 0x36, 0x1e, 0x79, 0x8e, 0x13, 0x95, 0x8a, 0x6c, 0xc1, 0xcb,
		0x2e, 0xdc, 0xaa, 0x12, 0xae, 0x79, 0x46, 0xe5, 0xd1, 0x0b, 0x74, 0xae,
		0x8c, 0x1e, 0x35, 0xad, 0xc6, 0x1b, 0xe8, 0x26, 0x8a, 0xe7, 0x72, 0x9e,
		0x53, 0x45, 0x09, 0xfb, 0x2f, 0x5f, 0xe3, 0x3d, 0x5a, 0x50, 0x26, 0x59,
		0x94, 0x30, 0x22, 0xe5, 0xd9, 0xd1, 0xae, 0x4c, 0x25, 0xcf, 0x63, 0x63,
		0x98, 0x8c, 0xb7, 0x54, 0xc5, 0x1f, 0xa9, 0x5a, 0x4c, 0xaf, 0x8e, 0xbc,
		0x1b, 0x90, 0xf6, 0xee, 0x34, 0x3e, 0x8d, 0xff, 0x63, 0x2d, 0x1e, 0x27,
		0x3c, 0xdf, 0xd0, 0xed, 0x67, 0x13, 0xdb, 0xc9, 0xe9, 0x78, 0xd4, 0x24,
		0x18, 0xa1, 0x52, 0x82, 0xb8, 0x85, 0x8c, 0x2b, 0x98, 0x56, 0x5c, 0xe9,
		0x7d, 0xd0, 0x73, 0xe5, 0xaa, 0x25, 0xec, 0x64, 0xf5, 0x39, 0x82, 0x55,
		0xbe, 0x42, 0xad, 0x56, 0xb7, 0x97, 0x95, 0xa3, 0x34, 0xc9, 0x1d, 0x36,
		0x7a, 0xf6, 0x69, 0xb8, 0xaf, 0x5f, 0xa5, 0xf1, 0x5a, 0x90, 0x3c, 0xd9,
		0xc1, 0x93, 0x1a, 0x9e, 0x57, 0x32, 0x8b, 0x02, 0x92, 0x40, 0xb7, 0x5c,
		0x07, 0xeb, 0x87, 0x91, 0x56, 0xaf, 0xe6, 0xd3, 0x0d, 0x05, 0x51, 0xe9,
		0x99, 0x87, 0x91, 0xef, 0x53, 0xb4, 0x7d, 0xe8, 0x78, 0xd4, 0x54, 0x65,
		0x9c, 0x52, 0x49, 0xd6, 0x0c, 0x16, 0xe5, 0x3a, 0xe3, 0x69, 0xc9, 0xea,
		0xa4, 0xef, 0x32, 0xcc, 0x06, 0x01, 0x49, 0x89, 0x51, 0xb8, 0x0b, 0x39,
		0x4a, 0x94, 0xb8, 0xa3, 0x8f, 0x63, 0x2f, 0xe1, 0x1f, 0x21, 0x07, 0x41,
		0x54, 0xcd, 0x32, 0xfe, 0x29, 0x91, 0xa6, 0x0b, 0xd2, 0x5f, 0xfa, 0x5d,
		0x41, 0x73, 0x20, 0x29, 0xd5, 0x8e, 0x8b, 0x4f, 0x62, 0xca, 0xb3, 0x8c,
		0x2a, 0x05, 0xc2, 0x1d, 0xd0, 0x65, 0xd8, 0x84, 0x62, 0x40, 0xf2, 0x2a,
		0xca, 0x53, 0xfd, 0xef, 0x0b, 0x17, 0x7b, 0x59, 0x90, 0xc4, 0x94, 0x85,
		0x61, 0x1a, 0xc1, 0xaf, 0x98, 0xf6, 0x9f, 0x4a, 0xe5, 0x05, 0xdc, 0xc1,
		0x1d, 0xba, 0x11, 0x2f, 0x44, 0x99, 0xc3, 0xb9, 0x73, 0xa9, 0x95, 0x6d,
		0x12, 0x9d, 0xdb, 0x74, 0x56, 0xdc, 0x70, 0xc6, 0x9c, 0x54, 0x40, 0x31,
		0x22, 0x74, 0x9b, 0x73, 0x01, 0xd7, 0x5c, 0xd1, 0xcd, 0xc1, 0xe8, 0xef,
		0x44, 0x7b, 0x38, 0xbe, 0x04, 0x16, 0x3b, 0xc2, 0x18, 0xff, 0x3a, 0x65,
		0x3c, 0x07, 0x1b, 0x88, 0x36, 0xd5, 0xe6, 0x60, 0x49, 0x59, 0x3a, 0xdd,
		0x71, 0x8e, 0x39, 0xfa, 0x44, 0x6d, 0x96, 0x8a, 0xb2, 0xf8, 0x02, 0x36,
		0xa4, 0x64, 0xea, 0x3c, 0xd8, 0x72, 0x34, 0xb2, 0xe7, 0xa0, 0xd0, 0x92,
		0x73, 0x36, 0xb1, 0x32, 0xe3, 0x91, 0x23, 0xd8, 0x36, 0xe8, 0x03, 0xb7,
		0xd9, 0xba, 0x6b, 0x18, 0x95, 0xca, 0xef, 0x17, 0xc0, 0x30, 0x96, 0x77,
		0xb0, 0x24, 0x62, 0x0b, 0xea, 0x82, 0x8a, 0x9a, 0xb1, 0x01, 0x81, 0x1d,
		0x16, 0x1c, 0x01, 0xee, 0x13, 0x56, 0xa6, 0x90, 0xde, 0xc2, 0xd6, 0x77,
		0xed, 0x80, 0xac, 0x8b, 0x50, 0x06, 0x5a, 0x99, 0x4c, 0xd1, 0xed, 0xad,
		0x43, 0x9c, 0x65, 0x84, 0x32, 0x47, 0x95, 0x7b, 0x5a, 0x2c, 0xc9, 0xd6,
		0x3a, 0xcb, 0xad, 0x6c, 0x0c, 0xf2, 0xc6, 0x95, 0x13, 0xdb, 0xe1, 0x74,
		0xd8, 0x71, 0xac, 0xa0, 0xb2, 0x98, 0x30, 0x3a, 0x8b, 0xda, 0x04, 0x9c,
		0x0d, 0x55, 0x2a, 0xb5, 0xf7, 0xdb, 0x1b, 0x93, 0x2c, 0x54, 0xeb, 0x4f,
		0x2f, 0x87, 0xef, 0xfa, 0x1c, 0xee, 0x15, 0xe4, 0xb2, 0x1e, 0x7d, 0x23,
		0xd4, 0xa8, 0xfa, 0xa3, 0x04, 0xdd, 0x6e, 0xd1, 0x8d, 0x46, 0xce, 0x5a,
		0xbb, 0x2a, 0xa4, 0x12, 0x40, 0x32, 0x6c, 0xe7, 0xd2, 0xf4, 0x73, 0x94,
		0x74, 0x03, 0xc2, 0xed, 0x88, 0x6f, 0x41, 0xcf, 0x5e, 0xa8, 0x72, 0x64,
		0x69, 0x88, 0xf5, 0x10, 0xc2, 0x86, 0x33, 0xf2, 0xab, 0xd2, 0x9e, 0x77,
		0x63, 0x26, 0x71, 0xe5, 0xd5, 0xf6, 0x25, 0x98, 0xb5, 0x6d, 0x29, 0x3f,
		0xfc, 0x76, 0x02, 0xe4, 0x8e, 0xb3, 0xb4, 0xdd, 0x17, 0x17, 0xab, 0xe9,
		0x74, 0xb6, 0x58, 0x34, 0xdb, 0x20, 0x32, 0xb9, 0x40, 0x0c, 0x40, 0xd8,
		0xe4, 0xed, 0x78, 0xe4, 0xfe, 0xd6, 0xcc, 0x84, 0x33, 0x2e, 0x26, 0xe7,
		0x97, 0xab, 0x99, 0x9e, 0x2a, 0xfa, 0x7f, 0xc8, 0xcb, 0x0a, 0x06, 0xca,
		0xd8, 0x64, 0xf3, 0xa3, 0x49, 0xf3, 0xbd, 0xb6, 0xa5, 0xd4, 0x78, 0xf4,
		0x7c, 0xff, 0xf8, 0x74, 0x19, 0x84, 0xa3, 0xc0, 0xef, 0xc3, 0xc9, 0xd8,
		0xe7, 0xce, 0xc9, 0x0f, 0x91, 0xfd, 0x61, 0xf4, 0x82, 0x21, 0x61, 0xdb,
		0xc6, 0x0d, 0x97, 0xca, 0xb4, 0x86, 0x5f, 0x39, 0xdf, 0xcb, 0x66, 0xe7,
		0x68, 0x33, 0x07, 0x8d, 0x79, 0xf1, 0xc8, 0xcd, 0x68, 0x63, 0x90, 0x1c,
		0xa6, 0x7d, 0x7c, 0x11, 0xa4, 0x28, 0x7c, 0xba, 0xb4, 0xa7, 0x18, 0xc1,
		0x14, 0xab, 0x5c, 0x1a, 0xbf, 0xc7, 0x7f, 0x53, 0xfd, 0xef, 0x3c, 0xd8,
		0xe5, 0xc7, 0xbd, 0x97, 0x43, 0x98, 0xf3, 0xaf, 0xf8, 0x6d, 0x0d, 0x73,
		0x2a, 0xe2, 0x15, 0x29, 0x2a, 0x88, 0x82, 0x29, 0x2b, 0x32, 0x1b, 0x23,
		0x47, 0x1b, 0xf4, 0x0e, 0xba, 0xef, 0xdc, 0xfb, 0x7c, 0x7c, 0x66, 0xf1,
		0x05, 0xd6, 0x18, 0x9a, 0xbe, 0x2e, 0x15, 0x17, 0x2f, 0x46, 0x6a, 0x16,
		0x8f, 0x2e, 0x01, 0xb3, 0x06, 0xeb, 0x76, 0xf2, 0xe6, 0xe1, 0x7c, 0x35,
		0xbf, 0xbc, 0xf8, 0xed, 0x7a, 0x75, 0x75, 0x3e, 0xbb, 0xfd, 0x36, 0x7c,
		0xf3, 0xf0, 0x71, 0xbe, 0xfc, 0xed, 0x76, 0xf6, 0x79, 0xbe, 0x98, 0x7f,
		0xba, 0xfe, 0x27, 0x83, 0x7c, 0xab, 0x76, 0x67, 0x3f, 0x61, 0xfe, 0xb7,
		0x77, 0xd6, 0xd5, 0x93, 0xe2, 0xf2, 0x82, 0x4a, 0x24, 0x1f, 0x2a, 0x2f,
		0xd8, 0x2e, 0xdf, 0xa1, 0x3f, 0x1f, 0xd4, 0x75, 0x0d, 0x35, 0x11, 0xef,
		0x04, 0x59, 0xe7, 0x2a, 0x26, 0xd2, 0x5f, 0x10, 0xd5, 0x55, 0x3d, 0xa1,
		0xd3, 0x1a, 0x70, 0xc6, 0x54, 0xc5, 0xa3, 0xc7, 0xef, 0x46, 0x36, 0xb1,
		0xd6, 0x2b, 0x0e, 0x7c, 0xc5, 0x81, 0xaf, 0x38, 0xf0, 0x15, 0x07, 0x46,
		0x5d, 0x1c, 0x38, 0xea, 0x01, 0x72, 0x35, 0x88, 0xfa, 0x6b, 0x8f, 0xc6,
		0x47, 0xba, 0xbe, 0x41, 0x28, 0xf5, 0x23, 0x7f, 0xc1, 0xc8, 0x1d, 0x5c,
		0x92, 0x35, 0x30, 0x3b, 0x0c, 0x30, 0xa2, 0x98, 0x8f, 0x90, 0x5e, 0xf3,
		0xd4, 0x7e, 0x6b, 0x08, 0x25, 0xb0, 0x04, 0x43, 0xfe, 0xa0, 0x09, 0x7b,
		0xc6, 0x09, 0xc9, 0x6f, 0x39, 0xc9, 0x5c, 0x62, 0xbb, 0xe5, 0xa0, 0xee,
		0x44, 0x69, 0xab, 0x01, 0xa5, 0xc6, 0xd5, 0x8c, 0x27, 0x7b, 0x63, 0xc1,
		0x0e, 0xf2, 0x0b, 0xfe, 0x35, 0x37, 0xa0, 0xb1, 0x22, 0x21, 0x20, 0x77,
		0x9b, 0xbe, 0x2b, 0xd7, 0x3d, 0xcc, 0x01, 0xd2, 0xa7, 0x8f, 0xea, 0x48,
		0x0d, 0xcc, 0x60, 0xc3, 0x96, 0x87, 0x39, 0x6b, 0x2a, 0xc5, 0x5b, 0xd5,
		0x22, 0xfb, 0x5c, 0x09, 0x01, 0x37, 0xb6, 0x87, 0x2d, 0x44, 0xf1, 0x7b,
		0x81, 0x45, 0x4d, 0x12, 0x75, 0x01, 0x45, 0x1b, 0x19, 0xba, 0x58, 0x26,
		0xbc, 0x38, 0x10, 0x2b, 0x16, 0x4f, 0x71, 0xe1, 0xf6, 0xf8, 0xec, 0x09,
		0x25, 0xde, 0x9d, 0xc4, 0x3f, 0x9e, 0xc6, 0x27, 0x75, 0x06, 0xb9, 0xef,
		0x60, 0x3a, 0x54, 0x16, 0x63, 0xeb, 0xc4, 0xa9, 0xf0, 0x46, 0x15, 0xb0,
		0xe0, 0x43, 0x59, 0xb5, 0x61, 0x43, 0x99, 0x6e, 0x9c, 0x5a, 0xde, 0xdf,
		0xa4, 0xe5, 0x2c, 0xbd, 0xfe, 0x36, 0xa5, 0xeb, 0xb7, 0x46, 0xfb, 0xb6,
		0x2c, 0x65, 0x4d, 0x91, 0xc0, 0xf0, 0x5c, 0xfe, 0x58, 0xb7, 0x69, 0xd8,
		0x65, 0x01, 0x28, 0xa4, 0x95, 0xcb, 0x16, 0x76, 0xe7, 0x51, 0x30, 0x0a,
		0xd1, 0xb9, 0x6c, 0x4d, 0x92, 0xfd, 0x92, 0x5f, 0x12, 0xa9, 0x16, 0x65,
		0x92, 0x80, 0x94, 0x9b, 0x92, 0xd9, 0x3a, 0x7e, 0x94, 0x1d, 0x0c, 0x7a,
		0x1b, 0xc3, 0x0f, 0x95, 0x21, 0x0b, 0xa5, 0xe7, 0xcc, 0xf6, 0x30, 0xc1,
		0x3e, 0xf2, 0x91, 0xf1, 0x35, 0x61, 0x0b, 0x50, 0x0a, 0x43, 0x5b, 0x3f,
		0x49, 0x5a, 0x82, 0x7e, 0xd0, 0x3a, 0xc3, 0x3c, 0x25, 0xe5, 0xd8, 0x9c,
		0x3f, 0xe0, 0x5e, 0x10, 0x85, 0xa0, 0xb9, 0x72, 0x7e, 0x0b, 0x06, 0xdb,
		0x63, 0x02, 0xfd, 0xf5, 0xfb, 0x68, 0xcc, 0x9f, 0x7c, 0x4a, 0x10, 0xb9,
		0x47, 0x34, 0xbf, 0x03, 0xc6, 0x82, 0xfe, 0x91, 0x65, 0x24, 0xc7, 0x1c,
		0xfc, 0x5b, 0xe4, 0xd0, 0x60, 0x94, 0xa0, 0x6d, 0x3a, 0x2c, 0x19, 0xd9,
		0xd2, 0x64, 0x00, 0xf7, 0x05, 0x17, 0x2a, 0xba, 0x99, 0xdf, 0xcc, 0x2e,
		0xe7, 0xd7, 0x33, 0x8b, 0x58, 0xcf, 0xde, 0xfc, 0x1d, 0x92, 0x1d, 0x8f,
		0x8e, 0xde, 0x3c, 0x78, 0xce, 0xe7, 0xd9, 0xad, 0x06, 0xaf, 0xdf, 0x8e,
		0xa2, 0xff, 0x47, 0x49, 0xa9, 0xa2, 0xe1, 0xe6, 0x24, 0x1a, 0xa6, 0xc7,
		0xc3, 0xe3, 0x7f, 0xb8, 0x43, 0x34, 0xc6, 0x5d, 0xfc, 0xfa, 0xfe, 0x79,
		0x9b, 0x4f, 0xed, 0xe6, 0xfe, 0xaf, 0x89, 0x01, 0x02, 0xc5, 0x36, 0x1e,
		0xe1, 0xab, 0x0b, 0x70, 0xd3, 0x70, 0x47, 0x44, 0x1a, 0x19, 0x30, 0x8d,
		0x17, 0xd5, 0x4f, 0xf0, 0x41, 0x35, 0xdb, 0x8d, 0xb1, 0xce, 0x2f, 0xda,
		0xaf, 0xde, 0xfe, 0x86, 0x9b, 0x5b, 0x9e, 0xb2, 0xad, 0xd0, 0x3d, 0x79,
		0x8a, 0x72, 0x8d, 0x13, 0x6b, 0xd7, 0x69, 0xea, 0x66, 0xd7, 0x15, 0x8e,
		0x92, 0xa0, 0x73, 0x67, 0xd5, 0x12, 0xab, 0xee, 0x24, 0xe8, 0xda, 0x08,
		0x85, 0x68, 0x41, 0xb1, 0x01, 0x04, 0xd5, 0x90, 0x22, 0xe8, 0x36, 0x23,
		0x7c, 0xa6, 0x1f, 0x14, 0x2b, 0xec, 0x4c, 0xba, 0xbf, 0x85, 0x6f, 0xce,
		0x27, 0x45, 0xea, 0xaa, 0xca, 0xd3, 0x25, 0x9f, 0xe7, 0x29, 0xbd, 0xa3,
		0x69, 0x89, 0xe9, 0xe5, 0x06, 0x59, 0x87, 0xde, 0x67, 0xb3, 0xd1, 0xbe,
		0xf1, 0xec, 0x0f, 0x8a, 0xbc, 0xc7, 0x5c, 0xc7, 0x7d, 0x2f, 0x92, 0x9d,
		0x7e, 0x09, 0x79, 0x3d, 0x88, 0x4f, 0xe1, 0x4e, 0xa7, 0x20, 0xcd, 0xec,
		0xd6, 0x4f, 0xf3, 0x9c, 0x1d, 0xe6, 0x9b, 0x4e, 0xd1, 0x76, 0xc8, 0x4e,
		0x5e, 0xbf, 0x84, 0xa4, 0xfa, 0x84, 0x6c, 0x2b, 0x19, 0x10, 0xbc, 0x02,
		0x1a, 0xe1, 0xcc, 0xb2, 0x42, 0x1d, 0xac, 0x6e, 0x1e, 0x04, 0x76, 0x18,
		0x7d, 0x9e, 0xe8, 0x37, 0xac, 0x59, 0x59, 0xd6, 0x43, 0xd7, 0x70, 0xaf,
		0xae, 0x48, 0x5e, 0x56, 0xb9, 0x59, 0x7f, 0x1a, 0x21, 0x25, 0xd6, 0x69,
		0x16, 0x27, 0x18, 0x67, 0x51, 0x66, 0xf2, 0x20, 0x15, 0x64, 0x32, 0x6e,
		0x54, 0x72, 0x5c, 0x25, 0x96, 0x7f, 0x87, 0xd9, 0x97, 0x75, 0x5c, 0x45,
		0xd4, 0x3d, 0xca, 0x6c, 0xeb, 0xf3, 0x19, 0x55, 0x6d, 0xe9, 0xbc, 0x37,
		0x4f, 0x10, 0x17, 0xfc, 0x18, 0xe2, 0x82, 0xea, 0x7d, 0x11, 0xa6, 0x97,
		0x9b, 0x75, 0x41, 0x93, 0x37, 0xb1, 0xe9, 0xa8, 0xaf, 0x33, 0xad, 0x57,
		0xda, 0xfa, 0xe9, 0xcf, 0x35, 0xec, 0x11, 0xc7, 0xda, 0x01, 0xa8, 0x95,
		0x6b, 0x78, 0xb5, 0xd5, 0x09, 0x0b, 0x22, 0x50, 0x35, 0x6c, 0xc3, 0xf4,
		0x7f, 0x90, 0x36, 0xae, 0x69, 0xfb, 0xad, 0x21, 0x39, 0xb4, 0xa2, 0xfa,
		0x25, 0xfa, 0xef, 0x8e, 0xdb, 0x82, 0xa1, 0xf0, 0xd2, 0xeb, 0xda, 0x0f,
		0xce, 0xbe, 0x43, 0x9f, 0x7d, 0xb0, 0xc6, 0x48, 0x37, 0x8e, 0x51, 0x07,
		0xd3, 0xb5, 0xae, 0xce, 0xb1, 0x6e, 0x98, 0x9b, 0xb8, 0x86, 0xd3, 0x5b,
		0xb6, 0x35, 0x4a, 0xa9, 0x7e, 0x30, 0xd5, 0xdf, 0xca, 0x6a, 0x52, 0x28,
		0x68, 0x15, 0xf9, 0x42, 0xd5, 0xee, 0x9a, 0xd7, 0x9a, 0xb8, 0x4a, 0x7a,
		0x8c, 0x5d, 0xbb, 0xaf, 0x3d, 0xb8, 0x5e, 0xe8, 0xbf, 0x96, 0x95, 0x2f,
		0x3f, 0xae, 0x9b, 0x5c, 0xe8, 0x92, 0xa0, 0x7f, 0x07, 0xf0, 0xe6, 0x8f,
		0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xf1, 0x54, 0x8a, 0x39, 0x1c, 0x00,
		0x00,
	},
		"templates/jenkins/normal-job.xml",
	)
}

func templates_jenkins_pipeline_xml() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xa4, 0x53,
		0xcb, 0x6e, 0xdb, 0x30, 0x10, 0xbc, 0xfb, 0x2b, 0x08, 0xa1, 0x57, 0x53,
		0x4e, 0x0f, 0x45, 0x0e, 0x34, 0xd3, 0xc2, 0x6d, 0x80, 0x16, 0x7d, 0xa4,
		0x68, 0xda, 0x6b, 0xc1, 0x48, 0x6b, 0x9a, 0x2d, 0xb5, 0x24, 0xf8, 0xb0,
		0x1d, 0x04, 0xf9, 0xf7, 0xae, 0x1e, 0x4c, 0xa2, 0xf4, 0x01, 0x04, 0x3d,
		0x49, 0x9c, 0x99, 0x9d, 0x9d, 0xd5, 0x52, 0xe2, 0xec, 0xd8, 0x59, 0xb6,
		0x87, 0x10, 0x8d, 0xc3, 0x75, 0x75, 0xc2, 0x57, 0x15, 0x03, 0x6c, 0x5c,
		0x6b, 0x50, 0xaf, 0xab, 0xaf, 0x97, 0xe7, 0xcb, 0xd3, 0xea, 0x4c, 0x2e,
		0x44, 0x04, 0xde, 0x1a, 0x75, 0xe5, 0x2c, 0xff, 0x01, 0xf8, 0xd3, 0x60,
		0xe4, 0xde, 0x78, 0xb0, 0x06, 0x81, 0xbf, 0xa6, 0x07, 0x19, 0x5c, 0x5f,
		0x4c, 0xc0, 0x37, 0x03, 0x07, 0xe6, 0x6d, 0xd6, 0x86, 0x0c, 0xdb, 0x89,
		0x5c, 0x16, 0xf9, 0x72, 0x64, 0x5e, 0xae, 0xf8, 0x29, 0x7f, 0x51, 0xc9,
		0x05, 0x63, 0x02, 0x55, 0x07, 0xf2, 0xe6, 0x86, 0xf1, 0x8f, 0xf4, 0xc2,
		0x6e, 0x6f, 0x45, 0x3d, 0x20, 0x3d, 0xb5, 0x35, 0x36, 0x41, 0x78, 0x73,
		0x84, 0x26, 0x27, 0x17, 0xa2, 0xdc, 0x2a, 0x1b, 0x41, 0xd4, 0x8f, 0xe1,
		0x7b, 0xe9, 0xe7, 0x0c, 0x19, 0xe6, 0xb2, 0x11, 0xea, 0x25, 0x3e, 0x38,
		0x0f, 0x21, 0x19, 0x88, 0xac, 0xb1, 0x2a, 0xc6, 0x75, 0xb5, 0xcb, 0x6d,
		0x74, 0xc8, 0x3b, 0x47, 0x39, 0x79, 0x1f, 0xfc, 0xd9, 0xc5, 0xa8, 0xb9,
		0x7e, 0x6f, 0x62, 0xaa, 0xea, 0xa1, 0xac, 0x71, 0x9d, 0x77, 0x08, 0x98,
		0xbe, 0x78, 0x68, 0x86, 0x66, 0x04, 0x3e, 0xf1, 0x83, 0x7c, 0x5f, 0x6e,
		0x1e, 0xba, 0x8c, 0x26, 0xff, 0x9c, 0x7d, 0xa0, 0xb7, 0x26, 0xc4, 0xf4,
		0xce, 0x5d, 0xcd, 0x25, 0x77, 0xe8, 0x18, 0xa5, 0xfe, 0xff, 0x2c, 0xa2,
		0xfe, 0x7d, 0x48, 0x81, 0xee, 0xd3, 0xb6, 0xd4, 0x45, 0x79, 0xb2, 0xa2,
		0x6c, 0x33, 0xa4, 0xd7, 0xc4, 0x9d, 0x3b, 0xbc, 0xd2, 0x3a, 0x80, 0x56,
		0x09, 0xda, 0xc2, 0x95, 0x05, 0xfc, 0x85, 0x2d, 0xe6, 0x1b, 0x67, 0x73,
		0x87, 0x64, 0x3d, 0x3a, 0x97, 0xe3, 0xe0, 0xeb, 0x68, 0x4f, 0xa8, 0x25,
		0x52, 0x26, 0xf2, 0x99, 0x4e, 0x77, 0x1d, 0xf7, 0x2a, 0xa9, 0xfb, 0xeb,
		0xf0, 0x10, 0xea, 0x25, 0xd9, 0xb7, 0xd4, 0xef, 0x2d, 0xd2, 0xf2, 0xf7,
		0xca, 0xca, 0xe7, 0xa2, 0x7e, 0x84, 0x14, 0x9f, 0xcd, 0x4e, 0xa1, 0x86,
		0x99, 0x4f, 0x81, 0x7a, 0x89, 0xb2, 0xd6, 0x1d, 0x3e, 0x28, 0xcc, 0xca,
		0x5e, 0x06, 0xa3, 0x35, 0xfd, 0x25, 0x32, 0x85, 0x4c, 0xca, 0x3f, 0x31,
		0x7d, 0x05, 0x4d, 0x0a, 0x47, 0x7f, 0x3e, 0xed, 0x27, 0xd2, 0xfd, 0x79,
		0xf2, 0x76, 0xe4, 0xe2, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x40,
		0xf2, 0x3f, 0x94, 0x03, 0x00, 0x00,
	},
		"templates/jenkins/pipeline.xml",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func assetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"templates/jenkins/multi-job.xml":  templates_jenkins_multi_job_xml,
	"templates/jenkins/normal-job.xml": templates_jenkins_normal_job_xml,
	"templates/jenkins/pipeline.xml":   templates_jenkins_pipeline_xml,
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
func assetDir(name string) ([]string, error) {
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
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"templates": &_bintree_t{nil, map[string]*_bintree_t{
		"jenkins": &_bintree_t{nil, map[string]*_bintree_t{
			"multi-job.xml":  &_bintree_t{templates_jenkins_multi_job_xml, map[string]*_bintree_t{}},
			"normal-job.xml": &_bintree_t{templates_jenkins_normal_job_xml, map[string]*_bintree_t{}},
			"pipeline.xml":   &_bintree_t{templates_jenkins_pipeline_xml, map[string]*_bintree_t{}},
		}},
	}},
}}
