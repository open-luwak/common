package common

import (
	"errors"
)

var DefaultLicenseManager LicenseHandler

var DefaultEventManager EventHandler

type LicenseHandler interface {
	GetLicenseCiphertext(code string, version string) (string, error)
}

type EventHandler interface {
	SendNotification(eventSource string, eventType string, eventData any) error
}

func GetLicenseCiphertext(code string, version string) (string, error) {
	if DefaultLicenseManager == nil {
		return "", errors.New("licenseManager is nil")
	}
	return DefaultLicenseManager.GetLicenseCiphertext(code, version)
}

func SendNotification(eventSource string, eventType string, eventData any) error {
	if DefaultEventManager == nil {
		return errors.New("eventManager is nil")
	}
	return DefaultEventManager.SendNotification(eventSource, eventType, eventData)
}
