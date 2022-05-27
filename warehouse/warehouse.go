package warehouse

import (
	"errors"
	"fmt"
)

var (
	// ErrNotEnoughStock is returned when there isn't sufficient stock in the warehouse.
	ErrNotEnoughStock = errors.New("warehouse: not enough stock")
)

type Warehouse[T comparable] struct {
	inventory map[T]int
}

func NewWarehouse[T comparable]() *Warehouse[T] {
	return &Warehouse[T]{
		inventory: make(map[T]int),
	}
}

func (w *Warehouse[T]) Add(t T, qty int) {
	w.inventory[t] += qty
}

func (w *Warehouse[T]) HasInventory(t T) bool {
	return w.inventory[t] > 0
}

func (w *Warehouse[T]) Remove(t T, qty int) error {
	tQty := w.inventory[t]
	if tQty < qty {
		return fmt.Errorf("%w: required %v but %v in stock", ErrNotEnoughStock, qty, tQty)
	}

	w.inventory[t] -= qty
	return nil
}
