# golang-learning
A small repository to journal my learnings while experimenting with golang

## Learnings
1. Golang exports by capitalising first letter in a field/function name. e.g.,
    ```go
    package main
    
    type Rectangle struct {
        width float64 //not exported
        Height float64 //exported
    }
    
    func sum() { } //not exported
    
    func Area() { } //exported
    
    func main() {
        sum() //accessible within this go package, but not from other packages 
        Area() //accessible from other files
    }
    ```