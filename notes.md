# How to start a new Go project

Steps:
1. Create a new directory and navigate into it:
    ```bash
    mkdir my-go-project
    cd my-go-project
    ```
2. Initialize a new Go module:
    ```bash
    go mod init .
    ```
   A `go.mod` file will be created in your project directory.
3. Create a new Go file (e.g., `main.go`) and add your code:
    ```go
    package main

    import "fmt"
    func main() {
         fmt.Println("Hello, World!")
    }
    ```
4. Run your Go program:
    ```bash
    go run .
    ```
   You should see the output: `Hello, World!`
5. To build your Go program into an executable, use:
    ```bash
    go build . -o my-go-project
    ```
   This will create an executable file named `my-go-project` in your directory. You can run it with:
    ```bash
    ./my-go-project
    ```
6. To manage dependencies, you can use `go get` to add new packages to your project. For example:
    ```bash
    go get github.com/some/package
    ```
   This will add the package to your `go.mod` file and download it to your project.
7. To tidy up your `go.mod` file and remove any unused dependencies, run:
    ```bash
    go mod tidy
    ```
    This will ensure that your `go.mod` and `go.sum` files are clean and only contain the necessary dependencies for your project.
