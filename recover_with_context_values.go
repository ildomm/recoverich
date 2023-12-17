package recoverich

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"unsafe"
)

// RecoverWithContextValues recovers from a panic and logs the error and stack trace.
// It also logs the context values.
func RecoverWithContextValues(ctx context.Context) {
	if err := recover(); err != nil {

		const size = 64 << 10
		stackTrace := make([]byte, size)
		stackTrace = stackTrace[:runtime.Stack(stackTrace, false)]

		switch x := err.(type) {
		case string:
			err = errors.New(x)
		default:
			err = fmt.Errorf("unknown panic: %w", x.(error))
		}

		log.Printf("ERROR: %v", err)
		log.Printf("ERROR: Stacktrace dump ***\n%s\n*** end\n", stackTrace)

		values := dumpValues(ctx)
		log.Printf("ERROR: Context values dump ***\n%v\n*** end\n", values)
	}
}

func dumpValues(ctx context.Context) []string {
	return iterateOverParents(ctx)
}

func iterateOverParents(ctx context.Context) []string {
	values := make([]string, 0)

	for {
		parent := getParentContext(ctx)

		if parent == nil {
			break
		}

		_values := extractValuesFromContext(ctx)
		values = append(values, _values...)

		ctx = parent
	}

	return values
}

func getParentContext(ctx context.Context) context.Context {
	ctxValue := reflect.ValueOf(ctx)
	if ctxValue.Kind() == reflect.Ptr && !ctxValue.IsNil() {
		elemValue := ctxValue.Elem()
		if elemValue.Kind() == reflect.Struct {
			contextField := elemValue.FieldByName("Context")
			if contextField.IsValid() && contextField.Type().AssignableTo(reflect.TypeOf((*context.Context)(nil)).Elem()) {
				return contextField.Interface().(context.Context)
			}
		}
	}
	return nil
}

func extractValuesFromContext(ctx context.Context) []string {

	values := make([]string, 0)
	ctxValue := reflect.ValueOf(ctx)

	if ctxValue.Kind() == reflect.Ptr && !ctxValue.IsNil() {
		elemValue := ctxValue.Elem()

		if elemValue.Type().String() == "context.valueCtx" {

			keyField := elemValue.FieldByName("key")
			valueField := elemValue.FieldByName("val")
			valueField = reflect.NewAt(valueField.Type(), unsafe.Pointer(valueField.UnsafeAddr())).Elem()

			if keyField.IsValid() && valueField.IsValid() {
				tracked := fmt.Sprintf("Key: %v, Type: %v, Value: %v\n", keyField, reflect.TypeOf(valueField.Interface()), valueField)
				values = append(values, tracked)
			}

			return values
		}
	}

	return values
}
