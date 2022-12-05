# go-delegate
A delegate implementation of Go, similar to a [delegate in C#](https://learn.microsoft.com/en-us/dotnet/csharp/programming-guide/delegates/)

Example
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


