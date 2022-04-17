package suspect

import (
	"github.com/BurntSushi/toml"
	"github.com/gavv/httpexpect/v2"
	inbucket "github.com/inbucket/inbucket/pkg/rest/client"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type Suspect struct {
	api  *httpexpect.Expect
	db   *pgx.Conn
	mail *inbucket.Client
	t    *testing.T
}

type config struct {
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
	var conf config
	_, err := toml.DecodeFile("./supabase/config.toml", &conf)
	require.NoError(t, err)

	ensureTestEnvironment(t)

	api := newApiClient(t, conf)
	db := newDbConn(t, conf)
	mail := newMailClient(t, conf)

	return &Suspect{api, db, mail, t}
}

func (s *Suspect) Wait(dur int) *Suspect {
	s.t.Helper()
	s.t.Logf("Waiting for %d seconds", dur)
	time.Sleep(time.Duration(dur) * time.Second)
	return s
}
