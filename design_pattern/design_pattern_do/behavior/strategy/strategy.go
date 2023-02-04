package strategy

import (
	"fmt"
	"io/ioutil"
	"os"
)

type StorageStrategy interface {
	Save(name string, data []byte) error
}

type fileStorage struct {
}

func (f fileStorage) Save(name string, data []byte) error {
	return ioutil.WriteFile(name, data, os.ModeAppend)
}

type encryptFileStorage struct {
}

func (e encryptFileStorage) Save(name string, data []byte) error {
	// 加密
	data, ok := encryptData(data)
	if ok != nil {
		return fmt.Errorf("encryData error")
	}
	return ioutil.WriteFile(name, data, os.ModeAppend)
}

func encryptData(data []byte) (res []byte, err error) {
	// todo 加密数据
	return data, nil
}

var storageMap = map[string]StorageStrategy{
	"file":         fileStorage{},
	"encrypt_file": encryptFileStorage{},
}

func newStorage(name string) (StorageStrategy, error) {
	storage, ok := storageMap[name]

	if !ok {
		return nil, fmt.Errorf("no extists!")
	}

	return storage, nil
}
