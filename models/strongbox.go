package models

type Strongbox struct {
	GeneralId      string
	AvailableAfter []byte
	DecryptionKey  []byte
	Data           []byte
}
