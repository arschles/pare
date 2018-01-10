function build() {
    var exit = cmd("go", "build", "-o", "pare", ".")
    if (exit != 0) {
        return error(1, "build failed!")
    }
    return success()
}
addTarget("build", build)

function test() {
    var exit = cmd("go", "test", "./...")
    if (exit != 0) {
        return error(1, "tests failed")
    }
    return success
}
addTarget("test", test)
