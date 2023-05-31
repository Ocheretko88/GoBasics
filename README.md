# Go Programming Portfolio

This repository contains examples of Go (Golang) programming concepts and techniques. Each script is self-contained and demonstrates a different aspect of Go programming.

## Scripts Included

1. main.go - Workers!
2. variables.go - Working with Variables
3. flow_control.go - Flow Control Statements
4. functions.go - Functions and Recursion
5. array_slices.go - Arrays and Slices
6. pointers.go - Pointers
7. structures.go - Structures
8. maps.go - Maps
9. json_parsing.go - JSON Parsing
10. error_handling.go - Error Handling
11. concurrency.go - Goroutines and Channels
12. http_server.go - HTTP Server
13. database.go - Interacting with a Database
14. reflection.go - Using Reflection
15. interfaces.go - Defining and Implementing Interfaces

## Running the Scripts

Before running the scripts, you need to install Go on your machine.

### Install Go

1. Download Go from the official website: https://golang.org/dl/.

2. After downloading, install Go. For Mac, open terminal and type:

    ```bash
    tar -C /usr/local -xzf goX.Y.Z.tar.gz

Replace goX.Y.Z with the version you downloaded.

3. Add the Go binary path to your PATH environment variable by adding the following line to your shell profile file (e.g., ~/.bash_profile or ~/.bashrc):

    ```bash
    export PATH=$PATH:/usr/local/go/bin
    ```
4. Restart your terminal or run source ~/.bash_profile or source ~/.bashrc, and then you can check the Go version by running:

    ```bash
    go version

For Linux, the installation process is similar. Refer to the official Go installation guide for detailed instructions: https://golang.org/doc/install#install.

### Run Go Files
Once Go is installed, you can run Go scripts using the go run command. Navigate to the directory containing your Go script and use the go run command followed by the filename:

    ```bash
    go run main.go

This command compiles and runs your Go program in one step.
