# Go Foundation

## Magesh Kuppan

## Schedule
- Commence      : 9:00 AM
- Tea Break     : 10:30 AM (20 mins)
- Lunch Break   : 12:30 PM (1 hr)
- Tea Break     : 3:00 PM (20 mins)
- Wind up       : 5:00 PM

## Methodology
- No powerpoint
- 100% code driven class

## Repo
- https://github.com/tkmagesh/cisco-go-nov-2023

## Software Requirements
- Go Tools (https://go.dev/dl)
- Visual Studio Code (https://code.visualstudio.com)
- Go Extension for VSCode (https://marketplace.visualstudio.com/items?itemName=golang.Go)

## Why Go?
- Simple Language
    - ONLY 25 keywords
    - No access modifiers (private/public/protected etc)
    - No classes (ONLY structs)
    - No inheritence (ONLY composition)
    - No reference types (Everything is a value)
    - No pointer arithmatic
    - No exceptions (ONLY errors)
    - No try..catch..finally construct
    - No implicit type conversion
- Close to hardware
    - Compiled to machine code
    - Extremely fast startup
    - Very lightweight (no need for VM, App Servers etc)
    - Tools for cross-compilation
- Concurrency
    - Builtin scheduler
    - Lightweight concurrency model (goroutine)
    - 1 goroutine = ~4KB (whereas 1 OS Thread = ~2MB)
    - Concurrency features are built in the language
        - go keyword, chan datatype, channel operator (<-), range, select-case
    - SDK support
        - sync package
        - sync/atomic packages

## Run a go file
- go run <filename.go>

## Create a build
- go build <filename.go>
- go build -o <bindary_name> <filename.go>

## Cross Compilation
- To get the list of platforms supports
    >go tool dist list
- To get the list of environment variables
    >go env
- To get the value of given environment variables
    >go env <var_name_1> <var_name_2> ....
    ex:
    >go env GOOS GOARCH
- To set the environment variables
    >go env -w <var_name_1>=<value_1> <var_name_2>=<value_2> ...
- To cross compile
    - Set the appropriate values for the GOOS & GOARCH env variables and create a build
    >go env -w GOOS=windows GOARCH=amd64
    >go build <file_name.go>
    - One can also combine the above steps
    >GOOS=windows GOARCH=amd64 go build <file_name.go>

## Data Types
- string
- bool
- integer types
    - int8
    - int16
    - int32
    - int64
    - int
- unsigned integer types
    - uint8
    - uint16
    - uint32
    - uint64
    - uint
- floating point types
    - float32
    - float64
- complex number types
    - complex64 (real [float32] + imaginary [float32])
    - complex128 (real [float64] + imaginary [float64])
- type aliases
    - byte (alias for uint8)
    - rune (alias for int32)(unicode code point)

## Variable Declarations
- using "var" keyword
- using ":=" 

## Variable Scope
- package scope
    - Cannot use ":="
    - Can have unused variables
- function scope
    - Can use both ":=" & "var"
    - Cannot have unused variables 

## Constants
- can have unused constants at both package & function scope

## Constructs
1. if else
2. for
3. switch case

## Functions
- Can have more than one return results (named result)
- Variadic function
- Anonymous function
- Higher Order functions
    - Assign functions as values to variables
    - Pass functions as arguments
    - Return functions as return values

## Errors
- Errors are just values
- Errors are not "thrown" but returned from functions
- Errors implement the "error" interface (prescribed)
- Errors can be created using the following:
    - errors.New()
    - fmt.Errorf()

## Panic & Recovery
**Panic** - the state of the application where the application execution cannot proceed further

## Collections
- Only 3 types (Array, Slice & Map)

### Array
- Fixed sized collection of data of the same type
- Memory allocation and initialization happens at the declaration itself

### Slice
- Varying sized collection of data of the same type
- We do not mention the size of the list in declaration
- Use **make** function to customize the memory allocation 
- len -> allowed # of values that can be accessed from the underlying array
- capacity -> overall memory allocated to the underlying array (initialized + uninitialized)
![image slices](./images/slices.png)

### Map
- collection of key/value pairs

## Modularity
Application is organized interms of "modules" & "packages"

### Module
- Any code that need to be versioned and deployed together (application)
- Typically it is a folder with the **go.mod** file
- go.mod
    - name of the module (preferrably the repo path)
    - the targetted go runtime version
    - dependencies
- How to create a module
    > go mod init <module_name>
- How to execute a module
    > go run .
- How to create a build
    > go build .
    > go build -o <binary_name> .

### Package
- Internal organization of code in a module
- Typically a folder
- All the code across the files in a folder are considered to be ONE entity
