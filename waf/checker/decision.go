package checker

const (
	ActionAllow   = "allow"
	ActionWatch   = "watch"
	ActionBlock   = "block"
	ActionDefault = ActionAllow
)

// Scroe is range from 0 to 100.
var CreditScore uint

// Block Message.
const (
	BlockByIP      = "block by IP"
	BlockByIPIDC   = "block by IP IDC"
	BlockByUA      = "block by UA"
	BlockByURI     = "block by URI"
	ActionWatchMsg = "watch"
)
