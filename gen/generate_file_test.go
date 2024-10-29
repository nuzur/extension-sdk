package gen

import (
	"context"
	"errors"
	"os"
	"path"
	"testing"

	"github.com/nuzur/extension-sdk/client"
	"github.com/stretchr/testify/assert"
)

func TestGenerateFile(t *testing.T) {
	err := GenerateFile(context.Background(), FileRequest{
		ExecutionUUID: "123",
		OutputFile:    "/folder/test.txt",
		TemplateName:  "test",
		Data: struct {
			Name string
		}{Name: "nuzur"},
		Funcs:           nil,
		DisableGoFormat: true,
	})

	assert.NoError(t, err)
	rootDir := client.CurrentPath()
	finalOutput := path.Join(rootDir, "executions", "123", "folder", "test.txt")
	if _, err := os.Stat(finalOutput); errors.Is(err, os.ErrNotExist) {
		assert.NoError(t, err)
	}

	os.RemoveAll(path.Join(rootDir, "executions"))
}
