package suspect

import (
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func ensureTestEnvironment(t *testing.T) {
	assert.NoError(t, startDb(), "failed to start test environment")
	t.Cleanup(func() {
		assert.NoError(t, stopDb(), "failed to stop test environment")
	})
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
