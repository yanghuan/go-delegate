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

func f3(args ...interface{}) {
	println("f3")
}

func main() {
	d := Delegate{}
	d = d.Combine(f1)
	d = d.Combine(f2)
	d = d.Combine(f3)
	d.Invoke()
}
```

generics & one parameter
```
func f1(a int) {
	println("f1", a)
}

func f2(a int) {
	println("f2", a)
}

func f3(a int) {
	println("f3", a)
}

func main() {
	d := Action1[int]{}
	d = d.Combine(f1)
	d = d.Combine(f2)
	d = d.Combine(f3)
	d.Invoke()
}
```
