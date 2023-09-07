package features

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/stretchr/testify/mock"
)

type PgMockPool struct {
	mock.Mock
}

func (m *PgMockPool) Close() {
}

func (m *PgMockPool) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	args := m.Called(ctx, sql)
	return args.Get(0).(pgconn.CommandTag), args.Error(1)
}
