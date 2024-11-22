package core

var lg Logger

// SetLogger allows consumers to set their own logger
func WithLogger(customLogger Logger) {
	lg = customLogger
}

// Example function that uses the logger
func DoSomething() {
	lg.Info("Doing something important")
}
