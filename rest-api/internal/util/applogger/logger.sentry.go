package applogger

type sentryLogger struct {
	name string
}

func (logger sentryLogger) I(message any, args ...any) {}
func (logger sentryLogger) E(message any, args ...any) {}
func (logger sentryLogger) D(message any, args ...any) {}

func newSentryLogger(name string) sentryLogger {
	return sentryLogger{
		name: name,
	}
}
