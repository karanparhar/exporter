package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfigFromFile(t *testing.T) {
	content := []byte(`
		{	
			"urls" : ["someurl"]
	}`)

	tmpFileName := createTemporaryFile(t, content)
	defer os.Remove(tmpFileName)

	config, err := ReadConfigFile(tmpFileName)
	require.NoError(t, err, "Expected no error")

	assert.Equal(t, "someurl", config.Urls[0])

}

func TestNewConfigFromFileWhenFileMissing(t *testing.T) {
	_, err := ReadConfigFile("/some/path/which/does/not/exist")
	require.Error(t, err, "Expected error for missing file")

}

func createTemporaryFile(t *testing.T, content []byte) string {
	name := "config.json"
	tmpFile, err := ioutil.TempFile(".", name)
	require.NoError(t, err, "Expected no error")

	_, err = tmpFile.Write(content)
	require.NoError(t, err, "Expected no error writing to temporary file")

	err = tmpFile.Close()
	require.NoError(t, err, "Expected no error closing file")

	return tmpFile.Name()
}
