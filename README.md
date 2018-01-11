# Pare - a Modern Build Tool for Modern Software Development

[![CircleCI](https://circleci.com/gh/arschles/pare.svg?style=svg)](https://circleci.com/gh/arschles/pare)

_This project is alpha. Interface and features may change. This message will be removed
when the project reaches beta._

Pare is a build and workflow tool for modern software develompent. It focuses on making it easy 
to run _all_ your development tools when you need them, so you don't have to write hacky 
Makefiles or learn complicated new build systems.

Instead, you just write Javascript to define your workflow. 

# What is a workflow?

The Javascript you write isn't just plain Javascript. Pare provides powerful functions to use in 
your scripts that let you do things like this:

- Execute commands
- Return errors with custom messages and exit codes
- Execute commands concurrently (TODO)
- Execute commands in Docker containers (TODO)

All this functionality means that you can execute plain builds and tests
(like `go build` or `npm test`), but you can also do more complex things like:

- Run unit tests, and fail immediately if the tests failed
- Launch the development database
- Run the frontend webpack development server
- Run the backend API server
- Wait until all these services launched properly
- Run your integration tests against the frontend
- Exit 0 only if everything came up properly and the integration tests passed

# I'm interested, how do I install this thing?

The `pare` CLI is distributed as a self-contained binary. You don't have to download any
dependencies or run any special installer. Binaries are distributed for Linux (64 bit),
Windows (64 bit), and Mac OS X (64 bit). To install, download the appropriate CLI
for your system (see the links below) and put it in your executable path:

- [Linux (64 bit)](https://storage.googleapis.com/pare-cli/pare_linux_amd64)
- [Mac OS X (64 bit)](https://storage.googleapis.com/pare-cli/pare_darwin_amd64)
- [Windows (64 bit)](https://storage.googleapis.com/pare-cli/pare_windows_amd64.exe)

After you install the CLI, write your Javascript and run `pare run mytarget`

# Writing your build script

TODO