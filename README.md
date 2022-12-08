# go-delegate
A multicast delegate implementation of Go, similar to [delegate in C#](https://learn.microsoft.com/en-us/dotnet/csharp/programming-guide/delegates/)

## Example
no generics & variadic function
```
func f1(args ...interface{}) {
	println("f1")
}

func f2(args ...interface{}) {
	println("f2")
}

func main() {
	d := Delegate{}
	d = d.Combine(f1)
	d = d.Combine(f2)
	d.Invoke()
}
```

generics & parameters
```
func f() {
	println("f")
}

func f1(a int) {
	println("f1", a)
}

func f2(a, b int) {
	println("f2", a, b)
}

func main() {
    d := Action{}
    d = d.Combine(f)
    d.Invoke()    

    d1 := Action1[int]{}
    d1 = d.Combine(f1)
    d1.Invoke(1)
	
    d2 := Action2[int, int]{}
    d2 = d.Combine(f2)
    d2.Invoke(1, 2)
}
```

There is Action, Action1, ... Action7, that binds no parameters and up to 7 parameters.


