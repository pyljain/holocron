package distance

import "math"

func CalculateCosineSimilarity(a []float64, b []float64) float64 {
	vectorLength := len(a)
	numerator := 0.0
	denominatorA := 0.0
	denominatorB := 0.0

	for i := 0; i < vectorLength; i++ {
		numerator += a[i] * b[i]
		denominatorA += math.Pow(a[i], 2)
		denominatorB += math.Pow(b[i], 2)
	}

	return numerator / (math.Sqrt(denominatorA) * math.Sqrt(denominatorB))
}
