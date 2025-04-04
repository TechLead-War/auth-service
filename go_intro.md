In Go, a repository typically refers to a package or a collection of packages that are stored in a version control system like Git. 
It is a way to organize and manage the codebase for a project.<br>
A repository can contain multiple packages, and each package can have its own set of Go files. <br>
The repository structure helps in organizing the code, managing dependencies, and versioning the project.

go.mod is a file that defines the module’s properties and dependencies in a Go project.

go.sum is a file automatically generated by Go’s module system (go mod). 
It keeps a record of the cryptographic hashes of all dependencies used in your project.

When we run commands like go mod tidy, go build then GO verify the integrity of downloaded 
modules using checksums in go.sum.