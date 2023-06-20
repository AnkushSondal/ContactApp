package user

import (
	"contactapp/contact"
	contactinfo "contactapp/contactInfo"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID               uuid.UUID
	FirstName        string
	LastName         string
	IsAdmin          bool
	userName         string
	usersCreatedByMe []*User
	MyContacts       []*contact.Contact
}

func NewAdmin(firstName, lastName, username string) *User {
	return &User{
		ID:        uuid.NewV4(),
		FirstName: firstName,
		LastName:  lastName,
		userName:  username,
		IsAdmin:   true,
	}
}

func (u *User) NewUser(firstName, lastName, username string) (*User, error) {
	if !u.IsAdmin {
		return nil, errors.New(u.FirstName + " you are not an admin")
	}

	_, isUserExisit := findUser(u.usersCreatedByMe, username)
	if isUserExisit {
		return nil, errors.New("username already exists")
	}

	newUser := &User{
		ID:        uuid.NewV4(),
		FirstName: firstName,
		LastName:  lastName,
		userName:  username,
		IsAdmin:   false,
	}
	u.usersCreatedByMe = append(u.usersCreatedByMe, newUser)
	return newUser, nil
}

func findUser(userCreated []*User, username string) (*User, bool) {
	for i := 0; i < len(userCreated); i++ {
		if userCreated[i].userName == username {
			return userCreated[i], true
		}
	}
	return nil, false
}

func findContact(contactCreated []*contact.Contact, cname string) (*contact.Contact, bool) {
	for i := 0; i < len(contactCreated); i++ {
		if contactCreated[i].ContactName == cname {

			return contactCreated[i], true
		}
	}
	return nil, false
}

func (u *User) ReadUsersForAdmin() ([]*User, error) {
	if !u.IsAdmin {
		return nil, errors.New(u.FirstName + " you are not an admin")
	}
	fmt.Println("users created by", u.FirstName)
	return u.usersCreatedByMe, nil
}

func (u *User) UpdateUsersAdmin(username, field string, value string) (*User, error) {
	if !u.IsAdmin {
		return nil, errors.New(u.FirstName + "is not authorized to create a user ")
	}
	userToUpdate, isUserExisit := findUser(u.usersCreatedByMe, username)
	if !isUserExisit {
		return nil, errors.New("username does not exists")
	}
	switch field {
	case "FirstName":
		userToUpdate.FirstName = value
		return userToUpdate, nil
	case "LastName":
		userToUpdate.LastName = value
		return userToUpdate, nil

	default:
		return nil, errors.New("field not present")

	}
}

func (u *User) DeleteUsersAdmin(username string) error {
	if !u.IsAdmin {
		return errors.New(u.FirstName + "is not authorized to create a user ")
	}
	_, isUserExisit := findUser(u.usersCreatedByMe, username)
	if !isUserExisit {
		return errors.New("username does not exists")
	}

	for i := 0; i < len(u.usersCreatedByMe); i++ {
		if u.usersCreatedByMe[i].userName == username {
			fmt.Println("user", u.usersCreatedByMe[i].FirstName, "is deleted")
			u.usersCreatedByMe = append(u.usersCreatedByMe[:i], u.usersCreatedByMe[i+1:]...)
			return nil
		}
	}
	return nil
}

// CRUD on Contact by User
func (u *User) CreateContact(contactName string) (*contact.Contact, error) {
	if u.IsAdmin {
		return nil, errors.New(u.FirstName + "is not authorized to create a contact ")
	}
	_, isContactExist := findContact(u.MyContacts, contactName)
	if isContactExist {
		return nil, errors.New("contact already exists")
	}
	newContact := contact.NewContact(contactName)
	u.MyContacts = append(u.MyContacts, newContact)
	return newContact, nil
}

func (u *User) ReadContact() (*contact.Contact, error) {
	if u.IsAdmin {
		return nil, errors.New(u.FirstName + "is not authorized to create a contact ")
	}
	fmt.Println("All Contacts")
	for i := 0; i < len(u.MyContacts); i++ {
		fmt.Println(u.MyContacts[i])
	}
	return nil, nil
}

func (u *User) UpdateContact(oldContactName, newContactname string) (*contact.Contact, error) {
	if u.IsAdmin {
		return nil, errors.New(u.FirstName + "is not authorized to update a contact ")
	}
	contactObj, isOldContactExist := findContact(u.MyContacts, oldContactName)
	if !isOldContactExist {
		return nil, errors.New(oldContactName + "contact does not exist")
	}

	_, isNewContactExist := findContact(u.MyContacts, newContactname)
	if isNewContactExist {
		return nil, errors.New(newContactname + "contact already exist")
	}
	contactObj.UpdateContact(newContactname)

	return contactObj, nil
}

func (u *User) DeleteContact(contactName string) error {
	if u.IsAdmin {
		return errors.New(u.FirstName + "is not authorized to delete a contact ")
	}
	_, isContactExist := findContact(u.MyContacts, contactName)
	if !isContactExist {
		return errors.New("contact does not exist")
	}

	for i := 0; i < len(u.usersCreatedByMe); i++ {
		if u.MyContacts[i].ContactName == contactName {
			fmt.Println("Contact", u.MyContacts[i], "is deleted")
			u.MyContacts = append(u.MyContacts[:i], u.MyContacts[i+1:]...)
			return nil
		}
	}
	return nil

}

// CRUD on ContactInfo

func (u *User) CreateContactInfo(contactname, ctype, cvalue string) (*contactinfo.ContactInfo, error) {
	if u.IsAdmin {
		return nil, errors.New(u.FirstName + "is not authorized to create a contact info ")
	}
	contactObj, isContactExist := findContact(u.MyContacts, contactname)
	if !isContactExist {
		return nil, errors.New("contact does not exist")
	}
	//fmt.Println("contact info for", contactObj.ContactName)
	return contactObj.CreateContactInfo(ctype, cvalue), nil
}

func (u *User) UpdateContactInfo(contactName, oldCType, newCtype, newCValue string) (*contact.Contact, error) {
	if u.IsAdmin {
		return nil, errors.New(u.FirstName + "is not authorized to update a contact ")
	}
	contactObj, isContactExist := findContact(u.MyContacts, contactName)
	if !isContactExist {
		return nil, errors.New(contactName + "contact does not exist")
	}
	_, err := contactObj.UpdateContactInfo(oldCType, newCtype, newCValue)
	if err != nil {
		return nil, errors.New("cant update info")
	}
	
	return contactObj, nil
}

func (u *User) DeleteContactInfo(contactname, ctype string) (*contact.Contact, error) {
	if u.IsAdmin {
		return nil, errors.New(u.FirstName + "is not authorized to update a contact ")
	}
	contactObj, isContactExist := findContact(u.MyContacts, contactname)
	//fmt.Println("delete contact info>>>>>>>>>>>", contactObj, isContactExist)
	if !isContactExist {
		return nil, errors.New(contactname + "contact does not exist")
	}

	fmt.Println(contactObj.DeleteContactInfo(ctype))
	// contactObj.MyContactInfo, _ = contactObj.DeleteContactInfo(ctype)
	err := contactObj.DeleteContactInfo(ctype)
	if err != nil {
		return nil, errors.New("cant delete info")
	}
	return contactObj, nil
}
