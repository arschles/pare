# Pare

[![CircleCI](https://circleci.com/gh/arschles/pare.svg?style=svg)](https://circleci.com/gh/arschles/pare)

Pare is a build tool that focuses on getting your work done quickly and efficiently.

# Installing

The `pare` CLI is distributed as a self-contained binary. You don't have to download any
dependencies or run any special installer. Binaries are distributed for Linux (64 bit),
Windows (64 bit), and Mac OS X (64 bit). To install, download the appropriate CLI
for your system (see the links below) and put it in your executable path:

- [Linux (64 bit)](https://storage.googleapis.com/pare-cli/pare_linux_amd64)
- [Mac OS X (64 bit)](https://storage.googleapis.com/pare-cli/pare_darwin_amd64)
- [Windows (64 bit)](https://storage.googleapis.com/pare-cli/pare_windows_amd64.exe)

# Configuring

You configure Pare with a very simple [TOML](https://github.com/toml-lang/toml) file. It looks
like this:

```toml
commands = ["echo command 1", "echo command 2"]
```

The `pare` CLI looks for a `pare.toml` file in the current working directory when it runs.

If you don't know how to write TOML, don't worry - there's nothing more to the file than
the above. Put each command in inside double-quotes (`"`), separate them by commas (`,`), and
surround the entire list with brackets (`[` and `]`).

