# golang-and-serverless-intro

`Basics:`

- Statically typed and compiled programming language

- Known for its simplicity, efficiency and strong support for concurrent programming

_Go originated from the mantra how can we enable individuals to write software easier - it provides one of the best standard libraries ._

## Spinning up a project

The start a new project, the command is `go mod init {mod name}`. E.g: `go mod init goBasics`.

It creates a `go.mod` file which is comparable to some extent to `package.json`. The most important bits are:

`module`: the module name. When it's meant to be a public module, the convention of the module name is `{where it is hosted}`/`{author / organisation}`/`{project name}`. Eg: `github.com/gabrielcedran/golang-and-serverless-intro`. **this is the project identity**

`go`: the go version it is based on

`require`: external dependencies

### Go files

The bare minimum components that go files must have are `package` and `func`.

**package and func main** every time a new package is created, these two pieces are mandatory - they act as the entry point for the entire application.

Running go: `go run main.go`

_by design, go applications can have multiple entry points (mains). It allows multiple microservices defined like in a list. The only requirement is that they are in separate directories. `go run {subPath}/main.go`_

### Variables

Explicit declaration:

```go
    var name string = "Mary"
```

Implicit declaration (type inferred):

```go
    name := "Mary"
    age := 20
    total := 10.0
```

**Zero value concept:**

Go has the concept of zero value, which means that declared but not assigned variables have default values (some data structure does equate to `nil`):

```go
    var myText string
    var myBool bool
    var myInt int

    fmt.Printf("My text %s, my bool %t my int %d\n", myText, myBool, myInt)
```

This will print `My text  , my bool false, my int 0`
