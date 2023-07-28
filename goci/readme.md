## Readme 

- we are going to the run pattern , where the main function only handles the command line arguments and the rest of the code is in the run function


## Handling error 

- Create the new error file 
- Error is the interface that has Error() string method 
- Define the new type of error


### Testing the tool

- Create a new directory under testdata and create a file with the working and non working file


## Defining the pipline 

- Create the custom type `struct` with the constructor on the step.go and create the execute method and replace the run function with the pipline

## Handling output from external program 

- Create a new file called extrastep and create the new struct with the constructor and create the execute method and replace the run function with the pipline
- Create the new interface

```go
type executer interface{
    exectue() (string,error)
}
```