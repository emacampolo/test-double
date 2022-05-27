package warehouse_test

import (
	"testing"

	"test-double/warehouse"

	"github.com/stretchr/testify/require"
)

func TestWarehouseHasInventory(t *testing.T) {
	// Given
	w := warehouse.NewWarehouse[string]()

	// When
	w.Add("x", 1)

	// Then
	b := w.HasInventory("x")
	require.True(t, b)
}

func TestWarehouseRemove(t *testing.T) {
	// Given
	w := warehouse.NewWarehouse[string]()
	w.Add("x", 5)

	// When
	err := w.Remove("x", 2)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	require.True(t, w.HasInventory("x"))

	// When
	err = w.Remove("x", 3)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	require.False(t, w.HasInventory("x"))
}

func TestWarehouseRemoveErrNoStock(t *testing.T) {
	// Given
	w := warehouse.NewWarehouse[string]()
	w.Add("x", 5)

	// When
	err := w.Remove("x", 2)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	require.True(t, w.HasInventory("x"))

	// When
	err = w.Remove("x", 4)

	// Then
	require.ErrorIs(t, err, warehouse.ErrNotEnoughStock)
}
