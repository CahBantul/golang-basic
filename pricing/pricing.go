package pricing

func calculateDiscount(basePrice float64) float64 {
	if basePrice > 1_000_000 {
		return basePrice * 0.2
	} else if basePrice >= 500_000 {
		return basePrice * 0.1
	}
	return 0
}

func CalculateFinalPrice(basePrice float64) float64 {
	discount := calculateDiscount(basePrice)

	return basePrice - discount
}
