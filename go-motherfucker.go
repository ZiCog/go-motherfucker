package main

import (
	"fmt"
	"time"
)

// Strings, Go source is utf-8
var greekLowerCase string = "αβγδεζηθικλμνξοπρστυφχψω"
var greekUpperCase string = "ΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥΦΧΨΩ"


// Constants
const Π float32 = 3.14159

// Enumerations, "iota" starts at zero an increments each use
const (
	FIRST = iota + 1
	SECOND
	THIRD
	FOURTH
)

// Variable types
var a int = 01   // Octal, int is natural size for machine
var b int8 = 0x1 // Hexadecimal
var c int16 = 2  // Decimal
var d int32 = 3
var e int64 = 4
var f uint8 = 5
var g uint16 = 6
var h uint32 = 7
var i uint64 = 8

var j float32 = 3.1e2 // No "float" type
var k float64 = 9

var l complex64 = 10 + 0i
var m complex128 = 5 + 5i

var n bool

var o error // Special type for errors

// Raw string literals
var p string = `A raw string literal
can span many lines`

var q string = "Normal string" + p // Use + to concatenate strings

// Runes, for unicode
var r rune = 'd'

// MyError is an error implementation that includes a time and message.
type MyError struct {
	When time.Time
	What string
}

// Provide an Error funtion so MyError can be used as an error type
func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

// How to change a char in an immutable string 
// Change the nth character of string s to c
// Returns modified string and error
func tweakCharacter(s string, n int, c rune) (string, error) {
	// Check length of string
	if n < len(s) {
		runes := []rune(s)        // Convert s to array of runes 
		runes[n] = c              // Tweak the nth character to c
		return string(runes), nil // Conver to new string and return it
	}
	return "", MyError{time.Now(), "Index out of range"}
}

func main() {
	// Declare and initialize combined, note the "Duck" typing 
	α := "α"
	β := "β"
	_, z := 2, 3 // Assign to _ is discarded.
	_ = z

	// Cannot mix types
	//d = h
	msg := "go motherfucker!"
	msg, err := tweakCharacter(msg, 0, 'G')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msg)
	fmt.Println(greekLowerCase)
	fmt.Println(greekUpperCase)
	fmt.Printf("Alpha = %v\n", α)
	fmt.Printf("Beta  = %v %v\n", β, SECOND)

	// "if" and "switch" statements can have an intializer, like "for" in C
	if a = 2; 1 < 2 {
		a = 1
	} else {
		a = 2
	}
	// Go has a goto, yay! 
	goto there
there:	
	// Forever loop, note breaks can have loop labels
outter:	for {
		// Normal C style for loop, note the "break"
		for a = 1; a < 10; a++ {
			// For loop in the style of while
inner:			for a < 10 {
				break inner
				continue
			}
			break outter
		}
		break
	}
	
	// Loops can loop over slices, arrays, stings, maps and channels with "range"
	// N.B. In this case the type of char is rune and the returned positions are NOT 0, 1, 2
	for pos, char := range "aΦx" {
		fmt.Printf("Character 0x%x starts at byte position %d\n", char, pos)
	}
	
	// switch can have an expression to be evaluated
	switchExpression := "hello"			
	switch switchExpression {
		case "hello":
			fmt.Println("Case: hello")
			fallthrough			// Won't fall through unless told
		case "bye":
			fmt.Println("Case: bye")
		case "fish", "dog", "ant":		// Comma separated list of cases
		default:
			fmt.Println("Case: default")	// If all else fails do this
	}
	
	// switch can be expressionless, first true case is executed
	// Can be used like if-else-if-else chain
	switch {
		case switchExpression == "first":	// We can compare strings easily
			fmt.Println("First case")
		case switchExpression == "second":
			fmt.Println("Second case")

		default:
			print("No case")		// print and println are built in functions
			println("")
	}
}
