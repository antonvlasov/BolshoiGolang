package storage

import (
	"go.uber.org/zap"
	"strconv"
)

type Storage struct {
	innerString map[string]string
	logger      *zap.Logger
}

func InitStorage() (Storage, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return Storage{}, err
	}
	defer logger.Sync()
	logger.Info("storage initialized")
	return Storage{
		innerString: make(map[string]string),
		logger:      logger,
	}, nil
}

func (storage Storage) Set(key, value string) {
	defer storage.logger.Sync()
	storage.innerString[key] = value
	storage.logger.Info("Set key value")
}

func (storage Storage) Get(key string) *string {
	defer storage.logger.Sync()
	out, ok := storage.innerString[key]
	storage.logger.Info("Get key value")
	if ok {
		return &out
	}
	return nil
}

func (storage Storage) GetKind(key string) string {
	defer storage.logger.Sync()
	out, ok := storage.innerString[key]
	storage.logger.Info("Get kind key value")
	if !ok {
		return "N"
	}
	if _, err := strconv.ParseFloat(out, 64); err == nil {
		return "D"
	}
	if _, err := strconv.Atoi(out); err == nil {
		return "D"
	}
	return "S"
}
