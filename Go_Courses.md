# INSTALL GO

https://go.dev/dl/  
```sh
    #  DOWNLOAD  go1.22.2.linux-386.tar.gz  X86 (NOT ARM)
```
https://go.dev/doc/install
```sh
 # onnce  go1.22.2.linux-386.tar.gz   DOWNLOADED to desktop
 rm -rf /usr/local/go && tar -C /usr/local -xzf /home/ariel/Desktop/go1.22.2.linux-amd64.tar.gz
 # change /home/ariel/Desktop/ to your file path
 export PATH=$PATH:/usr/local/go/bin
 go version
```
## GO TOOLS dlv & gopls
```sh
go install -v github.com/go-delve/delve/cmd/dlv@latest
go install -v golang.org/x/tools/gopls@latest

```
## IF ERROR FOR TOOLS
https://stackoverflow.com/questions/54082459/fatal-error-bits-libc-header-start-h-no-such-file-or-directory-while-compili
```sh
go install -v github.com/go-delve/delve/cmd/dlv@latest
runtime/cgo
# runtime/cgo
In file included from _cgo_export.c:3:
/usr/include/stdlib.h:26:10: fatal error: bits/libc-header-start.h: No such file or directory
   26 | #include <bits/libc-header-start.h>
      |          ^~~~~~~~~~~~~~~~~~~~~~~~~~
compilation terminated.

```

```sh
 sudo apt-get install gcc-multilib
 sudo apt-get install libx11-dev:i386 libx11-dev
```


# MANUAL 
```sh
	go help
	go help <command>
	go help build 
	go build -o <output> <[build flags> <packages>
    go build -o test      -x            "/home/ariel/Desktop/GO/Derek_Banas/C1_Output.g
```


# DOCS
https://go.dev/doc/effective_go     Intoduction
https://go.dev/doc/faq              FAQ
https://pkg.go.dev/std              Standard library
https://pkg.go.dev/fmt@go1.22.2                 std::format
https://pkg.go.dev/github.com/gin-gonic/gin     Gin Web framework
# TUTORIAL LIST

1) Derek_Banas         Golang Tutorial     3:49:14     https://www.youtube.com/watch?v=YzLrWHZa-Kc
2) bootdotdev           Go Programming     9:32:46     https://www.youtube.com/watch?v=un6ZyFkqFKo
3) Tech_With_Tim      Golang Tutorials     4:48:06     https://www.youtube.com/playlist?list=PLzMcBGfZo4-mtY_SE3HuzQJzuj4VlUG0q
4) Net Ninja    (Golang) for Beginners     3:10:28     https://www.youtube.com/playlist?list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM
5) Michael_Van_Sickle   Go Programming     6:39:57     https://www.youtube.com/watch?v=YS4e4q9oBaU
6) Nana           Golang for Beginners     3:24:58     https://www.youtube.com/watch?v=yyUHQIec83I
7) Alex_Mux              Learn GO Fast     1:07:52     https://www.youtube.com/watch?v=8uiZC0l4Ajw
8) Learn Go by Building 11 Projects        8:24:39     https://www.youtube.com/watch?v=jFfo23yIWac

# Derek_Banas           3:49:14     
https://www.youtube.com/watch?v=YzLrWHZa-Kc
*   00:00   Intro
*   01:44   Package
*   02:15   Import
*   02:42   Alias
*   03:19   Comments
*   03:40   Main
*   04:27   User Input
*   06:00   Error Handling
*   06:35   Blank Identifier
*   07:33   Variables
*   10:19   Data Types
*   12:12   Casting
*   12:30   Casting Strings
*   17:33   If Conditional
*   20:56   Strings
*   29:19   Runes
*   31:15   Printf
*   33:15   Time
*   34:39   Math
*   44:35   For Loop
*   46:25   While Loop
*   53:16   Range
*   54:23   Arrays
* 1:01:18   Slices
* 1:08:29   Functions
* 1:10:20   Return Multiple
* 1:11:05   Function Errors
* 1:13:11   Varadic Functions
* 1:14:35   Passing Arrays
* 1:17:50   Pointers
* 1:21:35   Pass Array Pointers
* 1:26:09   File IO
* 1:37:11   Command Line
* 1:43:03   Packages / Modules
* 1:52:40   Maps
* 1:59:27   Generics
* 2:00:38   Constraints
* 2:04:24   Structs
* 2:12:04   Composition
* 2:16:05   Defined types
* 2:21:54   Associate Methods
* 2:24:50   Protecting Data
* 2:25:12   Getter / Setter
* 2:31:00   Encapsulation
* 2:33:42   Interfaces
* 2:38:39   Concurrency / GoRoutines
* 2:40:49   Sleep
* 2:42:10   Channels
* 2:45:16   Mutex / Lock
* 2:51:13   Closures
* 2:53:13   Passing Functions
* 2:55:50   Recursion
* 2:58:59   Regular Expressions
* 3:07:00   Automated Testing
* 3:16:42   Web app
* 3:38:44   Templates / HTML
* 3:45:16   Installation

# bootdotdev            9:32:46 
https://www.youtube.c/watch?v=un6ZyFkqFKo
*   0:00:00   Intro
*   0:03:17   Ch  1. Why write Go?
*   0:27:39   Ch  2. Variables
*   0:51:11   Ch  3. Functions
*   1:16:58   Ch  4. Structs
*   1:34:36   Ch  5. Interfaces
*   2:00:26   Ch  6. Errors
*   2:22:01   Ch  7. Loops
*   2:48:21   Ch  8. Slices
*   3:39:54   Ch  9. Maps
*   4:06:19   Ch 10. Advanced functions
*   4:31:03   Ch 11. Pointers
*   4:48:02   Ch 12. Local development
*   5:31:43   Ch 13. Channels & concurrency
*   6:07:38   Ch 14. Mutexes
*   6:30:56   Ch 15. Generics
*   6:38:38   Ch 16. Quiz
*   6:43:13   P1. RSS aggregator project
*   6:53:43   P2. Chi router
*   7:11:37   P3. Postgres database
*   7:39:10   P4. Authentication w/ API keys
*   8:18:28   P5. Many to many relationships
*   8:39:13   P6. Aggregation worker
*   9:05:28   P7. Viewing blog posts  
  
# Net Ninja             3:10:28 

* 1   4:40 - Introduction & Setup
* 2  13:54 - Your First Go File
* 3  12:26 - Variables, Strings & Numbers
* 4  10:57 - Printing & Formatting Strings
* 5  12:31 - Arrays & Slices
* 6   8:45 - The Standard Library
* 7   8:50 - Loops
* 8  10:22 - Booleans & Conditionals
* 9   8:03 - Using Functions
* 10  6:21 - Multiple Return Values
* 11  6:49 - Package Scope
* 12  9:05 - Maps
* 13  9:09 - Pass By Value
* 14  5:38 - Pointers
* 15 10:17 - Structs & Custom Types
* 16  7:21 - Receiver Functions
* 17 14:52 - Receiver Functions with Pointers
* 18  7:34 - User Input
* 19  6:09 - Switch Statements
* 20  4:53 - Parsing Floats
* 21  5:37 - Saving Files
* 22  5:37 - Interfaces


# Tech_With_Tim         4:48:06     
https://www.youtube.com/playlist?list=PLzMcBGfZo4-mtY_SE3HuzQJzuj4VlUG0q
*   14:47   #1 - An Introduction to Go Programming
*   8:25    #2 - Variables & Data Types
*   14:04   #3 - Assignment Expression & Implicit vs Explicit
*   10:47   #4 - Printing to Console & fmt
*   9:06    #5 - Console Input (Bufio Scanner) & Type Conversion
*   8:16    #6 - Arithmetic Operators & Math
*   7:31    #7 - Conditions & Boolean Expressions
*   13:53   #8 - Chained Conditionals (AND, OR, NOT)
*   12:39   #9 - If, Else If, Else
*   5:36    #10 - For Loops & While Loops
*   13:30   #11 - Switch Statement
*   15:32   #12 - Arrays
*   15:44   #13 - Slices
*   9:57    #14 - Range & Slice/Array Examples
*   14:29   #15 - Maps
*   14:33   #16 - Functions
*   17:45   #17 - Advanced Function Concepts & Function Closures
*   20:35   #18 - Mutable & Immutable Data Types
*   18:49   #19 - Pointers & Derefrence Operator (& and *)
*   13:41   #20 - Structs and Custom Types
*   12:19   #21 - Struct Methods
*   10:00   #22 - Interfaces


 
# Michael_Van_Sickle    6:39:57
https://www.youtube.com/watch?v=YS4e4q9oBaU
*   0:00:00   Introduction 
*   0:16:57   Setting Up Dev Environment
*   0:35:48   Variables
*   0:57:05   Primitives
*   1:26:29   Constants
*   1:47:53   Arrays and Slices
*   2:17:20   Maps and Structs
*   2:48:00   If and Switch Statements
*   3:21:17   Looping
*   3:41:34   Defer, Panic, and Recover
*   4:03:57   Pointers
*   4:21:30   Functions
*   4:57:59   Interfaces
*   5:33:57   Goroutines
*   6:05:10   Channels
  
# Nana                  3:24:58
https://www.youtube.com/watch?v=yyUHQIec83I
*      0:00 - Intro & Course Overview
*   INTRODUCTION TO GO
*     02:47 - What is Go? Why Go? How it's different?
*     06:50 - Characteristics of Go and Go Use Cases
*   GO SYNTAX & CONCEPTS
*     08:59 - Local Setup - Install Go & Editor
*     12:54 - Write our First Program & Structure of a Go File
*     22:02 - Variables & Constants in Go
*     30:43 - Formatted Output - printf 
*     33:43 - Data Types in Go
*     45:18 - Getting User Input
*     47:19 - What is a Pointer?
*     53:55 - Book Ticket Logic
*     57:16 - Arrays & Slices
*   1:11:12 - Loops in Go
*   1:24:24 - Conditionals (if / else) and Boolean Data Type
*   1:39:33 - Validate User Input
*   1:54:02 - Switch Statement
*   1:58:37 - Encapsulate Logic with Functions
*   2:22:36 - Organize Code with Go Packages
*   2:35:39 - Scope Rules in Go
*   2:37:16 - Maps
*   2:53:20 - Structs
*   3:02:15 - Goroutines - Concurrency in Go
*   3:23:51 - Congratulations!


# Alex_Mux              1:07:52
https://www.youtube.com/watch?v=8uiZC0l4Ajw
*    0:00   Introduction to Golang
*    6:25   Constants Variables and Basic Data Types
*   13:14   Functions and Control Structures
*   19:30   Arrays, Slices, Maps and Loops
*   26:36   Strings, Runes and Bytes
*   31:06   Structs and Interfaces
*   35:18   Pointers
*   40:06   Goroutines
*   47:10   Channels
*   52:56   Generics
*   55:47   Building an API!

# Learn Go  by Building 11 Projects
https://www.youtube.com/watch?v=jFfo23yIWac
1.  0:00:00  Build A Simple Web Server With Golang - https://github.com/AkhilSharma90/simp...
2.  0:20:34  Build A CRUD API With Golang       - https://github.com/AkhilSharma90/Gola...
3.  1:07:14  Golang With MySQL Book Management System - https://github.com/AkhilSharma90/Gola...
4.  2:30:57  Simple SlackBot To Calculate Age   - https://github.com/AkhilSharma90/GO-S...
5.  2:44:12  Golang Slackbot for File Uploading - https://github.com/AkhilSharma90/GO-S...
6.  3:01:45  Email Verifier Tool With Golang    - https://github.com/AkhilSharma90/GO-E...
7.  3:24:32  AWS Lambda With Golang             - https://github.com/AkhilSharma90/Simp...
8.  3:50:12  CRM with Golang Fiber              - https://github.com/AkhilSharma90/go-b...
9.  4:34:34  HRMS with Golang Fiber             - https://github.com/AkhilSharma90/go-b...
10. 5:44:25 Complete Serverless Stack           - https://github.com/AkhilSharma90/Gola...
11. 7:28:55  A.I Bot with Wolfram, Wit.ai and golang - https://github.com/AkhilSharma90/AI-B...
    