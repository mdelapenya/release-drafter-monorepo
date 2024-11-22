package main

var lg Logger

func main() {
	WithLogger(NewDefaultLogger())

	c := &Core{}

	c.Info("Hello, world!")
}

type Core struct{}

func (c *Core) Info(msg string) {
	lg.Info(msg)
}

// SetLogger allows consumers to set their own logger
func WithLogger(customLogger Logger) {
	lg = customLogger
}

// Example function that uses the logger
func DoSomething() {
	lg.Info("Doing something important")
}
