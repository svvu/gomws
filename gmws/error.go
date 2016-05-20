package gmws

// Error represents the error message from the API.
type Error struct {
	// Error type. Values: Sender, Server.
	Type string `json:"Type"`
	// Amazon error code.
	Code string `json:"Code"`
	// Text explain the error.
	Message string `json:"Message"`
	// Detail about the error.
	Detail string `json:"Detail"`
}
