package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (tUnit TemperatureUnit) String() string {
	if tUnit == Celsius {
		return "°C"
	}
	return "°F"
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (temp Temperature) String() string {
	return fmt.Sprintf("%d %s", temp.degree, temp.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (sUnit SpeedUnit) String() string {
	if sUnit == KmPerHour {
		return "km/h"
	}
	return "mph"
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (speed Speed) String() string {
	return fmt.Sprintf("%d %s", speed.magnitude, speed.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (metData MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", metData.location, metData.temperature, metData.windDirection, metData.windSpeed, metData.humidity)
}
