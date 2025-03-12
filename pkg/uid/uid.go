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
