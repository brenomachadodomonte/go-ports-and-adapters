package application

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := Product{}
	product.Name = "Smartphone"
	product.Status = DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := Product{}
	product.Name = "Smartphone"
	product.Status = ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Smartphone"
	product.Status = DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestProduct_GetId(t *testing.T) {
	product := Product{}

	id := uuid.NewV4().String()
	product.ID = id

	require.Equal(t, product.GetId(), id)
}

func TestProduct_GetName(t *testing.T) {
	product := Product{}

	name := "Smartphone"
	product.Name = name

	require.Equal(t, product.GetName(), name)
}

func TestProduct_GetPrice(t *testing.T) {
	product := Product{}

	price := 10.0
	product.Price = price

	require.Equal(t, product.GetPrice(), price)
}

func TestProduct_GetStatus(t *testing.T) {
	product := Product{}

	product.Status = DISABLED
	require.Equal(t, product.GetStatus(), DISABLED)

	product.Status = ENABLED
	require.Equal(t, product.GetStatus(), ENABLED)
}
