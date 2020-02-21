package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCleanup(t *testing.T) {
	tDir, err := ioutil.TempDir("", "golang-leipzig-*")
	require.NoError(t, err)
	t.Cleanup(func() {
		t.Log("removing ", tDir)
		require.NoError(t, os.RemoveAll(tDir))
	})

	require.True(t, true)
}
