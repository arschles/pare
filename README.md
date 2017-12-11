# Pare - a Modern Build Tool for Modern Software Development

[![CircleCI](https://circleci.com/gh/arschles/pare.svg?style=svg)](https://circleci.com/gh/arschles/pare)

Pare is a build tool for modern software develompent. It focuses on making it easy to run _all_ your development
tools when you need them, so you don't have to write hacky Makefiles or learn complicated new build systems.

You can configure Pare with a single one-line file, or you can build more advanced workflows.

# What does this do?

Imagine you're building a website that needs a bunch of daemons running (microservices!). Normally
you'd open up a few terminal windows and run each one in a new window. But that's a pain!

Pare can run all of your daemons with a single command. Just tell it what to run in a simple 
[configuration file](#configuring) and run `pare`. The command takes over from there.

# Installing

The `pare` CLI is distributed as a self-contained binary. You don't have to download any
dependencies or run any special installer. Binaries are distributed for Linux (64 bit),
Windows (64 bit), and Mac OS X (64 bit). To install, download the appropriate CLI
for your system (see the links below) and put it in your executable path:

- [Linux (64 bit)](https://storage.googleapis.com/pare-cli/pare_linux_amd64)
- [Mac OS X (64 bit)](https://storage.googleapis.com/pare-cli/pare_darwin_amd64)
- [Windows (64 bit)](https://storage.googleapis.com/pare-cli/pare_windows_amd64.exe)

# Configuring

The `pare` CLI looks for a `pare.toml` file in the current working directory when it runs. This file
holds the commands that `pare` should run.

Pare configuration files are simple. They are written in [TOML](https://github.com/toml-lang/toml) and usually only 
have one line.

Here's an example file:

```toml
commands = ["echo command 1", "echo command 2"]
```

There's nothing more to it! All the commands in the list you give will be executed in parallel, and `pare`
will exit after they all finish.

>If you don't know how to write TOML, don't worry - there's nothing more to the file than the above. Put each command in inside double-quotes (`"`), separate them by commas (`,`), and surround the entire list with brackets (`[` and `]`).

