package test2
import "fmt"
func Add(a,b int) int {
    return a+b
}
func sub(a,b int) int {
	return a-b
}
func multi(a,b int) int {
	return a*b
}
func divi(a,b int) int {
	if b != 0 {
		return a / b
	}
	fmt.Println("Error: Division by zero")
	return 0
}
func Arithomatic() int {
  var a,b,choice int; var value =0
  fmt.Print("Enter two number:-")
  fmt.Scanf("%d%d ",&a,&b)
  fmt.Printf("\n1.Addtion +")
  fmt.Printf("\n2.Substration -")
  fmt.Printf("\n3.Multiplication X")
  fmt.Printf("\n4.Division /\n")
  fmt.Printf("5._Exit /\n")
  fmt.Printf("Enter the choice:-")
  fmt.Scanf("%d ",&choice)

  switch choice {
	case 1:
	   value=Add(a,b) 
		
	case 2:
		 value=sub(a,b)
	case 3:
		value=multi(a,b)
	case 4:
		if b != 0 {
		 value=divi(a,b)
		}
	default:
		fmt.Println("Invalid choice. Please try again.")

  }
  return value
}