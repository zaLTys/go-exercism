package learn

import (
	"strings"
	"testing"
)

func TestSetupHello(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"normal", "Paul", "Hello, Paul!"},
		{"empty", "", "Hello, world!"},
		{"spaces", " ", "Hello, world!"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := setupHello(tc.in)
			if got != tc.want {
				t.Fatalf("setupHello(%q) = %q; want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestSetupParseGoVersion(t *testing.T) {
	tests := []struct {
		in         string
		wantMajor  int
		wantMinor  int
		wantParsed bool
	}{
		{"go version go1.22.3 darwin/arm64", 1, 22, true},
		{"go1.21.0", 1, 21, true},
		{"go version deez nuts", 0, 0, false},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			maj, min, ok := setupParseGoVersion(tc.in)
			if ok != tc.wantParsed {
				t.Fatalf("setupParseGoVersion(%q) ok=%v; want %v", tc.in, ok, tc.wantParsed)
			}
			if ok && (maj != tc.wantMajor || min != tc.wantMinor) {
				t.Fatalf("setupParseGoVersion(%q)=(%d,%d); want (%d,%d)",
					tc.in, maj, min, tc.wantMajor, tc.wantMinor)
			}
		})
	}
}

func TestSetupModulePathFromGoMod(t *testing.T) {
	goMod := strings.TrimSpace(`
module example.com/mymodule

go 1.22

require (
	example.com/dep v1.2.3
)
`)
	if got := setupModulePathFromGoMod(goMod); got != "example.com/mymodule" {
		t.Fatalf("setupModulePathFromGoMod(...)=%q; want %q", got, "example.com/mymodule")
	}
}
