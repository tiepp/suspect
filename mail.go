package suspect

import (
	"github.com/inbucket/inbucket/pkg/rest/client"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

const mailBoxName = "test"

func newMailClient(t *testing.T, conf config) *client.Client {
	mail, err := client.New("http://localhost:" + strconv.FormatUint(uint64(conf.Inbucket.Port), 10))
	assert.NoError(t, err)
	t.Cleanup(func() {
		t.Log("Purge mailbox")
		assert.NoError(t, mail.PurgeMailbox(mailBoxName))
	})
	return mail
}

func (s *Suspect) Mail(a func(*client.Client)) *Suspect {
	a(s.mail)
	return s
}
