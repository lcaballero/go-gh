package conf

type PR struct {
	Dest   string `long:"dest-branch" short:"d" description:"Destination branch to merge changes into." required:"1"`
	Source string `long:"src-branch" short:"s" description:"Source branch of changes to merge from." required:"2"`
}

func (p PR) IsValid() bool {
	return p.Dest != "" && p.Source != ""
}
func (p PR) CmdName() string {
	return "pr"
}
