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

func TestExample(t *testing.T) {
	testCreds := suspect.UserCredentials{
		Email:    "t@s.t",
		Password: "12345678",
	}
	testProfileId := 123
	testProfileName := "Test Name"

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
