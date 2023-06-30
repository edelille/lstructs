# Lam-Structures
This is a useful test to publish to the golang packages and allows users to pull from here

Please refer to all of the functions and data structures in https://pkg.go.dev/github.com/edelille/lstructs


# How to use

1. Import via `"github.com/edelille/lstructs"` on top of the file
2. Grab it using `$ go get github.com/edelille/lstructs`
3. Update using `$ go mod tidy`


# Example Code
```go
package main

import (
    "fmt"

    ls "github.com/edelille/lstructs"
)

func main() {
    // Initialize the data structure
    stack := ls.NewStack[int]()

    // Push onto the Stack!
    stack.Push(5)

    // Peek on top of the Stack!
    fmt.Println(stack.Peek())   // Should print 5

    // Pop off the top!
    fmt.Println(stack.Pop())   // Should print 5, and removes 5 off of stack

    fmt.Println(stack.Size())   // Should print 0
    fmt.Println(stack.IsEmpty())   // Should print true
}
```
