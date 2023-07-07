## Organizing Your Code

1. Put the Business Logic in the different ​package and cli in the different package



todo
​ 	├── cmd
​ 	│   └── todo
​ 	│       ├── main.go
​ 	│       └── main_test.go
​ 	├── go.mod
​ 	├── todo.go
​ 	└── todo_test.go

Todo.go represents the code for todo packages as library 

```bash
​​mkdir​​ ​​-p​​ ​​cmd/todo
```
### Creating todo package

- Create two data structure one struct for the todo and one for the list of todos
- Attach the methods to the list of todos to add, complete , delete
- Create a methods which saves the files and get the files from the disk
- Write a test 

### Create a CLI 
- Add the error handling with fmt.Fprintln(os.Stderr, err)

### Write a test for the CLI
- Use the go build tool to compile the program into a binary file.
- Execute the binary file with different arguments and assert its correct behavior.

## Write the CLI 
- Use flag package to parse the command line arguments and switch statement to parse the command line 
- Use the flag.Usage function to display the usage of the program when the user provides invalid arguments.
- Improve the output with the stringer interface
- Add the env variable to the program

## Parsing the value from the stdin 

- Create a helper function which checks the length of arguments and if the length is greater than 0 then return the strings 

## Understanding 

- When developing command-line tools, it’s a good practice to use the standard error (STDERR) output instead of the standard output (STDOUT) to display error messages as the user can easily filter them out if they desire.

1. What is vardic function

Varadic function is a function which takes variable number of arguments

```go
func Add(nums ...int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}
Add(1, 2, 3...) // 6
```


```go 
	*l = append(ls[:i-1], ls[i:]...)
```
Append function takes the second paramter as vardiac

- ... is the ellipsis notation in Go, which is used to unpack a slice
