package integers

// named return are pretty interesting. By having that, you just have a return statement at the end of the code.
// Also, this comment with become godoc :)
func Add(firstOperand, secondOperand int) (sum int) {
	sum = firstOperand + secondOperand
	return
}

