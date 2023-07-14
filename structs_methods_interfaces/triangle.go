package geometry

type Triangle struct {
	Base   float64
	Height float64
}

func (r Triangle) Area() float64 {
	return r.Base * r.Height * 0.5
}
