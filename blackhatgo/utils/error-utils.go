package utils

import (
	"fmt"
	"log"
)

// =============================================================================
// FATAL ERROR HANDLERS (Program termination)
// =============================================================================

// Must panics on error
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

// Mustf panics with formatted message on error
func Mustf[T any](value T, err error, format string, args ...any) T {
	if err != nil {
		panic(fmt.Sprintf(format+":%v", append(args, err)...))
	}
	return value
}

// fatal logs fatal error and exits - use for unrecoverable errors
func Fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// fatalf logs fatal error with context and exits
func FatalF(err error, format string, args ...any) {
	if err != nil {
		log.Fatalf(format+":%v", append(args, err)...)
	}
}

func FatalOn(condition bool, message string, args ...any) {
	if condition {
		log.Fatalf(message, args...)
	}
}

// =============================================================================
// ERROR LOGGING (Non-terminating)
// =============================================================================

// logError logs error and returns it - use in return statements
func LogError(err error) error {
	if err != nil {
		log.Println("ERROR:", err)
	}
	return err
}

func LogErrorF(err error, format string, args ...any) error {
	if err != nil {
		log.Printf("ERROR: "+format+":%v", append(args, err)...)
	}
	return err
}

func LogWarn(err error) error {
	if err != nil {
		log.Println("WARN:", err)
	}
	return err
}

func LogWarnF(err error, format string, args ...any) error {
	if err != nil {
		log.Printf("WARN:"+format+":%v", append(args, err)...)
	}
	return err
}

// =============================================================================
// CONDITIONAL EXECUTION HELPERS
// =============================================================================

// check logs and exits on error - use for critical operations
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// checkf logs with context and exits on error
func CheckF(err error, format string, args ...any) {
	if err != nil {
		log.Fatalf(format+": %v", append(args, err)...)
	}
}

// ignore silently ignores errors - use sparingly and document why
func Ignore(_ error) {
	// Explicitly ignore error
}

// ignoref logs ignored errors for debugging
func IgnoreF(err error, format string, args ...any) {
	if err != nil {
		log.Printf("IGNORED: "+format+": %v", append(args, err)...)
	}
}

// =============================================================================
// RETURN HELPERS
// =============================================================================

// orPanic returns value or panics - use for initialization
func OrPanic[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

// orZero returns value or zero value of type T
func OrZero[T any](value T, err error) T {
	if err != nil {
		var zero T
		return zero
	}
	return value
}

// orDefault returns value or provided default
func OrDefault[T any](value T, err error, defaultValue T) T {
	if err != nil {
		return defaultValue
	}
	return value
}

// =============================================================================
// SPECIALIZED HELPERS
// =============================================================================

// tryClose safely closes closers and logs errors
func TryClose(closer interface{ Close() error }, context string) {
	if closer != nil {
		if err := closer.Close(); err != nil {
			log.Printf("ERROR closing %s: %v", context, err)
		}
	}
}

// tryCloseQuiet closes without logging - returns error for caller to handle
func TryCloseQuiet(closer interface{ Close() error }) error {
	if closer != nil {
		return closer.Close()
	}
	return nil
}
