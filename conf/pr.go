package conf

import "encoding/json"

// PR contains the parameters required for a PR POST.
//
// link: https://developer.github.com/v3/pulls/#create-a-pull-request
type PR struct {
	Owner         string
	Repo          string
	Title         string
	Head          string
	Base          string
	Body          string
	Ticket        string
	CurrentBranch string
	ShowHint      bool
	ShowJson      bool
	ShowSummary   bool
	Verbose       bool
	Interactive   bool
}

func LoadPr(ctx ValueContext) PR {
	c := ContextLoader{ctx}
	pr := PR{}
	c.String("owner", &pr.Owner)
	c.String("repo", &pr.Repo)
	c.String("title", &pr.Title)
	c.String("head", &pr.Head)
	c.String("base", &pr.Base)
	c.String("body", &pr.Body)
	c.String("ticket", &pr.Ticket)
	c.String("current-branch", &pr.CurrentBranch)
	c.Bool("show-hint", &pr.ShowHint)
	c.Bool("show-json", &pr.ShowJson)
	c.Bool("show-summary", &pr.ShowSummary)
	c.Bool("verbose", &pr.Verbose)
	c.Bool("interactive", &pr.Interactive)
	return pr
}

func (pr PR) ToJson() string {
	bin, err := json.MarshalIndent(pr, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(bin)
}
