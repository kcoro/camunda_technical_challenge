# CADMUNDA TECHNICAL CHALLENGE - IMPLEMENTED IN GO!
Author: Kevin Coronato<br>
Date: 04/15/2021
<br>

# About the Go Implementation
I have included an additional implementation of the coding exercise in Go to demonstrate the simplicity and beauty of the Go language,
 as well as the power of it's standard library.<br>
In fact this Go implementation has ZERO third party dependencies.  Go's standard library includes packages for JSON as well as a production ready Http server.<br>
Go projects have a simple project structure, as you can see this project contains a single source file: main.go, and a single module file: go.mod.<br>
I have provided a pre-built binary targeting Windows for users who do not have Go installed on their machine.<br>
In the event of third party dependencies there is no complicated package manager or build tools.<br>
Simply import the public url of the desired third party package in an import statement, go.mod takes care of the rest!

# Getting Started
Note: This program uses the Go runtime, to build and run the exe you will need Go installed.
- If you don't have Go installed, you can run the provided prebuilt windows binary if you are on a windows machine.

Build and Run:

```bash
go mod tidy
go run main.go
```

Run pre-built binary on Windows:
```bash
.\client_windows.exe
```
