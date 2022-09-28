package entity

import (
    "math/rand"
    "time"
)

// init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}


func RandomNum() int{
	return rand.Intn(10)
}

