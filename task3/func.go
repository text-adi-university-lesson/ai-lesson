package main

func stepFunction(s float64, activation float64) float64 {
	if s >= activation {
		return 1
	}
	return 0
}

func logicalAND(a, b float64) float64 {
	w1, w2 := 1, 1
	s := a*float64(w1) + b*float64(w2)
	t := 1.5
	return stepFunction(s, t)
}

func logicalOR(a, b float64) float64 {
	w1, w2 := 1, 1
	s := a*float64(w1) + b*float64(w2)
	t := 0.5
	return stepFunction(s, t)
}

func logicalNOT(a float64) float64 {
	w := -1.5
	t := float64(-1)
	s := a * w
	return stepFunction(s, t)
}

func logicalXOR(x1, x2 float64) float64 {
	w11, w21 := float64(1), float64(-1)
	s1 := x1*w11 + x2*w21
	y1 := stepFunction(s1, 0.5)

	w12, w22 := float64(-1), float64(1)
	s2 := x1*w12 + x2*w22
	y2 := stepFunction(s2, 0.5)

	w31, w32 := 1, 1
	s3 := y1*float64(w31) + y2*float64(w32)

	return stepFunction(s3, 0.5)

}
