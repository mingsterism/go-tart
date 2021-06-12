package tart

type Natr struct {
	n   int64
	atr *Atr
}

func NewNatr(n int64) *Natr {
	return &Natr{
		n:   n,
		atr: NewAtr(n),
	}
}

func (a *Natr) Update(h, l, c float64) float64 {
	tr := a.atr.Update(h, l, c)
	if c == 0 {
		return 0
	}
	return tr / c * 100.0
}

func NatrArr(h, l, c []float64, n int64) []float64 {
	out := make([]float64, len(c))

	a := NewNatr(n)
	for i := 0; i < len(c); i++ {
		out[i] = a.Update(h[i], l[i], c[i])
	}

	return out
}