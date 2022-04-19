package suspect

import (
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

const testDbName = "suspecttest"

func ensureTestEnvironment(t *testing.T) {
	assert.NoError(t, startDb(), "failed to start test environment")
	t.Cleanup(func() {
		assert.NoError(t, stopDb(), "failed to stop test environment")
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

func resetTestDb() error {
	return runCmd("db", "reset")
}

func stopDb() error {
	return runCmd("stop")
}

func startDb() error {
	return runCmd("start")
}

func runCmd(arg ...string) error {
	cmd := exec.Command("supabase", arg...)
	return cmd.Run()
}
