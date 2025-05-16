package safeos

import (
	"os"
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestSafeOS(t *testing.T) {
	root := &Root{Dir: "./tmp"}
	defer root.Delete()
	assert.NoError(t, root.CreateDir("./safeos"))
	publicFile := "./safeos/public.txt"
	assert.NoError(t, root.CreateFile(publicFile, []byte("test")))
	secretFile := "./secret.txt"
	f, _ := os.Create(secretFile)
	defer func(){
		assert.NoError(t, f.Close())
		assert.NoError(t, os.Remove(secretFile))
	}()
	assert.NotNil(t, f)
	_, err := f.Write([]byte("secret"))
	assert.NoError(t, err)
	b, err  := root.ReadFile("../secret.txt")
	assert.Empty(t, b)
	assert.Error(t, err)
	b, err = root.ReadFile(publicFile)
	assert.NoError(t, err)
	assert.Equal(t, string(b), "test")
	assert.NoError(t, root.DeleteFile(publicFile))
	fileInfo, err := root.Stat(publicFile)
	assert.True(t, os.IsNotExist(err))
	assert.Nil(t, fileInfo)
}
