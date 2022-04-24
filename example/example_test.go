package example

import (
	"context"
	"github.com/gavv/httpexpect/v2"
	"github.com/inbucket/inbucket/pkg/rest/client"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"github.com/tiepp/suspect"
	"net/http"
	"testing"
)

var (
	testCreds = suspect.UserCredentials{
		Email:    "t@s.t",
		Password: "12345678",
	}
	testCreds2 = suspect.UserCredentials{
		Email:    "t2@s.t",
		Password: "12345678",
	}
	testProfileId   = 123
	testProfileName = "Test Name"
)

func TestSignup(t *testing.T) {
	suspect.NewSuspect(t).
		Api(suspect.AssertSignUp(testCreds)).
		Mail(func(mail *client.Client) {
			box, err := mail.ListMailbox("test")
			assert.NoError(t, err)
			assert.Empty(t, box)
		}).
		Api(suspect.AssertUser).
		Api(func(api *httpexpect.Expect) *httpexpect.Expect {
			api.POST("/rest/v1/profile").
				WithJSON(map[string]interface{}{"id": testProfileId, "name": testProfileName}).
				Expect().
				Status(http.StatusCreated)
			return api
		}).
		Wait(1).
		Db(func(db *pgx.Conn) {
			var name string
			err := db.QueryRow(context.Background(), "SELECT name FROM profile WHERE id = $1", testProfileId).
				Scan(&name)
			assert.NoError(t, err)
			assert.EqualValues(t, testProfileName, name)
		}).
		Api(func(api *httpexpect.Expect) *httpexpect.Expect {
			api.GET("/rest/v1/profile").
				Expect().Status(http.StatusOK).
				JSON().Array().First().Object().
				ValueEqual("id", testProfileId).
				ValueEqual("name", testProfileName)
			return api
		}).
		Api(suspect.AssertSignOut)
}

func TestSignIn(t *testing.T) {
	suspect.NewSuspect(t).
		Db(func(db *pgx.Conn) {
			_, err := db.Exec(context.Background(),
				"INSERT INTO auth.users (instance_id, id, aud, role, email, encrypted_password, email_confirmed_at, invited_at, confirmation_token, confirmation_sent_at, recovery_token, recovery_sent_at, email_change_token_new, email_change, email_change_sent_at, last_sign_in_at, raw_app_meta_data, raw_user_meta_data, is_super_admin, created_at, updated_at, phone, phone_confirmed_at, phone_change, phone_change_token, phone_change_sent_at, email_change_token_current, email_change_confirm_status, banned_until, reauthentication_token, reauthentication_sent_at) VALUES ('00000000-0000-0000-0000-000000000000', '692818a0-c0ff-4bfe-88b7-ac35d943bb46', 'authenticated', 'authenticated', 't@s.t', '$2a$10$thBmVA21xFXsHmJUODmhcOvaUNBTJVsAO5JzpDAj76wsUIT5zmNiG', '2022-04-23 17:17:56.411797 +00:00', null, '', null, '', null, '', '', null, '2022-04-23 17:17:56.414820 +00:00', '{\"provider\": \"email\", \"providers\": [\"email\"]}', '{}', false, '2022-04-23 17:17:56.405233 +00:00', '2022-04-23 17:17:56.405237 +00:00', null, null, '', '', null, '', 0, null, '', null);")
			assert.NoError(t, err)
		}).
		Api(suspect.AssertSignIn(testCreds)).
		Api(suspect.AssertUser).
		Api(suspect.AssertSignOut)
}

func TestSlice(t *testing.T) {
	s1 := suspect.NewSuspect(t)
	s2 := s1.Slice()
	s1.Api(suspect.AssertSignUp(testCreds))
	s2.Api(suspect.AssertSignUp(testCreds2))
	s1.Api(suspect.AssertUser)
	s2.Api(suspect.AssertUser)
}
