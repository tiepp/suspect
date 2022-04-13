package suspect

import (
	"github.com/BurntSushi/toml"
	"github.com/gavv/httpexpect/v2"
	inbucket "github.com/inbucket/inbucket/pkg/rest/client"
	"github.com/jackc/pgconn"
	"github.com/stretchr/testify/require"
	"testing"
)

type Suspect struct {
	Api  *httpexpect.Expect
	Db   *pgconn.PgConn
	Mail *inbucket.Client
	T    *testing.T
}

type Config struct {
	Api struct {
		Port uint
	}
	Db struct {
		Port uint
	}
	Inbucket struct {
		Port uint
	}
}

func NewSuspect(t *testing.T) *Suspect {
	var conf Config
	_, err := toml.DecodeFile("./supabase/config.toml", &conf)
	require.NoError(t, err)

	ensureTestEnvironment(t)

	api := newApiClient(t, conf)
	db := newDbConn(t, conf)
	mail := newMailClient(t, conf)

	return &Suspect{api, db, mail, t}
}
