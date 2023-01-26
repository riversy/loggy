package connection

type Status int

const (
	Pending Status = iota
	Connecting
	Connected
	Downloading
	Following
	Closed
	Failed
)
