package music

import (
	"github.com/justinas/alice"
)

var (
	// ErrorChain middleware
	ErrorChain alice.Chain
)

// Start func to start server
func Start() {
	StartServer()
}
