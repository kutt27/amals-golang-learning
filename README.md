# amals-golang-learning
A repository I used to push the code that I used to learn Golang.

# Revision portion for learning Golang
---

## Some of the Important points

1. **Concurrency**: Go has built-in support for concurrency through goroutines and channels.
2. **Simplicity**: Go emphasizes simplicity and readability, with a clean and minimalistic syntax.
3. **Strongly Typed**: Go is a statically typed language, which means that variable types are checked at compile time. However, it also supports type inference, allowing you to omit type declarations in many cases.
4. **Garbage Collection**: Go features automatic garbage collection, which helps manage memory allocation and deallocation.
5. **Compiled Language**: Go is a compiled language. Therefore it's faster compared to some other languages.
6. **Cross-Platform**: Go supports cross-platform development.
7. **Standard Library**: Go comes with a rich standard library that provides support for common tasks such as I/O operations, networking, cryptography, and more. This reduces the need for third-party libraries for many tasks.
8. **Error Handling**: Go encourages explicit error handling through the use of multiple return values, where functions return both a result and an error value.
9. **Static Linking**: Go produces statically linked binaries by default, which means that all dependencies are included in the executable. This simplifies deployment, as you don't need to worry about distributing or managing external dependencies.

# Go Syntax

A Go file consists of the following parts:

- Package declaration
- Import packages
- Functions
- Statements and expressions

Example:

```go
package main
import "fmt"

func main(){
	fmt.Println("Hello World")
}
```

**Some of the points:**

- Ending of the statement can be ';' or by hitting enter. Compilation will remove ';' from code.
- Curly braces should also come in the first line and should not be on a new line. It will give error "_syntax error: unexpected semicolon or newline before {_"

=>
```go
func main
{
	fmt.Println("Hello World")
}
```

# GO Comments

Single-line comments start with two forward slashes (`//`).
Multi-line comments start with `/*` and ends with `*/`.

# GO Variables

### Creating Variables

**With var keyword:**

```go
var variablename type = value;
```

**With := sign:**

```go
variablename := value;
```

**Note:** In this case, the type of the variable is **inferred** from the value (means that the compiler decides the type of the variable, based on the value).

**Note:** It is not possible to declare a variable using `:=`, without assigning a value to it.

```go
package main  
import ("fmt")  
  
func main() {  
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred
```

In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type
### Difference Between var and :=
|var|:=|
|---|---|
|Can be used **inside** and **outside** of functions|Can only be used **inside** functions|
|Variable declaration and value assignment **can be done separately**|Variable declaration and value assignment **cannot be done separately** (must be done in the same line)|
### Go Multiple Variable Declaration

```go
package main  
import ("fmt")  
  
func main() {  
  var a, b, c, d int = 1, 3, 5, 7
}
```

# GO Constants

If a variable should have a fixed value that cannot be changed, you can use the `const` keyword.

The `const` keyword declares the variable as "constant", which means that it is **unchangeable and read-only**.

```go
const CONSTNAME type = value
```

**Note**: Constants can be declared both inside and outside of a function.

# GO Output Functions

Go has three functions to output text:

- `Print()`
- `Println()`
- `Printf()`

### The Print() Function

```go
package main  
import ("fmt")  
  
func main() {  
  var i,j string = "Hello","World"  
  
  fmt.Print(i)  
  fmt.Print(j)  
}
```

Output:

```text
HelloWorld
```

### The Println() Function

The `Println()` function is similar to `Print()` with the difference that a whitespace is added between the arguments, and a newline is added at the end.

### The Printf() Function

The `Printf()` function first formats its argument based on the given formatting verb and then prints them.

Here we will use two formatting verbs:

- `%v` is used to print the **value** of the arguments
- `%T` is used to print the **type** of the arguments

```go
package main  
import ("fmt")  
  
func main() {  
  var i string = "Hello"  
  var j int = 15  
  
  fmt.Printf("i has value: %v and type: %T\n", i, i)  
  fmt.Printf("j has value: %v and type: %T", j, j)  
}
```

Output:

```text
i has value: Hello and type: string  
j has value: 15 and type: int
```

# GO Arrays

### Declaration

```go
var array_name = [length]datatype{values} // length is defined

or 

var array_name = [...]datatype{values} // length is inferred

// also we could use
array_name := [length]datatype{values}

or 

array_name := [...]datatype{values}
```

Other things are the same like other languages, such as access, changing elements, etc.

### Initialize Only Specific Elements

```go
package main  
import ("fmt")  
  
func main(){
	variable1 = [3]int{1:10, 2:30};

	fmt.Print(variable1);
}

```

Output:
```text
0 10 30
```

# Slices in Go

Like arrays, slices are also used to store multiple values of the same type in a single variable.

However, unlike arrays, the length of a slice can grow and shrink as you see fit.

```go
package main  
import ("fmt")  
  
func main() {  
  myslice1 := []int{}  
  fmt.Println(len(myslice1))  
  fmt.Println(cap(myslice1))  
  fmt.Println(myslice1)  
  
  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}  
  fmt.Println(len(myslice2))  
  fmt.Println(cap(myslice2))  
  fmt.Println(myslice2)  
}
```

**Output**:
```text
0  
0  
[]  
4  
4  
[Go Slices Are Powerful]
```

### Creating a Slice from an Array

```go
var myarray = [length]datatype{values} // An array  
myslice := myarray[start:end] // A slice made from the array
```

Example:

```go
var arr1 = []string{"Go", "Slices", "Are", "Powerful"} 
var mySlice := arr1[2:4];
```

### Creating a Slice using `make()` Function

The `make()` function can also be used to create a slice.

```go
slice_name := make([]data_type, length, capacity)
```
Pending: Modification of Slice
# Operators in GO

### Arithmetic Operators

| Operator | Name | Description | Example |
| ---- | ---- | ---- | ---- |
| + | Addition | Adds together two values | x + y |
| - | Subtraction | Subtracts one value from another | x - y |
| * | Multiplication | Multiplies two values | x * y |
| / | Division | Divides one value by another | x / y |
| % | Modulus | Returns the division remainder | x % y |
| ++ | Increment | Increases the value of a variable by 1 | x++ |
| -- | Decrement | Decreases the value of a variable by 1 | x-- |
### Assignment Operators

![[Pasted image 20240222142236.png]]
Referred from: https://www.w3schools.com/go/go_assignment_operators.php

### Comparison Operators

![[Pasted image 20240222142346.png]]

# GO Conditions

- Less than `<`
- Less than or equal `<=`
- Greater than `>`
- Greater than or equal `>=`
- Equal to `==`
- Not equal to `!=`
- Logical AND `&&`
- Logical OR `||`
- Logical NOT `!`

### if-else Statement

Syntax:
```go
if condition{  
  // code to be executed if condition is true
}
```

**if-else**

```go
if condition{
	// code to be executed
} else {
	// else code to be executed
}
```

**Example**:

```go
package main  
import ("fmt")  
  
func main() {  
  x:= 20  
  y:= 18  
  if x > y {  
    fmt.Println("x is greater than y");
  } else {
	  fmt.Print("y is greater than x");
  }
}
```

#### Wrong Approach

```go
package main  
import ("fmt")  
  
func main() {  
  x:= 20  
  y:= 18  
  if x > y {  
    fmt.Println("x is greater than y");
  } 
  else {
	  fmt.Print("y is greater than x");
  }
}
```

```text
 syntax error: unexpected else, expecting }
```

Also their is else-if and nested if cases. Putting it for later review.

# GO Switch Statement

**Single case:**

```go
switch expression {
	case 1:
		// code expression
	case 2:
		//code expression
	default:
		//code expression
}
```

**Multi-case:**

```go
switch expression {
	case x, y:
		//code block to refer the flow
	case p, q:
		//code block to refer the flow
	default:
		//code block to refer the flow
}
```

# GO Loops

It's the same as C language. The only difference is that they don't contain the conditions inside the brackets.

**Syntax**:
```go
for expression1; expression2; expression3{
	// code for execution
}
```

There is no `while` statement in golang.

