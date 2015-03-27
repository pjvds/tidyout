package main

import "github.com/pjvds/tidy"

func main() {
	log := tidy.GetLogger()
	log.V(tidy.INFO).With("foo", "bar").Write("info message")

	log.Error("error")
	log.With("week", 8).Warn("warning entry")
	log.Info("info")
	log.Debug("debug")

	log.WithFields(tidy.Fields{
		"foo": "bar",
		"baz": 42,
	}).Info("hello world")

	log.Fatal("\\o/")
}
