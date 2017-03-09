package conf

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
