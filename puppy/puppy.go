package puppy

import (
	"fmt"

	"github.com/makikorp/animal/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof!, Woof!, Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func v11() {
	fmt.Println("this is version 1.1.0")
}
