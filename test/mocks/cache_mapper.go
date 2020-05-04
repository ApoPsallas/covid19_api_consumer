package mocks

import "time"

type MockCacheMapper struct {
	MapperStringResult string
	MapperIntResult    int64
	MapperError        error
}

func (m MockCacheMapper) Exists(key string) (int64, error) {
	return m.MapperIntResult, m.MapperError
}

func (m MockCacheMapper) Get(key string) (string, error) {
	return m.MapperStringResult, m.MapperError
}

func (m MockCacheMapper) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return m.MapperStringResult, m.MapperError
}
