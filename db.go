package suspect

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func newDbConn(t *testing.T, conf config) *pgx.Conn {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:"+strconv.FormatUint(uint64(conf.Db.Port), 10)+"/postgres?sslmode=disable")
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, conn.Close(ctx))
	})
	return conn
}

func (s *Suspect) Db(a func(conn *pgx.Conn)) *Suspect {
	a(s.db)
	return s
}
