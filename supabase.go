package suspect

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os/exec"
	"testing"
)

const testDbName = "suspect_test"

func ensureTestEnvironment(t *testing.T) {
	require.NoError(t, createTestDb(), "failed to create test database")
	require.NoError(t, checkoutTestDb(), "failed to checkout test database")
	t.Cleanup(func() {
		assert.NoError(t, checkoutMainDb(), "failed to checkout main database")
		assert.NoError(t, deleteTestDb(), "failed to delete test database")
	})
}
func createTestDb() error {
	return runCmd("db", "branch", "create", testDbName)
}

func checkoutTestDb() error {
	return runCmd("db", "switch", testDbName)
}

func checkoutMainDb() error {
	return runCmd("db", "switch", "main")
}

func deleteTestDb() error {
	return runCmd("db", "branch", "delete", testDbName)
}

func runCmd(arg ...string) error {
	cmd := exec.Command("supabase", arg...)
	return cmd.Run()
}
