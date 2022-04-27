package main

import "fmt"

type InotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChanel() string
}

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notificaion via SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotoficaionSender{}
}

type SMSNotoficaionSender struct {
}

func (SMSNotoficaionSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotoficaionSender) GetSenderChanel() string {
	return "Twilio"
}

type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notificaion via Email")
}
func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}
func (EmailNotificationSender) GetSenderChanel() string {
	return "SES"
}

func getNotificationFactory(notificationType string) (InotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("No notification type")
}

func sendNotification(f InotificationFactory) {
	f.SendNotification()
}

func getMethod(f InotificationFactory) {
	println(f.GetSender().GetSenderMethod())
}
func getChanel(f InotificationFactory) {
	println(f.GetSender().GetSenderChanel())
}
func main() {

	smsFactory, _ := getNotificationFactory("SMS")
	emiailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emiailFactory)

	getMethod(smsFactory)
	getChanel(smsFactory)
	getMethod(emiailFactory)
	getChanel(emiailFactory)
}
