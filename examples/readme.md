## Context in golang

https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go

- Context allows to cancellation propagation , Let's try this to understand with the example, lets say you run a shop and someone orders a food and you send your employee to buy the ingredients and the customer cancels the order, now you have to cancel the order of the ingredients as well, so you have to propagate the cancellation to the employee as well, this is what context does, it propagates the cancellation to the goroutines as well.


## Dummy context 

- Use context.TODO function is the one way to create an empty context
- Use context.Background function is the other way to create an empty context


## Using Data withing context

