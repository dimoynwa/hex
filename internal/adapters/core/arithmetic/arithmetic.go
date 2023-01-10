package arithmetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (ad Adapter) Addition(a, b int32) (int32, error) {
	return a + b, nil
}

func (ad Adapter) Subtraction(a, b int32) (int32, error) {
	return a - b, nil
}

func (ad Adapter) Multiplication(a, b int32) (int32, error) {
	return a * b, nil
}

func (ad Adapter) Division(a, b int32) (int32, error) {
	return a / b, nil
}
