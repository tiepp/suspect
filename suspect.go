package suspect

import (
	"github.com/gavv/httpexpect/v2"
	inbucket "github.com/inbucket/inbucket/pkg/rest/client"
	"github.com/jackc/pgconn"
	"testing"
)

type Suspect struct {
	Api  *httpexpect.Expect
	Db   *pgconn.PgConn
	Mail *inbucket.Client
	T    *testing.T
}

type Config struct {
	ApiKey     string
	ApiBaseUrl string
	DbUrl      string
	MailUrl    string
	Debug      bool
}

func NewSuspect(t *testing.T, conf Config) *Suspect {
	ensureTooling(t)

	api := newApiClient(t, conf)
	db := newDbConn(t, conf)
	mail := newMailClient(t, conf)

	return &Suspect{api, db, mail, t}
}
