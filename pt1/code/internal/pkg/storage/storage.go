package storage

import (
	"strconv"

	"go.uber.org/zap"
)

type Storage struct {
	innerString map[string]string
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
		innerString: make(map[string]string),
		logger:      logger,
	}, nil
}

func (r Storage) Set(key, value string) {
	r.innerString[key] = value

	r.logger.Info("key set", zap.Any("Storage", r.innerString))
	r.logger.Sync()
}

func (r Storage) Get(key string) *string {
	res, ok := r.innerString[key]
	if !ok {
		return nil
	}
	return &res
}

func (r Storage) GetKind(key string) string {
	value, ok := r.innerString[key]
	if !ok {
		return ""
	}
	_, err := strconv.Atoi(value)
	if err == nil {
		return "D"
	}
	return "S"
}
