//go:build tools

package tools

import (
	_ "github.com/air-verse/air"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/magefile/mage"
	_ "github.com/pressly/goose/v3/cmd/goose"
	_ "golang.org/x/vuln/cmd/govulncheck"
)
