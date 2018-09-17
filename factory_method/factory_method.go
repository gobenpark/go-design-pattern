package factory_method

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:

		//reutnr new(DebitCardPM), nil

		// 크레딧 카드로 바꿈
		// 메시지 내용이 다르다고 테스트를 수정하면 안됨  테스트코드 커플링을 만들수 있기때문에
		return new(CreditCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}

type CreditCardPM struct{}

func (c *CreditCardPM) Pay(amount float32) string {

	//return fmt.Sprintf("%#0.2f paid using new credit card implementation\n", amount)
	//메시지의 내용을 바꿈
	return fmt.Sprintf("%#0.2f paid using debit card (new)\n", amount)
}
