package suspect

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os/exec"
	"testing"
)

const testDbName = "suspect_test"

func ensureTooling(t *testing.T) {
	require.NoError(t, createTestDb())
	require.NoError(t, checkoutTestDb())
	t.Cleanup(func() {
		assert.NoError(t, checkoutMainDb())
		assert.NoError(t, deleteTestDb())
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
	// TODO: make pzth configurable
	cmd.Dir = "../"
	return cmd.Run()
}
