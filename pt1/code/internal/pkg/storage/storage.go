package storage

import (
	"go.uber.org/zap"
)

type Value struct {
	s string
	d int
	a any
	b bool
}

type Storage struct {
	innerString map[string]Value
	logger      *zap.Logger
}

func NewStorage() (Storage, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return Storage{}, err
	}

	defer logger.Sync()
	logger.Info("created new storage")


	return Storage{
		inner:  make(map[string]string),
		logger: logger,
	}, nil
}

func (r Storage) Set(key, value string) {
	r.inner[key] = value

	r.logger.Info("key set", zap.Any())
	r.logger.Sync()
}

func (r Storage) Get(key string) *string {
	res, ok := r.inner[key]
	if !ok {
		return nil
	}

	return &res
}

func sum[T int64 | uint64](x, y T) T {
	return x + y
}
