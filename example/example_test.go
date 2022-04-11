package example

import (
	"github.com/tiepp/suspect"
	"testing"
)

func TestExample(t *testing.T) {
	suspect.NewSuspect(t, suspect.Config{Debug: true}).
		SignUp(suspect.UserCredentials{
			Email:    "test@suspect.io",
			Password: "P@ssw0rd"}).
		SignOut()
}
