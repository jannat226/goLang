
//ques1
// package main
// import "fmt";
// func main() {
// 	var x int = 10
// 	var y float64 = 5.5
// 	result := x + y
// 	fmt.Println("Result", result)
// } 
//invalid operation: x + y (mismatched types int and float64)

//ques2
// package main
// import "fmt"
// func main() {
// var x int = 10
// var y float64 = 5.5
// result := float64(x) + y
// fmt.Println("Result:", result)
// }
//invalid operation: x + y (mismatched types int and float64)

//ques3
// package main
// import "fmt"
// func main() {
// 	for i := 0; i<5; i++
// 	fmt.Println(i)
// }
// syntax error: unexpected newline, expected { after for clause

//ques4
// package main
// import "fmt"
// func main() {
// 		for i := 0; i< 5; i++ {
// 		fmt.Println(i)
// 	}
// }

//ques5
// package main
// import "fmt"
// func main() {
// 	var x int
// 	for x < 5 {
// 		fmt.Println(x)
// 	}
// } 
//-> infinte 0
//ques6
//package main
// import "fmt"
// 	func main() {
// 	var x int
// 	for x < 5 {
// 		fmt.Println(x)
// 		x++
// 	}
// }

//ques7
// package main
// import "fmt"
// func main() {
// 	var i int
// 	for i := 0; i<3; i++ {

// 		fmt.Println(i)
// 	}
// 	fmt.Println(i)
// }
//0 1 2 0

//ques8 
// package main
// import "fmt"
// func main() {
// 	var i int;
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(i);
// 	};
// 	fmt.Println(i);
// }

//ques9
// package main
// import "fmt"
// func main() {
// 	var num int = 5
// 	if num % 2 == 0 {
// 		fmt.Println("Even")
// 	} else
// 	fmt.Println("Odd")
// }
//syntax error: else must be followed by if or statement block

//ques10
// package main;
// import "fmt";
// func main() {
// 	var num int = 5
// 	if num % 2 == 0 {

// 		fmt.Println("Even")
// 	} else {
// 		fmt.Println("Odd")
// 	}
// }

//ques 11
// package main
// import "fmt"
// func main() {
// 	var i int
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(i)
// 	}
// }

//ques 13
// package main
// import "fmt"
// func main() {
// 	const x int = 5
// 	x = 10
// 	fmt.Println(x)
// }

//ques 14
// package main
// import "fmt"
// func main() {
// 	var x int = 5
// 	x = 10
// 	fmt.Println(x)
// }

//ques 15
// package main
// import "fmt"
// func main() {
// 	var num int = 5
// 	if num < 10
// 		fmt.Println("Less than 10&quot;)
// 	else {
// 		fmt.Println(&quot;Greater than or equal to 10&quot;)
// 	}
// }

//ques 16 
// package main
// import "fmt"
// func main() {
// 	var num int = 5
// 		if num < 10
// 			fmt.Println("Less than 10")
// 		else {
// 			fmt.Println("Greater than or equal to 10")
// 		}
// }

//ques 17
package main
import "fmt"
func main() {
for i := 0; i < 3; i++ {
	if i % 2 == 0
		fmt.Println("Even")
	else
		fmt.Println("Odd")
	}
}

