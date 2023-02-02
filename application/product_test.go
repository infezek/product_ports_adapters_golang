package application_test

import (
	"testing"

	"github.com/infezek/product_ports_adapters_golang/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10
	err := product.Enable()
	t.Log(err)
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Error(t, err, "the price must be greater than zero to enable the product")
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 0
	err := product.Disable()
	t.Log(err)
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Error(t, err, "the price must be zero in order to have the product disabled")
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())

}
