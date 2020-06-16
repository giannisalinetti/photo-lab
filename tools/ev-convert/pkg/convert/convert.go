package convert

import "math"

//LuxToEv100 returns the EV100 value from a Lux input
func LuxToEv100(lux float64) float64 {
	return math.Log2(lux / 2.5)
}

//Ev100TLux returns the Lux value from a EV100 input
func Ev100ToLux(ev float64) float64 {
	return math.Pow(2, ev) * 2.5
}
