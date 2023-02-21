package models

type Strongbox struct {
	AvailableUntil string
	AvailableAfter string
	DecryptionKey  string
	Data           string
}
