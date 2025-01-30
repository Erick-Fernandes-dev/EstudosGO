package main

import "fmt"

// Interface do processador
type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

// Implementações concretas
type MercadoPago struct{}

func (m *MercadoPago) ProcessPayment(amount float64) {
	println("Processing payment with MercadoPago", amount)
}

type PagarMe struct{}

func (p *PagarMe) ProcessPayment(amount float64) {
	println("Processing payment with PagarMe", amount)
}

// Interface do factory
type PaymentFactory interface {
	CreatePaymentProcessor() PaymentProcessor
}

// Implementações concretas da factory
type MercadoPagoFactory struct{}

func (m *MercadoPagoFactory) CreatePaymentProcessor() PaymentProcessor {
	return &MercadoPago{}
}

type PagarMeFactory struct{}

func (p *PagarMeFactory) CreatePaymentProcessor() PaymentProcessor {
	return &MercadoPago{}
}

// Responsável por criar o processador de pagamento
// Logo, não é necessário saber qual é o processador de pagamento
// apenas a factory
func ProcessPayment(factory PaymentFactory, amount float64) {
	processor := factory.CreatePaymentProcessor()
	processor.ProcessPayment(amount)
}

func main() {

	fmt.Println("Processing payment with PagarMe")
	pagarmeFactory := &PagarMeFactory{}
	ProcessPayment(pagarmeFactory, 100.00)

	println("---------------------------------------------------------")

	fmt.Println("Processing payment with MercadoPago")
	mercadoPagoFactory := &MercadoPagoFactory{}
	ProcessPayment(mercadoPagoFactory, 200.00)

}

// O que é Factory Method?
// É um padrão de projeto criacional que fornece uma interface para criar objetos em uma superclasse,
// mas permite que as subclasses alterem o tipo de objetos que serão criados.
