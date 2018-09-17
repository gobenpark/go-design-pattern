package factory_method

import (
	"strings"
	"testing"
)

func TestFactoryMethodCash(t *testing.T) {

	payment, err := GetPaymentMethod(Cash)

	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}

	t.Log("LOG:", msg)

	payment, err = GetPaymentMethod(DebitCard)

	if err != nil {
		t.Error("A payment method of type 'DebitCard' must exist")
	}

	msg = payment.Pay(22.30)

	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}

	t.Log("LOG:", msg)

}
