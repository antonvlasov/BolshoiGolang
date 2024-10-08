package storage

import (
	"strconv"

	"go.uber.org/zap"
)

type Value struct {
	v string
	t Kind
}

type Storage struct {
	inner  map[string]Value
	logger *zap.Logger
}

func NewStorage() (Storage, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return Storage{}, err
	}

	defer logger.Sync()

	logger.Info("created new storage")

	return Storage{
		inner:  make(map[string]Value),
		logger: logger,
	}, nil
}

func (r Storage) Set(key, value string) {
	switch kind := getType(value); kind {
	case KindInt:
		r.inner[key] = Value{v: value, t: kind}
	case KindString:
		r.inner[key] = Value{v: value, t: kind}
	case KindUndefined:
		r.logger.Error(
			"undefined value type",
			zap.String("key", key),
			zap.Any("value", value),
		)
	}

	r.logger.Info("key set", zap.Any("value", value))
	r.logger.Sync()
}

func (r Storage) Get(key string) *string {
	res, ok := r.get(key)
	if !ok {
		return nil
	}

	return &res.v
}

func (r Storage) get(key string) (Value, bool) {
	res, ok := r.inner[key]
	if !ok {
		return Value{}, false
	}

	return res, true
}

type Kind string

const (
	KindInt       Kind = "D"
	KindString    Kind = "S"
	KindUndefined Kind = "UN"
)

func getType(value string) Kind {
	var val any

	val, err := strconv.Atoi(value)
	if err != nil {
		val = value
	}

	switch val.(type) {
	case int:
		return KindInt
	case string:
		return KindString
	default:
		return KindUndefined
	}
}

func sum[T int64 | uint64](x, y T) T {
	return x + y
}
