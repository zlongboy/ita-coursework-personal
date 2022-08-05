package main

import (
	"fmt"
	//"github.com/gin-gonic/gin"
	// "strings" // package with Go string methods
	"strconv" // Itao() method converts numeric values to strings, Atoi(), ParseFloat() converts strings to numerics
	// "io" //io.Reader, Read() for reading data
	// "math/rand"
	"math"
	// "time" // Sleep() method like setTimeout?
	"golang.org/x/tour/pic"
)

func main() {
	fmt.Println("Answer: ", Sqrt(64))

	testDefer()

	varNotes()

	testStruct()

	pic.Show(Pic)

	tp := sampleStruct{10, "Strikeouts"}
	// a := &tp
	// a.prec() // In this case interprests a as a* as well
	tp.prec() // Interprets tp as &tp because the function is a pointer receiver

	etest("") // Catching errors

	rouChan()
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
			zero value of pointer is nil
			Vars can still be undefined if you try to call them before declaring
		Types (note aliases):
			byte = uint8
			rune = int32
		If unspecified, bits will convert to be compatible with system
		Converting var x to type T: T(x)
		Name exported variables/functions with uppercase first letter

		If vars are explicity typed and not compatible, math functions cannot be performed with them (e.g. multiply int and int64) --> must convert to standard
		When converting from string to numeric --> check returned value! Methods will return error if string is not a number
		Math: unary operators (+/-) use to change sign of var. divide operator (/) on two ints floors output, unless explicity converted to float
		Combining math with assignment (e.g. +=)

		Boolean logic: || (or), && (and), ! (not)
		Switch statements without conditions => switch true
		No key found when searching a map will return 0

		Slices / Arrays / Maps
			Slice is of variable length but takes up more memory
			Changing elements of a slice modifies the corresponding elements (all slices that share underlying array will be affected)
			len(), cap() - capacity of slice is number of elements in underlying array
			zero value = nil
			Slice can be of type slice e.g. sliceOfSlices := [][]int{ }
			What if underlying array is not long enought to hold new appended slice values? A new underlying array will be created!
				QUESTION: What happens to the old underlying array, memory address
			Map is unordered so you can't search by index
			Check if key is in map, if true, assign key value to var
				elem, ok := m[key]

		Seed values for random integer generation (rand.Intn() method) -- use time as a seed otherwise method will generate same integer each run
			rand.Seed(time.Now().UnixNano())

		Functions
			func (var receiver) name(arg argType) returnType { }
			Specify return type after argument parens on declaration
			Can return multiple values!
			defer keyword (defer to end of execution block) -- executed last-in first-out at end of wrapping function execution
			*Closures
				QUESTION: Functions that return functions? TODO: https://go.dev/tour/moretypes/26
			Receivers: Methods for Go // Good explanation: https://medium.com/@adityaa803/wth-is-a-go-receiver-function-84da20653ca2
				The var declaration (of type struct (or not)) that goes before the function name is the variable that this function is bound to
				--> Once bound, can call func as a method of that var
				Using value or pointer?
					Avoid copying large data on each function call
					Stay consistent -- all methods on a var should use same implementation

		Structs
			Custom type
			Think a combination of a Node class (object/constructor)

		Interfaces
			Another custom type (like structs)
			A collection of methods
			To implement an interface, all methods in the collection must be used
			QUESTION: Giving methods to structs? Why not just define the method for the struct?

		Addresses / Pointers
			Addresses: Where in memory a variable is stored
				Print address (fmtPrint(&[var]))
				&var returns the the address of var
				We can then save that address in a pointer (points to the address) ptr = &var
			Pointers: variables that store addresses
				Asterisk denotes underlying value, e.g. *ptr => value at address ptr (which is &var) so value of var
			Usage
				Changing variables outside of scope (read, set value through pointer)
				When passing a value into a function it is not mutable,
					Need to instead pass in the pointer and then edit the underlying value at that pointer in order to change the value
					Otherwise you are just changing a local (local to the function) copy and original value is not changed
				QUESTION: using with structs?

		Loops - for loop only
			Parts: init (e.g. declare vars needed for loop); condition (do until); post (what to do after iteration, e.g. i++) { }
			break => exits, goes to end of loop
			continue => goes back to top of loop
			Keyword: range
				Used almost as a for each/for in the loop will iterate over slice/map
				for [index], [copy of element] := range [slice]
					can replace index or element with _, if uneeded

		Concurrency
			Goroutines - threads
				go f(x, y) <-- x and y in current thread, f in a new thread
			Channels <-
				Different conduits, usage: sync goroutines
				define: ch := make(chan int, buffer)
				send: ch <- [value to be sent]
				Buffer, block if full (specifies capacity of chan)
			Select, block until one of cases can run (if multiple, chosen randomly <-- not good)
				default, runs if nothing else can run (prevents total blocking)

		*TOUR OF GO: CONCEPTS OVER MY HEAD
			Closures
			Type assertions and switches
			goroutines: mutual exclusion

	*/

	// Using short declaration wrap value in type declaration (if no type specified Go will infer)
	i := int(20001)

	a := [5]string{"royal", "air", "maroc", "sliced", "off"}

	//s := []string{"royal", "air", "maroc"} // Slice literal, creates array under the hood and makes a slice that references
	// s := make([]int, 5, 10) // Another way to create a slice, pass in lenght and capacity
	var s []string = a[1:4] // Create a slice from array (includes first index but not last)

	// Map -  unordered. Access via keys, e.g. m["model"]
	m := map[string]int{"model": 747, "captain": 1, "delta": 8} // GOOGLE: Println map formats in alphabetical order?

	s = append(s, "Casablanca") // Appending values to a slice

	j := int64(i)                         // Converting int types
	fmt.Printf("Converted type: %T\n", j) // Print data types

	fmt.Println("Slice:", s, "Slice length:", len(s), "Map:", m)
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

// Playing around with deferring & pointer dereferencing
func testDefer() {
	v := "Prone"
	fmt.Println("First:", v)
	// defer fmt.Println("Defer:", v) // Doesn't work because it evaluates argument at execution line (just doesn't ex)
	wait(&v)
	fmt.Println("Defer:", v) // Works because now executes after wait execution -- is there no asynchronisity in Go?
}

// Dereference: update value of var outside scope
func wait(t *string) {
	// time.Sleep(2 * time.Second)
	*t = "Slept"
}

// Struct playground
func testStruct() {
	type NewStruct struct {
		x int
		y string
	}

	n := NewStruct{23, "Help me"}
	np := &n

	fmt.Println("New Struct:", n)
	fmt.Println("Struct property value:", n.x)
	fmt.Println("value using pointer:", np.x) // Don't have to explicity dereference when referring to struct (e.g. *np.x)

}

// Addresses of slices and arrays
func sa() {
	a := [3]int{0, 1, 2}
	var s []int = a[:]
	fmt.Printf("Slice address: %p and then Array address: %p", &s, &a) // QUESTION: why not same address?
	fmt.Println()
	fmt.Printf("slice element 1: %p and array element 1: %p", &s[1], &a[1]) // Same address!

	//s = append(s, 3) // expanding slice longer than underlying array
	//a[0] = 1 // changing value of underlying array, after new underlying array has been created for slice

	//fmt.Println(s)
	//fmt.Println(a)
	//fmt.Printf("Slice address: %p and then Array address: %p", &s, &a)

}

// Looping with slices and arrays exercise
func Pic(dx, dy int) [][]uint8 {
	ss := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		s := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			s[j] = uint8(i ^ j)
		}
		ss[i] = s
	}
	//fmt.Println(ss)
	return ss
}

// Pointer receiver function example (method)
type sampleStruct struct {
	v int
	z string
}

func (p *sampleStruct) prec() {
	fmt.Printf("v is %d and z is %v\n", p.v, p.z)
	p.z = "Walks"
	fmt.Printf("v is %d and z is %v\n", p.v, p.z)
}

// Wouldn't work because it's just changing the local copy of p.name
// type person struct {
// 	name string
//    }
//
//	func main() {
// 	p := person{"Richard"}
// 	rename(p)
// 	fmt.Println(p)
//    }
//
//	func rename(p person) {
// 	p.name = "test"
//    }

// Does work because changing the underlying value (*) of the pointer to p.name
//  func main() {
// 	p := person{"Richard"}
// 	rename(&p)
// 	fmt.Println(p)
//    }
//
//	func rename(p *person) {
// 	p.name = "test"
//    }

// Source: https://medium.com/@meeusdylan/when-to-use-pointers-in-go-44c15fe04eac

// Errors example
func etest(t string) {
	i, err := strconv.Atoi(t)
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)
}

// Goroutines and channels example
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func rouChan() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
