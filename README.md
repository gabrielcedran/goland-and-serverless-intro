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

### Arrays and Slices

Arrays are declared with a fixed size and can only hold one type: `animals := [2]string{}`. It can also be declared and assigned at the same time: `animals2 := [2]string{"dog", "cat"}`.

Slices are like lists (or flexible arrays) that have dynamic size and can have value appended or removed. To create a slice, just omit the size of the array `animals := []string{}`

_removing elements_: from go 121, there is a helper package called slices:

```go
     import "slices"
     ...
     arr = slices.Delete(arr, 1, 2) // from and to index to remove
```

Prior 121:

```go
    arr = append(arr[0:1], arr[2:3]...) // basically append 2 arrays based on the original one removing the wanted element - this is what the slices helper does.
```

_`...` is the variadic function, which is a function that receives an arbitrary number of parameters (like vargs in java)_

### Loops

Go has only `for` loops. However it's possible to use it in a `do... while` fashion as well:

```go

    counter := 0
    for {
        fmt.Println(counter)
        counter++
    }

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for index, animal := range animals2 {
		fmt.Printf("My index %d, my animal %s\n", index, animal)
	}

	for index := range 10 {
		fmt.Println(index)
	}
```

### Structs

The simplest definition is a type that holds data and allows that data to be passed around applications. Along with interfaces, it's also Go's solution to OOP.

_go does not have a concept of classes and chooses composition over inheritance_

When instantiating the struct, Go doesn't force you to provide all the values due to the zero value approach.

#### Receivers

It's like class methods in other programming languages (Textbook definition: a method implemented on a type).

This is where the composition over inheritance comes from in Go as it's possible to compose structs with different structs.

### Pointers and References

Go passes variables and structs either as copy or reference. To pass as reference, simply prefix the parameter type with `*`.

In order to check the allocated memory address for any variable, just use ampersand `&varName`.

#### Declaring variables as pointers

To declare a variable as a pointer to the memory address, just use ampersand:

```go
    a = 1
    b := &a

    fmt.Println(b) // prints the memory address
```

To change the value of `b` (the value stored in that memory address), it's necessary to do what is called `dereferencing`:

```go

    *b = 10
    fmt.Println(*b)
    // here a is also 9 - b is pointing to the memory address that a is also pointing

```

### Interfaces

An interface is a collection of method signatures.

Example of declaration:

```go
type UserStore interface {
	DoesUserExist(username string) (bool, error)
	InsertUser(User types.RegisterUser) error
}
```

In Go, there is no need to have anything implementing that interface. Instead, it's just a matter of having structs satifying the interface's definition (with receivers) and then they can be used directly
(somewhat how typescript relies on structural types rather than nominal types).

## AWS and CDK (Cloud Development Kit)

CDK is a open-source framework built by AWS which allows developers to define infrastructure as code (avoiding having to go to aws dashboard / ui).

CDK uses CloudFormation service in order to provision the necessary infra. CloudFormation basically uses flat file templates to define what needs to be created (CDK Code -> Cloud Formation -> Rest of AWS suite).

### Initiating a CDK project

`cdk init app --language go`

`go get` to download the dependencies

**notes:**

1. jsii is a framework built by aws to transpile other programming languages into typescript - to allow them to communicate with typescript, which is the native language of aws (when creating a project in TS, it's not necessary)
2. `defer` in go means "execute this line of code within this function (scope) after everything else has finished"
3. `app` is the base component where all the other pieces of infrastructure is bound (where the app came from, where it was deployed to, etc) - it's applies a concept of `constructs`.
4. `stack` is a collection of individual infrastructure (AWS Lamnda, AWS DynamoDB, API Gateway, etc) - it's bound to the app
5. `env` is the env variable to configure the application

_apps can have multiple stacks and multiple resources_

#### Local development environment

It's hard to spin up a local development environment, however it's feasible. There are some tools meant to ease this process like [SST](https://v2.sst.dev/live-lambda-development) ([other resources](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-using-debugging.html))

### Lambdas

_Lambda is a serverless compute service that is triggered by an event (API Gateway, S3, etc). It automatically handles scaling and abstracts a lot of the complexity of using a server-based architecture._

To create a new lambda:

1. create a new directory and init it as a regular go module `go mod init {module_name}`
2. add the dependency to aws lambda lib for go: `go get github.com/aws/aws-lambda-go/lambda`
3. create the handler function and then pass it as the parameter to the `lambda.Start` function
4. in the cdk config, add the created lambda to the stack.

_for more details check the commit_

_to ease development, create a Makefile with the process to create the bootstrap executation and then zip it_

To deploy:

1. build the project with the command `GOOS=linux GOARCH=amd64 go build -o bootstrap` (this is a hard requirement from AL runtime "Amazon Linux") _the bootstrap artifact is actually an executable_
2. zip the binary generated app (bootstrap) `zip function.zip bootstrap` - function.zip is an arbitrary name, but it has to match the name defined in the stack resource.
3. `cdk diff`
4. `cdk deploy`

### DynamoDB

It's a fully managed NoSQL (key-value storage) database.

Set up:

1. add the dependency `go get github.com/aws/aws-sdk-go/service/dynamodb`
2. add the session dependency `go get github.com/aws/aws-sdk-go/aws/session`
3. create and add the necessary tables into the stack (refer to `user-registration.go`)
4. grant permission to the lambda function to access the tables (refer to `user-registration.go`)
5. create session and necessary logic (refer to `database.go`)

_Amazon has a concept of session when interacting with different infrastructure/service._

_To view the table's content, open amazon console -> DynamoDB -> Tables -> Look for the table -> Explore table items (top right corner)_

### API Gateway

Besides working as an entry point and routing, api gateway also has some security features like config to only accept requests from one region (or not from one region).

Set up:

1. ~add the dependency `github.com/aws/aws-sdk-go/service/apigateway`~ _apparently not necessary due to the cdk sdk having already been added?_
2. create a new api gateway resource and add it to the stac (refer to `user-registration.go`)
3. create an api lambda integration and assign the lambda (refer to `user-registration.go`)
4. create a new route, define details (path, method, etc) and assign the api lambda integration that we created in the previous step (refer to `user-registration.go`)
