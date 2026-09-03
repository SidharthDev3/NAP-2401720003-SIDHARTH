package utility

type Celsius float64

func (c Celsius) ToFarehheit() float64 {
	return float64(c)*1.8 + 32
}
