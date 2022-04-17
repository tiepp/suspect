package suspect

import (
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"strconv"
	"testing"
)

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var cred = UserCredentials{
	Email:    "t@s.t",
	Password: "12345678",
}

// Taking advantage of fixed key. This is likely to break in the future.
// https://github.com/supabase/cli/blob/7fa402cd5a95d6a83e32f82113de449656a080e2/internal/start/start.go#L77
const apiKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZS1kZW1vIiwicm9sZSI6ImFub24ifQ.625_WdcF3KHqz5amU0x2X5WWHP-OEs_4qj0ssLNHzTs"

func newApiClient(t *testing.T, conf config) *httpexpect.Expect {
	printers := []httpexpect.Printer{httpexpect.NewCompactPrinter(t)}
	printers = append(printers, httpexpect.NewDebugPrinter(t, true))

	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:" + strconv.FormatUint(uint64(conf.Api.Port), 10),
		Client:   &http.Client{Jar: httpexpect.NewJar()},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: printers,
	})
}

func (s *Suspect) Api(a func(*httpexpect.Expect) *httpexpect.Expect) *Suspect {
	s.api = a(s.api)
	return s
}

func AssertSignUp(api *httpexpect.Expect) *httpexpect.Expect {
	r := api.POST("/auth/v1/signup").
		WithQuery("apikey", apiKey).
		WithJSON(cred).
		Expect().
		Status(http.StatusOK)
	accessToken := r.JSON().Object().Value("access_token").String().Raw()
	api = api.Builder(func(r *httpexpect.Request) {
		r.WithQuery("apikey", apiKey)
		r.WithHeader("Authorization", "Bearer "+accessToken)
	})
	return api
}

func AssertSignOut(api *httpexpect.Expect) *httpexpect.Expect {
	api.POST("/auth/v1/logout").
		Expect().
		Status(http.StatusNoContent)
	return api
}

func AssertUser(api *httpexpect.Expect) *httpexpect.Expect {
	api.GET("/auth/v1/user").
		Expect().
		Status(http.StatusOK)
	return api
}
