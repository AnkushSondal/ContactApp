package main

import (
	"contactapp/user"
	"fmt"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("recovered from", r)
		}
	}()

	admin1 := user.NewAdmin("Ankush", "Sondal", "AS")
	// user1, err := admin1.NewUser("Sanjeev", "Yadav", "SY")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(user1)

	// user2, err := admin1.NewUser("Adita", "Yadav", "AY")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(user2)

	// fmt.Println(admin1)

	// fmt.Println(admin1.ReadUsersForAdmin())
	// fmt.Println(admin1.UpdateUsersAdmin("SY", "FirstName", "Sanjiiiiiv"))
	// fmt.Println(admin1.ReadUsersForAdmin())
	// fmt.Println(user1)
	// fmt.Println(admin1.DeleteUsersAdmin("SY"))

	// fmt.Println(user1.CreateContact("ankush"))
	// // fmt.Println("After Updation :")
	// // fmt.Println(user1.UpdateContact("ankush","sondal"))
	// // fmt.Println(user1.ReadContact())

	// // fmt.Println("After Deletion :")
	// // fmt.Println(user1.DeleteContact("sondal"))
	// // // fmt.Println(user1.ReadContact())

	// fmt.Println(user1.CreateContactInfo("ankush", "mobile", "123456"))
	// yash, _ := admin1.NewUser("yash", "shah", "yashshah")
	// yash.CreateContact("Ankush")
	// yash.CreateContactInfo("Ankush", "Mobile", "99898989")
	// fmt.Println(yash.MyContacts[0].MyContactInfo[0])
	// // yash.UpdateContactInfo("Ankush","Mobile","type","Work")
	// yash.DeleteContactInfo("Ankush", "Mobile")
	// fmt.Println(yash.MyContacts[0].MyContactInfo[0])

	user3, err := admin1.NewUser("Ankit", "Singh", "AnkSin")
	if err != nil {
		panic(err)
	}
	fmt.Println("User Created : ", user3.FirstName)
	fmt.Println(user3)
	user3.CreateContact("Ankush")
	fmt.Println("contact Created : ", user3.MyContacts[0].ContactName)
	fmt.Println(user3.MyContacts[0])

	// user3.CreateContact("Sanjeev")
	// fmt.Println(user3.MyContacts)
	// fmt.Println("contact Created : ", user3.MyContacts[0].ContactName,",", user3.MyContacts[1].ContactName )

	user3.CreateContactInfo("Ankush", "Mobile", "123456")
	//fmt.Println(user3)
	fmt.Println("contact info Created : ", user3.MyContacts[0].MyContactInfo[0].ContactInfoType, user3.MyContacts[0].MyContactInfo[0].ConttactInfoValue)
	fmt.Println(user3.MyContacts[0].MyContactInfo[0])
	fmt.Println(user3.MyContacts[0])

	contactObj, _ := user3.UpdateContactInfo("Ankush", "Mobile", "Email", "ankush@ankush.com")

	// fmt.Println("contact info updated : ", user3.MyContacts[0].MyContactInfo[0].ContactInfoType, user3.MyContacts[0].MyContactInfo[0].ConttactInfoValue)
	// fmt.Println(user3.MyContacts[0].MyContactInfo[0])
	// fmt.Println(user3.MyContacts[0])
	fmt.Println("contactObj UPDATE:>>>>>>>>>>>>", contactObj.MyContactInfo[0])
	fmt.Println("user3.MyContacts", user3.MyContacts[0].MyContactInfo[0])

	contactObj2, err := user3.UpdateContact("Ankush", "Sondal")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("contactObj:>>>>>>>>>>>>", contactObj2)
	fmt.Println("user3.MyContacts", user3.MyContacts[0])

	user3.DeleteContactInfo("Sondal", "Email")
	//fmt.Println("contactDeleteObj>>>>>>>>>>>>>>>>>>>",contactDeleteObj)
	fmt.Println(user3.MyContacts[0])

}
