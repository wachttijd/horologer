package models

type Strongbox struct {
	GeneralId      string
	AvailableAfter []byte
	DecryptionKey  []byte
	Integrity      []byte
	Data           []byte
}
