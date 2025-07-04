package utils

import (
	"math/rand"
	"time"
)

var secretNumber int
var random *rand.Rand // rand.New returns a pointer

// seed the random sequence and generate the first secret number
func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
	GenerateSecretNumber()
}

// generates a number between 1 to 5
func GenerateSecretNumber() {
	secretNumber = random.Intn(5) + 1
}

func GetSecretNumber() int {
	return secretNumber
}
