/*
Package flatten flattens arrays and slices.

Features

    - Accept any number of arguments.
    - Process complex or simple nested arrays or slices.
    - Process any type of array or slice.
    - Has no external dependency
    - Code follows the standard Golang format

Usage

    import "github.com/thalesfsp/flatten"

analysis

flatten pass a rigid set of analysis, including `gosec`. To analyze the
code, just run:

	make analyze

_Note that you will have to have such tools previously installed. The command
will print out a list._

Tests
Code has 100% coverage. To test, just run:

	make test

Benchmark
All functions are benchmarked. To benchmark, just run:

	make benchmark

Documentation

flatten code is well documented, please check the comments.
*/
package flatten

import (
	"fmt"
	"reflect"
)

// flattenDeep do the flatten process. It's responsible to check the type of
// the element of an array/slice and then decide if should append to `flatArray`
// or call it self.
func flattenDeep(
	flatArray []interface{},
	valueOfInterface reflect.Value,
) []interface{} {
	// `.Kind()` is responsible to get the real type of `valueOfInterface`,
	// "at this point" is unknown/anonymous (=== interface). Special care is
	// taken for cases where the type of the interface is another `interface`,
	// in which case it's necessary to obtain the concrete value (`.Elem()`)
	if valueOfInterface.Kind() == reflect.Interface {
		valueOfInterface = valueOfInterface.Elem()
	}

	// TDT, aka The Decision Tree :)
	// Should it call it self or just append?
	if valueOfInterface.Kind() == reflect.Array ||
		valueOfInterface.Kind() == reflect.Slice {
		for i := 0; i < valueOfInterface.Len(); i++ {
			// Recursion
			flatArray = flattenDeep(flatArray, valueOfInterface.Index(i))
		}
	} else {
		flatArray = append(flatArray, valueOfInterface.Interface())
	}

	return flatArray
}

// Flatten is a proxy to `flattenDeep` (private), wrapping and allowing it to
// process anything. It also provides an additional layer of protection, by not
// allowing empty interfaces to be processed, which could cause an cyclic
// problem (infinity recursion)
func Flatten(anything ...interface{}) ([]interface{}, error) {
	if len(anything) <= 0 {
		return nil, fmt.Errorf("Error: there's nothing to flatten")
	}

	return flattenDeep(nil, reflect.ValueOf(anything)), nil
}
