package contact

import (
	contactinfo "contactapp/contactInfo"
	"errors"
	// "fmt"
	
	uuid "github.com/satori/go.uuid"
)

type Contact struct {
	ID            uuid.UUID
	ContactName   string
	MyContactInfo []*contactinfo.ContactInfo
}

func NewContact(contactname string) *Contact {
	NewContact := &Contact{
		ID:          uuid.NewV4(),
		ContactName: contactname,
	}
	return NewContact
}

func findContactInfo(contactInfoCreated []*contactinfo.ContactInfo, ctype string) (*contactinfo.ContactInfo, bool) {
	for i := 0; i < len(contactInfoCreated); i++ {
		if contactInfoCreated[i].ContactInfoType == ctype {
			return contactInfoCreated[i], true
		}
	}
	return nil, false
}

func (c *Contact) UpdateContact(newContactname string) {
	c.ContactName = newContactname
	//return c
}

// for Contact Info

func (c *Contact) CreateContactInfo(ctype, cvalue string) *contactinfo.ContactInfo {
	newContactInfo := contactinfo.NewContactInfo(ctype, cvalue)
	c.MyContactInfo = append(c.MyContactInfo, newContactInfo)
	return newContactInfo
}

func (c *Contact) UpdateContactInfo(oldCType, newCType, newCValue string) (*contactinfo.ContactInfo, error) {
	contactInfoObj, isContactInfoExist := findContactInfo(c.MyContactInfo, oldCType)
	if !isContactInfoExist {
		return nil, errors.New("contactinfo type doesnot exists")
	}
	contactInfoObj.UpdateContactInfo(newCType, newCValue)
	//fmt.Println("SEE HERE >>>>>>>>>>>>>>>>>>>>>>>>>>>",contactInfoObj.ContactInfoType)
	return contactInfoObj, nil
}

func (c *Contact) DeleteContactInfo(ctype string) error {
	_, isContactInfoExist := findContactInfo(c.MyContactInfo, ctype)
	if !isContactInfoExist {
		return errors.New("contact info does not exist")
	}

	for i := 0; i < len(c.MyContactInfo); i++ {
		if c.MyContactInfo[i].ContactInfoType == ctype {
			//fmt.Println("Contact Info", c.MyContactInfo[i], "is deleted")
			c.MyContactInfo = append(c.MyContactInfo[:i], c.MyContactInfo[i+1:]...)
			return nil
		}
	}
	return nil
}
