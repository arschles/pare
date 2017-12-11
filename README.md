# Pare - a Modern Build Tool for Modern Software Development

[![CircleCI](https://circleci.com/gh/arschles/pare.svg?style=svg)](https://circleci.com/gh/arschles/pare)

Pare is a build tool for modern software develompent. It focuses on making it easy to run _all_ your development
tools when you need them, so you don't have to write hacky Makefiles or learn complicated new build systems.

You can configure Pare with a single one-line file, or you can build more advanced workflows.

# What does it do?

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

The `pare` CLI looks for a `pare.toml` file in the current working directory when it runs. This file holds the commands that `pare` should run.

Pare configuration files are simple. They are written in 
[TOML](https://github.com/toml-lang/toml) and usually only have one line.

Here's an example file:

```toml
[targets.MY_TARGET.commands.COMMAND_1]
exec = "echo command 1"
[targets.MY_TARGET.command.COMMAND_2]
exec = "echo command 2"
```

When you run `pare run `, all the commands you put in the list 
(between the `[` and `]`) will be executed in parallel, and `pare` will exit after 
they all finish. Please see the [config file reference](#config-file-reference) below
for details on the file format.

>If you don't know how to write TOML, don't worry - there's nothing more to the file than the above. Put each command in inside double-quotes (`"`), separate them by commas (`,`), and surround the entire list with brackets (`[` and `]`).

# Config File Reference

All config files should be named `pare.toml`. This section outlines all the possible values
that Pare supports in the config files. You'll need to be familiar with 
[TOML](https://github.com/toml-lang/toml) to understand this section.

## `targets`

This is a [TOML table](https://github.com/toml-lang/toml#table) that contains one or more
`Target` rows in it.

### `targets.TARGET_NAME`

This is a [TOML table](https://github.com/toml-lang/toml#table) that represents a Pare build
target. It contains one or more `Command` rows in it. In your config file, replace `TARGET_NAME`
with your real target name.

#### `targets.TARGET_NAME.commands`

This is a [TOML table](https://github.com/toml-lang/toml#table) that represents the commands
to go in the target.

##### `targets.TARGET_NAME.commands.COMMAND_NAME`

This is set of [key/value pairs](https://github.com/toml-lang/toml#keyvalue-pair) that
define what a single command in the target should do. The possible keys are:

- `exec`: a [TOML string](https://github.com/toml-lang/toml#string) that says what Pare
should run
    - This value is required
- `crash`: a [TOML boolean](https://github.com/toml-lang/toml#user-content-boolean) that says
whether Pare should crash if this command exits with a code other than `0` (i.e. a failure).
    - If the command does fail, then Pare will exit with a code of `1`. Other commands might still 
    get a chance to execute even if this command fails, so don't rely on this for flow control.
    - This value defaults to `false`
- `directory`: a [TOML string](https://github.com/toml-lang/toml#string) that tells
Pare what directory to run the command in. When the command runs, this will be the current
working directory of the command.
    - This value defaults to the current working directory that `pare` is executed in (i.e. `.`)
