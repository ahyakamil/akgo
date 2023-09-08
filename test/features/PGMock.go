package features

import (
	"akgo/db"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/mock"
)

type PgMockPool struct {
	mock.Mock
}

func (m *PgMockPool) Query(ctx context.Context, sql string, arguments ...interface{}) (pgx.Rows, error) {
	result, err := m.Query(ctx, sql)
	return result, err
}

func (m *PgMockPool) Begin(ctx context.Context) (pgx.Tx, error) {
	return new(MockTx), nil
}

func (m *PgMockPool) Close() {
}

func (m *PgMockPool) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	result := m.Called(ctx, sql)
	return result.Get(0).(pgconn.CommandTag), result.Error(1)
}

type MockTx struct {
}

func (m MockTx) Begin(ctx context.Context) (pgx.Tx, error) {
	return &MockTx{}, nil
}

func (m MockTx) BeginFunc(ctx context.Context, f func(pgx.Tx) error) (err error) {
	return nil
}

func (m MockTx) Commit(ctx context.Context) error {
	return nil
}

func (m MockTx) Rollback(ctx context.Context) error {
	return nil
}

func (m MockTx) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	return 0, nil
}

func (m MockTx) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return nil
}

func (m MockTx) LargeObjects() pgx.LargeObjects {
	return m.LargeObjects()
}

func (m MockTx) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
	return nil, nil
}

func (m MockTx) Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error) {
	return db.Pg.Exec(ctx, sql, arguments)
}

func (m MockTx) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return nil, nil
}

func (m MockTx) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return nil
}

func (m MockTx) QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error) {
	return nil, nil
}

func (m MockTx) Conn() *pgx.Conn {
	return nil
}
