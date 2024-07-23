## Fan Out and Fan In 

- Fan out creating multiple go routines handle the incoming work 
- Fan in is the process of combining the results of multiple go routines

### Steps 

1. Define channels(One for task and one result channel) 
2. Fan out
3. Worker Function 
4. Fan in 
5. Sync and collect results 