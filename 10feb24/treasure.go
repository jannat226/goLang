// Scenario:
//A team of archaeologists exploring an ancient temple stumbles upon a series of sealed chambers, each containing valuable artifacts and clues to unlock the next chamber. These chambers are arranged in a linear fashion, with each chamber containing a unique artifact and a riddle or puzzle to solve. The team must use their wits and knowledge of ancient civilizations to navigate through the chambers and retrieve the ultimate treasure hidden within the final chamber.
//The team possesses an array named chambers where each element represents a chamber. Each chamber is represented by an object with the following properties:artifact: A description of the valuable artifact found within the chamber.puzzle: A riddle or puzzle presented to the team to unlock the next chamber.solved: A boolean value indicating whether the puzzle of the chamber has been solved or not.
//The team's task is to create a function exploreTemple(chambers) that takes in the array of chambers and iterates through them, solving each puzzle sequentially to progress to the next chamber. Once the team successfully solves the puzzle of the final chamber, the function should return the ultimate treasure hidden within.
//Create the function exploreTemple(chambers) and provide an example of how it would be used with a sample array of chambers.
package main

import (
	"fmt"
	
)
type chamber struct {
	Artifact,Puzzle string 
	Solved bool
}
// func exploreTemple(){
	
// }
func main() {
    var noOfChambers int
    var artifact, puzzle string
    var solved bool

    fmt.Println("Enter the number of chambers")
    fmt.Scanln(&noOfChambers)

    chambers := make([]chamber, noOfChambers)

    for i := 0; i < noOfChambers; i++ {
        fmt.Println("Enter the chamber details")
        fmt.Println("Artifacts")
        fmt.Scanln(&artifact)
        fmt.Println("Puzzle")
        fmt.Scanln(&puzzle)
        fmt.Println("Solved Status")
        fmt.Scanln(&solved)
        chambers[i] = chamber{Artifact: artifact, Puzzle: puzzle, Solved: solved}
    }

    
}