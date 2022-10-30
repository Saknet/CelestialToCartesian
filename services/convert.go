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

func AddClusters() []model.Celstial {
	clusters := make([]model.Celstial, 0)
	virgo := model.Celstial{
		Name:            "Virgo",
		Distance:        16.4,
		Right_ascension: model.Right_ascension{Hour: 12, Minute: 30, Second: 49.42338},
		Declination:     model.Declination{Degree: 12, Minute: 23, Second: 28.0439, Sign: true},
	}

	clusters = append(clusters, virgo)

	fornax := model.Celstial{
		Name:            "Fornax",
		Distance:        20.23,
		Right_ascension: model.Right_ascension{Hour: 3, Minute: 38, Second: 30.00},
		Declination:     model.Declination{Degree: 35, Minute: 27, Second: 0.01, Sign: false},
	}

	clusters = append(clusters, fornax)

	antlia := model.Celstial{
		Name:            "Antila",
		Distance:        38.11,
		Right_ascension: model.Right_ascension{Hour: 10, Minute: 28, Second: 53.588},
		Declination:     model.Declination{Degree: 35, Minute: 36, Second: 19.98, Sign: false},
	}

	clusters = append(clusters, antlia)

	hydra := model.Celstial{
		Name:            "Hydra",
		Distance:        57,
		Right_ascension: model.Right_ascension{Hour: 10, Minute: 36, Second: 42.8},
		Declination:     model.Declination{Degree: 27, Minute: 31, Second: 42, Sign: false},
	}

	clusters = append(clusters, hydra)

	centaurus := model.Celstial{
		Name:            "Centaurus",
		Distance:        35,
		Right_ascension: model.Right_ascension{Hour: 12, Minute: 48, Second: 49.3},
		Declination:     model.Declination{Degree: 41, Minute: 18, Second: 40, Sign: false},
	}

	clusters = append(clusters, centaurus)

	norma := model.Celstial{
		Name:            "Norma",
		Distance:        67.8,
		Right_ascension: model.Right_ascension{Hour: 16, Minute: 15, Second: 32.8},
		Declination:     model.Declination{Degree: 60, Minute: 54, Second: 30, Sign: false},
	}

	clusters = append(clusters, norma)

	pavo_Indus := model.Celstial{
		Name:            "Pavo-Indus",
		Distance:        63.7731,
		Right_ascension: model.Right_ascension{Hour: 21, Minute: 7, Second: 52.2},
		Declination:     model.Declination{Degree: 47, Minute: 10, Second: 44, Sign: false},
	}

	clusters = append(clusters, pavo_Indus)

	coma := model.Celstial{
		Name:            "Coma",
		Distance:        94.43,
		Right_ascension: model.Right_ascension{Hour: 13, Minute: 0, Second: 8.1},
		Declination:     model.Declination{Degree: 27, Minute: 58, Second: 37, Sign: true},
	}

	clusters = append(clusters, coma)

	perseus_Pisces := model.Celstial{
		Name:            "Perseus Pisces",
		Distance:        68.2,
		Right_ascension: model.Right_ascension{Hour: 3, Minute: 19, Second: 48.1},
		Declination:     model.Declination{Degree: 41, Minute: 30, Second: 42, Sign: true},
	}

	clusters = append(clusters, perseus_Pisces)

	return clusters

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
