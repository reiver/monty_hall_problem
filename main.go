package main


import (
	"fmt"
	"math/rand"
	"time"
)


const (
	maxInterations = 100000
)

func main() {

	// Initialize the pseudo-random number generator.
	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))


	// Without switching.
	numCorrect := 0
	for iterationNumber:=0; iterationNumber<maxInterations; iterationNumber++ {

		// Randomly put the prize behind one of the three doors.
		//
		// The doors are represented as a []bool of length 3.
		// So the 3 doors are:
		//
		// * doors[0]
		// * doors[1]
		// * doors[2]
		//
		// A value of false means that there is not a prize behind the
		// door. A value of true means that there is a prize behind
		// the door.
		//
		// So, for example, if we had:
		//
		//	[]bool{false, false, true}
		//
		// Then door #0 does NOT have a prize behind it.
		// Door #1 does NOT have a prize behind it.
		// Door #2 does have a prize behind it.
		//
		doors := []bool{false, false, false}
		doors[randomness.Intn(len(doors))] = true

		// Pick a door at random.
		index := randomness.Intn(len(doors))
		value := doors[index]

		// See if there was a prize behind the door.
		//
		// If there was, then we picked the correct
		// door, so note it.
		//
		// The 'numCorrect' variable stores the number
		// of times we correctly guessed the correct door.
		if value {
			numCorrect++
		}
	}
	percentCorrect := 100.0 * float64(numCorrect) / float64(maxInterations)
	fmt.Printf(" [Never Switch] %d correct out of %d. I.e., %f%% correct.\n", numCorrect, maxInterations, percentCorrect)


	// With switching.
	numCorrect = 0
	for iterationNumber:=0; iterationNumber<maxInterations; iterationNumber++ {

		// Randomly put the prize behind one of the three doors.
		//
		// The doors are represented as a []bool of length 3.
		// So the 3 doors are:
		//
		// * doors[0]
		// * doors[1]
		// * doors[2]
		//
		// A value of false means that there is not a prize behind the
		// door. A value of true means that there is a prize behind
		// the door.
		//
		// So, for example, if we had:
		//
		//	[]bool{false, false, true}
		//
		// Then door #0 does NOT have a prize behind it.
		// Door #1 does NOT have a prize behind it.
		// Door #2 does have a prize behind it.
		//
		doors := []bool{false, false, false}
		doors[randomness.Intn(len(doors))] = true

		// Pick a door at random.
		index := randomness.Intn(len(doors))
		value := doors[index]

		// To make things easier for what we are going to
		// do next, we remove the door we just picked.
		doors = append(doors[:index], doors[index+1:]...)

		// Switch.
		index2 := randomness.Intn(len(doors))
		value2 := doors[index2]
		doors = append(doors[:index2], doors[index2+1:]...)
		if false == value2 {
			value = doors[0]
		} else {
			value = value2
		}

		// See if there was a prize behind the door.
		//
		// If there was, then we picked the correct
		// door, so note it.
		//
		// The 'numCorrect' variable stores the number
		// of times we correctly guessed the correct door.
		if value {
			numCorrect++
		}
	}
	percentCorrect = 100.0 * float64(numCorrect) / float64(maxInterations)
	fmt.Printf("[Always Switch] %d correct out of %d. I.e., %f%% correct.\n", numCorrect, maxInterations, percentCorrect)
}
