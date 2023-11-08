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