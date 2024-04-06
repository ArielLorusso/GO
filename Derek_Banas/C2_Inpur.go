// 04:27	User Input	https://www.youtube.com/watch?v=YzLrWHZa-Kc

package main

import (
	"fmt" 	 // format  (Print)
	"bufio" // buffer
	"os" 	 // operation sys
	"log" 	 // operation sys
)

var pl = fmt.Println // macro


func main () {
	pl("Enter your name :")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n') // read till '\n'
	if ( err== nil ) {
		pl("Hello",name) 
	} else {  
		log.Fatal(err) 
	}
}