# Pare

[![CircleCI](https://circleci.com/gh/arschles/pare.svg?style=svg)](https://circleci.com/gh/arschles/pare)

_This project is alpha. Interface and features may change. This message will be removed
when the project reaches beta._

Pare is a build and workflow tool for modern software develompent. It focuses on making it easy 
to run _all_ your development tools when you need them, so you don't have to write hacky 
Makefiles or learn complicated new build systems.

Instead, you just write Javascript to define your workflow. 

# What's a workflow?

A workflow is all the things you need to develop your software, including:

- Build
- Test
- Run development servers
- Watch files in the background
- Run a release

Pare enables all of this functionality by providing powerful functions to use in your Javascript
build definition files. These functions let you do things like this:

- Execute commands
- Return errors with custom messages and exit codes
- Execute commands concurrently (TODO)
- Execute commands in Docker containers (TODO)

# Example Workflows

All this functionality means that you can execute plain builds and tests (like `go build` or 
`npm test`), but you can also build more complex workflows like this one:

- Run unit tests, and fail immediately if the tests failed
- Launch the development database
- Run the frontend webpack development server
- Run the backend API server
- Wait until all these services launched properly
- Run your integration tests against the frontend
- Exit 0 only if everything came up properly and the integration tests passed

# Installing

The `pare` CLI is distributed as a self-contained binary. It includes everything
you need - even the Javascript runtime for your build scripts.

Binaries are distributed for Linux (64 bit), Windows (64 bit), and Mac OS X 
(64 bit). To install, download the appropriate CLI for your system (see the 
links below) and put it in your executable path:

- [Linux (64 bit)](https://storage.googleapis.com/pare-cli/pare_linux_amd64)
- [Mac OS X (64 bit)](https://storage.googleapis.com/pare-cli/pare_darwin_amd64)
- [Windows (64 bit)](https://storage.googleapis.com/pare-cli/pare_windows_amd64.exe)

After you install the CLI, write your Javascript and run `pare run mytarget`

# Writing your build script

A build script is a collection of Javascript functions and calls to `addTarget`. Here is an
example build script that's adapted from [pare's own build script](./pare.js):

```javascript
function build() {
    var binaryOutput = "./pare"
    var exit = cmd("go", "build", "-o", binaryOutput, ".")
    if (exit != 0) {
        return error(1, "build failed!")
    }
    return success("binary has been output to " + binaryOutput)
}
addTarget("buildcli", build)
```

In this script, we define a `build` function, which calls `go build` to compile the `pare` binary.
After we define the function, we call `addTarget`, which exposes the function to the `pare run`
command. In this script, we tell Pare to expose a `buildcli` target that will run the `build`
function.

In other words, when `pare run buildcli` is called on the command line, the `build` function 
in the Javascript file will be called.

# Javascript Reference

Pare uses the [otto](https://github.com/robertkrimen/otto) Javascript interpreter to execute
Javascript. It's simple Javascript - think closer to ECMAScript 5 than 6 or 7.

On top of the standard Javascript, Pare adds a few extra functions. These are what you should
use to make your build script powerful. Here they are:

## `cmd`

This function runs a command on the host. It pipes `STDOUT` from the command to the host's
standard out, and similarly pipes `STDERR` from the command to the host's standard error.

Call this function by splitting your command into individual pieces:

```javascript
var exit = cmd("npm", "test")
```

The return value of `cmd` will be the exit code of the command.

## `error`

This function returns an error that you define. Call it by providing a numeric exit
code and a string description:

```javascript
return error(1, "the thing failed!")
```

You should always want to put `return` in front of this function because pare 
knows how to convert the return value of `error` into nicely formatted output.

## `success`

This function prints a custom message indicating a successful run
(Pare will print it in green text where it can) and causes your program to `exit 0`.

You should always put a `return` in front of this function because Pare knows
how to convert the return value of `success` into nicely formatted output.

## `addTarget`

This function tells Pare about a new target that it should expose on the command line.
Pass it a string name and a function:

```javascript
func build() {
    ...
}

addTarget("build", build)
```

In the above example, you'll be able to execute `pare run build` on the command line, and Pare
will execute the `build` function for you. You can name your targets anything you want,
even if they are not the same as the actual Javascript function.