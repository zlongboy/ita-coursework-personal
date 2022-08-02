package main

import (
	"fmt"
	//"github.com/gin-gonic/gin"
	// "strings" // package with Go string methods
	// "strconv" // Itao() method converts numeric values to strings, Atoi(), ParseFloat() converts strings to numerics
	// "time"
	// "math/rand"
	"math"
)

func main() {
	fmt.Println("Answer: ", Sqrt(64))
	varNotes()
}

func varNotes() {
	/*
		Go standards typically use shorter variable names
		Printf() verbs:
			%T - print data type of var
			%v - print var
			%q - print quoted value
			%d - interpolate number into a string
			%.2f - print float with two decimal points
		fmt.Sprint() method converts float to string
		fmt docs: https://pkg.go.dev/fmt

		Can only use short var declaration (:= with implicit type) in function scope, not global/package level
		Go values can never be undefined (e.g. bool resolves to false, int resolves to 0, and string resolves to "", if not defined)
			Vars can still be undefined if you try to call them before declaring
		Types (note aliases)
			byte = uint8
			rune = int32
		If unspecified bits will convert to compatible with system
		Converting var x to type T: T(x)
		Name exported variables with uppercase first letter
		Keyword: range <-- GOOGLE: is this some sort of index? Used in for loops similar to a for in loop

		If vars are explicity typed and not compatible, math functions cannot be performed with them (e.g. multiply int and int64) --> must convert to standard
		When converting from string to numeric --> check returned value! Methods will return error if string is not a number
		Math: unary operators (+/-) use to change sign of var. divide operator (/) on two ints floors output, unless explicity converted to float
		Combining math with assignment (e.g. +=)

		Boolean logic: || (or), && (and), ! (not)
		Switch statements without conditions => switch true
		No key found when searching a map will return 0

		Seed values for random integer generation (rand.Intn() method) -- use time as a seed otherwise method will generate same integer each run
			rand.Seed(time.Now().UnixNano())

		Functions
			specify return type after argument parens on declaration
			Can return multiple values!
			defer keyword (defer to end of execution block)

		Addresses / Pointers
			Addresses: Where in memory a variable is stored
			Print address (fmtPrint(&[var]))
			Pointers: variables that store addresses (*)
			QUESTION: why would you use dereferencing?
				- Changing variables outside of scope? Changing pointer (value at address)

		Loops - for loop only
			init (e.g. declare vars needed for loop); condition (do until); post (what to do after iteration, e.g. i++) { }
			break - exits, goes to end of loop
			continue - goes back to top of loop

	*/
	// Using short declaration wrap value in type declaration (if no type specified Go will infer)
	i := int(20001)
	// Slices - more flexible arrays (arrays are of fixed length) but requiring more memory
	s := []string{"royal", "air", "maroc"}

	// Map -  unordered. Access via keys, e.g. m["model"]. (GOOGLE: Println formats in alphabetical order?)
	m := map[string]int{"model": 747, "captain": 1, "delta": 8}

	// Appending values to a slice
	s = append(s, "Casablanca")

	// Converting int types and printing var type
	j := int64(i)
	fmt.Printf("%T\n", j)

	fmt.Println(i, s, m, len(s))
}

// Sqrt function (Newton's method) from Tour of Go exercises
func Sqrt(x float64) float64 {
	z := 1.00
	c := 0
	for {
		c++
		p := z
		fmt.Println(c, "-", "p: ", p)
		z -= (z*z - x) / (2 * z)
		fmt.Println(c, "-", "z: ", z)
		if math.Abs(p-z) > 0 {
			continue
		} else {
			break
		}
	}
	return z
}
