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
1. You can declare variable using `var` or walrus operator `:=`
   ```go
   package main

   var normal string = "normal" //normal assignment

   func increment() int {
       number := 1 //walrus operator
       return number + 1
   }
   ```

## Testing

1. File name should be `xxx_test.go`
1. Test function must start with `Test`
1. Test function must take 1 argument only `t *testing.T`