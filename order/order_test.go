package order_test

import (
	"testing"

	"test-double/order"
	"test-double/warehouse"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

/*
State verification tests, asserts against the warehouse's state.
*/

func TestOrderIsFilledIfEnoughInWarehouse(t *testing.T) {
	// Given
	w := warehouse.NewWarehouse[order.Product]()
	w.Add(order.Book, 50)
	o := order.Order{Product: order.Book, Quantity: 50}

	// When
	o.Fill(w)

	// Then
	require.True(t, o.IsFilled())
	require.False(t, w.HasInventory(order.Book))
}

func TestOrderIsDoesNotRemoveIfNotEnough(t *testing.T) {
	// Given
	w := warehouse.NewWarehouse[order.Product]()
	w.Add(order.Book, 50)
	o := order.Order{Product: order.Book, Quantity: 51}

	// When
	o.Fill(w)

	// Then
	require.False(t, o.IsFilled())
	require.True(t, w.HasInventory(order.Book))
}

/*
Behaviour verification tests, we check to see if the order made the correct calls on the warehouse.
*/

type WarehouseMock struct {
	mock.Mock
}

func (m *WarehouseMock) HasInventory(p order.Product) bool {
	args := m.Called(p)
	return args.Bool(0)
}

func (m *WarehouseMock) Remove(p order.Product, qty int) error {
	args := m.Called(p, qty)
	return args.Error(0)
}

func TestOrderIsFilledIfEnoughInWarehouse_(t *testing.T) {
	// Given
	w := WarehouseMock{}
	w.On("HasInventory", order.Book).Return(true)
	w.On("Remove", order.Book, 50).Return(nil)
	o := order.Order{Product: order.Book, Quantity: 50}

	// When
	o.Fill(&w)

	// Then
	require.True(t, o.IsFilled())
	w.AssertExpectations(t)
}

func TestOrderIsDoesNotRemoveIfNotEnough_(t *testing.T) {
	// Given
	w := WarehouseMock{}
	w.On("HasInventory", order.Book).Return(false)
	o := order.Order{Product: order.Book, Quantity: 51}

	// When
	o.Fill(&w)

	// Then
	require.False(t, o.IsFilled())
	w.AssertExpectations(t)
}
