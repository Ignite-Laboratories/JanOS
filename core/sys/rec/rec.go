package rec

import (
	"fmt"
	"os"
)

// Verbose sets whether the system should emit more verbose recordings or not.
var Verbose bool

// Verbosef prepends the provided string format with a name identifier and then prints it to the console, but only if Verbose is true.
func Verbosef(name string, format string, a ...any) {
	if Verbose {
		fmt.Printf("[%v] %v", name, fmt.Sprintf(format, a...))
	}
}

// Printf prepends the provided string format with a mnameodule identifier and then prints it to the console.
func Printf(name string, format string, a ...any) {
	fmt.Printf("[%v] %v", name, fmt.Sprintf(format, a...))
}

// Fatalf prepends the provided string format with a name identifier, prints it to the std.Err, and then calls os.Exit(1).
func Fatalf(name string, format string, a ...any) {
	fmt.Printf("[%v] %v", name, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// FatalfCode prepends the provided string format with a name identifier, prints it to the std.Err, and then calls os.Exit(exitCode).
func FatalfCode(exitCode int, name string, format string, a ...any) {
	_, _ = fmt.Fprintf(os.Stderr, "[%v] %v", name, fmt.Sprintf(format, a...))
	os.Exit(exitCode)
}
