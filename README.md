# flatten

* [Overview](#pkg-overview)
* [Features](#pkg-features)
* [Usage](#pkg-usage)
* [Code analysis](#pkg-analysis)
* [Tests](#pkg-tests)
* [Benchmark](#pkg-benchmark)
* [Documentation](#pkg-documentation)
* [Example](#pkg-example)

## <a name="pkg-overview">Overview</a>
Package **flatten** flattens arrays and slices.

## <a name="pkg-features">Features</a>

- Accept any number of arguments.
- Process complex or simple nested arrays or slices.
- Process any type of array or slice.
- Has no external dependency
- Code follows the standard Golang format

## <a name="pkg-usage">Usage</a>

```
import "github.com/thalesfsp/flatten"
```

## <a name="pkg-analysis">Code analysis</a>
**flatten** pass a rigid set of analysis, including `gosec`. To analyze the 
code, just run:

	make analyze
	
_Note that you will have to have such tools previously installed. The command 
will print out a list._

## <a name="pkg-tests">Tests</a>
Code has 100% coverage. To test, just run:

	make test

## <a name="pkg-benchmark">Benchmark</a>
All functions are benchmarked. To benchmark, just run:

	make benchmark

## <a name="pkg-documentation">Documentation</a>
**flatten** code is well documented, please check the comments.

## <a name="pkg-example">Example</a>

Code:

```
// The following code demonstrates how to flatten a 4D array and assumes
// that the flatten package has already been imported.

// Multidimensional array (4D)
const rolls = 4
const columns = 4
var fourD [rolls][columns]int

// Generator/filler
for roll := 0; roll < rolls; roll++ {
    for column := 0; column < columns; column++ {
        fourD[roll][column] = (roll + 1) + (column * 2)
    }
}

result, err := Flatten(fourD)
if err != nil {
    panic(err)
}

fmt.Println(result)
```

Expected output:

`[1 3 5 7 2 4 6 8 3 5 7 9 4 6 8 10]`