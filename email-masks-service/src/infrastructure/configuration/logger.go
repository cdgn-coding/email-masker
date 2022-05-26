package configuration

type Logger interface {
	Debug(format string, v ...any)
	Log(format string, v ...any)
	Info(format string, v ...any)
	Warn(format string, v ...any)
}
