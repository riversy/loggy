package connection

import "io"

type Connection struct {
	Status     Status
	StatusCh   chan Status
	Uri        string
	KeyPath    string
	RemotePath string
	Command    string
	SshOut     io.Reader
}

// Init initializes state of the object
func (c Connection) Init() error {
	c.StatusCh <- c.Status

	return nil
}

// GetName converts some internal data into the name
func (c Connection) GetName() string {
	return c.Uri
}

type DownloadConnection struct {
	*Connection
	TargetPath string
}

type ConnInterface interface {
	Init() error
	GetName() string
}

// NewDownloadConnection provides a new Connection instance
func NewDownloadConnection(
	uri string,
	keyPath string,
	remotePath string,
	command string,
	targetPath string,
) *DownloadConnection {
	state := Status(Pending)

	conn := &DownloadConnection{
		Connection: &Connection{
			Status:     state,
			StatusCh:   make(chan Status),
			Uri:        uri,
			KeyPath:    keyPath,
			RemotePath: remotePath,
			Command:    command,
		},
		TargetPath: targetPath,
	}

	err := conn.Init()
	if err != nil {
		panic(err)
	}

	return conn
}
