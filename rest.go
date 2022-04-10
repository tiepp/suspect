package suspect

import (
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"testing"
)

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const authUri = "/auth/v1"
const restUri = "/rest/v1"

// Taking advantage of fixed key. This is likely to break in the future.
// https://github.com/supabase/cli/blob/7fa402cd5a95d6a83e32f82113de449656a080e2/internal/start/start.go#L77
const apiKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZS1kZW1vIiwicm9sZSI6ImFub24ifQ.625_WdcF3KHqz5amU0x2X5WWHP-OEs_4qj0ssLNHzTs"

func newApiClient(t *testing.T, conf Config) *httpexpect.Expect {
	printers := []httpexpect.Printer{httpexpect.NewCompactPrinter(t)}
	if conf.Debug {
		printers = append(printers, httpexpect.NewDebugPrinter(t, true))
	}

	api := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  conf.ApiBaseUrl,
		Client:   &http.Client{Jar: httpexpect.NewJar()},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: printers,
	})
	return api
}

func (s *Suspect) SignUp(cred UserCredentials) *Suspect {
	r := s.Api.POST(authUri+"/signup").
		WithQuery("apikey", apiKey).
		WithJSON(cred).
		Expect().
		Status(http.StatusOK)
	accessToken := r.JSON().Object().Value("access_token").String().Raw()
	s.Api = s.Api.Builder(func(r *httpexpect.Request) {
		r.WithQuery("apikey", apiKey)
		r.WithHeader("Authorization", "Bearer "+accessToken)
	})
	return s
}

func (s *Suspect) SignOut() *Suspect {
	s.Api.POST(authUri + "/logout").
		Expect().
		Status(http.StatusNoContent)
	return s
}
