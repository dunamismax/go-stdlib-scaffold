//go:build tools

package tools

import (
	_ "github.com/magefile/mage"
	_ "golang.org/x/vuln/cmd/govulncheck"
)
