package uid

import (
	"github.com/google/uuid"
	"goyave.dev/goyave/v5/util/errors"
)

// Generate 生成唯一ID
func Generate() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", errors.New(err)
	}

	return id.String(), nil
}

// Validate 验证给定字符串是否为有效UUID
func Validate(uid string) bool {
	return uuid.Validate(uid) == nil
}
