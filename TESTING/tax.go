package tax

import "time"

func CalculateTax(amout float64) float64 {
	if amout <= 0 {
		return 0
	}
	if amout >= 1000 && amout < 20000 {
		return 10
	}
	if amout >= 20000 {
		return 20.0
	}

	return 5.0
}

func CalculateTaxWithSleep(amout float64) float64 {
	time.Sleep(time.Millisecond)

	if amout == 0 {
		return 0
	}

	if amout >= 1000 {
		return 10.0
	}

	return 5.0
}
