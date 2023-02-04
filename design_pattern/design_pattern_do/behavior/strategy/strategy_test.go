package strategy

import "testing"

import "github.com/stretchr/testify/assert"

func Test_newStorage(t *testing.T) {
	data, sensitive := getData()

	sensitiveType := "file"

	if sensitive {
		sensitiveType = "encrypt_file"
	}

	encrypt, err := newStorage(sensitiveType)

	assert.NoError(t, err)

	assert.NoError(t, encrypt.Save("test.text", data))

}

func getData() (data []byte, sensitive bool) {
	return []byte("sss"), false
}
