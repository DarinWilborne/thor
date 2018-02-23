// Code generated by go-bindata.
// sources:
// compiled/Authority.abi
// compiled/Authority.bin-runtime
// compiled/ERC223Receiver.abi
// compiled/ERC223Receiver.bin-runtime
// compiled/Energy.abi
// compiled/Energy.bin-runtime
// compiled/Executor.abi
// compiled/Executor.bin-runtime
// compiled/Params.abi
// compiled/Params.bin-runtime
// compiled/Token.abi
// compiled/Token.bin-runtime
// DO NOT EDIT!

package gen

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

var _compiledAuthorityAbi = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x95\xcb\x4e\xf3\x30\x10\x85\xdf\x65\xd6\x5e\xfd\x3f\x62\x91\x5d\x05\x88\x15\x1b\x58\x56\x55\xe5\xc4\x53\x61\x29\x9d\xa9\xe2\x71\x68\xa8\xfa\xee\x28\xe0\x5c\xd5\x90\x04\x51\xda\x5d\x2b\x1d\xe7\x7c\x67\x2e\xf6\xf2\x00\x09\x93\x13\x4d\x02\x91\x64\x1e\x15\x58\xda\x79\x71\x10\x2d\x57\x0a\x48\x6f\x11\x22\x20\x2d\x36\xc7\x3b\xf6\x24\xa0\x80\xbd\x04\xc5\xa1\x12\x80\x02\x29\x76\xe5\x2f\x6f\x49\x6e\x6f\xe0\xb8\x52\xb0\xd3\x85\x8e\x53\x84\x68\xa3\x53\x87\x0a\x9c\x68\xc1\x27\x2f\x3a\xb6\xa9\x95\x02\x22\xc8\x2d\xbe\x35\x47\x37\x9e\x12\xb1\x4c\x70\x54\x93\xa8\x92\xcb\xf0\x84\xd3\x35\x50\x6d\xba\xd6\xc6\x64\xcd\xf1\xf2\x1f\x3a\xf7\x69\x1d\x14\x06\xb5\x97\x57\xce\xec\x3b\x76\xb8\x27\xc1\x11\x53\x25\xfa\x31\xe2\x00\xa1\xaa\x05\xd6\x20\x49\xe9\x57\x8b\xe2\x42\xd0\xfd\xff\xd7\x8e\xf1\x35\x0e\x0b\x63\xc6\x8a\x1f\x33\xa7\x13\x4b\x3f\x23\x5d\x6f\x20\x66\xd4\xbf\xb4\xf5\xee\x34\x75\x6a\x9d\xa0\xe9\xb3\x4f\xaa\x8c\x1a\x98\xbb\x50\xb5\xf3\xec\xc1\x58\x4f\xfb\xed\x7a\xf9\x26\xfb\x70\xea\x8b\xa5\x9d\xbd\x65\x0d\xd6\x7a\xd2\x10\x5f\xe7\x26\xf6\xbb\xf6\x8c\x5b\xce\xf1\x2a\xf6\xac\x8f\xf6\x88\xf2\xb0\xc7\xc4\x0b\x67\x63\x7c\xed\x78\xe7\x7e\x15\xf0\xaf\x99\x34\x31\x15\x5b\xf6\xee\x54\xb7\x2d\x19\xdc\xa3\xa9\x78\x47\xaf\xe1\x5a\x1f\x3e\x35\xe7\x5a\x5e\xb4\x26\x3a\xa8\x30\x47\x92\x5f\xa6\x6c\x0c\xef\x3b\xcf\x59\xd7\x72\xf5\x11\x00\x00\xff\xff\xb4\x4e\x3e\xa1\x5a\x08\x00\x00")

func compiledAuthorityAbiBytes() ([]byte, error) {
	return bindataRead(
		_compiledAuthorityAbi,
		"compiled/Authority.abi",
	)
}

func compiledAuthorityAbi() (*asset, error) {
	bytes, err := compiledAuthorityAbiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Authority.abi", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledAuthorityBinRuntime = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\x09\x6e\xc3\x38\x0c\xfc\xd2\xf0\x14\xf5\x1c\x1d\xd6\xff\x9f\xb0\x90\x65\xa7\x89\xd3\x24\x2d\xd0\xdd\x8d\x51\xc0\x66\xa9\xd1\x90\x1c\x8a\x72\xcc\x47\x61\xec\x80\x8a\x13\x9c\x80\x28\x96\x5c\xc6\xf1\x73\x6c\x70\x80\x51\x1c\x80\x18\x94\xdc\x05\xc6\x55\xa0\x1c\x44\xba\xaf\x19\x96\x02\x2e\x70\x77\x2a\xb5\x2f\x6b\x3d\xac\x9c\x5a\x4e\xa3\xd8\xb2\x76\x5e\x56\x67\x46\xaa\x7c\x20\x0c\x39\xac\x6a\x35\x2a\xd5\xdd\x4a\x9c\x97\x35\x69\xab\xa5\xf9\x42\xa0\xa4\xcb\x9a\x3d\x6f\x64\x45\x96\x35\x1f\x08\xa5\x6e\xa9\x78\x1a\xcb\x5a\xed\xb0\xb6\xaa\x8d\xb1\x2d\x6b\x3f\x10\x9a\x68\x43\x6c\x0b\x97\x21\x96\xac\xce\x38\x03\xa3\x5b\x15\x25\x9b\xdc\xf2\xcc\xc8\xcd\x3a\x2d\x85\xa7\x3f\xb9\xf9\xf4\x57\x18\x79\x1a\x97\x5f\x46\x26\xf2\xa0\x99\x5d\x06\x68\xf9\x05\x32\x41\xf2\x8c\xf7\x86\xdf\x5e\xe1\xd7\x89\x7f\x7a\xf5\x7e\xf5\x1a\xe4\x3b\x6e\x39\x2a\x04\xd9\xeb\x68\xe4\x73\x75\xf0\x5c\x0d\x7c\x21\x8c\xed\x82\x40\x64\x2f\x11\xc0\x2a\x73\x95\xd4\xf8\x8a\x32\x83\x8c\xec\x07\x31\x91\xe8\x75\x2f\x8d\x37\x6c\xa5\xe1\x6e\x17\xde\x77\x91\xb5\x4b\x08\x28\xd3\xcc\xa6\xf1\x97\x2e\x27\x43\x45\x3c\xfe\x17\x8e\xf9\x06\xca\x2f\x58\xa5\xf1\x2b\x56\xaa\xe3\xae\x02\x94\xaf\xf9\x7b\x57\x81\x33\x7f\x6a\xf9\x1e\xa3\xe1\x17\x35\x20\x58\x7d\x60\xd0\x9f\xf8\x6f\xe9\x49\x89\x17\xb4\x1f\xeb\x90\xf1\xa4\x8f\x1d\xdd\x9a\x2d\x74\x20\xe3\x7c\x13\x5c\xf7\xb9\x3f\x17\xa6\xc7\xc9\x66\xee\x39\xf7\xde\x77\xa5\xc7\x0a\x9e\x67\x0b\x78\xf1\x83\x62\x5f\x71\x72\x0c\x81\x04\xed\x9c\x52\x40\xea\xd4\xc5\x64\xea\x72\x61\xca\xcd\xad\x40\x06\x9d\x1e\xe9\x41\x7f\x86\xf9\x9c\xa8\x46\x99\xe6\xf7\x8a\xe6\xfb\x48\xce\xf3\xe2\xdf\x8e\xa4\xc5\xa7\x48\x7a\x7e\x1b\xc9\xfe\x75\xe5\x2f\xf2\x6c\x23\x3d\x10\x05\x7c\x8f\xf8\x22\xfe\xe3\x14\x8d\x97\xb5\xbc\x9b\x0c\xb7\x4c\x84\x4e\xa7\x23\x03\xdf\xa8\x70\xe6\x25\x78\x61\xf0\xcf\x32\x24\x96\x3f\x64\x48\xbc\x7c\xca\xd0\xe9\xf9\xd8\xff\xf1\x4d\xe4\x69\xa4\xbc\xf5\x4d\x42\x72\xeb\xa9\xa4\x1e\xb5\xa7\xb4\xe5\xb4\x59\x8b\x5d\x30\x63\xd3\x91\xbc\xf5\x5a\xb7\x52\xc4\xd5\xd5\x37\x8e\x12\x9e\xad\xa4\x08\x3d\xb3\x73\xdf\x65\x85\xed\xd6\x39\x99\x97\xfa\xce\xef\x9d\xc5\x8b\x7e\x3a\xe7\x5e\xd8\x7d\x15\xfc\xcf\xab\xe0\x9f\xaa\xa0\xd4\x3e\x54\x41\xb9\xbf\xaf\xc2\x52\xcf\xd3\x3b\xe7\x94\x29\xbb\x21\x53\x56\x43\xe6\xb5\x72\x66\x67\x57\x36\x32\x82\xd6\xdf\xff\xde\xad\x9a\xaf\xe7\xef\x53\x16\x2a\xfe\xb8\x5b\xf5\xb1\xff\xaf\x9e\xc1\xe4\x64\x01\xb2\x3c\x6f\x6f\x3a\x60\xc9\x10\xbb\xe2\xed\xc6\x7b\xd4\xcf\x1d\x7f\xde\xc7\x42\x42\x7e\xd5\xf3\xf6\x46\x6d\x7c\xa8\x4d\x4e\xb5\x9d\xba\xd3\x9f\x65\xdc\xec\xd3\xf9\x68\xfe\xf1\x7c\x3c\x3d\xd3\xc3\x7c\x8b\xa7\x49\x39\xbb\x3f\xe6\x5d\xb5\x96\x16\x16\x35\x86\x74\x8f\xcd\x2c\x15\xe9\x4d\x6c\xcb\x01\x68\x0a\xb6\x5e\x46\x23\x2f\x2c\x39\xab\xa3\x67\xd4\x68\x86\x36\x44\x82\xcf\x9b\xd2\xab\x89\x5b\xf8\xaa\xef\xf7\x33\xf5\xbf\x9a\x44\xf7\x33\x15\x28\xf3\x4a\xc1\xa9\xa4\x92\x58\x60\xb3\x6c\xa3\x24\x78\x6c\x30\xea\x85\x6a\x1f\x00\x57\x9e\x67\x14\xab\xcb\xd6\x61\xa5\x5b\xd3\xda\x8a\x16\x29\xa3\x2b\xf5\xcd\x94\x4b\x4b\x4d\xca\x7e\xff\xe4\xfc\x4f\x00\x00\x00\xff\xff\xf4\x06\x42\x55\x72\x0c\x00\x00")

func compiledAuthorityBinRuntimeBytes() ([]byte, error) {
	return bindataRead(
		_compiledAuthorityBinRuntime,
		"compiled/Authority.bin-runtime",
	)
}

func compiledAuthorityBinRuntime() (*asset, error) {
	bytes, err := compiledAuthorityBinRuntimeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Authority.bin-runtime", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledErc223receiverAbi = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xce\x41\x8a\x84\x30\x10\x85\xe1\xbb\xbc\x75\x56\x03\x33\x8b\x1c\x60\x76\x73\x02\x91\xa1\xa2\x25\x04\x63\x45\xcc\x4b\x43\x10\xef\xde\x34\xb4\x4a\x2f\x8b\xfa\xe0\xfd\xdd\x8e\x21\x5b\xa1\x18\xe1\x27\x49\x45\x1d\xa2\xad\x95\x05\xbe\xdb\x61\xb2\x28\x3c\xfe\xa7\x2d\x2f\x70\x60\x5b\x5f\xa7\x8c\xe3\xa6\xa5\xe0\x70\xb7\x78\x48\xaa\x7a\x93\x1a\x8d\x5f\xdf\x3f\x1f\x64\x14\xca\x2d\x42\xa3\x16\x1c\xbd\x3b\xff\xcc\xb3\xda\xaf\xa4\x14\x64\x98\xe1\x90\x2b\xdf\x1d\xbd\xc3\x2a\x4d\x42\xd2\xab\xb1\x50\xa8\x7f\x95\x12\x62\x8a\x6c\xf0\xb0\x6c\x27\xba\x36\xa6\x6a\x03\x63\x36\x1c\xfd\x33\x00\x00\xff\xff\x74\xa5\xc3\x0b\xe9\x00\x00\x00")

func compiledErc223receiverAbiBytes() ([]byte, error) {
	return bindataRead(
		_compiledErc223receiverAbi,
		"compiled/ERC223Receiver.abi",
	)
}

func compiledErc223receiverAbi() (*asset, error) {
	bytes, err := compiledErc223receiverAbiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/ERC223Receiver.abi", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledErc223receiverBinRuntime = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func compiledErc223receiverBinRuntimeBytes() ([]byte, error) {
	return bindataRead(
		_compiledErc223receiverBinRuntime,
		"compiled/ERC223Receiver.bin-runtime",
	)
}

func compiledErc223receiverBinRuntime() (*asset, error) {
	bytes, err := compiledErc223receiverBinRuntimeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/ERC223Receiver.bin-runtime", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledEnergyAbi = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xc1\x6e\xdb\x30\x0c\xfd\x17\x9d\x7d\x1a\xb6\x62\xc8\xad\x1b\xb6\x9e\x8a\x01\x4d\x6f\x45\x51\xd0\x36\xd3\x6a\x90\x25\x43\xa2\x9c\x18\x45\xff\x7d\x48\xe2\xd8\xce\x26\x59\x4a\xe7\xda\xbe\x19\x30\x25\xf1\x91\x8f\x7c\x94\x1e\x5e\x59\xa6\xa4\x21\x90\xc4\x56\xa4\x2d\x26\x8c\xcb\xd2\x92\x61\xab\x87\xc7\x84\x49\x28\x90\xad\x98\x04\xe2\x15\xde\x20\xdd\x2b\x02\xb1\xb6\x65\x29\x6a\x96\x30\x65\xa9\x31\x7d\x3d\x59\xb2\x84\x51\x5d\xee\xbf\x2c\x97\xf4\xe9\xcb\x15\x7b\x7b\x4c\x58\x09\x35\xa4\x02\xd9\x6a\x03\xc2\x60\xc2\x0c\x01\xe1\xad\x25\x48\xb9\xe0\x54\xb3\x15\xab\x38\x6e\xbb\xb5\x1b\x2b\x33\xe2\x4a\xb2\xb7\x24\xd2\xbf\x02\x43\xfe\x18\xd2\x5c\x3e\x47\xba\x53\x5a\x8d\x41\x77\x9a\xd5\xad\x3f\xed\xa1\x4f\x1a\x33\x8e\x15\xea\x6e\x0b\xc8\x73\x8d\xc6\x1c\x76\x38\x59\x41\xa1\xac\x24\x77\xc4\x1a\x1b\x28\x4b\xad\x2a\x0f\x36\x63\xb3\x6c\xbf\x67\xbb\x41\xaa\x94\x88\x04\x28\x95\x3c\x19\x5d\x18\xf5\xf6\xf8\x3d\x24\x07\x40\x07\x6b\xbe\x81\x00\x99\x05\x33\x34\x25\x63\x68\x56\x22\xfb\x99\xb3\xd1\xaa\x18\x66\x0d\xa9\xe1\xff\x99\xc6\x9c\xbb\x58\x95\x9c\xf1\x53\x55\xa8\xeb\x3b\x20\x1c\xb6\xc4\x5d\xc9\x35\x1c\x20\x9c\xd9\x5d\x7d\xee\x67\xda\xbc\x80\x3e\x4f\xef\xd8\x24\xf4\x47\x2c\x53\x92\x34\x64\x74\xed\xa6\x63\x0f\x8b\xc4\xed\x2d\x18\x72\x96\x65\x0f\x0b\xd2\xf7\x66\xcb\xd6\x7a\xd9\xb8\x42\x8c\x88\xe8\x33\xa4\x41\x9a\x0d\xea\x9f\x5a\x15\x27\xf4\xb3\x36\x9d\x8f\xac\x90\x0b\xe3\xb1\xa4\xe6\xdb\x79\x98\x63\xc6\x0b\x10\x26\xa6\x7b\x7d\x9d\x44\xf5\xb4\xbb\x99\xf4\xc4\x2c\xff\x6d\x0d\xdd\x68\xb5\xa5\x97\xa6\xf3\xcc\x50\x58\xd1\x75\x55\x84\x7b\xc5\x51\xe1\xd6\x93\x76\x0c\x9f\x1c\x3f\xa9\xad\x0c\xb8\x9b\x1e\x75\xf8\xd7\xc6\x4d\x9a\xb4\x95\xe9\xc9\x04\xb9\x3d\x3b\x58\xd5\xce\xa2\x76\xcc\x1a\xeb\x17\xd8\xcf\x79\x77\x58\x00\x97\x98\x2f\x49\xdc\x3d\x03\x53\x12\x97\x80\xbf\x29\xe7\x1a\xaa\x26\xe3\x5a\xa0\x88\x3a\x77\x9f\x07\x6b\xc3\x91\x90\xfe\x1e\x1f\x3d\x04\x9a\xba\x48\x95\x98\xfe\xe2\xe0\x0b\x6b\x74\x54\x5b\xb6\xcf\x1e\x5b\xbf\x4e\x8f\xa8\xc2\xcb\x54\xe0\x36\x09\x3f\x76\x98\x59\x52\x8b\x0a\x7f\x50\x8c\x8f\xde\x5f\x4f\x29\xc9\xde\x48\xe2\x02\x03\xf8\x3e\x41\xea\x7e\x47\x5c\xc3\xa2\x6f\x61\x71\x97\xb0\x56\x19\x1a\x09\x9c\x47\x19\xde\x39\x9e\x9f\x29\xc6\xfc\x1a\x7e\xf1\x90\xd5\xc3\x67\x4a\x94\x79\x60\x12\x03\x21\xd4\xd6\xff\x26\xa2\x0f\xc8\x8f\x39\x1c\x1b\x24\x48\x25\xeb\x42\x59\xe3\x22\x3f\x97\x39\xee\x30\x6f\x7f\xc5\x76\x13\x47\x1f\x69\x8c\xb1\x42\x49\x17\x1c\x7c\x8c\x7c\xb0\x08\xdd\xf6\xbe\x9a\xf4\xc0\x1a\xaa\x51\x5f\x24\x82\x35\xeb\x59\x18\x57\xc3\xeb\xe6\x21\x65\x8c\xe0\x85\x2f\x3a\x1e\x57\x95\xc8\xbd\xef\x24\xfe\x55\x71\xaf\x2b\xae\xbb\xd2\x18\x60\x07\xfa\x8e\x67\xc1\x85\x5c\x79\xaa\x40\xd8\xe1\x22\xb8\xef\xc6\x95\x51\x20\x0d\xf4\x1a\xcf\x0a\x7f\xeb\xf9\x2f\x60\xd7\x87\x57\x68\x10\xff\x00\x7b\xfc\x13\x00\x00\xff\xff\xe6\x2f\xe9\xc8\x3f\x18\x00\x00")

func compiledEnergyAbiBytes() ([]byte, error) {
	return bindataRead(
		_compiledEnergyAbi,
		"compiled/Energy.abi",
	)
}

func compiledEnergyAbi() (*asset, error) {
	bytes, err := compiledEnergyAbiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Energy.abi", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledEnergyBinRuntime = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x07\x96\x2c\xa7\x0e\xdd\x92\x02\x02\xb1\x1c\xe2\xfe\x97\xf0\x8f\x80\xea\xae\xd4\x61\xc6\xef\xfb\x79\xfa\x78\xdc\xa3\x57\x25\xa4\xab\x74\x01\x0f\xf6\x71\x20\xe4\x01\x1c\x7b\x04\x8f\x80\xd4\x25\x78\xee\xeb\xc7\x43\x03\x0f\x40\x90\x3c\x00\xb0\x80\x43\xef\x19\x5c\xc1\xd6\x1d\x2a\xa2\xb3\x77\xd8\x49\x50\xf0\x0c\xbe\xd7\xda\x80\xa7\x54\xe2\x92\x46\x69\x29\xe4\x25\x6d\x3c\xa5\x88\xb1\x94\x26\x53\x4a\xb8\x9e\x45\x45\x0f\xb5\xd6\x29\x65\x5d\xd2\xd6\x7c\x83\x46\x53\xea\xf2\x92\xf6\x96\x73\xf2\x7d\x4a\x55\xa6\x94\x90\x7c\x4a\x10\xa7\x34\xa5\x25\xe5\xac\x81\x36\xbd\x95\xa6\x94\x91\x4b\x13\x1f\xa6\xb4\xaf\x67\x1d\x52\x29\xbd\x4f\xbd\x4c\xcb\x5e\x9f\x95\x31\xa4\x25\xe5\x65\x6f\x80\x04\x4a\x3c\x71\x60\x69\x4b\x5a\x35\x00\x7b\x99\xd2\x50\x1f\xd2\x1e\x28\xcf\xd5\x38\x2d\x1b\xb4\xa4\x96\xd4\xfb\x29\x2d\x0b\xc9\x28\x55\x63\x76\x4b\x6f\xef\x4b\x9a\x5d\xc6\x5c\x26\x0e\x0e\x97\x86\x14\x41\x62\xc9\x79\x4a\x19\x97\xb4\x64\x57\x08\xda\x94\xca\xf2\x22\x23\x73\xd5\x3a\xbd\x70\xde\x4f\x69\x61\x57\x40\xdb\xb4\xd7\x85\xb2\x49\x93\x22\x4b\x9a\x52\x5d\x36\x54\xd6\x1e\x9c\x2e\x69\x09\x4b\x5a\x3d\xb5\xca\x6b\xb5\x56\x24\x48\xb6\x8c\x51\xe8\x55\x32\x3b\x94\x91\x27\x96\x5b\x0f\xa9\x49\x5c\xf0\x08\x82\x28\xde\x9e\x77\x20\x18\x41\xd1\x32\x92\x00\x70\x4a\x14\x22\x02\x47\xe8\xfc\xd4\xe4\xdd\x59\x93\x2f\x43\x93\x7f\x6a\x32\x1d\x0a\x4a\x42\x3a\xb4\x02\x2a\x2b\x2a\x0a\x3e\x57\x88\x28\xa0\x63\xd5\xf5\xb7\xbd\xc1\xca\xa6\x5b\xb2\x3d\x6f\x89\x6c\xfa\xd3\xc8\x45\x25\xb0\xf7\x79\xfc\x7f\xb3\x12\xed\x3d\x5b\x57\x60\x7e\xe2\xf8\x4f\x87\xdc\x03\x76\xf4\xba\xb4\x54\x59\x5a\x58\x87\x85\x66\xd3\xd0\x62\xe5\x87\x00\x90\x80\x31\xa2\x3f\x5a\x68\x9a\x23\x4d\xdd\xaf\x31\x69\xed\x84\x09\x81\xf8\x81\x62\x5a\x35\x0c\x3c\x2a\x5d\xd0\x7b\x20\xc7\xf6\x96\x48\xd8\x63\x8f\x82\xf2\x05\xfe\x44\x17\xfc\x5d\x78\xb9\x16\x82\x94\x11\xe1\xed\x6d\xc7\xf7\x79\x50\xc2\xfe\x29\xf1\x67\x7f\x46\x5c\x6e\xd6\x50\x44\x3f\x22\xb8\x7c\x73\xe6\x1b\x78\xfb\x1d\xfa\xe9\xc7\x83\x6e\x56\x79\x6a\xb6\x9e\x45\x7a\x5b\x33\xc2\x2f\xd7\x44\x08\xbd\xee\xad\xcf\xf2\x6d\x34\x6e\xad\x47\x88\x2e\xed\xf5\xd5\xfa\x4f\xf5\x75\xda\xe9\x63\x38\xdb\xc7\x50\x3d\x42\x0a\x75\x5f\x41\xbd\x47\x88\xb8\xcf\xc7\x97\x39\xc1\x74\xc9\x3f\x5d\x16\x99\x5e\x3d\xac\xee\xce\xf9\xf3\x3d\xce\x39\x1e\x34\xf9\xf8\xa3\x4c\xcc\xd1\xef\xdf\x56\xfd\xf6\xed\x8b\x1d\x05\xfa\x5e\x53\xba\xc4\xe7\x95\x47\xbb\xea\x3b\xf9\x52\xce\xfd\x91\xdb\x6b\x1d\x08\x05\xc3\x3e\x56\xc7\xe7\xbe\x8e\x9b\x83\x74\xdb\x4b\x4b\x80\x9d\x6d\x0e\xcf\xfe\xbd\xb7\xed\x58\xf1\x8e\xcb\x2f\x7a\x53\xc9\x07\x1d\x72\xce\x2f\xb3\xe0\x31\x3d\xb6\xa7\x02\xbe\xc9\xc2\x92\xdb\xfe\x59\x0d\xb7\x1a\xcb\xd1\xf6\x78\xc6\xe7\x4f\xf7\xa1\x0a\x71\xbf\x9e\xb1\x93\x5f\xe6\x64\xc5\x43\xcc\xfa\xd9\xbf\x1f\x68\x9a\x99\x69\x6f\xcf\xd9\x36\x90\x6e\x1e\x11\x39\x6e\x59\xa7\x6b\x5e\xaf\xec\x32\x1e\x99\x2c\xe3\x42\x17\xe7\xd5\xf7\x40\x64\x13\xab\x87\xe0\x25\x10\xfc\xe0\x67\xcc\xee\x31\x63\xe7\x2c\x5d\xb6\x1c\x2c\x67\x9b\xce\x7e\xf8\xf7\x98\xe8\xdb\x37\x37\xe6\x38\x41\x74\x1a\x6c\xfa\xaa\x13\x8a\xce\x6a\x42\xc8\xe6\x2f\x81\x4a\x04\x91\x08\x91\x22\x46\x08\x5d\x8b\xe4\x86\x4d\x5a\x6e\x25\x54\xc9\x15\x5d\x0f\xe8\x28\x54\x6c\xea\x3a\xd7\x6a\xbc\xb6\x40\x0f\x99\x28\x62\x93\x4c\x00\xa9\x68\x09\x85\xad\x8c\xe3\xd4\xf8\x86\xbf\x24\xb6\xe9\x0d\x68\x2b\x8e\x59\x3e\x18\xc3\x0e\x61\x63\xd8\x70\xf6\x72\xcf\xb7\xfd\x40\x66\xe3\x36\xb0\x3c\x35\x3e\xf3\x64\xec\x86\xc8\xc6\xd9\x81\xa6\x2d\xe0\x60\xbc\xb1\xd9\xa3\x0c\xac\x43\x9b\x06\x05\xce\x36\xf5\x6d\x1e\xc2\xb9\xff\x50\xf1\x92\x80\x3b\x6e\x4f\xd0\x61\x3e\x1e\x19\x89\x4c\xae\xb2\xbc\x51\x77\xf5\x84\x2f\x75\x83\x1e\x9d\xda\xbe\xc3\x57\x90\x20\x2f\xfc\xdf\x78\xaf\xfa\x57\x08\xec\xf6\x29\x0f\x24\xd4\x00\xdb\x10\xb8\xe9\x8d\x86\xcb\x83\xc7\x7d\x89\x50\xba\xf0\x83\x33\x42\x19\xdf\x22\x04\x72\xe3\xe1\x3d\x2e\xf2\x58\xb5\xe6\xbd\x4e\x47\xd7\x6e\xa2\x84\x1e\x8d\x89\xfa\x7e\xee\x7f\xbe\x27\x35\x56\x5e\x73\xb6\xb8\x0c\x8d\x01\x0e\xf3\xf7\x1e\xf5\x8d\xff\xab\xdf\x3e\x6f\xb0\xd6\x97\x58\x8b\xfa\x3d\xd6\xd1\xd0\x77\xb3\xcb\xab\x5b\x11\x74\xca\xc6\x74\x67\x7d\x5a\xc7\x5c\x71\xb9\xf8\xb9\x45\x4e\x1f\x91\x4b\x23\x72\xf0\x29\x72\x21\x5c\xe6\xf3\x29\x72\x41\xdb\x35\x72\xd7\x4e\x8f\x5e\x2f\x93\x0b\x7d\xe8\x8a\x25\x50\xd7\x50\x83\xcb\xc1\x97\xd0\x80\xc0\xa9\xcb\x49\x41\x7b\xf2\x62\x63\x24\x4a\x48\xb1\x08\x90\x6b\xd9\x33\x7a\x08\xd5\x75\x97\xba\xa4\x0a\x2a\xf6\x59\x3c\x9c\x94\x57\x3f\xdb\xe3\x72\x41\x63\xcc\x98\xd1\xe9\x0e\xe8\x81\x7d\x1b\x7b\x92\x43\xef\x79\x7c\xde\xf6\x9a\x47\xad\xb9\xbf\x5c\x6b\x7a\xe5\xff\xa7\x88\xa9\x0f\x9f\xbb\xd1\xb5\xb6\xac\x43\xdb\xb4\x1b\x35\x63\x95\xa1\x9a\x3f\x57\xc3\xb6\xff\x57\x67\xc8\xbc\xc4\x41\x5e\xe2\x40\xca\x87\x3a\xc0\x55\x03\xb4\xd5\xc0\x57\x79\xac\xed\xc2\x6f\xce\xa8\x74\xba\xc9\xe3\x8b\x3d\x96\xb3\x2c\xbd\x48\x13\x97\x7c\x68\x99\xa8\x2b\x72\x65\x2c\xa0\x3e\x7a\x8f\x45\xd4\x91\x54\x27\x14\x90\x5d\x95\x6e\x65\x49\x81\x21\x23\x85\xaa\xae\x15\xa5\x27\x16\x77\xde\xea\xd3\x4f\x78\xf8\xf9\xd8\x19\xef\xb2\x93\x7e\x94\x99\xf2\x97\x33\x33\xa6\xf3\xfe\xf5\x1c\x83\x98\xdd\x9f\x9a\x02\x9b\xc6\xcb\xfe\x2f\x36\xb5\x5c\x9c\xbd\xbd\x59\xb5\x64\xd3\x2b\x39\x72\x3c\xe1\xa9\x88\x32\xa6\x6c\xb2\x7d\xfb\x75\x6d\x05\x95\xc9\xa4\x9e\xec\xe5\xc2\xa4\x78\x30\x29\x16\x8a\xc6\xc9\x8c\xd9\x39\xa5\x08\x08\x6b\xa6\x24\xc6\xa5\x1b\x86\x75\x79\xb1\xc6\xc4\xe5\x6c\xe7\x7c\x3e\xf0\x0b\x5b\xdc\x17\xb6\xf8\xa3\x2d\xba\xac\xb1\xa8\xca\xec\xcf\x0f\x1b\x16\xc7\x3a\xa1\x82\x34\x79\xca\x7d\xae\x6d\xa7\x67\xff\x67\xc6\x95\xca\x79\xff\x79\xce\xa4\x54\xe3\x1f\xce\xa4\x0c\xf4\xb9\xdb\x6d\xe7\x84\xfa\x6d\x65\xb9\x51\xd3\x8f\x19\x44\xdf\x75\xb3\xec\x3e\xf9\x9f\xe5\xc6\xff\xd0\x5d\xa3\x96\xd9\x65\xa8\xd1\x17\x4f\x8d\x40\xd0\x71\x23\x8e\xd2\x03\x50\x73\xe4\x02\x53\x8f\x04\xa0\xb9\x9b\xcf\x39\x75\x57\x02\xb3\x96\x54\xea\xa0\x2b\xef\x4f\x1b\x13\xca\xf3\x4c\xef\x6d\x5f\xda\x4e\xb2\xe7\xa9\xe1\x5f\xec\x4b\xb9\x7d\xe2\xef\x05\xde\xf3\xf7\x4b\xe7\xd8\xef\x55\xbe\xe8\xcc\xff\x3d\x04\x4a\xd0\x8f\x7b\x56\xb7\xed\x59\x9d\xba\x2e\x3f\xda\xa7\x7e\xb1\x67\x5d\xdd\x90\xf9\xd8\x03\x3f\x61\xfa\x2f\x75\xa0\xd3\x9e\xef\x5b\x96\xb8\xdd\x70\x4c\x36\xf4\x83\x98\xff\x43\x7e\xf4\xd1\x9f\x7a\x3d\xff\x39\xd5\x40\xd5\x1b\x7e\xb4\xab\x01\x17\xf9\x84\xc1\xd1\x4e\x34\xde\xf8\x69\x3e\xcd\x3d\x8e\x92\xd0\xd6\x13\x6d\x5e\x3e\x73\x42\xaf\x73\xcf\x7a\xf4\xb2\xb0\x2e\x5e\x6b\xb5\x37\x32\x66\x9c\xc5\x28\x71\x5e\xfd\xde\x26\xfa\xf3\x66\x60\x4c\x81\x83\xe6\x4f\x9d\x2a\xfc\xe5\x3a\x6d\xd7\xf3\xdf\x53\x94\x9a\xc8\xdb\x28\x91\xdc\x4c\x3e\x86\x83\x97\x7f\xfb\xb4\xa0\xe5\xf3\xf9\xdb\xc5\xcb\xe3\xfc\xbf\xdf\xc1\x8c\x7f\x77\x88\xeb\x0e\xa9\xb5\xc1\x9b\x94\x95\x10\x50\x24\x6f\xf2\x3e\x4e\x50\x74\xed\x8d\x11\x27\xcf\x42\xa4\x2c\xe1\x65\x15\x8f\x1b\x49\x0d\xea\x55\x80\x7f\x55\xb7\x10\x69\xa1\xc3\x1b\x3a\x3f\xdc\xd3\x74\x97\x3f\xe0\xd4\xa5\x5c\x71\xfa\xe0\x93\x57\xaf\x6e\x9b\xea\xff\xbe\x4f\xd7\xfb\x9f\xb3\x4f\xc5\xdf\xc4\x1e\xa1\x57\xdb\xe1\x1c\xce\x6b\xec\x0f\x09\x77\x27\x6a\x9e\x0b\xb4\x06\x59\x93\xc5\x70\x6f\x97\xc4\x71\x9f\xdf\x7b\xb0\xac\x90\x28\x79\xeb\x52\xd8\x31\xda\x6f\xf3\xcd\xf6\x68\x80\x3b\xfe\x43\xf1\x99\x81\xee\xf3\x84\xb1\x6c\xbb\xb1\xea\x7a\x66\xf2\x60\x59\xca\xbb\xef\xeb\xd6\x55\x49\x11\x78\xdc\xd4\xfe\xfa\x86\x16\x21\xc4\xd7\x37\xb4\x86\xc5\x37\x37\xb4\x08\xc9\xff\xfa\x86\xd6\x3d\xa7\xe7\xe7\x0c\x41\x28\x1f\xce\x37\x6c\x98\xdd\xcc\xe6\xdb\x93\x55\xbd\xc1\x3c\xf4\x5a\x3b\x09\xa5\x8a\xb9\x51\xd1\x98\x7d\x2c\x94\xc1\x6b\x2f\x1c\xb4\xa6\x14\x85\x72\x0a\x1d\x3d\x17\x97\x10\x3d\x69\x17\x49\xae\x76\x21\xce\xdc\xfa\x76\xe2\xf7\xee\x4c\x7b\x9e\x68\x5b\xe6\x22\xf2\x8e\x37\x4e\x4c\xe4\xb2\xef\xda\xba\xe6\x89\x83\x8d\x15\xe2\xb8\x14\x4a\xe8\xc5\x53\x48\x21\x05\x62\x10\x25\x48\xa9\xe6\xc4\xe2\x43\x6e\xb5\x65\xaa\xd8\x73\x08\xb5\x53\x08\x89\x1c\x26\x1f\xa5\xf6\xaa\x44\x28\xd1\x85\x86\x9e\x81\x4a\xaa\xad\xb2\x8b\xb1\xb1\x41\x43\xf1\x7f\x01\x00\x00\xff\xff\xfe\x4b\x6e\x54\xee\x22\x00\x00")

func compiledEnergyBinRuntimeBytes() ([]byte, error) {
	return bindataRead(
		_compiledEnergyBinRuntime,
		"compiled/Energy.bin-runtime",
	)
}

func compiledEnergyBinRuntime() (*asset, error) {
	bytes, err := compiledEnergyBinRuntimeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Energy.bin-runtime", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledExecutorAbi = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\x8e\x05\x04\x00\x00\xff\xff\x29\xbb\x4c\x0d\x02\x00\x00\x00")

func compiledExecutorAbiBytes() ([]byte, error) {
	return bindataRead(
		_compiledExecutorAbi,
		"compiled/Executor.abi",
	)
}

func compiledExecutorAbi() (*asset, error) {
	bytes, err := compiledExecutorAbiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Executor.abi", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledExecutorBinRuntime = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x14\xc1\x0b\x11\x00\x21\x08\x04\xd0\x4a\xeb\x22\x1f\xe3\x20\x60\xff\x08\x37\x37\xcf\xf0\xdb\x50\x1a\x80\xc0\x6b\x20\x97\xa9\xd1\xd3\xd3\x29\xd0\x20\x8a\xd2\x92\xc3\xae\xb0\xec\x25\x50\x5b\xc9\xe0\xe9\xf0\xf2\xba\xb1\x45\xc5\x74\x1f\xde\x5e\xba\x63\xae\x8f\xbf\xc2\x13\x13\xde\x39\x00\xcf\x17\x00\x00\xff\xff\xd7\x93\xa9\x33\x6a\x00\x00\x00")

func compiledExecutorBinRuntimeBytes() ([]byte, error) {
	return bindataRead(
		_compiledExecutorBinRuntime,
		"compiled/Executor.bin-runtime",
	)
}

func compiledExecutorBinRuntime() (*asset, error) {
	bytes, err := compiledExecutorBinRuntimeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Executor.bin-runtime", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledParamsAbi = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x92\x41\x4b\xf4\x40\x0c\x86\xff\x4b\xce\x73\xda\x8f\xcf\x43\xef\xe2\xc9\x93\xc7\xa5\x48\xda\x66\x65\xb0\x9b\x29\x4d\x52\x77\x58\xfa\xdf\x45\xb7\xed\x80\xcc\x32\x22\xb2\x78\x2c\x7d\x27\x79\xf2\xf0\xee\xcf\xd0\x06\x16\x45\x56\xa8\x0e\xd8\x0b\x39\xf0\x3c\x98\x0a\x54\xfb\x33\x30\x1e\x09\x2a\x78\x7e\xa5\x08\x0e\x34\x0e\x1f\x5f\x4d\x54\x92\x7f\x3b\x98\x5d\x0a\x4c\xd8\x1b\xa5\x88\x79\xd6\xdd\xff\x3b\x98\x6b\xb7\x46\x84\x14\x1c\x04\xd3\x65\x78\xed\x60\xc0\x88\x4d\x4f\xdb\x62\x51\x54\x7a\x34\xc5\xc6\xf7\x5e\x23\x54\xc0\x81\xd7\xd0\x36\xfb\x60\xdc\xaa\x0f\xfc\xb9\x3f\xc1\xeb\x68\x59\xf6\x3c\x7a\xe2\x62\x54\x3f\xd1\xc3\x17\xba\xed\x79\xfe\xa6\x32\xf9\xe4\xe9\xed\xc7\xcc\x57\x7c\x27\xe8\x97\xbf\x80\x9b\x71\x78\x7f\xa2\xd6\x34\x8c\x25\x38\xec\xba\x91\x44\x7e\x13\xee\x6a\x79\x4b\xdd\x2d\x57\xf7\x72\xde\xd3\x6d\x0b\x9c\xd6\xd3\xad\xa5\x22\x07\x8e\xc7\x60\x92\xb3\xea\xb9\xa3\x13\x75\x2b\x6f\x49\xf2\x16\x5f\x26\x7d\x5b\xfa\x45\xf7\xf2\x9f\x26\x62\x85\xb9\x7e\x0f\x00\x00\xff\xff\xd3\x6a\xf5\x71\xaf\x04\x00\x00")

func compiledParamsAbiBytes() ([]byte, error) {
	return bindataRead(
		_compiledParamsAbi,
		"compiled/Params.abi",
	)
}

func compiledParamsAbi() (*asset, error) {
	bytes, err := compiledParamsAbiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Params.abi", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledParamsBinRuntime = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x54\x6d\xae\x63\x2b\x0c\xdb\x92\x93\x90\x90\x2c\x07\x38\xb0\xff\x25\x3c\x05\x7a\xda\x77\x3b\xd5\xed\xfc\x9a\x22\x55\x3a\x56\x3e\x6c\x27\x60\xc8\x53\xa0\x6c\x40\x11\x23\x18\x01\x3a\xb5\x9a\xac\xc7\xcf\x30\x61\x00\xa3\x19\x00\x51\x14\x32\x13\xae\xb2\x4a\x14\x38\x51\xc9\x1c\x13\xad\x0e\x13\x6e\x6d\xf5\xa1\xf5\xa0\x75\x1e\xd4\x67\x6b\xd6\x06\x0e\xda\xec\xa0\x6d\xf4\x32\x18\xf3\xa0\x99\x95\x68\x17\x19\x97\xcf\x76\xd0\xd9\x0f\x3a\xa4\x0c\xf8\xd4\x8d\x12\x8a\x56\xed\xc9\xc7\xb1\x2e\xed\x52\x48\x37\x8b\x64\xfe\x44\x37\x83\xb1\x95\xa9\x81\xf3\x9f\x40\x54\xd5\xb4\x03\xaf\x2c\x8f\xf7\xac\x28\x8f\x2c\x02\x0b\x65\xfc\x76\x89\x02\x4e\xe9\x15\x03\x74\x10\x47\x10\x24\xb0\xe4\x55\xaf\xd3\xaf\xf5\x76\xff\x3b\x76\xd4\xf7\xd8\xb1\x32\xaa\xf9\xab\xab\xed\x6e\xed\x31\x85\xec\x16\x44\xf6\x37\x4c\x96\x7d\xf5\x83\xdb\xf5\x3f\x3e\x84\xf5\x99\x4f\xdf\x2e\x08\xde\xb9\xe4\x2e\xdc\x73\xcc\xbc\x9b\x71\xf2\x4a\x7e\x9b\x19\xbd\xb6\x89\xec\xb5\x4f\xe0\xa3\x01\x05\x3b\xe3\xd6\xe1\x02\x71\xda\x2c\xaa\x43\x3a\xe9\xe1\xa6\xd7\x1b\x37\x1e\xa6\x0d\xb2\xe8\x8e\xf8\x39\x7f\x45\x9e\xbb\x6a\x4e\x4f\x3f\xf0\x17\xf9\x13\xa3\x72\x57\x8c\x1f\xf3\xf9\xac\xff\xde\x58\x17\x97\x5b\xff\xf3\xce\x3c\x95\xfb\xb6\xe5\x56\xec\x59\x9b\x03\xc1\x7b\x8a\xc5\xf9\xe1\xd7\xf6\x02\x5f\xbd\x98\xf2\xcd\x8b\x55\x7e\x7a\xe1\x12\xa4\xa8\x8b\x7d\x0a\x17\x5b\xe9\x88\x2e\x1d\x34\xaf\xf0\xda\x49\x64\x2e\x5e\x14\x52\x24\x3a\x6b\x1b\xbd\x59\xd3\x69\xb1\x98\x62\x7a\xd8\x1a\x71\x51\x8d\x74\xd1\xf9\xdb\x7d\x68\xbc\xdd\xb7\xec\x9c\x2c\x02\x67\x9f\x81\xcf\x1e\xde\xef\x46\x3a\xf8\x79\x8b\x3e\x7a\x5a\xfe\xf0\x94\xf2\x76\x1c\x4f\xff\x6a\xab\xd8\xfd\x8b\x93\x1c\xf1\xeb\x56\x49\x1c\xad\x4f\x85\xf1\xd0\xad\x5f\x34\xff\xa3\x7b\xc3\xeb\xfd\x7d\x7b\x57\x28\x68\xbf\x2a\xa4\xfc\x4e\x55\x40\x23\x53\xe3\xda\x6a\xab\x2c\x50\x67\x70\x63\xc1\xe5\xa3\x5d\x56\xd1\xf3\x69\x36\x74\x1a\x24\xd5\xe8\xd2\xab\x11\xcc\x6c\x2d\xab\x4b\x0a\xcd\x7c\xc8\xfd\x1a\x66\x52\xea\x72\x1d\x29\x83\xe3\xbf\x00\x00\x00\xff\xff\xc7\x52\x5f\xd3\x88\x06\x00\x00")

func compiledParamsBinRuntimeBytes() ([]byte, error) {
	return bindataRead(
		_compiledParamsBinRuntime,
		"compiled/Params.bin-runtime",
	)
}

func compiledParamsBinRuntime() (*asset, error) {
	bytes, err := compiledParamsBinRuntimeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Params.bin-runtime", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledTokenAbi = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x95\x3f\x4f\xc3\x30\x10\xc5\xbf\xcb\xcd\x99\x90\x60\xc8\xc6\xc2\x86\x18\x60\xab\x2a\x74\x49\x2e\xc8\x92\x73\x67\xd9\xe7\x94\xa8\xea\x77\x47\x2d\x69\x5c\x41\xfe\x54\x05\x04\xac\xf1\x3b\xdf\x7b\xfe\x5d\xec\xd5\x16\x4a\xe1\xa0\xc8\x0a\x79\x8d\x36\x50\x06\x86\x5d\xd4\x00\xf9\x6a\x0b\x8c\x0d\x41\x0e\xcf\xc1\x11\x57\xe4\x21\x03\xed\xdc\xfe\x0b\x56\x95\xa7\x10\x60\x97\x25\x51\x8b\x36\x52\x92\x44\xc3\x7a\x75\x7d\x03\xbb\x75\x76\x94\xa0\x73\x5e\xda\xbd\x46\xa2\x7e\x6c\x12\x62\x59\xee\xb7\x1c\x36\x28\x44\xec\xa1\xda\x61\x87\x85\xa5\xc1\x60\x50\x54\xba\x8f\x8a\x85\xb1\x46\x3b\xc8\x81\x85\x8f\xa2\xa1\xbc\x8e\x5c\xaa\x11\x3e\x78\x4c\x21\xd5\xc7\xd3\x8c\xc9\x9c\x8a\xa2\x7d\x8c\xce\xd9\x6e\xca\x60\xbf\x36\x12\x70\xd9\x62\x6b\x68\xb3\x68\x6e\x9a\x40\xed\xa5\x99\x3f\x7e\x95\xaf\xe2\x51\x8f\x1c\x6a\xf2\x77\xef\xbd\xfe\x0c\xa3\x94\x41\x36\x3c\x3a\x85\x29\x43\x81\x16\xb9\xa4\x87\x7a\x3c\x40\xbf\xfc\x4b\x10\xbf\x0f\xd1\xff\xc2\x93\x9d\x73\x93\x9c\x5c\x13\xd6\xca\xa6\xc7\x34\x92\xd2\x53\x83\x86\x0d\xbf\xfc\x00\x45\x64\xe1\xae\x91\x18\xc6\x30\x1a\xae\xe8\x95\xaa\xe3\x01\x2c\xff\x9a\x13\x05\x53\x63\x30\xc8\xfb\xd6\xe7\x8f\xc5\x53\x1a\x8b\x5e\x44\x2d\xb1\x5e\x1c\x69\x06\xe4\x44\xc5\xec\x0b\x71\x79\xb0\xdb\xc3\x8b\x81\xf6\x53\xb0\xf5\x5b\x00\x00\x00\xff\xff\x73\x6f\x06\xdc\xba\x06\x00\x00")

func compiledTokenAbiBytes() ([]byte, error) {
	return bindataRead(
		_compiledTokenAbi,
		"compiled/Token.abi",
	)
}

func compiledTokenAbi() (*asset, error) {
	bytes, err := compiledTokenAbiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Token.abi", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledTokenBinRuntime = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func compiledTokenBinRuntimeBytes() ([]byte, error) {
	return bindataRead(
		_compiledTokenBinRuntime,
		"compiled/Token.bin-runtime",
	)
}

func compiledTokenBinRuntime() (*asset, error) {
	bytes, err := compiledTokenBinRuntimeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/Token.bin-runtime", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"compiled/Authority.abi": compiledAuthorityAbi,
	"compiled/Authority.bin-runtime": compiledAuthorityBinRuntime,
	"compiled/ERC223Receiver.abi": compiledErc223receiverAbi,
	"compiled/ERC223Receiver.bin-runtime": compiledErc223receiverBinRuntime,
	"compiled/Energy.abi": compiledEnergyAbi,
	"compiled/Energy.bin-runtime": compiledEnergyBinRuntime,
	"compiled/Executor.abi": compiledExecutorAbi,
	"compiled/Executor.bin-runtime": compiledExecutorBinRuntime,
	"compiled/Params.abi": compiledParamsAbi,
	"compiled/Params.bin-runtime": compiledParamsBinRuntime,
	"compiled/Token.abi": compiledTokenAbi,
	"compiled/Token.bin-runtime": compiledTokenBinRuntime,
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
	"compiled": &bintree{nil, map[string]*bintree{
		"Authority.abi": &bintree{compiledAuthorityAbi, map[string]*bintree{}},
		"Authority.bin-runtime": &bintree{compiledAuthorityBinRuntime, map[string]*bintree{}},
		"ERC223Receiver.abi": &bintree{compiledErc223receiverAbi, map[string]*bintree{}},
		"ERC223Receiver.bin-runtime": &bintree{compiledErc223receiverBinRuntime, map[string]*bintree{}},
		"Energy.abi": &bintree{compiledEnergyAbi, map[string]*bintree{}},
		"Energy.bin-runtime": &bintree{compiledEnergyBinRuntime, map[string]*bintree{}},
		"Executor.abi": &bintree{compiledExecutorAbi, map[string]*bintree{}},
		"Executor.bin-runtime": &bintree{compiledExecutorBinRuntime, map[string]*bintree{}},
		"Params.abi": &bintree{compiledParamsAbi, map[string]*bintree{}},
		"Params.bin-runtime": &bintree{compiledParamsBinRuntime, map[string]*bintree{}},
		"Token.abi": &bintree{compiledTokenAbi, map[string]*bintree{}},
		"Token.bin-runtime": &bintree{compiledTokenBinRuntime, map[string]*bintree{}},
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

