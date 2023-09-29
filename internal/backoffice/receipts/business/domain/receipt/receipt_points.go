package receipt

type ReceiptPoints struct {
	value uint8
}

func NewReceiptPoints(pointsCalculator ...PointsCalculator) *ReceiptPoints {
	var accumulator uint8

	for _, v := range pointsCalculator {
		accumulator += v.CalculatePoints()
	}

	return &ReceiptPoints{
		value: accumulator,
	}
}

func (r *ReceiptPoints) Set(value uint8) {
	r.value = value
}

func (r ReceiptPoints) Value() uint8 {
	return r.value
}
