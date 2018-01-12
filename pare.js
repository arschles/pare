function build() {
    var binaryOutput = "./pare"
    var exit = cmd("go", "build", "-o", binaryOutput, ".")
    if (exit != 0) {
        return error(1, "build failed!")
    }
    return success("binary has been output to " + binaryOutput)
}
addTarget("build", build)

function test() {
    var exit = cmd("go", "test", "./...")
    if (exit != 0) {
        return error(1, "tests failed")
    }
    return success("build succeeded")
}
addTarget("test", test)

function buildAndTest() {
    var grp = new Group(new Step("build", build), new Step("test", test))
    // run each step in the order they were defined in the group.
    // This is the equivalent to calling 'cmd' multiple times for each step.
    //
    // runSeq will return after all steps complete or there is a failure. It
    // will return a run result, which looks like this:
    //
    // {
    //     "build": { "message": "the build succeeded!", "runtime": 310 },
    //     "test": {"message": "the test failed!", "code": 1, "runtime": 10}
    // }
    //
    // the returned value is an object where each key is the name of a step.
    // the values in the top-level dictionary are dictionaries themselves. It
    // also has a few methods defined on it:
    //
    // - isFailed - true if one or more steps failed, false otherwise
    // - failures - a dictionary of just the failures
    //
    // If the step was a success, the dictionary will have the success message 
    // and the runtime of the step in milliseconds.
    //
    // If the step was a failure, the dictionary will have the failure message,
    // exit code, and the runtime of the step until it failed, in milliseconds
    //
    // you can inspect this value and do whatever else you like with it, but you
    // can also return it and pare will interpret it as follows:
    // 
    // - it will print the outputs of each step, in the order they were defined
    //   in the group
    // - if all steps passed, it will return with an exit code of 0
    // - if one or more steps failed, it will return with the exit code of the
    //   first failed step in the original group
    var runResult = group.runSeq()

    // start running all steps concurrently and return a run reference that you 
    // can use to keep track of what's going on. steps will really be running
    // concurrenrtly, not on the event loop. If you machine has multiple CPUs --
    // multi-core for example -- your steps will be running in parallel.
    //
    // because steps are running concurrently, logs would not be easy to
    // read if each step just piped all of its output to STDOUT and STDERR as
    // they do in runSeq(). Pare doesn't fix this problem completely, but it
    // does buffer outputs from each step so that logs aren't completely 
    // obfuscated. In a future version, Pare may support writing step outputs
    // to a file
    var runRef = group.runAll()

    // wait until all steps completed and return a run result. this run result
    // looks like this:
    // 
    // {
    //     "build": { "message": "the build succeeded!", "runtime": 310 },
    //     "test": {"message": "the test failed!", "code": 1, "runtime": 10}
    // }
    //
    // the returned value is a dictionary where each key is the name of a step.
    // the values in the top-level dictionary are dictionaries themselves.
    //
    // If the step was a success, the dictionary will have the success message 
    // and the runtime of the step in milliseconds.
    //
    // If the step was a failure, the dictionary will have the failure message,
    // exit code, and the runtime of the step until it failed, in milliseconds
    //
    // you can inspect this value and do whatever else you like with it, but you
    // can also return it and pare will interpret it as follows:
    // 
    // - it will print the outputs of each step, in the order they were defined
    //   in the group
    // - if all steps passed, it will return with an exit code of 0
    // - if one or more steps failed, it will return with the exit code of the
    //   first failed step in the original group
    var runResult = runRef.wait()
}
