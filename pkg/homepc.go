package pkg

type HomePC struct {
	CPU int
}

func (pc *HomePC) GetPrice() float64 {
	return 10.0 + float64(pc.CPU)
}
