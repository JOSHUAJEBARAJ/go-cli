## File Walk 

- root - the root directory to start the walk
- list - List files found by the tool, which is default operation 
- ext - Extension you want to search for 
- size - Min size 
## Developing the application 
- create the two files  main.go and actions.go which performs the all operations 
- Main.go has the main and run function 
- going to define the flags as the struct 
- create the actions.go has the filter out and list the files functionality 

## Task 2 Testing with the Table Driven Test
1. Create the test file with the name of actions_test.go 
2. Create the test cases as anonymous slice struct and use the t.Run to run the test cases
3. Create the testdata 

```bash
mkdir -p testdata/dir2 
echo "Just a test file" > testdata/dir.log
touch testdata/dir2/empty.sh
```

## Deleting a matched file 

- Add the del flag to the struct
- Add the delete function to the actions.go

## Testing with the Hep of Test Helper
- 


## Logging a deleted file 

- use stdout to provide the output 
- use the log package to log the deleted file
- Update the delete function with the additional parameter of the logger
- Add the logger to the struct as the type of io.Writer interface
- When the user provided the log flag then create the log file and pass it to the log.New function which takes the io.Writer interface