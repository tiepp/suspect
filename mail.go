package suspect

import (
	"github.com/inbucket/inbucket/pkg/rest/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

const mailBoxName = "test"

func newMailClient(t *testing.T, conf Config) *client.Client {
	mail, err := client.New(conf.MailUrl)
	assert.NoError(t, err)
	t.Cleanup(func() {
		t.Log("Purge mailbox")
		assert.NoError(t, mail.PurgeMailbox(mailBoxName))
	})
	return mail
}

func (s *Suspect) LastEmailText() (text string) {
	box, err := s.Mail.ListMailbox(mailBoxName)
	if err == nil && len(box) > 0 {
		message, _ := s.Mail.GetMessage(mailBoxName, box[0].ID)
		text = message.Body.Text
	}
	assert.NoError(s.T, err)
	assert.NotEmpty(s.T, text)

	return
}
