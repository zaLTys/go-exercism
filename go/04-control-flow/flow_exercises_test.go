package learn

import "testing"

func TestFlowClassifyHTTPStatus(t *testing.T) {
	tests := []struct {
		name string
		code int
		want string
	}{
		{"100 informational", 100, "informational"},
		{"101 informational", 101, "informational"},
		{"200 success", 200, "success"},
		{"201 created", 201, "success"},
		{"301 redirect", 301, "redirect"},
		{"404 not found", 404, "client_error"},
		{"500 internal", 500, "server_error"},
		{"503 unavailable", 503, "server_error"},
		{"600 unknown", 600, "unknown"},
		{"0 unknown", 0, "unknown"},
		{"99 unknown", 99, "unknown"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := flowClassifyHTTPStatus(tc.code); got != tc.want {
				t.Fatalf("flowClassifyHTTPStatus(%d)=%q; want %q", tc.code, got, tc.want)
			}
		})
	}
}

func TestFlowBuildGradeReport(t *testing.T) {
	tests := []struct {
		name   string
		scores []int
		want   string
	}{
		{"mixed grades", []int{95, 82, 67}, "A,B,D"},
		{"nil slice", nil, ""},
		{"empty slice", []int{}, ""},
		{"boundary values", []int{100, 90, 89, 60, 59}, "A,A,B,D,F"},
		{"single score", []int{73}, "C"},
		{"all F", []int{10, 20, 0}, "F,F,F"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := flowBuildGradeReport(tc.scores); got != tc.want {
				t.Fatalf("got %q; want %q", got, tc.want)
			}
		})
	}
}

func TestFlowFirstMatch(t *testing.T) {
	items := []string{"banana", "avocado", "apple", "apricot"}

	t.Run("match found", func(t *testing.T) {
		v, ok := flowFirstMatch(items, "ap")
		if !ok || v != "apple" {
			t.Fatalf("got (%q,%v); want (\"apple\",true)", v, ok)
		}
	})
	t.Run("no match", func(t *testing.T) {
		v, ok := flowFirstMatch(items, "zz")
		if ok || v != "" {
			t.Fatalf("got (%q,%v); want (\"\",false)", v, ok)
		}
	})
	t.Run("nil slice", func(t *testing.T) {
		v, ok := flowFirstMatch(nil, "a")
		if ok || v != "" {
			t.Fatalf("got (%q,%v); want (\"\",false)", v, ok)
		}
	})
	t.Run("first item matches", func(t *testing.T) {
		v, ok := flowFirstMatch(items, "ban")
		if !ok || v != "banana" {
			t.Fatalf("got (%q,%v); want (\"banana\",true)", v, ok)
		}
	})
}
