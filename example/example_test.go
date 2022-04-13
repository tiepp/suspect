package example

import (
	"github.com/tiepp/suspect"
	"testing"
)

func TestExample(t *testing.T) {
	suspect.NewSuspect(t).
		SignUp(suspect.UserCredentials{
			Email:    "test@suspect.io",
			Password: "P@ssw0rd"}).
		SignOut()
}
