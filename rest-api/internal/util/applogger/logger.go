package applogger

type Logger interface {
	i(string, ...string)
	e(string, ...string)
	d(string, ...string)
}
