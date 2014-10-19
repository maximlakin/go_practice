#Go Notes

##Overview

Go is designed for public repo integration (ex Github).  Point $GOPATH to your projects

### Setup

For a project in /src/&lt;repo&gt;/&lt;user&gt;/my_project

* Compile/install: `$ go install`
* run with `$ $GOPATH/bin/my_project`
* Build libraries `$ go build`

Import library with:

```
import (
	…
	"<repo>/<user>/my_library"
)
```

## Golang

* It is strongly typed and garbage-collected
* Lower-case production names are used to identify lexical tokens. Non-terminals are in CamelCase.
* Comments `// line` and `/* section */`
* When two or more consecutive named function parameters share a type, you can omit the type from all but the last: `x, y int`
* variables, type is last: `var c, python, java bool`
* short declarations inside functions: `c, python, java := true, false, "no!"`
* basic types:

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

* The expression T(v) converts the value v to the type T: `i := 42 f := float64(i)`
* Constants: `const Pi = 3.14`
* Go has only one looping construct, the for loop: 

```
	// for loop
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    // while loop
    for sum < 1000 {
        sum += sum
    }
    // forever
    for {
    }
```

* Conditionals:

```
    if x < 0 {
        return sqrt(-x) + "i"
    }
    
    //Like for, the if statement can start with a short statement to execute before the condition:
    
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
    //Variables declared inside an if short statement are also available inside any of the else blocks
    }
```

* A struct is a collection of fields:

```
type Vertex struct {
    X int
    Y int
}
…
v := Vertex{1, 2}
fmt.Println(v.X)
```

* Arrays: `var a [2]string`
* Slices: `[]T` points to array, `s[lo:hi]` goes from `lo` to `hi`
* Making Slices: `a := make([]int, 5)  // len(a)=5`
* The range form of the for loop iterates over a slice or map.  To skip index: `for _, value := range pow {`
* Maps:

```
type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }
    fmt.Println(m["Bell Labs"])
}
// or map literal:
var m = map[string]Vertex{
    "Bell Labs": Vertex{
        40.68433, -74.39967,
    },
    "Google": Vertex{
        37.42202, -122.08408,
    },
}
```

* Function closure:

```
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}
//ex fibonacci:
func fibonacci() func() int {
}
func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
```
* Switch: A case body breaks automatically, unless it ends with a fallthrough statement:

```
 	switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        // freebsd, openbsd,
        // plan9, windows...
        fmt.Printf("%s.", os)
    }
```

* Go does not have classes. 

```
// methods
type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

func main() {
    f := MyFloat(-math.Sqrt2)
    fmt.Println(f.Abs())
}
```

* There are two reasons to use a pointer receiver. First, to avoid copying the value on each method call (more efficient if the value type is a large struct). Second, so that the method can modify the value that its receiver points to.
* An interface type is defined by a set of methods:

```
type Abser interface {
    Abs() float64
}
```
* Go routines: `go f(x,y,z)`
* Channels:

```
ch := make(chan int)
ch := make(chan int, 100) //buffered
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
v, ok := <-ch //test for closed channels
for i := range ch { //receive until channel closes
```

* Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.





