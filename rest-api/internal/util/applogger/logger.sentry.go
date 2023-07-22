package applogger

type sentryLogger struct {
	name string
}

func (logger sentryLogger) I(message any, args ...string) {}
func (logger sentryLogger) E(message any, args ...string) {}
func (logger sentryLogger) D(message any, args ...string) {}

func newSentryLogger(name string) sentryLogger {
	return sentryLogger{
		name: name,
	}
}
