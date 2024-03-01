package main
import "fmt"
func main() {

	// Creating and initializing empty map
	// Using var keyword
	//var map_1 map[int]int

	// // Checking if the map is nil or not
	// if map_1 == nil {
	// 		fmt.Println("True")
	// } else {	
	// 	fmt.Println("False")
	// }

	// // Creating and initializing a map
	// // Using shorthand declaration and
	// // using map literals
	// map_2 := map[int]string{
	
	// 		90: "Dog",
	// 		91: "Cat",
	// 		92: "Cow",
	// 		93: "Bird",
	// 		94: "Rabbit",
	// }
	
	//fmt.Println("Map-2: ", map_2)

	var map2 map[int]int
	if map2 == nil{
		fmt.Println("Map is empty")
	}else{
		fmt.Println("Map is not empty")
	}
	map2 = map[int]int{
		1: 23,
		2: 47,
		3: 22,
		4: 06,
	}
	fmt.Println("Map2 is", map2)
	for _, value := range map2 {
		fmt.Println(value)
	}
}
