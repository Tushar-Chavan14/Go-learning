package main

import (
	"fmt"
)

type Notifier interface {
	SendNotification() error
}

type Email struct {
	recipent string
	subject  string
	body     string
}

type pushNotification struct {
	deviceToken string
	message     string
}

type SMS struct {
	phNo    string
	message string
}

func (e Email) SendNotification() error {
	fmt.Println("sending message to", e.recipent, "with subject", e.subject, "and message", e.body)

	return nil
}

func (s SMS) SendNotification() error {
	fmt.Println("sending message to", s.phNo, "and message", s.message)
	return nil
}

func (p pushNotification) SendNotification() error {
	fmt.Println("sending message to device", p.deviceToken, "and message", p.message)
	return nil
}

func main() {

	email := Email{recipent: "sadanandgadwal@gmail.com", subject: "do some work", body: "hello reasearch about PWA"}
	sms := SMS{phNo: "+917854778415", message: "hey! whats up."}
	push := pushNotification{deviceToken: "dxcy1432", message: "hello from push notification"}

	sendNotification(email)
	sendNotification(sms)
	sendNotification(push)

}

func sendNotification(n Notifier) {
	n.SendNotification()
}
