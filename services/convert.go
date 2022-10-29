package services

import (
	"math"

	model "github.com/saknet/celestialtocartesian/models"
)

func Convert(celestial model.Celstial) model.Cartesian {

	sign := 1.0

	if !celestial.Declination.Sign {
		sign = sign - 2.0
	}

	a := float64(celestial.Right_ascension.Hour)*15 + float64(celestial.Right_ascension.Minute)*0.25 + celestial.Right_ascension.Second*0.004166
	b := (float64(celestial.Declination.Degree) + float64(celestial.Declination.Minute)/60 + (celestial.Declination.Second / 3600)) * sign

	cartesian := model.Cartesian{}
	cartesian.Name = celestial.Name
	cartesian.X = celestial.Distance * math.Cos(toRadians(b)) * math.Cos(toRadians(a))
	cartesian.Y = celestial.Distance * math.Cos(toRadians(b)) * math.Sin(toRadians(a))
	cartesian.Z = celestial.Distance * math.Sin(toRadians(b))

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

	procyon := model.Celstial{
		Name:            "Procyon",
		Distance:        11.402,
		Right_ascension: model.Right_ascension{Hour: 7, Minute: 39, Second: 18.11950},
		Declination:     model.Declination{Degree: 5, Minute: 13, Second: 29.9552, Sign: true},
	}

	nearestStars = append(nearestStars, procyon)

	tauCeti := model.Celstial{
		Name:            "Tau Ceti",
		Distance:        11.912,
		Right_ascension: model.Right_ascension{Hour: 1, Minute: 44, Second: 4.08338},
		Declination:     model.Declination{Degree: 15, Minute: 56, Second: 14.9262, Sign: false},
	}

	nearestStars = append(nearestStars, tauCeti)

	return nearestStars

}

func toRadians(angle float64) float64 {
	return angle * (math.Pi / 180)
}
