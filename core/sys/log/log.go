package log

import (
	"fmt"
	"os"
)

// Verbose sets whether the system should emit more verbose logs or not.
var Verbose bool

// Verbosef prepends the provided string format with a module identifier and then prints it to the console, but only if core.Verbose is true.
func Verbosef(module string, format string, a ...any) {
	if Verbose {
		fmt.Printf("[%v] %v", module, fmt.Sprintf(format, a...))
	}
}

// Printf prepends the provided string format with a module identifier and then prints it to the console.
func Printf(module string, format string, a ...any) {
	fmt.Printf("[%v] %v", module, fmt.Sprintf(format, a...))
}

// Fatalf prepends the provided string format with a module identifier, prints it to the console, and then calls core.ShutdownNow(1).
func Fatalf(module string, format string, a ...any) {
	fmt.Printf("[%v] %v", module, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// FatalfCode prepends the provided string format with a module identifier, prints it to the console, and then calls core.ShutdownNow(exitCode).
func FatalfCode(exitCode int, module string, format string, a ...any) {
	fmt.Printf("[%v] %v", module, fmt.Sprintf(format, a...))
	os.Exit(exitCode)
}
