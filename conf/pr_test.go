package conf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ContextValues map[string]interface{}

func (cv ContextValues) IsSet(name string) bool {
	_, ok := cv[name]
	return ok
}

func (cv ContextValues) Int(name string) int {
	n, ok := cv[name].(int)
	if ok {
		return n
	}
	return 0
}

func (cv ContextValues) String(name string) string {
	s, ok := cv[name].(string)
	if ok {
		return s
	}
	return ""
}

func (cv ContextValues) Bool(name string) bool {
	b, ok := cv[name].(bool)
	if ok {
		return b
	}
	return false
}

func Test_Load_Pr(t *testing.T) {
	var vals ContextValues = map[string]interface{}{
		"owner":          "batman",
		"repo":           "justice",
		"title":          "title",
		"head":           "head",
		"base":           "base",
		"body":           "body",
		"ticket":         "ticket",
		"current-branch": "current-branch",

		"show-hint":    true,
		"show-json":    true,
		"show-summary": true,
		"verbose":      true,
		"interactive":  true,
	}
	pr := LoadPr(vals)
	assert.Equal(t, "batman", pr.Owner)
	assert.Equal(t, "justice", pr.Repo)
	assert.Equal(t, "title", pr.Title)
	assert.Equal(t, "head", pr.Head)
	assert.Equal(t, "base", pr.Base)
	assert.Equal(t, "body", pr.Body)
	assert.Equal(t, "ticket", pr.Ticket)
	assert.Equal(t, "current-branch", pr.CurrentBranch)

	assert.True(t, pr.ShowHint)
	assert.True(t, pr.ShowJson)
	assert.True(t, pr.ShowSummary)
	assert.True(t, pr.Verbose)
	assert.True(t, pr.Interactive)
}
