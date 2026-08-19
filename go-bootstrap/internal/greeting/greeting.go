// Package greeting contains application-domain logic.
package greeting

// Message returns a greeting for name.
func Message(name string) string {
	return "hello, " + name
}
