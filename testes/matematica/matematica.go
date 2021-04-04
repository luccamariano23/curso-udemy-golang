package matematica

func Media(numero ...float64) float64 {
	total := 0.0
	for _, num := range numero {
		total += num
	}
	media := total / float64(len(numero))
	return media
}
