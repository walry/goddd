package appctx

import (
	"context"
	"youras/pkg/uuid"
)

type _ctxKey struct{}

type UserCtx struct {
	traceId  string
	userId   string
	clientIp string
}

func (m *UserCtx) TraceId() string {
	return m.traceId
}

//func(m *MyContext)SetTraceId(traceId string) {
//	m.traceId = traceId
//}

func FromContext(ctx context.Context) *UserCtx {
	if v, ok := ctx.Value(_ctxKey{}).(*UserCtx); ok {
		return v
	}
	return &UserCtx{traceId: uuid.ShortUuid()}
}

func WithContext(ctx context.Context, value *UserCtx) context.Context {
	return context.WithValue(ctx, _ctxKey{}, value)
}

func InjectCtx(ctx context.Context) context.Context {
	return WithContext(ctx, FromContext(ctx))
}
