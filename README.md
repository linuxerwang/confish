# confish
A config parser in golang for Confish Configuration Format.

## Install

To update the generated cfg.go (optional):

```bash
$ go tool yacc -o cfg.go -p Cfg cfg.y
```

Go get confish:

```bash
$ go get github.com/linuxerwang/confish
```

## Parse confish file

```go
import "github.com/linuxerwang/confish"

type MyConf struct {
	...
}

conf := &MyConf{}
err := confish.ParseFile("path/to/myconf.conf", conf)
...
```

## Write confish file

```go
import "github.com/linuxerwang/confish"

type MyConf struct {
	...
}

conf := &MyConf{}
err := confish.WriteFile("path/to/myconf.conf", conf, "my-conf")
...
```

## The Confish Configuration Format

A confish file maps to a Go struct directly through the Go struct field tags.

In Go:
```go
type Book struct {
	Name string `cfg-attr:"name"`
	ISBN string `cfg-attr:"isbn"`
}
```

In config file:
```
book {
	name: "The way to Go"
	isbn: "123456789"
}
```

The confish parser uses the "cfg-attr" tags to locate the corresponding fields
in structs. Confish supports primitive types and collection types.

To specify a field of struct type:

```
	name {
		...
	}
```

To specify a field of slice of struct:

```
	name {
		...
	}

	name {
		...
	}

	name {
		...
	}
```

To specify a normal field:

```
	name: value
```

Primitive types: string, int, int32, int64, float32, float64, bool.
Collection types: slice of primitives, map of primitives.

Confish supports two types of slices: simple slice and struct slice.

Elements in simple slice must be primitive types.
Elements in struct slice must be struct types.

Keys in map must be int, int32, int64, or string.
Values in map must be int, int32, int64, bool, or string.

Comments start with "#" and stop at the end of line.

### Examples

*Example for embedded structs*

```go
type DisplayInfo struct {
	Model string `cfg-attr:"model"`
	Size  string `cfg-attr:"size"`
}

type KeyboardInfo struct {
	KeyCount int    `cfg-attr:"key-count"`
	Layout   string `cfg-attr:"key-layout"`
}

type Laptop struct {
	DI DisplayInfo  `cfg-attr:"display"`
	KI KeyboardInfo `cfg-attr:"keyboard"`
}
```

```
laptop {
	display {
		model: "Dell 1234"
		size: "1920x1080"
	}

	keyboard {
		key-count: 87
		key-layout: "compact"
	}
}
```

*Example for simple slice*

```go
type Router struct {
	Whitelist []string `cfg-attr:"whitelist"`
	Blacklist []string `cfg-attr:"blacklist"`
}
```

```
router {
	whitelist: [
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.twitter.com",
	]
	blacklist: [
		"http://www.test.com",
		"http://www.fishing.com",
	]
}
```

*Example for slice of structs*

```go
type BookShelf struct {
	Category string  `cfg-attr:"category"`
	Books    []*Book `cfg-attr:"book"`
}
```

```
bookshelf {
	category: "Computer Technology"

	book {
		name: "The way to Go"
		isbn: "123456789"
	}

	book {
		name: "The Go Programming Language"
		isbn: "987654321"
	}
}
```

*Example for maps*

```go
type PriceMap struct {
	Prices map[string]float32 `cfg-attr:"prices"`
}
```

```
price-map {
	prices: {
		"Red": 12.49,
		"Green": 10.99,
		"Blue": 13.99,
	}
}
```

*Example for deep struct nesting*

```go
type A struct {
	Name string `cfg-attr:"name"`
}

type B struct {
	ARef A `cfg-attr:"a"`
}

type C struct {
	BRef B `cfg-attr:"b"`
}

type D struct {
	CRef C `cfg-attr:"c"`
}

```

```
d {
	c {
		b {
			a {
				name: "John Doe"
			}
		}
	}
}
```
