package contactinfo

import uuid "github.com/satori/go.uuid"

type ContactInfo struct {
	ID                uuid.UUID
	ContactInfoType   string
	ConttactInfoValue string
}

func NewContactInfo(contactType, contactValue string) *ContactInfo {
	return &ContactInfo{
		ID:                uuid.NewV4(),
		ContactInfoType:   contactType,
		ConttactInfoValue: contactValue,
	}
}

func (ci *ContactInfo) UpdateContactInfo(ctype, cvalue string) {
	ci.ContactInfoType = ctype
	ci.ConttactInfoValue = cvalue
	//return ci
}
