its-structer
==============

its-structer is code generator to create matchers for structs.

```
$ go run github.com/youta-t/its/structer --help
Usage of .../structer:
structer is a matcher generator for any struct.
This is designed to be used as go:generate.

It generates a file with same name as a file having go:generate directive.
The new file, has "Matcher" and "Spec" types, is placed in "./gen_structer" directory (by default).

  -as-package
        handle -source as package path
  -dest string
        directory where new file to be created at (default "./gen_structer")
  -p    alias of -as-package
  -s value
        alias of -struct
  -source string
        recognise source as package path. If not set, use environmental variable GOFILE.
  -struct value
        Struct names to generate Matcher. Repeatable. By default, all structs are target.
```

Typical Usage
-------------

its-structer is designed to be used as

```go
//go:generate go run github.com/youta-t/its/structer
```


When you do `go generate ./...`, its-structer creates "Matcher" and "Spec" for each structs in the file.

For example, given code that

```go
//go:generate go run github.com/youta-t/its/structer -s MyStruct -dest gen
//go:generate gofmt -w ./gen/type.go
package example

import "time"

type MyStruct struct {
	Name      string
	Value     []int
	Timestamp time.Time
}
```

then, we get

- `type MyStructSpec`
- `func ItsMyStruct(spec MyStructSpec) its.Matcher[example.MyStruct]`

in `gen` package aside of the file invoking structer.

As the type shows you, they are used as

```go
gen.ItsMyStruct(
    gen.MyStructSpec{ ... }
).
    Match(example.MyStruct{ ... }).
    OrError(t)
```

`gen.MyStructSpec` has fields having same name as original's, but their types are wrapped by `its.Matcher`.

```go
type MyStructSpec struct {
	Name      its.Matcher[string]
	Value     its.Matcher[[]int]
	Timestamp its.Matcher[time.Time]
}
```

To use it, set mathcers as you like and pass `ItsMyStruct`.
`ItsMyStruct` creates a mathcer testing for each field by each matcher.
When all matchers are passed, `ItsMyStruct` is passed.

Advanced Usage
---------------

its-structer can create specs & matchers from package name.

Try that:

```go
//go:generate go run github.com/youta-t/its/structer -source k8s.io/api/core/v1 -as-package -s PodSpec
```

If you have `k8s.io/api/core/v1` in your modcache, its-structer find it and generates matchers for you.

`-s` flag lists struct name to create spec & matcher pair.
