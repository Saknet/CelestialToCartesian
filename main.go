package main

import (
	"fmt"

	service "github.com/saknet/celestialtocartesian/services"
)

func main() {

	nearestClusters := service.AddClusters()

	for i := 0; i < len(nearestClusters); i++ {

		cartesianStar := service.Convert(nearestClusters[i])
		fmt.Println("Object: ", cartesianStar)
		fmt.Println("Object name: ", cartesianStar.Name)
		fmt.Println("Object x: ", cartesianStar.X)
		fmt.Println("Object y: ", cartesianStar.Y)
		fmt.Println("Object z: ", cartesianStar.Z)

	}

}
