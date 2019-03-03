// Go's `sort` package implements sorting for builtins
package main

import "fmt"
import "sort"

func main() {

	// Sort methods are specific to the builtin type;
	// here's an example for strings.
	strs := []string{"Shim", "Bernard", "Joy"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// We can also use `sort` to check if a slice is
	// already in sorted order.
	strSorted := sort.StringsAreSorted(strs)
	fmt.Println("Strings sorted: ", strSorted)

	// An example of sorting `int`s.
	ints := []int{8, 2, 0, 1}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// We can also use `sort` to check if a slice is
	// already in sorted order.
	intSorted := sort.IntsAreSorted(ints)
	fmt.Println("Integers sorted: ", intSorted)
}
