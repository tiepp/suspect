package suspect

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func newDbConn(t *testing.T, conf Config) *pgconn.PgConn {
	ctx := context.Background()
	db, err := pgconn.Connect(ctx, "postgres://postgres:postgres@localhost:"+strconv.FormatUint(uint64(conf.Db.Port), 10)+"/postgres")
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, db.Close(ctx))
	})
	return db
}
