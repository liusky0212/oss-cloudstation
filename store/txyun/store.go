package txyun

import (
	txyunStore "oss-station/store"
)

type TxyunStore struct {
}

var (
	//对象是否实现接口的约束
	_ txyunStore.Uploader = &TxyunStore{}
)

func NewTxyunStore() (*TxyunStore, error) {
	return &TxyunStore{}, nil
}

func (s *TxyunStore) Upload(bucketName string, objectKey string, fileName string) error {
	return nil
}
