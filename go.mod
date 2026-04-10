// --- GO MODULE CONFIGURATION (go.mod) ---
// This file defines the root of your Go project and manages its dependencies.

// 1. module: The unique name/path of your project.
module learn-go

// 2. go: The minimum Go version required to build this project.
go 1.26.2


// OTHER DIRECTIVES YOU CAN GIVE:

// 3. require: Lists all external dependencies.
//    Example:
//    require (
//        github.com/gin-gonic/gin v1.9.1
//        golang.org/x/crypto v0.10.0
//    )

// 4. replace: Swaps a dependency with another version or a local path.
//    (Useful for testing local changes in a library).
//    Example:
//    replace github.com/my/project => ../local-path

// 5. exclude: Prevents specific versions of a module from being used.
//    Example:
//    exclude github.com/buggy/pkg v1.2.3

// 6. retract: Tells users that a specific version you released is broken.
//    Example:
//    retract v1.0.0 // Contains a critical security bug

// 7. toolchain: Specifies a specific Go toolchain version to use.
//    Example:
//    toolchain go1.22.1

