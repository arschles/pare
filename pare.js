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

