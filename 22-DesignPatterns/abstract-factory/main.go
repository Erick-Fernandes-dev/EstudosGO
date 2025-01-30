package main

// Abstract Factory

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

type NotificationService interface {
	SendNotification(message string)
}

// Interface da abstract factory
type BillingFactory interface {
	CreatePaymentProcessor() PaymentProcessor
	CreateNotificationService() NotificationService
}

// Implementação concreta da abstract factory para MercadoPago

type MercadoPagoProcessor struct{}

func (m *MercadoPagoProcessor) ProcessPayment(amount float64) {
	fmt.Printf("Processing payment with MercadoPago: %.2f \n", amount)
}

type MercadoPagoNotification struct{}

func (p *MercadoPagoNotification) SendNotification(message string) {
	fmt.Println("Sending notification with MercadoPago: ", message)
}

// Implementação concreta da abstract factory para PagarMe

type PagarMeProcessor struct{}

func (p *PagarMeProcessor) ProcessPayment(amount float64) {
	fmt.Printf("Processing payment with PagarMe: %.2f \n", amount)
}

type PagarMeNotification struct{}

func (p *PagarMeNotification) SendNotification(message string) {
	fmt.Println("Sending notification with PagarMe: ", message)
}

// Implementação concretas da abstract factory

type MercadoPagoFactory struct{}

func (m *MercadoPagoFactory) CreatePaymentProcessor() PaymentProcessor {
	return &MercadoPagoProcessor{}

}

func (m *MercadoPagoFactory) CreateNotificationService() NotificationService {
	return &MercadoPagoNotification{}
}

type PagarMeFactory struct{}

func (p *PagarMeFactory) CreatePaymentProcessor() PaymentProcessor {
	return &PagarMeProcessor{}
}

func (p *PagarMeFactory) CreateNotificationService() NotificationService {
	return &PagarMeNotification{}
}

func ProcessPayment(factory BillingFactory, amount float64, message string) {

	// Logo, não é necessário saber qual é o processador de pagamento e nem
	// o serviço de notificação, apenas implementamos a factory

	paymentProcessor := factory.CreatePaymentProcessor()
	paymentProcessor.ProcessPayment(amount)

	notificationService := factory.CreateNotificationService()
	notificationService.SendNotification(message)

}

func main() {
	fmt.Println("=== PagarMe ===")
	pagarMeFactory := &PagarMeFactory{}
	ProcessPayment(pagarMeFactory, 100.00, "Payment successful")

	fmt.Println("=== MercadoPago ===")
	mercadoPagoFactory := &MercadoPagoFactory{}
	ProcessPayment(mercadoPagoFactory, 200.00, "Payment successful")

}

// O que é Abstract Factory?
// É um padrão de projeto criacional que fornece uma interface para criar famílias de objetos relacionados ou dependentes sem especificar suas classes concretas.
// Isso permite que uma classe delegue a responsabilidade de instanciar objetos para subclasses.
// O objetivo é criar objetos sem especificar a sua classe concreta.

// Quando usar?
// Quando você precisa criar uma família de objetos relacionados ou dependentes sem especificar suas classes concretas.
// Quando você quer evitar acoplamento entre classes.
// Quando você quer expor uma interface para criar objetos sem especificar a sua classe concreta.
