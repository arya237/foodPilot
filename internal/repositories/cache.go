package repositories

import (
	"errors"
	"time"
)

type Cache interface {
	Set(key string, code string, ttl time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
}

var (
	ErrNotFoundOTP = errors.New("برای کاربر کدی پیدا نشد")
	ErrDeleteOTP = errors.New("کد مورد نظر پاک نشد یا وجود ندارد")
	ErrInvalidExpire = errors.New("زمان انقضا کد اشتباه است")
)