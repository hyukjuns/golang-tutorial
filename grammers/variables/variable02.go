package variables

import "fmt"

func Var2() {

	var (
		name      string = "machine"
		height    int32
		weight    float32
		isRunning bool
	)

	height = 250
	weight = 350.56
	isRunning = true

	fmt.Println("name: ", name, "height: ", height, "weight: ", weight, "isRunning: ", isRunning)

}
