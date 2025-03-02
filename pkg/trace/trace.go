package trace

import (
	"github.com/google/uuid"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/util/errors"
)

// MetaKey Trace Key
const MetaKey = "X-Trace-ID"

// Generate 生成TraceID
func Generate() (string, error) {
	uid, err := uuid.NewV7()
	if err != nil {
		return "", errors.New(err)
	}

	return uid.String(), nil
}

// Get 获取TraceID
func Get(request *goyave.Request) string {
	return request.Request().Header.Get(MetaKey)
}
