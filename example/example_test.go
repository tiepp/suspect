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
	testId := 123
	testName := "Test Name"

	suspect.NewSuspect(t).
		Api(suspect.AssertSignUp).
		Mail(func(mail *client.Client) {
			box, err := mail.ListMailbox("test")
			assert.NoError(t, err)
			assert.Empty(t, box)
		}).
		Api(suspect.AssertUser).
		Api(func(api *httpexpect.Expect) *httpexpect.Expect {
			api.POST("/rest/v1/profile").
				WithJSON(map[string]interface{}{"id": testId, "name": testName}).
				Expect().
				Status(http.StatusCreated)
			return api
		}).
		Wait(1).
		Db(func(db *pgx.Conn) {
			var name string
			err := db.QueryRow(context.Background(), "SELECT name FROM profile WHERE id = $1", testId).
				Scan(&name)
			assert.NoError(t, err)
			assert.EqualValues(t, testName, name)
		}).
		Api(func(api *httpexpect.Expect) *httpexpect.Expect {
			api.GET("/rest/v1/profile").
				Expect().Status(http.StatusOK).
				JSON().Array().First().Object().
				ValueEqual("id", testId).
				ValueEqual("name", testName)
			return api
		}).
		Api(suspect.AssertSignOut)
}
