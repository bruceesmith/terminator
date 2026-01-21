module github.com/bruceesmith/terminator

go 1.25

tool (
	github.com/gojp/goreportcard/cmd/goreportcard-cli
	golang.org/x/vuln/cmd/govulncheck
	honnef.co/go/tools/cmd/staticcheck
)

require (
	github.com/BurntSushi/toml v1.6.0 // indirect
	github.com/gojp/goreportcard v0.0.0-20250418060254-1060522058eb // indirect
	golang.org/x/exp/typeparams v0.0.0-20260112195511-716be5621a96 // indirect
	golang.org/x/mod v0.32.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/sys v0.40.0 // indirect
	golang.org/x/telemetry v0.0.0-20260116145544-c6413dc483f5 // indirect
	golang.org/x/tools v0.41.0 // indirect
	golang.org/x/tools/go/expect v0.1.1-deprecated // indirect
	golang.org/x/tools/go/packages/packagestest v0.1.1-deprecated // indirect
	golang.org/x/vuln v1.1.4 // indirect
	honnef.co/go/tools v0.6.1 // indirect
)
