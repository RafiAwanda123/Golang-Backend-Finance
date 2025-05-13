package ml

// Prediksi dengan Regresi Linear
func PredictLinearRegression(sales []int) float64 {
	n := len(sales)
	if n < 2 {
		return float64(sales[0]) // Fallback ke Naive jika data < 2
	}

	var sumX, sumY, sumXY, sumX2 float64
	for i := 0; i < n; i++ {
		x := float64(i + 1) // Periode waktu (1, 2, 3, ...)
		y := float64(sales[i])
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
	}

	// Hitung slope (b) dan intercept (a)
	slope := (float64(n)*sumXY - sumX*sumY) / (float64(n)*sumX2 - sumX*sumX)
	intercept := (sumY - slope*sumX) / float64(n)

	// Prediksi periode berikutnya (n+1)
	nextPeriod := float64(n + 1)
	return intercept + slope*nextPeriod
}

// Prediksi dengan Pendekatan Naive (menggunakan rata-rata)
func PredictNaive(sales []int) float64 {
	if len(sales) == 0 {
		return 0
	}

	total := 0
	for _, s := range sales {
		total += s
	}
	return float64(total) / float64(len(sales))
}
