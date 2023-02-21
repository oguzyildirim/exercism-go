package space

type Planet string

const earthRotation = 31557600

func Age(ageInSec float64, planet Planet) float64 {
	Planets := map[Planet]float64{
		"Earth":   earthRotation,
		"Mercury": earthRotation * 0.2408467,
		"Venus":   earthRotation * 0.61519726,
		"Mars":    earthRotation * 1.8808158,
		"Jupiter": earthRotation * 11.862615,
		"Saturn":  earthRotation * 29.447498,
		"Uranus":  earthRotation * 84.016846,
		"Neptune": earthRotation * 164.79132,
	}
	return calcAge(Planets[planet], ageInSec)
}

func calcAge(planetRotation float64, ageInSec float64) float64 {
	return ageInSec / planetRotation
}
