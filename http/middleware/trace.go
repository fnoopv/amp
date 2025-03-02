package middleware

import (
	"github.com/fnoopv/amp/pkg/trace"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/util/errors"
)

// TraceMiddleare 全局Trace
type TraceMiddleare struct {
	goyave.Component
}

// Handle TraceID处理逻辑
func (tr *TraceMiddleare) Handle(next goyave.Handler) goyave.Handler {
	return func(response *goyave.Response, request *goyave.Request) {
		id, err := trace.Generate()
		if err != nil {
			panic(errors.New(err))
		}
		request.Header().Set(trace.MetaKey, id)
		response.Header().Set(trace.MetaKey, id)
		next(response, request)
	}
}
