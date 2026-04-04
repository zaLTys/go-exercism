package learn

import "testing"

func TestModulesJoinImportPath(t *testing.T) {
	if got := modulesJoinImportPath("example.com/m", ""); got != "example.com/m" {
		t.Fatalf("got %q; want %q", got, "example.com/m")
	}
	if got := modulesJoinImportPath("example.com/m", "auth/token"); got != "example.com/m/auth/token" {
		t.Fatalf("got %q; want %q", got, "example.com/m/auth/token")
	}
}

func TestModulesWithMajorSuffix(t *testing.T) {
	if got := modulesWithMajorSuffix("example.com/m", 1); got != "example.com/m" {
		t.Fatalf("got %q; want %q", got, "example.com/m")
	}
	if got := modulesWithMajorSuffix("example.com/m", 2); got != "example.com/m/v2" {
		t.Fatalf("got %q; want %q", got, "example.com/m/v2")
	}
	if got := modulesWithMajorSuffix("example.com/m", 3); got != "example.com/m/v3" {
		t.Fatalf("got %q; want %q", got, "example.com/m/v3")
	}
}

func TestModulesParseRequireLine(t *testing.T) {
	mod, ver, ok := modulesParseRequireLine("require example.com/dep v1.2.3")
	if !ok || mod != "example.com/dep" || ver != "v1.2.3" {
		t.Fatalf("got (%q,%q,%v)", mod, ver, ok)
	}
	_, _, ok = modulesParseRequireLine("nonsense")
	if ok {
		t.Fatalf("expected ok=false for invalid line")
	}
}
