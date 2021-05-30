# go-curry

_"Currying is the process of transforming a function that takes multiple arguments in a tuple as its argument, into a function that takes just a single argument and returns another function which accepts further arguments, one by one, that the original function would receive in the rest of that tuple."_ 
[Haskell wiki](https://wiki.haskell.org/Currying)

Recently I was reading an interesting article about ["Functional programming in Go with generics"](https://ani.dev/2021/05/25/functional-programming-in-go-with-generics/) and it mentions that Go not support currying functions, which is true, so I decided to try to implement it with my small experience in this language. 
It was quite fun to do it, and it works!! Is it fluent go code tho? Obviously not, but I'm accepting suggestions of how to improve it.

Differently from the definition, wrapped functions can accept one or more arguments.

## Examples:

**Curry a function**
```go
func crossMultiply(a, b, c float64) (float64, error) {
	if a == 0 {
		return 0, errors.New("can not divide by zero")
	}
	return (b * c) / a, nil
}

// ...
var a, b, c float64 = 100, 420, 10
cross, _ := curry.Func(crossMultiply)
partial, _ := cross(a)
// when all arguments are passed the list of output values is returned
_, out := partial(b, c)

result := out[0].Float()
err := out[1]
if !err.IsNil() {
	log.Fatal(err.Interface().(error))
}

fmt.Printf("result: %f\n", result)
```

**Curry a method**
```go
type AnyStruct struct {
	
}
func (s AnyStruct) Sum(a, b int) int {
	return a + b
}

// ...
s := AnyStruct{}
sum, _ := curry.Method(s, "Sum")
partial, _ := sum(20)
_, out := partial(22)

result := out[0].Int()

fmt.Printf("result: %d\n", result)
```