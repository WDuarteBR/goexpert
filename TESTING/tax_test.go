package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amout := 500.0
	expected := 5.0

	result := CalculateTax(amout)

	if expected != result {
		t.Errorf("Expected: %f, but got: %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amout, expected float64
	}

	table := []calcTax{
		{5000.0, 10.0},
		{800.0, 5.0},
		{8000.0, 10.0},
		{0, 0},
	}

	for _, item := range table {
		result := CalculateTax(item.amout)

		if item.expected != result {
			t.Errorf("Expected: %f, but got: %f", item.expected, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTaxWithSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTaxWithSleep(500.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}

	for _, amout := range seed {
		f.Add(amout)
	}

	f.Fuzz(func(t *testing.T, amout float64) {
		result := CalculateTax(amout)
		if amout <= 0 && result != 0 {
			t.Errorf("Experado : 0, mas recebido %f", result)
		}

		if amout > 20000 && result != 20 {
			t.Errorf("Experado : 20, mas recebido %f", result)
		}
	})

}
