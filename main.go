package main

import (
	"fmt"

	service "github.com/saknet/celestialtocartesian/services"
)

func main() {

	nearestStars := service.AddTestObjects()

	for i := 0; i < len(nearestStars); i++ {

		cartesianStar := service.Convert(nearestStars[i])
		fmt.Println("Object: ", cartesianStar)
		fmt.Println("Object name: ", cartesianStar.Name)
		fmt.Println("Object x: ", cartesianStar.X)
		fmt.Println("Object y: ", cartesianStar.Y)
		fmt.Println("Object z: ", cartesianStar.Z)

	}

}
