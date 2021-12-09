package tasks

// tools is a list of Go tools to install to avoid polluting global modules.
// Gotools module already sets up most of the basic go tools.
var toolList = []string{ //nolint:gochecknoglobals // ok to be global for tooling setup
	"github.com/goreleaser/goreleaser@v0.174.1",
	"github.com/golangci/golangci-lint/cmd/golangci-lint@master",
	"github.com/dustinkirkland/golang-petname/cmd/petname@master",
	"mvdan.cc/gofumpt@latest",
	"github.com/segmentio/golines@latest",
}
