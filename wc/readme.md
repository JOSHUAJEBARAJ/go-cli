The objective is to build the simple cli file 

- Learned the writing test and flag parsing and how to check for existence of the flag 

## Learned 
- Modules are group of packages and main act as the entry point for the application 


## Task 1 Building the basic CLI Tool 

- We are going to the use the buffio package to read the input from the console in the count function and return the number of line and it accepts the io.Reader interface

## Task 2 Testing the Basic word count 

- We are going to use the testing package and create the input to the count function using the bytes.NewBufferString and pass the string to the count function and check the output
- We can use the go test -v to run the test 

## Task 3 Adding the flag parsing
1. Create the l flag of type of l bool and add parse it to the flag.Parse() function
2. Now update the count function signature to accept the l bool flag and inside the function check if the flag is true then return the number of lines else return the number of words
3. Now update the test to pass the flag to the count function and check the output

## Compiling for the different OS
1. We can use the go build command to build the binary for the different OS
```bash
GOOS=linux go build -o wc
```