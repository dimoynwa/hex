package api

import (
	"hex/internal/ports"
)

type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DbPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{db: db, arith: arith}
}

func (adapter Adapter) GetAddition(a, b int32) (int32, error) {
	result, err := adapter.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = adapter.db.AddToHistory(result, "addition")

	if err != nil {
		return 0, err
	}

	return result, nil
}

func (adapter Adapter) GetSubtraction(a, b int32) (int32, error) {
	result, err := adapter.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = adapter.db.AddToHistory(result, "subtraction")

	if err != nil {
		return 0, err
	}

	return result, nil
}

func (adapter Adapter) GetMultiplication(a, b int32) (int32, error) {
	result, err := adapter.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = adapter.db.AddToHistory(result, "multiplication")

	if err != nil {
		return 0, err
	}

	return result, nil
}

func (adapter Adapter) GetDivision(a, b int32) (int32, error) {
	result, err := adapter.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = adapter.db.AddToHistory(result, "division")

	if err != nil {
		return 0, err
	}

	return result, nil
}
