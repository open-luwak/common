package common

var NotifyEvent func(eventSource string, eventType string, eventData any) error
