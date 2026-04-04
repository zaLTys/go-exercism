package learn

import (
	"strings"
	"testing"
)

func TestInterfacesNotifyAll(t *testing.T) {
	notifiers := []interfacesNotifier{
		interfacesEmailNotifier{Sender: "noreply@acme.com"},
		interfacesSMSNotifier{Provider: "Twilio"},
	}
	results := interfacesNotifyAll(notifiers, "alice", "Hello!")
	if len(results) != 2 {
		t.Fatalf("got %d results; want 2", len(results))
	}
	wantEmail := "Email from noreply@acme.com to alice: Hello!"
	if results[0] != wantEmail {
		t.Fatalf("email: got %q; want %q", results[0], wantEmail)
	}
	wantSMS := "SMS via Twilio to alice: Hello!"
	if results[1] != wantSMS {
		t.Fatalf("sms: got %q; want %q", results[1], wantSMS)
	}
}

func TestInterfacesNotifyAllEmpty(t *testing.T) {
	results := interfacesNotifyAll(nil, "bob", "test")
	if len(results) != 0 {
		t.Fatalf("nil notifiers: got %d results; want 0", len(results))
	}
}

func TestInterfacesDescribeAny(t *testing.T) {
	if got := interfacesDescribeAny(7); got != "int:7" {
		t.Fatalf("got %q; want %q", got, "int:7")
	}
	if got := interfacesDescribeAny("hi"); got != "string:hi" {
		t.Fatalf("got %q; want %q", got, "string:hi")
	}
	if got := interfacesDescribeAny(true); got != "unknown" {
		t.Fatalf("got %q; want %q", got, "unknown")
	}
}

func TestInterfacesReadAllUpper(t *testing.T) {
	got, err := interfacesReadAllUpper(strings.NewReader("go rocks"))
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if got != "GO ROCKS" {
		t.Fatalf("got %q; want %q", got, "GO ROCKS")
	}
}
