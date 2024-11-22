package core

var Log Logger

// SetLogger allows consumers to set their own logger
func WithLogger(customLogger Logger) {
	Log = customLogger
}

// Example function that uses the logger
func DoSomething() {
	Log.Info("Doing something important")
}
