<!-- README.md -->

# Recoverich Module

The `recoverich` library provides functions to recover from panics, log errors, and track values during a panic situation. This README provides an explanation of each public method along with examples of how to use them.

## `Recover()`

This is the most basic version of the recover function. It recovers from a panic, logs the error, and prints a stack trace.

### Example:

```go
package main

import (
	"github.com/ildomm/recoverich"
)

func main() {
    defer recoverich.Recover()

    // Your code that may panic
    panic("Something went wrong!")
}
```

## `RecoverWithTrackedValues(values ...interface{})`

This function enhances the basic recovery by allowing you to log additional tracked values along with the error and stack trace.

### Example:

```go
package main

import (
	"github.com/ildommm/recoverich"
)

func main() {
    defer recoverich.RecoverWithTrackedValues("Track this value", 42)

    // Your code that may panic
    panic("Something went wrong!")
}
```

Results in the following log output:
```
ERROR: Something went wrong!
ERROR: Stacktrace dump ***
goroutine 1 [running]:
runtime/debug.Stack(0xc000010150, 0x10, 0x40)
	/usr/local/go/src/runtime/debug/stack.go:24 +0xa2
github.com/ildommm/recoverich.RecoverWithTrackedValues(0xc000010150, 0x10, 0x40)
	/path/to/your/code/recoverich.go:24 +0x11d
panic(0x4ba880, 0xc000010160)
	/usr/local/go/src/runtime/panic.go:965 +0x1b9
main.main()
	/path/to/your/code/main.go:8 +0x95
*** end
ERROR: Tracked values dump ***
string_0: Track this value
int_1: 42
*** end

```

## `RecoverWithContextValues(ctx context.Context)`
This function is designed for use with the context package. It recovers from a panic, logs the error and stack trace, and also logs values associated with the context.

### Example:
```go
package main

import (
    "context"
    "fmt"
    "github.com/ildomm/recoverich"
)

func main() {
    ctx := context.WithValue(context.Background(), "exampleKey", "exampleValue")

    defer recoverich.RecoverWithContextValues(ctx)

    // Your code that may panic
    panic("Something went wrong!")
}
```
Note: Ensure that your context values are compatible with reflection, and the Key and Value fields are accessible. This example assumes a simple case; you may need to adapt it for more complex scenarios.

Results in the following log output:
```
ERROR: Something went wrong!
ERROR: Stacktrace dump ***
goroutine 1 [running]:
runtime/debug.Stack(0xc000010150, 0x10, 0x40)
	/usr/local/go/src/runtime/debug/stack.go:24 +0xa2
github.com/ildommm/recoverich.RecoverWithContextValues(0xc000010150, 0x10, 0x40)
	/path/to/your/code/recoverich.go:22 +0x11d
panic(0x4ba880, 0xc000010160)
	/usr/local/go/src/runtime/panic.go:965 +0x1b9
main.main()
	/path/to/your/code/main.go:11 +0x95
*** end
ERROR: Context values dump ***
Key: "exampleKey", Type: string, Value: exampleValue
*** end

```

Feel free to explore and integrate these functions into your projects to enhance error recovery and debugging capabilities.