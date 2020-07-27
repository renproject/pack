# `📦 pack`

[![GitHub](https://github.com/renproject/pack/workflows/test/badge.svg)](https://github.com/renproject/pack/workflows/test/badge.svg)
[![Coverage](https://coveralls.io/repos/github/renproject/pack/badge.svg?branch=master)](https://coveralls.io/github/renproject/pack?branch=master)
[![Report](https://goreportcard.com/badge/github.com/renproject/pack)](https://goreportcard.com/badge/github.com/renproject/pack)

[Documentation](https://godoc.org/github.com/renproject/pack)

In distributed systems, message passing often requires marshaling/unmarshaling messages in order to send them over the network. This is also useful for writing messages to logs, for persistence and debugging. The `📦 pack` library defines a generic application binary interface for serialising values without losing their type information. It supports:

- [x] `Bool`,
- [x] `U8`, `U16`, `U32`, `U16`, `U32`, `U64`, `U128`, `U256`,
- [x] `String`, `Bytes`, `Bytes32`, `Bytes65`,
- [x] `Struct`,
- [ ] `List`, and
- [x] custom types.

## Values

The primary interface for working with `📦 pack` is the `Value` interface. All values implement this interface, allowing them to be marshaling to binary and JSON, and to expose their type information.

```go
import (
    "fmt"

    "github.com/renproject/pack"
)

func main() {
    point := pack.NewStruct(
        "x", pack.NewU64(42),
        "y", pack.NewU64(100),
    )
    fmt.Printf("type: %v", point.Type())
}
```

## Types

By default, values do not marshal their types. This means that values cannot be unmarshalled, because there it is not always possible to know what type you are looking at without additional context. Consider the following examples:

```go
import (
    "fmt"

    "github.com/renproject/pack"
    "github.com/renproject/surge"
)

func main() {
    x := pack.NewU64(1)
    xData, _ := surge.ToBinary(x)

    y := pack.NewString("1")
    yData, _ := surge.ToBinary(x)

    fmt.Printf("x: %v", pack.Bytes(xData))
    fmt.Printf("y: %v", pack.Bytes(yData))
}
```

Looking only at the binary representation of the values, it is impossible to distingish between the types of these values when unmarshalling. This presents a common problem in distributed systems: how do my services tell each other about the type context? Well, with `📦 pack` we use the `Typed` value:

```go
import (
    "encoding/json"
    "fmt"

    "github.com/renproject/pack"
)

func main() {
    typed := pack.NewTyped(
        "x", pack.NewU64(1),
        "y", pack.NewString("1"),
    )
    fmt.Printf("type: %v", typed.Type())

    typedData, _ := json.MarshalIndent(typed, "", "  ")
    fmt.Printf("json: %v", string(typedData))
}
```

Now, we can see that the type information of our value has also been marshalled. In the case of JSON, the type information favours being verbose, so that it is easily debuggable by humans. However, the binary representation is much more compact. In practice, most services in distributed systems should use binary marshalling, unless they are in debug mode (binary marshalling is not only more compact, but it is also faster to marshal).

## Kinds

Types are not always simple. In the case of integers, there is minimal information that we need to know: what kind of integer is it? The only answers are `U8`, `U16`, `U32`, `U64`, `U128`, and `U256`. However, structs and lists are more complex data types and the same question has an infinite possible answers. This is where _kinds_ are useful. The kind of a value can be thought of as the "type of the type". We can understand this better with a few examples:

```go
import (
    "fmt"

    "github.com/renproject/pack"
)

func main() {
    x := pack.NewStruct("foo", pack.NewString("bar"))
    y := pack.NewStruct("bar", pack.NewBool(true))

    fmt.Printf("x type: %v", x.Type())
    fmt.Printf("y type: %v", y.Type())

    fmt.Printf("x kind: %v", x.Type().Kind())
    fmt.Printf("y kind: %v", y.Type().Kind())
}
```

We can see from this example that, although `x` and `y` have different _types_, they both have the same _kind_; they are both structs. It turns out that the existence of kinds is necessary when marshaling/unmarshaling type information, but you should very rarely need to explicitly use kinds.

## Custom Types

It is often convenient to use strongly-typed values at the language level (e.g. define/use custom Go structs). Using `📦 pack`, we can `Encode` and `Decode` to/from custom structs:

```go
import (
    "fmt"

    "github.com/renproject/pack"
)

type Foo struct {
    X pack.U64 `json:"x"`
    Y pack.U64 `json:"y"`
}

func main() {
    foo := Foo {
        X: pack.NewU64(1),
        Y: pack.NewU64(2),
    }
    bar := pack.NewStruct(
        "x", pack.NewU64(3),
        "y", pack.NewU64(4),
    )
    
    packed, err := pack.Encode(foo)
    if err != nil {
        panic(err)
    }

    fmt.Printf("foo type: %v", packed.Type())
    fmt.Printf("bar type: %v", bar.Type())
    
    if err := pack.Decode(&foo, bar); err != nil {
        panic(err)
    }
    
    fmt.Printf("foo.X: %v", foo.X)
    fmt.Printf("foo.Y: %v", foo.Y)
}
```

## Contribution

Built with ❤ by Ren.
