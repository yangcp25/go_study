package template

import "testing"
import "github.com/stretchr/testify/assert"

func Test_smg_Valid(t *testing.T) {
	tel := NewTelecomSms()
	err := tel.Send("test", "1239999")
	assert.NoError(t, err)
}
