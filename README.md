# deptask
*Simple dependable async tasks with monitoring.*

# problem

Modern services in microservice architecture might have lots of services they fetch data from. 
Therefore, process of handling user request can be splited into many *tasks* like this:
<img width="900" alt="Screen Shot 2022-05-09 at 16 46 45" src="https://user-images.githubusercontent.com/32338211/167424814-eaf7f55a-4737-49fa-a244-ada39dd3b6dc.png">

Of course, in Go we have tools to handle concurrent tasks gracefully, like *goroutines*, *wait groups*, etc. 
However, they do not have any errors logging, traces collecting and other usefull stuff. 
Implementing it of every task manually results in lots of code duplication.

# solution

Introduce new *task* primitive to enable middlewares when writing concurrent workflows in Go. 

Useful middlewares examples:
- write task execution time metric  
- log error in task
- create span of a task.
