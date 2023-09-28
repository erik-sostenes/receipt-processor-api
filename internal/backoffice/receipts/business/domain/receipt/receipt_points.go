package receipt

type ReceiptPoints struct {
	value uint8
}

func NewReceiptPoints(value uint8) (ReceiptPoints, error) {
	return ReceiptPoints{
		value: value,
	}, nil
}

func (r ReceiptPoints) Value() uint8 {
	return r.value
}
