package math

import (
	"math"
)

func CosineSimilarity(a, b []float64) float64 {
	var magnitudeA, magnitudeB float64
	var dotProduct float64
	if len(a) != len(b) {
		panic("Vectors must have the same length")
	}

	// Calculate the dot product of the vectors.

	for i := range a {
		dotProduct += a[i] * b[i]
		magnitudeA += a[i] * a[i]
		magnitudeB += b[i] * b[i]
	}

	// Calculate the magnitude of each vector.

	// Calculate the cosine similarity.
	if magnitudeA == 0 || magnitudeB == 0 {
		return 0.0 // Handle cases with zero magnitude.
	}
	return dotProduct / (math.Sqrt(magnitudeA) * math.Sqrt(magnitudeB))
}
