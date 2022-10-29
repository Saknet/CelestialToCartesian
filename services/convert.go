package services

import (
	"math"

	model "github.com/saknet/celestialtocartesian/models"
)

func Convert(celestial model.Celstial) model.Cartesian {

	sign := 1

	if !celestial.Declination.Sign {
		sign = sign - 2
	}

	a := float64(celestial.Right_ascension.Hour)*15 + float64(celestial.Right_ascension.Minute)*0.25 + celestial.Right_ascension.Second*0.004166
	b := (math.Abs(float64(celestial.Declination.Degree)) + float64(celestial.Declination.Minute)/60 + (celestial.Declination.Second / 3600)) * float64(sign)

	cartesian := model.Cartesian{}
	cartesian.Name = celestial.Name
	cartesian.X = (celestial.Distance * math.Cos(a)) * math.Cos(b)
	cartesian.Y = (celestial.Distance * math.Cos(b)) * math.Sin(a)
	cartesian.Z = celestial.Distance * math.Sin(b)

	return cartesian

}

func AddTestObjects() []model.Celstial {
	nearestStars := make([]model.Celstial, 0)
	siriusA := model.Celstial{
		Name:            "Sirius A",
		Distance:        8.7094,
		Right_ascension: model.Right_ascension{Hour: 6, Minute: 45, Second: 8.917},
		Declination:     model.Declination{Degree: 16, Minute: 42, Second: 58.02, Sign: false},
	}

	nearestStars = append(nearestStars, siriusA)

	rigilKentaurus := model.Celstial{
		Name:            "Rigil Kentaurus",
		Distance:        4.37,
		Right_ascension: model.Right_ascension{Hour: 14, Minute: 39, Second: 36.49400},
		Declination:     model.Declination{Degree: 60, Minute: 50, Second: 2.3737, Sign: false},
	}

	nearestStars = append(nearestStars, rigilKentaurus)

	return nearestStars

}
