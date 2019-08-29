package main

import (
	"fmt"
	"sort"
	"time"
)

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

type FuncResult struct {
	index  int
	result interface{}
}

func PromiseAll(fns ...func() (interface{}, error)) ([]interface{}, error) {

	count := len(fns)
	ch := make(chan FuncResult, count)
	chErr := make(chan error)

	var frs []FuncResult

	for i, fn := range fns {

		go func(goFn func() (interface{}, error), index int) {
			fmt.Println("function loop run")
			r, err := goFn()

			if err != nil {
				chErr <- err
				return
			}

			fr := FuncResult{
				index:  index,
				result: r,
			}

			ch <- fr
		}(fn, i)
	}

	for {
		select {
		case err := <-chErr:
			return nil, err
		case rs := <-ch:
			frs = append(frs, rs)
			fmt.Println("function append")

			if len(frs) == count {
				sort.Slice(frs, func(i, j int) bool {
					return frs[i].index < frs[j].index
				})

				var r []interface{}
				for _, v := range frs {
					r = append(r, v.result)
				}

				return r, nil
			}
		}
	}
}

func main() {
	result, e := PromiseAll(
		func() (interface{}, error) {
			time.Sleep(5 * time.Second)
			return nil, nil
		},
		func() (interface{}, error) {
			return 1 + 2, nil
		},
		func() (interface{}, error) {
			return 6 * 6, nil
		},
		func() (interface{}, error) {
			time.Sleep(1 * time.Second)
			return true, nil
		},
		func() (interface{}, error) {
			return "result", nil
		},
	)

	if e != nil {
		panic(e)
	}

	for _, v := range result {
		fmt.Println(v)
	}
}
