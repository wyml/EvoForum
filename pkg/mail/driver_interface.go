package mail

type Driver interface {
	// Send Check Code
	Send(email Email, config map[string]string) bool
}
