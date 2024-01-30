package recoverich

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"unsafe"
)

// RecoverWithContextValues recovers from a panic and logs the error and stack trace.
// It also logs the context values.
func RecoverWithContextValues(ctx context.Context) {
	if err := recover(); err != nil {

		print(err)

		// Pile the values formatted
		log.Print("Context values ***\n")
		printPile(gatherContextValues(ctx))
		log.Print("***")
	}
}

// gatherContextValues iterates over the context and extracts the tracked values
func gatherContextValues(ctx context.Context) []string {
	return gatherParentContextValues(ctx)
}

// gatherParentContextValues iterates over the context and extracts the tracked values
// from the context and its parents
func gatherParentContextValues(ctx context.Context) []string {
	values := make([]string, 0)

	for {
		parent := parentContext(ctx)

		if parent == nil {
			break
		}

		_values := extractContextValues(ctx)
		values = append(values, _values...)

		ctx = parent
	}

	return values
}

// parentContext returns the parent context
func parentContext(ctx context.Context) context.Context {
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

// extractContextValues extracts the tracked values from the context
func extractContextValues(ctx context.Context) []string {
	values := make([]string, 0)
	ctxValue := reflect.ValueOf(ctx)

	if ctxValue.Kind() == reflect.Ptr && !ctxValue.IsNil() {
		elemValue := ctxValue.Elem()

		if elemValue.Type().String() == "context.valueCtx" {

			keyField := elemValue.FieldByName("key")
			valueField := elemValue.FieldByName("val")
			valueField = reflect.NewAt(valueField.Type(), unsafe.Pointer(valueField.UnsafeAddr())).Elem()

			if keyField.IsValid() && valueField.IsValid() {
				tracked := fmt.Sprintf("Name: %v, Type: %v, Value: %v", keyField, reflect.TypeOf(valueField.Interface()), valueField)
				values = append(values, tracked)
			}

			return values
		}
	}

	return values
}
