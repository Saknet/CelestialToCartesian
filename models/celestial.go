package models

type Celstial struct {
	Name            string
	Distance        float64
	Right_ascension Right_ascension
	Declination     Declination
}

type Declination struct {
	Degree int64
	Minute int64
	Second float64
	Sign   bool
}

type Right_ascension struct {
	Hour   int64
	Minute int64
	Second float64
}
