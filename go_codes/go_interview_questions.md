What are string literals?
---
A string literal is a string constant formed by concatenating characters. 
The two forms of string literal are raw and interpreted string literals.    

Raw string literals are written within backticks (foo) and are filled with uninterpreted UTF-8 characters

Interpreted string literals are strings, written within double quotes

---
What are packages in a Go program?
---
a package is a collection of related Go source files that are compiled together

it provides a way to organize code into reusable modules

Packages help manage dependencies and promote code reusability by allowing developers to import and use them in their projects

---
What form of type conversion does Go support? Convert an integer to a float.
---
Go supports explicit type conversion to satisfy its strict typing requirements.

---
How do you stop go a routine?
---
you can use a cancellation pattern using a channel 
where you can send a value to this channel 

and this channel is waiting for a value when it gets the value it can return thus stopping the go routine

```
package main
func main() {
  quit := make(chan bool)
  go func() {
    for {
        select {
        case <-quit:
            return
        default:
           
        }
  }
}()

quit <- true
}
```

---
How do you check a variable type at runtime?
---
The Type Switch is the best way to check a variable’s type at runtime.

The Type Switch evaluates variables by type rather than value

```
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Double %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know  type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```

---
How do you concatenate strings
---
The easiest way to concatenate strings is to use the concatenation operator (+)
this plus operator

---
Explain the steps of testing with Golang.
---
write a small testcase for this question


---
What are function closures
---
Function closures is a function value that references variables from outside its body. 
The function may access and assign values to the referenced variables.

```
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

---
How do we perform inheritance with Golang
---
there is no inheritance in Golang because it does not support classes.

Inheritance is a mechanism that allows one class to inherit properties or behaviors from another class

you can mimic inheritance behavior

by using embedded struct

type Animal struct {
}

type Dog struct {
    Animal
}

---
Explain Go interfaces. What are they and how do they work?
---
Values of interface type can hold any value that implements those methods

For example, you could implement a geometry interface that defines 
that all shapes that use this interface must have an implementation of area() and perim().

type geometry interface {
    area() float64
    perim() float64
}

---
Print all permutations of a slice characters or string
---

---
Swap the values of two variables without a temporary variable
---
```
package main

import "fmt"

func main() {
   fmt.Println(swap())
}

func swap() []int {
      a, b := 15, 10
   b, a = a, b
   return []int{a, b}
}
```

---
Reverse the order of a slice
---
```
package main

import "fmt"

func reverse(sw []int) {
        for a, b := 0, len(sw)-1; a < b; a, b = a+1, b-1 {
                sw[a], sw[b] = sw[b], sw[a]
        } 
}

func main() { 
    x := []int{3, 2, 1} 
    reverse(x)
    fmt.Println(x)
}
```

---
Format a string without printing it
---
  s := fmt.Sprintf("Size: %d MB.", 85)
  fmt.Println(s)

---
Sum of Squares (concrruency)
---
```
package main
import "fmt"

func SumOfSquares(c, quit chan int) {
  y := 1
  for {
    select {
    case c <- (y*y):
      y++
    case <-quit:
      return
    }
  }
}

func main() {
  mychannel := make(chan int)
  quitchannel:= make(chan int)
  sum:= 0
  go func() {
    for i := 1; i <= 5; i++ {
      sum += <-mychannel
    }
    fmt.Println(sum)
    quitchannel <- 0
  }()
  SumOfSquares(mychannel, quitchannel)
}
```

---
What and How are pointers used in Go?
---
pointers are variables that store the memory address of another variable

They allow for efficient data manipulation and the ability to modify values without copying them

---
What is a workspace in go
---
a workspace consists of 3 directories 
src Contains Go source files organized in packages
pkg Contains compiled package objects
bin Contains executable binaries

---
What is the purpose of a GOPATH environment variable?
---
it is the path to the go workspace directory

---
how to convert golang array to dynamic array
---
```
// Declare a fixed-size array
arr := [5]int{1, 2, 3, 4, 5}

// Convert the array to a slice
dynamicArray := arr[:] // This creates a slice referencing the entire array

// Print the original array and the slice
fmt.Println("Array:", arr)
fmt.Println("Slice:", dynamicArray)

// You can now append to the slice
dynamicArray = append(dynamicArray, 6, 7) // Append new elements
fmt.Println("Updated Slice:", dynamicArray)
```

---
array vs slice 
---
array 
-  array have fixed size
- array are value types 
- array size cannot be changed  

slice 
- slice have dynamic size it can grow and shrink in size 
- slice are reference types
- golang provides builtin functions for slice ex
	1.	append(): Adds one or more elements to the end of a slice and returns the updated slice.
	2.	copy(): Copies elements from one slice to another, returning the number of elements copied.
	3.	len(): Returns the number of elements in a slice.
	4.	cap(): Returns the capacity of a slice, which is the total number of elements it can hold before needing to allocate more memory.
	5.	make(): Creates a new slice with a specified length and capacity.

```
slice := []int{1, 2, 3}
slice = append(slice, 4)     // append example
n := copy(slice, []int{5, 6}) // copy example
length := len(slice)         // len example
capacity := cap(slice)       // cap example
newSlice := make([]int, 5, 10) // make example
```
---
what is make
---
make is a built-in function that initializes and allocates memory for reference types


---
How does Go handle dependencies?
---
Go handles dependencies using modules, allowing developers to specify and manage package versions through 
a go.mod file, which defines module dependencies and ensures reproducible builds.

---
What makes Go compile quickly?
---
There are three main reasons for the compiler’s speed.

all imports must be explicitly listed at the beginning of each source file, 
so the compiler does not have to read and process an entire file to determine its dependencies

the dependencies of a package form a directed acyclic graph, and because there are no cycles, 
packages can be compiled separately and perhaps in parallel.

the object file for a compiled Go package records export information not just for the package itself, 
but for its dependencies too. When compiling a package, the compiler must read one object file 
for each import but need not look beyond these files.

syntax is very easy

Go imports dependencies once for all files, so the import time doesn't increase exponentially with project size.

Simpler linguistics means interpreting them takes less computing.

---
elaborated version
---
Why Go compiles faster:

- All imports are explicitly listed at the beginning of each source file, so the compiler does not have to read and process an entire file to determine its dependencies.
- The dependencies of a package form a directed acyclic graph, and because there are no cycles, packages can be compiled separately and perhaps in parallel.
- Changes only trigger recompilation of affected files and their dependencies.
- Go intentionally avoids complex language features like inheritance, and operator overloading. This reduces parsing complexity.
- Go's compiler leverages multiple cores for parallel compilation of independent files, maximizing CPU utilization.
- Unused import in Golang is an error and will only import the packages needed. It makes sure the compilation time is not increased due to unused packages.

---
Does Go support method overloading?
---
golang does not support method overloading

---
What is Method Overloading?
---
Method overloading refers to the ability to define multiple functions or methods with the same name but different parameters (different type, number, or both). The correct version to invoke is determined by the compiler based on the arguments passed.

---
Does Golang support method overridding?
---
Yes, Golang (Go) supports method overriding
Method overriding occurs when a type provides its own implementation of a method that is 
already defined in another type, typically a parent or embedded type. 
This allows the new method to replace or extend the behavior of the existing method.

method overriding is primarily achieved through type embedding and interfaces

---
what is select 
---
The select statement in Go is used to wait on multiple channel operations. 
It’s similar to a switch statement but specifically for channels. 
This allows your program to handle multiple channel communications concurrently 
and decide which one to proceed with based on which channel is ready.

Use select when you have multiple channel operations and you want your goroutine to respond to whichever channel is ready first. It’s especially useful for:
-	Multiplexing: Handling multiple channel inputs without blocking.
-	Timeouts: Implementing timeouts using time.After.
-	Non-blocking Operations: Performing operations without getting stuck waiting for a channel.

---
What is the Go Scheduler?
---
-	Definition: The Go scheduler is a component of Go’s runtime that manages the execution of goroutines. 
- It multiplexes thousands of goroutines onto a smaller number of OS threads, optimizing resource usage and performance.
- Purpose: Its primary role is to handle the creation, scheduling, and execution of goroutines, ensuring that they run efficiently without overwhelming the system’s resources.