package main

import "fmt"

type PaymentProcessor interface {
	Pay()
}

// O go já reconhece que o PagarMe implementa a interface PaymentProcessor
// por isso não precisamos declarar que o PagarMe implementa a interface PaymentProcessor
type PagarMe struct{}

func (p *PagarMe) Pay() {
	fmt.Println("Paying with PagarMe")
}

type MercadoPago struct{}

func (m *MercadoPago) Pay() {
	fmt.Println("Paying with MercadoPago")
}

func SimplePaymentFactory(provider string) PaymentProcessor {
	switch provider {
	case "pagarme":
		return &PagarMe{}
	case "mercadopago":
		return &MercadoPago{}
	default:
		return nil
	}
}

// ========================================================================================================

func main() {

	processor := SimplePaymentFactory("pagarme")
	processor.Pay()

	processor = SimplePaymentFactory("mercadopago")
	processor.Pay()
}

// O que é Simple Factory?
// É um padrão de projeto criacional que fornece uma interface para criar objetos em uma superclasse,
// mas permite que as subclasses alterem o tipo de objetos que serão criados.
