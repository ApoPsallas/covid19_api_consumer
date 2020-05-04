package repository

import "time"

type Cache interface {
	Set(k string, v interface{}, expiration time.Duration) (string, error)
	Get(k string) (string, error)
	Exists(k string) (int64, error)
}
