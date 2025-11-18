package internal

type TraceValue string

const (
	TraceOff      TraceValue = "off"
	TraceMessages TraceValue = "messages"
	TraceVerbose  TraceValue = "verbose"
)

type WorspaceFolder struct {
	// The associated URI for this workspace folder.
	URI string `json:"uri"`
	// The name of the workspace folder. Used to refer to this workspace in the
	// user interface.
	Name string `json:"name"`
}

type ClientServerInfo struct {
	// The name of the client/server as defined by the client/server.
	Name string `json:"name"`
	// The client/server's version as defined by the client/server.
	Version string `json:"version,omitempty"`
}
