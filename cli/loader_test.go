package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Load_Token(t *testing.T) {
	token, err := NewLoader().LoadToken("./.files/.go-gh-token")

	assert.Nil(t, err)
	assert.Equal(t, "1209801297017091j091j0fhslasnflajflskj90", token)
}

func Test_Update_From_Ini(t *testing.T) {
	locals, err := NewLoader().UpdateFromIni("./.files/.go-gh")

	assert.Nil(t, err)
	assert.Equal(t, "Work", locals.Active)
	assert.Equal(t, "1221212312312314121211231423425312341242", locals.Current.Token)
	assert.Equal(t, "https://github.schq.secious.com/api/v3/", locals.Current.BaseUrl)
	assert.Equal(t, "Lucas-Caballero", locals.Current.Username)
}
