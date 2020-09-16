// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// database/migrations/000001_init_schema.down.sql
// database/migrations/000001_init_schema.up.sql
// database/migrations/000002_divide_interaction_read_state.down.sql
// database/migrations/000002_divide_interaction_read_state.up.sql
// database/migrations/000003_rename_prefSet_to_engine.down.sql
// database/migrations/000003_rename_prefSet_to_engine.up.sql
// database/migrations/000004_add_active_feed_to_user.down.sql
// database/migrations/000004_add_active_feed_to_user.up.sql
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

	info := bindataFileInfo{name: "000001_init_schema.down.sql", size: 28, mode: os.FileMode(420), modTime: time.Unix(1597804349, 0)}
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

	info := bindataFileInfo{name: "000001_init_schema.up.sql", size: 2016, mode: os.FileMode(420), modTime: time.Unix(1597804349, 0)}
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

	info := bindataFileInfo{name: "000002_divide_interaction_read_state.down.sql", size: 106, mode: os.FileMode(420), modTime: time.Unix(1597804349, 0)}
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

	info := bindataFileInfo{name: "000002_divide_interaction_read_state.up.sql", size: 166, mode: os.FileMode(420), modTime: time.Unix(1597804349, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000003_rename_prefset_to_engineDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _000003_rename_prefset_to_engineDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000003_rename_prefset_to_engineDownSql,
		"000003_rename_prefSet_to_engine.down.sql",
	)
}

func _000003_rename_prefset_to_engineDownSql() (*asset, error) {
	bytes, err := _000003_rename_prefset_to_engineDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000003_rename_prefSet_to_engine.down.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1598572155, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000003_rename_prefset_to_engineUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xc1\x4a\xc4\x30\x14\x45\xf7\xf9\x8a\xbb\x6c\x65\xfe\x60\x56\x9d\xfa\x2a\xc1\x4e\x3a\x26\x29\x38\xab\x12\x26\x4f\x27\x8b\x66\x86\x24\x2a\xf8\xf5\x42\x3b\xba\x10\x04\x97\x8f\x73\x78\x5c\x4e\xab\xa9\xb1\x04\xdb\xec\x7a\x82\xec\xa0\x06\x0b\x7a\x96\xc6\x1a\x70\x7c\x0d\x91\xb3\x00\x2a\x04\x8f\x9d\x7c\x30\xa4\x65\xd3\xe3\xa0\xe5\xbe\xd1\x47\x3c\xd2\x71\x23\x00\xe0\x2d\x73\x9a\x82\x47\x88\x65\x79\xa0\xc6\xbe\x87\xa6\x8e\x34\xa9\x96\xcc\xc2\x73\x15\x7c\xbd\xea\xd1\xcd\x8c\x77\x97\x4e\x67\x97\x7e\xfc\x15\xe5\x4b\x2a\x7f\xa1\xe2\x52\x99\xbc\x2b\x8c\x12\x66\xce\xc5\xcd\x57\x7c\x84\x72\x5e\x4e\x7c\x5e\x22\xaf\x22\x47\xff\x1f\x6d\x54\xf2\x69\x24\x54\xb7\xf1\x9b\x65\x56\x2d\x80\x7a\x2b\x84\x54\x86\xb4\x85\x54\x76\xf8\xee\x00\x43\x3d\xb5\x16\x77\xe8\xf4\xb0\xc7\x35\xf1\x0b\x27\x8e\x27\x9e\x32\x97\xbc\x15\xe2\x5e\x0f\x87\x5b\xc8\xdf\xf0\x2b\x00\x00\xff\xff\xf7\x36\x9c\x34\x66\x01\x00\x00")

func _000003_rename_prefset_to_engineUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000003_rename_prefset_to_engineUpSql,
		"000003_rename_prefSet_to_engine.up.sql",
	)
}

func _000003_rename_prefset_to_engineUpSql() (*asset, error) {
	bytes, err := _000003_rename_prefset_to_engineUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000003_rename_prefSet_to_engine.up.sql", size: 358, mode: os.FileMode(420), modTime: time.Unix(1598572155, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000004_add_active_feed_to_userDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _000004_add_active_feed_to_userDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000004_add_active_feed_to_userDownSql,
		"000004_add_active_feed_to_user.down.sql",
	)
}

func _000004_add_active_feed_to_userDownSql() (*asset, error) {
	bytes, err := _000004_add_active_feed_to_userDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000004_add_active_feed_to_user.down.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1598572155, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000004_add_active_feed_to_userUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x91\xc1\x6e\xea\x30\x10\x45\xf7\xfe\x8a\x59\x82\x94\x3f\x60\x65\xc0\xef\xc9\xaa\x09\xd4\x24\x52\x59\x59\xae\x3d\x14\x2f\x30\x91\xed\x50\xb5\x5f\x5f\xc5\x49\xd3\x8a\xb6\xd0\x4d\x97\xd1\x9c\x99\xdc\x7b\xbc\x90\x8c\x56\x0c\x2a\x3a\x17\x0c\xf8\x3f\x28\xd7\x15\xb0\x07\xbe\xad\xb6\xd0\x46\x0c\x6a\x8f\x68\x23\x99\x80\xb3\x30\xe7\xff\xb7\x4c\x72\x2a\x60\x23\xf9\x8a\xca\x1d\xdc\xb1\x5d\x41\xa0\x07\x9d\x05\xe7\x53\xde\x2f\x6b\x21\x20\xe0\x1e\x03\x7a\x83\x31\xcf\xe3\xc4\xd9\x69\x07\xa3\x7f\x72\x1e\xaf\xe1\x3d\x31\x2e\x78\x7d\x44\x38\xeb\x60\x0e\x3a\x8c\x7c\x37\x30\x01\x75\x42\xab\x74\x82\xe4\x8e\x18\x93\x3e\x36\xf0\xec\xd2\x21\x7f\xc2\xeb\xc9\x63\x4e\xd7\xd8\xdf\x60\x75\xc9\xef\x6b\x36\x19\xba\x14\xf9\xb7\x53\x32\x9d\x11\x72\x45\x51\x67\x47\xc5\xf6\x31\x9a\xe0\x9a\xe4\x4e\xfe\x96\xaa\xbc\x70\x43\x55\xef\xfc\xbd\x7e\x3c\xb5\xc1\xa0\x4a\x2f\xcd\xf7\x16\x86\xf9\xc5\xcd\x3f\xf3\x33\x14\x28\x3e\xe7\x2a\x3e\x42\xf4\xc6\xa8\xa8\x98\x1c\x84\xe5\xc7\x27\x00\x74\xb9\x84\xc5\x5a\xd4\xab\x12\xb4\x49\xee\x8c\x6a\xec\xfa\xa5\xd7\x0f\x27\x96\x72\xbd\xb9\xb8\xd1\x8c\xee\x54\xc4\x34\x7b\x0b\x00\x00\xff\xff\x87\x2f\x6e\x3f\xce\x02\x00\x00")

func _000004_add_active_feed_to_userUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000004_add_active_feed_to_userUpSql,
		"000004_add_active_feed_to_user.up.sql",
	)
}

func _000004_add_active_feed_to_userUpSql() (*asset, error) {
	bytes, err := _000004_add_active_feed_to_userUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000004_add_active_feed_to_user.up.sql", size: 718, mode: os.FileMode(420), modTime: time.Unix(1600027512, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x5b\x6f\xdb\x38\x16\x7e\xb6\x7e\x05\xc7\xc0\x14\xd6\x22\x6b\xeb\x7e\x31\x10\x60\x31\x6d\x17\xe8\xc3\x76\x80\x6d\xe7\x69\xb9\x30\x28\x91\x74\x85\x8d\x2d\x57\x92\x5b\xa6\x45\xfe\xfb\xe2\xe3\xa1\x13\xa7\xb9\x58\x4d\x33\xdb\x2d\x90\xc6\xa2\xc8\x73\xff\xbe\x73\xe8\x2c\x16\xec\x65\x2b\x15\x5b\xab\xad\xea\xc4\xa0\x24\xd3\x6d\xc7\x76\xa2\xfe\x8f\x58\x2b\xb6\x69\xd6\x9d\x18\x9a\x76\xdb\xb3\xea\x92\xad\xdb\xbf\x56\xcd\x56\x8a\x41\xb0\x57\xbf\xb3\xb7\xbf\xbf\x67\xaf\x5f\xbd\x79\x3f\x67\xb3\xbf\x5d\x9f\xf6\xbd\xc5\x82\xf5\xed\xbe\xab\x55\xbf\xc4\x67\xec\xae\x44\xaf\x16\x37\xa2\x16\x01\xfe\x85\xab\x66\xdb\x0c\xab\xbe\xfe\xa0\x36\x62\x2e\xdb\xcf\xdb\x79\xff\xf1\xe2\x3b\x8e\xec\x77\x27\x0f\x44\x2b\xd9\x7c\x6a\xa4\x5a\x35\xdb\x41\x75\xa2\xc6\xab\x55\xa7\x84\x5c\xf5\x83\x18\xd4\x38\xad\xa7\x84\x8c\xb0\x23\x5e\x75\x6a\x2b\x36\x6a\xb5\xeb\x94\x7e\xa7\x86\xd5\xd0\xae\xd4\x76\xdd\x6c\x47\x9a\xf0\xc8\xf9\x11\xda\x93\x95\x90\x72\x05\xc3\x3f\xa9\x95\x56\x4a\xe2\xf8\xbe\x57\xdd\x38\xe5\x0f\x1f\x7f\x5c\xb7\x2b\x95\xf9\xba\xf5\xee\x96\x93\xe7\x35\x9b\x5d\xdb\x0d\x6c\xe6\x4d\xa6\xd5\xe5\xa0\xfa\xa9\x37\x99\xd6\xed\x66\xd7\xa9\xbe\x5f\xac\xbf\x34\x3b\x2c\xe8\xcd\x80\x5f\x4d\x4b\xff\x2f\x9a\x76\x3f\x34\x17\x78\x68\xed\x81\x9d\x18\x3e\x2c\x74\x73\xa1\xf0\x01\x0b\xfd\xd0\x35\xdb\xb5\x7d\x37\x34\x1b\x35\xf5\x7c\xcf\xd3\xfb\x6d\xcd\x9c\x35\xff\x54\x42\xce\x6c\x05\xff\xeb\xdf\x50\x7b\xc6\x10\x57\x46\xc7\x7c\x36\x3b\xac\xaa\xae\x6b\x3b\x9f\x7d\xf5\x26\xeb\x2f\xf6\x89\x2d\xcf\x19\xac\x9a\xbf\x55\x9f\x21\x44\x75\x33\x6b\x36\x9e\x7f\xdb\x6b\xad\x3a\x2b\xd6\xf7\xbd\x49\xa3\xed\x81\x5f\xce\xd9\xb6\xb9\x80\x88\x49\xa7\x86\x7d\xb7\xc5\xe3\x19\xd3\x9b\x61\xfe\x1a\xd2\xf5\x6c\x0a\x41\xec\xd7\x8f\x4b\xf6\xeb\xa7\x29\x59\x62\x75\xf9\xde\xe4\xca\xf3\x26\x9f\x44\xc7\xaa\xbd\x66\xa4\x87\x94\x78\x93\x15\x99\x73\xce\x9a\x76\xfe\xb2\xdd\x5d\xce\x5e\x54\x7b\x7d\xc6\xd6\x5f\x7c\x6f\x52\x5f\xbc\x3e\x58\x3a\x7f\x79\xd1\xf6\x6a\xe6\x7b\xcf\x65\x0f\xc4\x90\xfc\x07\x04\xa9\xae\x23\xbb\xdd\x62\xb5\xd7\xf3\xdf\x60\xfa\xcc\x3f\xc3\x0e\xef\xca\xf3\x86\xcb\x9d\x62\xa2\xef\xd5\x80\x90\xef\xeb\x01\x52\xac\x7f\x2e\x1f\xde\xa4\xd9\xea\x96\xb1\xb6\x9f\xff\xbd\xb9\x50\x6f\xb6\xba\xbd\x3e\xe7\x52\x78\x58\x3f\x92\x60\x73\xc8\x98\x4b\xa3\x37\xe9\x9b\x2f\xf6\xb9\xd9\x0e\x59\xe2\x4d\x36\x20\x38\x76\x2d\xf4\x1f\xad\x54\x76\xf1\x7d\xb3\x51\x0c\x65\x32\xc7\x27\xe8\x59\x2c\xd8\x5b\xc8\x72\x2e\xa0\xb2\x6c\x18\xa8\x86\x66\xba\xf9\xd6\x08\xdf\xee\x9f\xf9\x4e\x35\x8c\xb9\x3e\x3b\xb7\x27\x49\xea\x3b\x58\x74\x2c\x15\x26\x3e\x22\x15\xfb\x67\x3e\x39\x70\x5b\xa8\x3d\x48\x42\xe1\xc8\x2d\xa1\x70\xf4\x11\xa1\xd8\x3f\xf3\x8f\xc3\x70\x5b\xb4\x3d\xfe\xb0\xe8\x46\x5f\xda\x68\x3d\xae\x01\xa1\x9c\xf9\x37\x61\xbd\xa3\xe2\x28\xd6\x6f\xfa\x57\x4d\x77\x4b\xcd\xe7\x0f\x6a\xf8\xa0\x3a\x26\x98\x6c\x3a\x55\x0f\x6d\x77\xf9\x88\x3a\x7b\x7e\xe6\xb3\xaa\x6d\x2f\xee\xba\xf2\xa2\xed\xe7\xf0\x03\x3a\x7e\x39\x67\xc1\x21\x15\x97\xfd\x2d\x95\x4d\xcf\xfa\xcb\xfe\x54\xec\xde\x5d\xf6\x94\x0f\xd5\x69\x51\xab\xaf\x57\x47\xfa\x5c\x71\x03\xaf\xab\xd5\xdd\x4e\xf5\xaa\xfd\xbc\x7d\xf7\xf1\x82\x9d\xbb\x1a\x9f\x4d\xb9\x09\x35\x37\x45\xc5\x4d\x50\x70\x13\x04\xf7\xff\x68\xcd\x8d\x8c\xb8\x91\x29\x37\x69\xca\x8d\x0e\xb8\xa9\x2b\x6e\xd2\x9c\x9b\xa8\xe0\x26\x11\xdc\x24\x92\x9b\xa8\xe6\xa6\x56\xdc\xd4\x9a\x9b\x34\xa6\x77\x75\x49\xfb\x92\x82\x9b\x48\x72\x93\x54\xf4\x3b\x8d\x68\xed\xf0\xfe\xb0\xb7\x2e\x48\x4e\x90\x71\x13\x24\xb7\x6d\xc0\x4f\x1d\x72\x13\x67\xdc\x84\x39\x37\x75\xc4\x4d\x58\x1f\xdb\x3a\x3d\xb0\xec\xc3\xde\x3b\x26\xb8\x8f\x61\x0f\x7c\x71\xc4\xd0\xde\x64\xf2\x48\x24\xcf\xbc\xc9\x64\xfa\xc8\x10\x31\x3d\xf3\x26\x3e\x32\x72\xc2\x26\x98\xf3\x17\xcb\x46\xc7\xe6\x58\x3a\xba\xe6\xfc\x93\x1e\x9d\xe2\xd7\x6b\x5a\xb4\xc4\xb6\x3c\xff\xb6\xb4\xbe\x82\x25\x96\xec\x71\x7f\x2c\x5d\x2c\x59\x54\x9c\xd9\x3a\x5d\x1e\x63\x78\x96\x44\x81\x6f\xd7\x81\xac\x25\x21\xef\x8f\x6d\x63\x66\x61\x5a\xe6\x45\x90\xc4\x49\x79\xc6\x02\xff\xca\x9b\x08\xe8\x7f\x61\x3d\xfe\x6a\xdd\x5c\x32\xe7\x2d\x8c\x5b\xda\xff\xaf\xae\x13\x22\xce\x4e\x55\xf6\x1f\xbb\xa7\xd6\x75\x95\x50\x4d\xa3\x66\x0b\xc5\x4d\x59\x71\x13\xd7\xdc\x84\x01\x37\xb1\xe6\x26\xd7\xf4\x5c\xa7\x54\x6b\x21\xea\xba\xe4\x26\xc3\xd9\x80\x1b\x91\x73\xa3\xf0\x5e\x70\xa3\xa1\xaf\xe4\x26\x4e\x09\x0f\x69\xc5\x4d\xe9\xea\x5c\x4a\x6e\xca\x98\x1b\x95\x72\x53\xa5\xdc\x44\x39\x37\x45\x49\xba\xe3\x80\x9b\xac\x20\x3d\x90\x05\x3c\xe5\x92\x1b\x2d\x48\x6e\x91\x90\x3d\xb2\xe0\x26\x14\x84\x91\x50\x71\x93\xe3\x73\xc2\x8d\x8e\xb9\x09\x2b\x6e\x64\xc9\x4d\xa5\xc9\xee\x38\x26\x1f\x21\x33\x4b\xb9\xc9\x22\x6e\x2a\x49\xb8\x81\xdd\xa9\xdb\x97\x29\xc2\x99\x48\x28\x06\x78\xa7\x25\x37\x15\xf0\x97\x73\xa3\x43\x92\x05\xfb\x95\xe4\x26\xcf\xb9\x49\x03\x6e\x72\xe0\x38\xe1\xa6\xd4\x0e\xf7\x01\xd9\x96\x83\x0f\x6a\xc7\x19\x9a\x7c\xc4\xfb\x54\x52\xcc\x94\x3b\x5b\x66\xdc\xe4\xca\xad\x03\xff\x9a\x78\x25\x41\x7c\x6a\x6e\x04\x62\x1d\x91\x0f\x38\x5f\xc2\x3e\xd8\x10\x52\x6c\x4b\xd8\xa5\xb8\x89\x2a\x6e\x32\xc1\x4d\x29\xb8\xa9\x10\xd3\x9c\xf6\x59\xfe\x50\xdc\xc8\x8a\x1b\x8d\xb3\xe0\xa0\x90\x7c\x85\xad\x32\xa4\xdc\x40\xb7\x42\x1e\x12\x6e\x8a\x90\xf2\x2d\xa0\x23\xe2\x26\x76\xdc\x04\x2e\x93\x8e\xf7\x8a\x88\x62\x02\x9d\xc2\xd9\x8a\x33\xb0\x1f\xfc\x95\x24\x54\x77\xc8\x61\x15\x53\x1e\x93\x98\x9b\xb4\xe4\x26\x0c\x29\x6f\xb0\x5d\x22\xe7\x25\xf9\x85\x35\xe8\x2a\x04\x37\x32\xa1\x7c\xdb\xcf\xa1\xe3\xc8\x94\x6c\x46\x4e\x92\x80\xf2\x24\xa0\xbf\xa0\x78\x22\x37\x1a\xf9\x2d\xb9\xd1\x25\x9d\x47\xec\x61\x5f\x1d\x53\xcc\x04\xea\x16\xf9\x89\xb9\x89\x73\x8a\x5d\x10\x72\x93\x16\x2e\x8e\x09\xd5\xa7\xdd\x53\x11\xdf\x23\x3e\xf0\xaf\x44\xec\x72\xf2\x4d\xc5\xb4\x0f\x36\xc1\x6f\xe9\x6c\x4a\x6b\x8a\x09\x6c\x2c\x6b\x8a\x77\xe1\x6a\x5b\x15\x84\x09\xe9\x7c\xcc\x24\xc5\x5c\x28\x6e\x34\xb0\x92\x11\x56\x54\x45\x35\x91\x23\xd6\x05\xc5\x31\x8a\xb9\x29\x62\x92\x9b\x07\xe4\x07\xe4\x07\x11\xe5\xb3\x40\x4d\x29\xea\x1f\x32\x77\xfd\x27\xa5\x73\x55\x44\xbe\x25\xda\xc9\x8e\xa8\x7f\x54\x82\x72\x2a\x90\x1b\x45\xb8\x82\x2f\xb0\x09\xf8\x0d\x25\xf9\xae\x5c\x0c\x60\xa7\xad\xab\x92\xe2\x1c\xa1\x76\x90\x6b\xf8\x90\x51\x7d\xc1\x8f\xac\xa2\x38\x59\x6c\x28\x7a\x87\x5a\x85\xfd\xf0\x3f\x0b\x48\x36\xea\x21\x42\x9c\x14\xd5\x3b\x3e\xdb\xdc\x22\x67\x11\x9d\x81\xff\xe8\x6f\xa8\xeb\xcc\xe5\x1f\x6b\xf0\x09\xfb\x0b\x49\xf5\x24\x5d\x3d\xa3\x66\x81\x05\x70\x06\x78\x0c\xeb\xa8\x75\xf8\x06\x1d\xb6\x96\x52\xb2\x0d\xd8\x54\xc8\x73\x46\x3e\x21\x2e\x45\x46\xba\x61\x37\x6c\x44\x9c\xb3\x98\x6a\x34\x2c\x29\x17\xa8\x51\x7b\x26\xa6\x9a\x00\x17\x20\xf6\xc0\x5b\x1a\x12\x87\x01\x6f\x75\x46\x75\x06\xcc\x20\x4f\xc0\x22\x6c\x45\x7c\x91\x23\xf0\x4e\x29\x89\xc7\xc0\x03\xb0\x15\x67\xed\x8c\xe0\xf4\xa2\xc7\x8b\x92\xf8\x09\x1c\x09\x5e\x53\x19\xd5\x1b\x38\x27\x73\x3c\x04\x1c\x60\x1d\x1c\x68\xe7\x0a\x37\x8b\x48\x45\xf5\x68\x79\x51\x13\x37\x21\xc6\xa9\xe3\xec\x40\x3b\x8c\xd5\x14\x2f\xe4\xee\xc0\xe1\xc8\x35\xb8\xae\x72\xdc\x13\x0b\xca\x2f\xb8\x20\x90\x94\x7b\xc4\x02\x76\x62\x2f\x74\x20\x8e\xe0\x74\xf0\x95\x8c\xc9\x26\xf0\x4f\xed\xb8\x19\x1c\x8a\x78\x95\x01\x71\x71\xe1\x38\x08\xfc\x80\xf8\x06\x35\xe1\x15\x39\x0e\x1c\x27\x83\x6f\xc0\x83\x45\x41\x35\x80\xb3\x79\xed\x66\xa7\x8a\x74\xe5\x09\xe9\x0a\x5d\x1d\x23\x56\x88\x2d\xf2\x08\x2e\x00\xf6\xc1\x3d\x98\xd7\xf0\x03\x3c\x00\x3f\x88\x03\x30\x0a\xdf\x80\xbf\x00\xf1\xa9\xc8\x7f\xcb\x11\xf8\x91\x54\x53\xd0\x05\x7c\xe0\x1c\x6a\x1a\x35\x89\x9c\x66\xce\x27\xe0\x19\x35\x09\x0e\x85\x2c\xf4\x25\xf8\x80\x9e\x82\x3a\x42\xbd\x01\x33\xe0\x18\xd4\x96\x7d\x1f\x91\xcf\x81\xe3\x0d\xe0\x1f\x7c\xa6\x6b\xea\x4d\x38\x03\x6e\xb2\x7d\xab\xa2\xbe\x19\xc4\x64\x33\x7a\x21\xfa\x1e\xea\x08\x1c\xf6\xed\x8c\x88\x1e\x6d\x7b\x5a\x46\xbc\x00\x8e\x0c\xf2\x31\x33\xa2\x9d\x23\x9e\x65\x42\xb4\x92\x1e\x9a\x0f\xe9\xfb\x8a\x93\xd3\xa1\x95\xf1\xc4\xd9\xf0\xd8\x93\x3f\x71\x32\x3c\x78\x72\x98\x0b\x83\x30\xfb\xe9\x93\xe1\x89\xef\xc9\x7e\xe4\x16\x04\xa6\x04\xdb\xa3\x4b\xe4\x21\xb1\x01\xd0\x88\x73\x78\x06\x1a\xd1\x21\xf0\xd9\x4e\x95\x60\xf1\xda\xa1\xa8\x74\x37\x20\x41\x13\x12\x10\x6a\x6f\x3c\x35\x75\x74\xec\x01\x92\x21\xcf\x4e\x58\x01\xe9\x81\x3c\xb0\xd5\x61\x4d\x83\xed\x72\xf7\x2e\x75\xcc\x14\xbb\x49\xc1\xdd\xbe\x22\xc7\xae\xd7\x3a\x4b\x62\xc5\x2a\x24\x1d\xe8\x6c\x51\x74\xc3\x24\xb0\x17\x48\x03\x53\x00\x8d\x76\x02\xad\x89\x4d\xed\x94\x16\xd2\x33\xba\x25\xf6\x25\x6e\x9f\x9d\xc2\xd0\xc1\x82\xbb\x28\x8c\x1d\x63\xe1\x2c\x26\x00\xec\x7f\xf0\xa6\x36\x2e\x67\x4f\x45\xe6\x38\xe9\x37\x68\x1d\xfb\x75\xed\x7d\x08\x1e\xa7\x6b\x3c\xaa\xbf\x2b\x32\xcf\x8a\xf4\xf1\x51\x70\xe8\x0f\x83\xff\x7b\xf0\x3f\xfd\xa2\x58\xd4\x04\x55\x7b\x89\x50\x34\x9c\xa1\xec\xd1\x8c\x82\x94\x06\x7a\x34\x4c\x0c\x17\x68\xe2\x80\x28\xe0\x6a\x61\x90\xd2\xf0\x83\xa1\x2b\xa8\xe8\x52\x84\xf3\x18\x44\xd1\x18\x31\x6c\x94\x18\x92\x13\x1a\x7a\x30\xc0\x5a\xbd\x21\x35\x44\x34\x77\xe8\x3d\x5c\x92\x30\xe0\x03\x82\x4a\x13\xec\xe3\x88\x60\x06\x78\x63\x38\xc2\xd0\x01\xe8\xda\xc6\x8f\x41\xc9\xd1\x02\x86\xb5\x83\x5c\x3b\x30\xe4\x44\x13\x80\x36\xec\x42\xc3\x55\x87\x41\x5c\xd2\x40\x81\x21\xc8\xd2\x5c\x4c\x31\x48\x1d\xfd\x21\x5e\x71\x48\xf6\x67\x09\x5d\xd6\xa0\x27\x73\x83\x32\x28\x0e\x3e\xe1\x12\x00\xfa\xc2\xe5\x33\x76\x71\x41\xcc\x21\xdb\x0e\x25\x39\x0d\x66\x68\xf0\xb0\x57\xba\x41\x1d\x17\xa3\xcc\xc5\x06\x97\x56\x7b\xb9\x96\x77\x69\x06\xd4\x85\x4b\x0a\x06\x82\xc2\x0d\xdd\x4f\xa5\x99\x1f\x6a\xff\x63\x64\x8f\xa7\x98\x87\x47\x84\x31\x7a\x9e\x8d\x5e\xfe\xb4\x31\x62\xac\xff\x07\x6a\xc9\x7e\x3e\xb5\xdc\xfa\xe3\x57\x7f\xfc\xc7\xaf\x1f\x19\x29\xee\x1b\x60\x1f\xda\x7f\x6f\x4d\x9f\x34\xeb\xa9\x05\x7d\x52\xf0\x4d\x35\x8f\xf8\xbb\xe2\x7d\xa5\x7c\x52\xc3\xf8\x3a\x1e\x1b\x85\x67\x2d\xe2\x51\x6e\xbb\x0a\x0e\xbe\xb7\x7e\x8b\x34\x8f\xc2\x34\xfd\x5f\xd4\xef\x8f\x75\x45\x0c\x79\x75\x48\x43\xa3\xbd\x96\xa1\x2b\x26\xf4\x75\x03\xd8\x5e\xbb\xaf\x7d\x70\x8d\xcb\xdc\xd7\x91\x5a\xd1\xd7\x21\x69\x46\xd7\x55\x2d\xfe\x1b\x00\x00\xff\xff\xf7\x25\x84\x82\x00\x20\x00\x00")

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

	info := bindataFileInfo{name: "bindata.go", size: 20480, mode: os.FileMode(420), modTime: time.Unix(1600215456, 0)}
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
	"000003_rename_prefSet_to_engine.down.sql":      _000003_rename_prefset_to_engineDownSql,
	"000003_rename_prefSet_to_engine.up.sql":        _000003_rename_prefset_to_engineUpSql,
	"000004_add_active_feed_to_user.down.sql":       _000004_add_active_feed_to_userDownSql,
	"000004_add_active_feed_to_user.up.sql":         _000004_add_active_feed_to_userUpSql,
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
	"000003_rename_prefSet_to_engine.down.sql":      &bintree{_000003_rename_prefset_to_engineDownSql, map[string]*bintree{}},
	"000003_rename_prefSet_to_engine.up.sql":        &bintree{_000003_rename_prefset_to_engineUpSql, map[string]*bintree{}},
	"000004_add_active_feed_to_user.down.sql":       &bintree{_000004_add_active_feed_to_userDownSql, map[string]*bintree{}},
	"000004_add_active_feed_to_user.up.sql":         &bintree{_000004_add_active_feed_to_userUpSql, map[string]*bintree{}},
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
