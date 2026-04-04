package learn

import "io"

// interfacesNotifier is an interface for sending notifications.
// Both types below must satisfy it implicitly — no "implements" keyword needed.
type interfacesNotifier interface {
	Notify(to, message string) string
}

type interfacesEmailNotifier struct{ Sender string }
type interfacesSMSNotifier struct{ Provider string }

// EXERCISE 1: Implement Notify() on both types so they satisfy interfacesNotifier.
// Then implement interfacesNotifyAll to send via all notifiers and collect results.
//
//   interfacesEmailNotifier{Sender: "noreply@acme.com"}.Notify("alice", "hi")
//     => "Email from noreply@acme.com to alice: hi"
//
//   interfacesSMSNotifier{Provider: "Twilio"}.Notify("+1234", "hi")
//     => "SMS via Twilio to +1234: hi"
//
// Real-world context: a notification service that fans out to multiple channels
// (email, SMS, push, Slack). Each channel implements the same interface, and
// the orchestrator doesn't care which concrete type it calls.
// Hint: you'll need "fmt" for fmt.Sprintf.

func (e interfacesEmailNotifier) Notify(to, message string) string {
	// TODO: implement
	return ""
}

func (s interfacesSMSNotifier) Notify(to, message string) string {
	// TODO: implement
	return ""
}

// interfacesNotifyAll sends a message via all notifiers and returns all results.
func interfacesNotifyAll(notifiers []interfacesNotifier, to, message string) []string {
	// TODO: implement
	return nil
}

// EXERCISE 2: Implement interfacesDescribeAny using a type switch.
// Rules:
//   - int    => "int:<n>"
//   - string => "string:<value>"
//   - default => "unknown"
//
// Real-world context: type switches are used in JSON/config parsers,
// event handlers, and middleware that dispatch on dynamic types.
// Hint: you'll need "fmt" for fmt.Sprintf.
func interfacesDescribeAny(v any) string {
	// TODO: implement type switch
	return ""
}

// EXERCISE 3: Implement interfacesReadAllUpper.
// Read everything from r, return the uppercase string.
//
// Real-world context: processing streams from any source (file, HTTP body,
// stdin) uniformly via io.Reader — the most important interface in Go.
// Hint: io.ReadAll reads everything; strings.ToUpper converts to uppercase.
func interfacesReadAllUpper(r io.Reader) (string, error) {
	// TODO: implement (hint: io.ReadAll + strings.ToUpper)
	return "", nil
}
