// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// database/migrations/000001_init_schema.down.sql
// database/migrations/000001_init_schema.up.sql
// database/migrations/000002_divide_interaction_read_state.down.sql
// database/migrations/000002_divide_interaction_read_state.up.sql
// database/migrations/bindata.go
package migrations

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __000001_init_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xf0\xcb\x57\x28\x4a\x4d\x2c\xce\xcf\x53\x28\xc9\x57\x48\x2d\x4b\x2d\x52\x48\xc9\x57\x28\xc9\xc8\x2c\x06\x04\x00\x00\xff\xff\xc1\x36\x17\xc2\x1c\x00\x00\x00")

func _000001_init_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_init_schemaDownSql,
		"000001_init_schema.down.sql",
	)
}

func _000001_init_schemaDownSql() (*asset, error) {
	bytes, err := _000001_init_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init_schema.down.sql", size: 28, mode: os.FileMode(420), modTime: time.Unix(1597460235, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000001_init_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\xcf\x8e\x9b\x3c\x10\x3f\x7f\x3c\xc5\x1c\x13\x29\x6f\xb0\xa7\xec\xca\xfb\x09\x35\x4d\x5b\x92\x48\xdd\x93\xe5\xb5\x27\x89\x55\x30\x68\x3c\xec\x4a\x7d\xfa\xca\x84\x10\xd8\x1a\xc8\x1e\x7a\xc4\xf3\x1b\xd9\xbf\x3f\x33\x00\x3c\x65\x62\xbd\x17\xb0\x5f\x3f\x6e\x04\xa4\xcf\xb0\xfd\xb6\x07\xf1\x33\xdd\xed\x77\x50\x7b\x24\x9f\x2c\xc0\x1a\x78\x4c\xff\xdf\x89\x2c\x5d\x6f\xe0\x7b\x96\x7e\x5d\x67\x2f\xf0\x45\xbc\xac\x12\x00\x2c\x94\xcd\xe1\x4d\x91\x3e\x2b\x6a\x9a\xb7\x87\xcd\x06\x0e\xdb\xf4\xc7\x41\x04\xc0\xd1\x92\x67\xe9\x54\x81\x7f\xa1\x42\x39\x57\x53\xd5\xf0\x82\xf1\x6a\xa5\xbc\x7f\x2f\xc9\x44\x8b\x4a\xb3\x7d\x43\x59\x11\x1e\x91\xd0\x69\x94\x1e\x39\x8a\xd4\x84\x8a\xd1\x48\xc5\xc0\xb6\x40\xcf\xaa\xa8\xe0\xdd\xf2\xb9\xf9\x84\xdf\xa5\xc3\xe1\xa3\x2a\x73\x37\x3e\x01\x58\x3e\x24\x93\x2a\x7b\xd2\x92\xbc\x97\x47\x44\xe3\x93\xc5\x7f\xd6\x40\x5c\x6a\xb6\x9c\xc7\x85\x30\xe8\x35\xd9\x8a\x6d\xe9\xae\xf5\x46\x5b\xeb\x7e\x75\xf8\x8b\x23\x83\xb6\x70\xa3\x9c\x02\x41\x8f\xee\x28\xd7\xce\xc5\x23\xb2\x3e\x4f\xeb\x72\xc1\xba\x53\xad\x4e\xd8\x7f\xe9\x09\x1d\x92\xe2\x92\xae\x87\xc9\xf2\x21\x99\x49\xa6\xf4\xf5\x6b\x47\x7b\x2e\xa6\x4d\x83\x35\x60\x1d\x43\x26\x9e\x45\x26\xb6\x4f\xa2\x4d\xf8\xc2\x9a\x65\xc0\xf8\xb2\x26\x8d\x11\xd4\xc0\xa1\x2b\xfa\x8e\xd4\x04\xd8\x45\xd1\x45\x7b\xff\xea\x76\xc9\x72\x86\xa2\x2e\x1d\xa3\x63\x69\x19\x8b\x39\x76\xc3\x97\x77\xee\xcd\x51\x68\xdb\xc6\x73\xd5\x02\x06\x11\xe9\xd7\x3f\x9d\xc8\x96\xd4\x68\x48\x23\x63\x36\xa9\x6e\x55\xbf\xe6\xd6\x9f\x67\x50\xaa\xe6\xf3\x2d\x5a\x4d\xde\x6a\x1b\xdf\x1a\xb6\x50\xa7\x0f\x8a\xdc\x8e\x6b\xca\xfb\x87\xed\xb0\x2c\x3a\xf5\x57\x0d\x99\x39\x63\x87\x1b\xc9\x27\x00\xd3\xe6\x0e\xc3\x1b\x33\x77\x90\x62\x80\xd1\x7d\x19\x0c\xa5\xf8\x0e\x04\xf0\xac\x88\x65\x90\x7c\x52\x4b\x00\x74\xe6\x1e\xd8\x55\x9d\x2e\xf9\xe1\x59\xcb\xcb\x3e\x9c\x50\xc7\x3a\x46\x0a\xbb\xfb\x73\x33\x3d\x2b\x4b\x7f\x9c\xa6\x9a\x06\x63\x77\x6d\x26\x54\x46\x7a\x0e\xa4\xa3\xff\x21\x24\x1d\x7a\x02\x0c\x0c\x6a\x5b\xa8\xfc\x5f\xff\x58\x62\xbb\xe5\x03\xc7\x10\xc4\x3f\x01\x00\x00\xff\xff\x13\x04\xa6\x28\xe0\x07\x00\x00")

func _000001_init_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_init_schemaUpSql,
		"000001_init_schema.up.sql",
	)
}

func _000001_init_schemaUpSql() (*asset, error) {
	bytes, err := _000001_init_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init_schema.up.sql", size: 2016, mode: os.FileMode(420), modTime: time.Unix(1597460235, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000002_divide_interaction_read_stateDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xc8\xcc\x2b\x49\x2d\x4a\x4c\x2e\xc9\xcc\xcf\x2b\xe6\x52\x50\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x48\xce\xcf\x2d\xc8\x49\x2d\x49\x4d\xb1\xe6\xe2\x22\x52\x4b\x71\x62\x59\x6a\x4a\x7c\x5a\x7e\x51\x7c\x4e\x62\x49\x6a\x91\x35\x20\x00\x00\xff\xff\x35\x90\x51\xe8\x6a\x00\x00\x00")

func _000002_divide_interaction_read_stateDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_divide_interaction_read_stateDownSql,
		"000002_divide_interaction_read_state.down.sql",
	)
}

func _000002_divide_interaction_read_stateDownSql() (*asset, error) {
	bytes, err := _000002_divide_interaction_read_stateDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_divide_interaction_read_state.down.sql", size: 106, mode: os.FileMode(420), modTime: time.Unix(1597460235, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000002_divide_interaction_read_stateUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcc\x41\x0e\xc2\x20\x10\x05\xd0\x3d\xa7\xf8\xf7\x70\x35\x95\xe9\xea\x0b\x89\xc2\xba\x21\xed\x98\x34\xa9\xc5\x00\xf1\xfc\x1e\x41\x2f\xf0\x84\x49\xef\x48\x32\x51\xb1\x9f\xc3\x5a\x59\xc7\x5e\xcf\xee\x00\xf1\x1e\xd7\xc8\x7c\x0b\x58\xeb\xeb\x7d\xd8\xb0\x0d\x53\x8c\x54\x09\x08\x31\x21\x64\x12\x5e\x67\xc9\x4c\x98\x85\x0f\xbd\x38\xf7\x1f\xd8\xcb\xc7\xb6\xe5\x59\xdb\x72\x94\x61\xed\x17\xfb\x0d\x00\x00\xff\xff\x71\xbc\xc4\x8e\xa6\x00\x00\x00")

func _000002_divide_interaction_read_stateUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_divide_interaction_read_stateUpSql,
		"000002_divide_interaction_read_state.up.sql",
	)
}

func _000002_divide_interaction_read_stateUpSql() (*asset, error) {
	bytes, err := _000002_divide_interaction_read_stateUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_divide_interaction_read_state.up.sql", size: 166, mode: os.FileMode(420), modTime: time.Unix(1597460235, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x5d\x6f\xdc\xc8\x11\x7c\x5e\xfe\x8a\xb9\x05\xce\x58\x06\xca\x2e\xbf\x3f\x16\x10\x10\x9c\xed\x00\x7e\x88\x0f\x88\x7d\x4f\xe9\x40\x18\x72\x66\x64\x22\xda\xa5\x4c\x72\xed\x96\x0d\xfd\xf7\xa0\xa6\x47\x3a\xe9\x6c\x4b\x49\x10\x03\x96\x76\x87\xd3\xdd\xd5\xd5\xd5\x45\xed\x76\xea\xe5\x68\xac\xba\xb4\x47\x3b\xe9\xc5\x1a\xe5\xc6\x49\x5d\xeb\xfe\x5f\xfa\xd2\xaa\xc3\x70\x39\xe9\x65\x18\x8f\xb3\xea\x6e\xd4\xe5\xf8\xe7\x6e\x38\x1a\xbd\x68\xf5\xea\x57\xf5\xf6\xd7\xf7\xea\xf5\xab\x37\xef\xb7\x6a\xf3\x97\xfb\xe8\x38\xda\xed\xd4\x3c\x9e\xa6\xde\xce\x7b\x7c\xc6\xed\x4e\xcf\x76\xf7\x7b\xaa\x5d\x82\x7f\xe9\xc5\x70\x1c\x96\x8b\xb9\xff\x60\x0f\x7a\x6b\xc6\xcf\xc7\xed\xfc\xf1\xea\xbf\x08\x39\x5d\x3f\x1b\x90\x5d\x98\xe1\xd3\x60\xec\xc5\x70\x5c\xec\xa4\x7b\x3c\xba\x98\xac\x36\x17\xf3\xa2\x17\xfb\x9f\x55\x7d\x2e\xc9\xd3\x38\x02\x61\xdb\xcb\x31\xfa\x96\xd4\x28\x1a\x0e\xd7\xe3\xb4\xa8\x4d\xb4\x5a\x77\x37\x8b\x9d\xd7\xd1\x6a\xdd\x8f\x87\xeb\xc9\xce\xf3\xee\xf2\xcb\x70\x8d\x03\x77\x58\xf0\x6b\x18\xe5\xe7\x6e\x18\x4f\xcb\x70\x85\x2f\xa3\x0f\xb8\xd6\xcb\x87\x9d\x1b\xae\x2c\x3e\xe0\x60\x5e\xa6\xe1\x78\xe9\x9f\x2d\xc3\xc1\xae\xa3\x38\x8a\xdc\xe9\xd8\xab\x80\xe6\xef\x56\x9b\x8d\x9f\xe3\x3f\xfe\x89\xb2\x67\xea\xa8\x0f\x56\x49\x58\xac\x36\x77\xa7\x76\x9a\xc6\x29\x56\x5f\xa3\xd5\xe5\x17\xff\x4d\xed\xcf\x15\x50\x6d\xdf\xda\xcf\x48\x62\xa7\x8d\x87\x8d\xef\xbf\x9c\x9c\xb3\x93\x4f\x1b\xc7\xd1\x6a\x70\x3e\xe0\xa7\x73\x75\x1c\xae\x90\x62\x35\xd9\xe5\x34\x1d\xf1\xf5\x4c\xb9\xc3\xb2\x7d\x8d\xec\x6e\xb3\x46\x22\xf5\xf3\xc7\xbd\xfa\xf9\xd3\x5a\x90\xf8\x5a\x71\xb4\xba\x8d\xa2\xd5\x27\x3d\xa9\xee\xe4\x94\xd4\x91\x22\xd1\xea\x42\xe0\x9c\xab\x61\xdc\xbe\x1c\xaf\x6f\x36\x2f\xba\x93\x3b\x53\x97\x5f\xe2\x68\xd5\x5f\xbd\xbe\x43\xba\x7d\x79\x35\xce\x76\x13\x47\xff\x2f\x3c\x48\x23\xf9\x7f\x90\xc8\x4e\x93\xe0\x0e\x87\xdd\xc9\x6d\x7f\x01\xf4\x4d\x7c\x86\x1b\xd1\x6d\x14\x2d\x37\xd7\x56\xe9\x79\xb6\x0b\x28\x3f\xf5\x0b\xb2\xf8\xfe\xc2\x3c\xa2\xd5\x70\x74\xa3\x52\xe3\xbc\xfd\xeb\x70\x65\xdf\x1c\xdd\x78\x1f\x17\x46\x78\x77\xfe\x20\x83\x9f\xa1\x52\x61\x8c\xd1\x6a\x1e\xbe\xf8\xef\xc3\x71\xa9\x8a\x68\x75\xc0\x9a\xab\xfb\xa4\x7f\x1b\x8d\xf5\x87\xef\x87\x83\x55\x90\xc9\x16\x9f\x50\x67\xb7\x53\x6f\x91\x2b\xb4\x00\x65\x79\x1a\x44\x43\x1b\x37\xfc\x11\x44\xec\xef\x6f\xe2\x50\x1a\x60\xee\x63\xb7\x3e\x52\xb2\xbe\x03\xa2\x87\x59\x01\xf1\x89\xac\xb8\xbf\x89\xa5\x81\xc7\x49\x7d\xa0\x24\x45\x23\x8f\x92\xa2\xd1\x27\x92\xe2\xfe\x26\x7e\x48\xc3\xe3\xd4\x3e\xfc\xc7\xa9\x07\x77\xe3\xd9\x7a\xba\x02\xa8\xdc\xc4\xbf\xd3\xfa\x4d\x89\x07\x5c\xbf\x99\x5f\x0d\xd3\xa3\x32\x9f\x3f\xd8\xe5\x83\x9d\x94\x56\x66\x98\x6c\xbf\x8c\xd3\xcd\x13\xe5\x7c\xfc\x26\x56\xdd\x38\x5e\x7d\xdb\xca\x8b\x71\xde\xa2\x0f\xd4\xf8\xe9\x5c\x25\x77\xa3\xb8\x99\x1f\x95\x1c\x66\x35\xdf\xcc\xcf\x71\xf7\xee\x66\x96\x79\xd8\xc9\xe9\xde\x7e\xbd\x7d\x50\x2f\x88\x1b\xfb\x7a\x71\xf1\xad\x5f\xbf\x1a\x3f\x1f\xdf\x7d\xbc\x52\xe7\x41\xe3\x9b\x35\x71\xea\x88\x9b\x8e\x38\x69\x88\x93\xe4\xfb\xff\x9d\x23\x36\x19\xb1\x29\x89\xcb\x92\xd8\x25\xc4\x7d\x47\x5c\xd6\xc4\x59\x43\x5c\x68\xe2\xc2\x10\x67\x3d\x71\x6f\x89\x7b\x47\x5c\xe6\xf2\xac\x6f\xe5\x5e\xd1\x10\x67\x86\xb8\xe8\xe4\x77\x99\xc9\xd9\xdd\xf3\xbb\xbb\x7d\x23\x79\x92\x8a\x38\x29\x1e\x63\xc0\xff\x3e\x25\xce\x2b\xe2\xb4\x26\xee\x33\xe2\xb4\x7f\x88\x75\x7d\xe7\xb2\x3f\xee\x3e\x38\xc1\xf7\x1c\xf6\xce\x2f\x1e\x38\x74\xb4\x5a\x3d\xc1\xe4\x59\xb4\x5a\xad\x9f\x78\x95\xae\xcf\xa2\x55\x8c\x89\x3c\x83\x09\x70\xfe\xe4\xdd\xe8\x21\x1c\x6f\x47\xf7\x9e\xff\x6c\x47\xcf\xf9\xeb\xbd\x2d\x7a\x63\xdb\x9f\xff\x51\x5a\x5f\xe1\x12\x7b\xf5\x74\x3f\xde\x2e\xf6\x2a\x6b\xce\xbc\x4e\xf7\x0f\x77\x78\x53\x64\x49\xec\xcf\xb1\x59\x7b\xd9\xbc\xdf\x8e\x03\x6f\xd2\xb2\xad\x8b\x2a\xc9\xf2\xf2\x4c\x25\xf1\x6d\xb4\xd2\xa8\xff\xc2\x77\xfc\xd5\xb7\xb9\x57\xa1\x5b\x80\xdb\xfb\x9f\xb7\xf7\x03\xd1\x67\xcf\x29\xfb\xb7\xeb\xff\x55\xd7\x5d\x21\x9a\x86\x66\x1b\x4b\xdc\x76\xc4\x79\x4f\x9c\x26\xc4\xb9\x23\xae\x9d\x7c\xef\x4b\xd1\x5a\x0a\x5d\xb7\xc4\x15\x62\x13\x62\x5d\x13\x5b\x3c\xd7\xc4\x0e\xf5\x5a\xe2\xbc\x94\x7d\x28\x3b\xe2\x36\xe8\xdc\x18\xe2\x36\x27\xb6\x25\x71\x57\x12\x67\x35\x71\xd3\x4a\xed\x3c\x21\xae\x1a\xa9\x83\x5c\xd8\xa7\xda\x10\x3b\x2d\x79\x9b\x42\xf0\x98\x86\x38\xd5\xb2\x23\xa9\x25\xae\xf1\xb9\x20\x76\x39\x71\xda\x11\x9b\x96\xb8\x73\x82\x3b\xcf\xa5\x47\xe4\xac\x4a\xe2\x2a\x23\xee\x8c\xec\x0d\x70\x97\xe1\x5e\x65\x65\xcf\x74\x21\x1c\xe0\x99\x33\xc4\x1d\xf6\xaf\x26\x76\xa9\xe4\x02\x7e\x6b\x88\xeb\x9a\xb8\x4c\x88\x6b\xec\x71\x41\xdc\xba\xb0\xf7\x89\x60\xab\xe1\x07\x7d\xf0\x0c\x27\x3d\xe2\x79\x69\x84\x33\x1b\x62\xdb\x8a\xb8\xb6\xe1\x1c\xfb\xef\xc4\x57\x0a\xf0\xd3\x13\x6b\x70\x9d\x49\x0f\x88\x6f\x81\x0f\x18\x52\xe1\xb6\x05\x2e\x4b\x9c\x75\xc4\x95\x26\x6e\x35\x71\x07\x4e\x6b\xb9\xe7\xfd\xc3\x12\x9b\x8e\xd8\x21\x16\x1e\x94\x4a\xaf\xc0\x6a\x52\x99\x0d\x6a\x5b\xcc\xa1\x20\x6e\x52\x99\xb7\x46\x8d\x8c\x38\x0f\xde\x04\x2f\x33\xc1\xf7\x9a\x4c\x38\x41\x4d\x1d\xb0\x22\x06\xf8\xe1\x5f\x45\x21\xba\xc3\x0c\xbb\x5c\xe6\x58\xe4\xc4\x65\x4b\x9c\xa6\x32\x37\x60\x37\x98\x79\x2b\x7d\xe1\x0c\xb5\x1a\x4d\x6c\x0a\x99\xb7\xff\x9c\x06\x8f\x2c\x05\x33\x66\x52\x24\x32\x27\x8d\xfa\x8d\xf0\x89\xd9\x38\xcc\xb7\x25\x76\xad\xc4\x83\x7b\xe0\xeb\x73\xe1\x4c\x43\xb7\x98\x4f\x4e\x9c\xd7\xc2\x5d\x92\x12\x97\x4d\xe0\xb1\x10\x7d\xfa\x3b\x9d\xf8\x3d\xf8\x41\x7f\x2d\xb8\xab\xa5\x37\x9b\xcb\x3d\x60\x42\xdf\x26\x60\x2a\x7b\xe1\x04\x18\xdb\x5e\xf8\x6e\x82\xb6\x6d\x23\x3b\x61\x42\x8f\x95\x11\xce\xb5\x25\x76\xd8\x95\x4a\x76\xc5\x76\xa2\x89\x1a\x5c\x37\xc2\x63\x96\x13\x37\xb9\xe4\xad\x13\xe9\x03\xf9\x93\x4c\xe6\xd9\x40\x53\x56\xde\x1f\xa6\x0e\xef\x9f\x52\xe2\xba\x4c\x7a\x2b\x5c\xc8\x9d\xc9\xfb\xa3\xd3\x32\x53\x8d\xd9\x58\xd9\x2b\xf4\x02\x4c\xd8\xdf\xd4\x48\xef\x36\x70\x00\x9c\x5e\x57\xad\xf0\x9c\x41\x3b\x98\x35\x7a\xa8\x44\x5f\xe8\xa3\xea\x84\x27\xbf\x1b\x56\x9e\x41\xab\xc0\x8f\xfe\xab\x44\x72\x43\x0f\x19\x78\xb2\xa2\x77\x7c\xf6\xb3\xc5\xcc\x32\x89\x41\xff\x78\xbf\x41\xd7\x55\x98\x3f\xce\xd0\x13\xee\x37\x46\xf4\x64\x82\x9e\xa1\x59\xec\x02\x3c\x03\x3e\x86\x73\x68\x1d\xbd\xa1\x86\xd7\x52\x29\xd8\xb0\x9b\x16\x73\xae\xa4\x27\xf0\xd2\x54\x52\x1b\xb8\x81\x11\x3c\x57\xb9\x68\x34\x6d\x65\x16\xd0\xa8\x8f\xc9\x45\x13\xf0\x02\x70\x8f\x7d\x2b\x53\xf1\x30\xec\x5b\x5f\x89\xce\xb0\x33\x98\x13\x76\x11\x58\xc1\x2f\x66\x04\xdf\x69\x8d\xf8\x18\x7c\x00\x58\x11\xeb\xff\x46\x08\x75\xf1\x8e\xd7\xad\xf8\x13\x3c\x12\xbe\x66\x2b\xd1\x1b\x3c\xa7\x0a\x3e\x84\x3d\xc0\x39\x3c\xd0\xff\x5d\x11\xfe\x16\x31\x56\xf4\xe8\x7d\xd1\x89\x37\x81\xe3\x32\x78\x76\xe2\xc2\x8e\xf5\xc2\x17\x66\x77\xe7\xe1\x98\x35\xbc\xae\x0b\xde\x93\x6b\x99\x2f\xbc\x20\x31\x32\x7b\x70\x01\x9c\xb8\x8b\x1a\xe0\x11\x9e\x0e\xbf\x32\xb9\x60\x82\xff\xf4\xc1\x9b\xe1\xa1\xe0\xab\x4d\xc4\x8b\x9b\xe0\x41\xb5\xa1\x7f\x07\x00\x00\xff\xff\xb6\x58\x2a\x20\x00\x10\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 12288, mode: os.FileMode(420), modTime: time.Unix(1597460290, 0)}
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
	"000001_init_schema.down.sql":                   _000001_init_schemaDownSql,
	"000001_init_schema.up.sql":                     _000001_init_schemaUpSql,
	"000002_divide_interaction_read_state.down.sql": _000002_divide_interaction_read_stateDownSql,
	"000002_divide_interaction_read_state.up.sql":   _000002_divide_interaction_read_stateUpSql,
	"bindata.go":                                    bindataGo,
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
	"000001_init_schema.down.sql":                   &bintree{_000001_init_schemaDownSql, map[string]*bintree{}},
	"000001_init_schema.up.sql":                     &bintree{_000001_init_schemaUpSql, map[string]*bintree{}},
	"000002_divide_interaction_read_state.down.sql": &bintree{_000002_divide_interaction_read_stateDownSql, map[string]*bintree{}},
	"000002_divide_interaction_read_state.up.sql":   &bintree{_000002_divide_interaction_read_stateUpSql, map[string]*bintree{}},
	"bindata.go":                                    &bintree{bindataGo, map[string]*bintree{}},
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
