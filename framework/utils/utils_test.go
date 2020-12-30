package utils_test

import (
	"testing"
	"github.com/rafaelcmd/encoder-service/framework/utils"
	"github.com/stretchr/testify/require"
)

func TestIsJson(t *testing.T) {
	json := `{
				"id": "525b5fd9-700d-4feb-89c0-415a1e6e148c",
				"file_path": "convite.mp4",
				"status": "pending"
			  }`
			  
	err := utils.IsJson(json)
	require.Nil(t, err)

	json = `raf`
	err = utils.IsJson(json)
	require.Error(t, err)
}