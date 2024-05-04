package pkg

type BasePC struct {
	CPU int
}

func (pc *BasePC) GetPrice() float64 {
	return 10.0 * float64(pc.CPU)
}
