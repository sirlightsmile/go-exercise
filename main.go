package main

//
// Go routines #2
//
//
// Implement PromiseAll function. The function implements concurrent execution of 1 or more functions and return
// result as an array. If any of the functions has error return an empty array and error (don't wait for other
// routines to finish). The result array must be sorted to match the order of the input functions.
// Refer to test file for example input and output.  All tests must pass.
// Implementation must be free of race conditions (go test --race).
//
// Hint: one of the possible implementation uses waitGroup, channel and range over channel.
//
//

func PromiseAll(fns ...func() (interface{}, error)) ([]interface{}, error) {

	count := len(fns)
	ch := make(chan interface{}, count)

	var result []interface{}

	for _, fn := range fns {
		r, err := fn()

		if err != nil {
			return nil, err
		}

		ch <- r
		result = append(result, <-ch)
	}

	return result, nil
}

func main() {
}
