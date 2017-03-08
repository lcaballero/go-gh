package conf

import (
	. "github.com/lcaballero/exam/assert"
	"testing"
)

func Test_Fork_003(t *testing.T) {
	t.Log("should return error if fork is invliad")
	f := Fork{}
	res, err := f.CreateFork(&Config{})
	IsNil(t, res)
	IsNotNil(t, err)
}

func Test_Fork_002(t *testing.T) {
	t.Log("new fork should be invalid")
	f := Fork{}
	IsFalse(t, f.IsValid())
}

func Test_Fork_001(t *testing.T) {
	t.Log("new fork should have name 'fork'")
	f := Fork{}
	IsEqStrings(t, f.CmdName(), "fork")
}
