package suspect

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/stretchr/testify/require"
	"testing"
)

func newDbConn(t *testing.T, conf Config) *pgconn.PgConn {
	ctx := context.Background()
	db, err := pgconn.Connect(ctx, "postgres://postgres:postgres@"+conf.DbUrl+"/postgres?sslmode=disable")
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, db.Close(ctx))
	})
	return db
}
