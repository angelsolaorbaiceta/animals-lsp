package internal

import "encoding/json"

type InitializeParams struct {
	WorkDoneToken string `json:"workDoneToken,omitempty"`
	// The process Id of the parent process that started the server.
	// Is null if the process has not been started by another process.
	// If the parent process is not alive, then the server should exit.
	ProcessID  *int              `json:"processId"`
	ClientInfo *ClientServerInfo `json:"clientInfo,omitempty"`
	Locale     string            `json:"locale,omitempty"`
	// The root URI of the workspace. Is null if no folder is open.
	RootURI               *string            `json:"rootUri"`
	InitializationOptions json.RawMessage    `josn:"initializationOptions,omitempty"`
	Capabilities          any                `json:"capabilities"`
	TraceValue            TraceValue         `json:"traceValue,omitempty"`
	WorkspaceFolders      *[]WorkspaceFolder `json:"workspaceFolders,omitempty"`
}

func makeInitializeResult() *InitializeResult {
	return &InitializeResult{
		Capabilities: ServerCapabilities{
			CompletionProvider: &CompletionOptions{
				// TODO: support work done progress?
				WorkDoneProgress: false,
				// TODO: learn what this means
				TriggerCharacters: []rune{},
				// TODO: learn what this means
				AllCommitCharacters: []rune{},
				// TODO: learn what this means
				ResolveProvider: false,
			},
			HoverProvider:      true,
			ReferencesProvider: true,
		},
		ServerInfo: &ClientServerInfo{
			Name:    "Animals Language Server",
			Version: "v0.1",
		},
	}
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   *ClientServerInfo  `json:"serverInfo,omitempty"`
}

type ServerCapabilities struct {
	CompletionProvider *CompletionOptions `json:"completionProvider,omitempty"`
	HoverProvider      bool               `json:"hoverProvider"`
	ReferencesProvider bool               `json:"referencesProvider"`
}

type CompletionOptions struct {
	WorkDoneProgress bool `json:"workDoneProgress"`

	// The additional characters, beyond the defaults provided by the client (typically
	// [a-zA-Z]), that should automatically trigger a completion request. For example
	// `.` in JavaScript represents the beginning of an object property or method and is
	// thus a good candidate for triggering a completion request.
	//
	// Most tools trigger a completion request automatically without explicitly
	// requesting it using a keyboard shortcut (e.g. Ctrl+Space). Typically they
	// do so when the user starts to type an identifier. For example if the user
	// types `c` in a JavaScript file code complete will automatically pop up
	// present `console` besides others as a completion item. Characters that
	// make up identifiers don't need to be listed here.
	TriggerCharacters []rune `json:"triggerCharacters"`

	// The list of all possible characters that commit a completion. This field
	// can be used if clients don't support individual commit characters per
	// completion item. See client capability
	// `completion.completionItem.commitCharactersSupport`.
	//
	// If a server provides both `allCommitCharacters` and commit characters on
	// an individual completion item the ones on the completion item win.
	AllCommitCharacters []rune `json:"allCommitCharacters"`

	// The server provides support to resolve additional information for a completion item.
	ResolveProvider bool `json:"resolveProvider"`
}
