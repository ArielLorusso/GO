// 	01:44 	Package, alias, comment, main 	https://www.youtube.com/watch?v=YzLrWHZa-Kc
/*
HELP :
	go help
	go help <command>
	go help build 
	go build [-o output] [build flags] [packages]

EXECUTE AS SCRIPT :
	go run           "/home/ariel/Desktop/GO/Derek_Banas/C1_Output.go
EXECUTE AS BINARY :
	go build -o test "/home/ariel/Desktop/GO/Derek_Banas/C1_Output.go"
	./test 
*/

package main

import (
	"fmt" 	//format  (print)
)

var pl = fmt.Println // macro


func main () {
	pl("Hello world")
}