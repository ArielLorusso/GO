// 10:19   Data Types   https://www.youtube.com/watch?v=YzLrWHZa-Kc

package main

import ("fmt" 	 // format  (Print)
		"reflect"
)
var pl = fmt.Println // macro


func main () {
	// int, float64, bool, string, rune
	// Default type 0, 0.0, false,
	pl (reflect.TypeOf(1))			// int
	pl (reflect.TypeOf(1.0))		// float64
	pl (reflect.TypeOf(false))		// bool
	pl (reflect.TypeOf('a'))		// int32
	pl (reflect.TypeOf("Hello"))	// string
}