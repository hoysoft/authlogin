package authlogin

import (
	"fmt"
	"github.com/mavricknz/ldap"
	//"reflect"
)

type UserAuth struct {
	Login    bool
	HasThumb bool
	Account  string
	Mail     string
	Thumb    string
	Name     string
	Err      error
}

//login check, input account, domain, passwd, server, port and base_dn for search
func saAuthCheck(
	account,
	domain,
	passwd,
	ldap_server string,
	ldap_port uint16,
	base_dn string) UserAuth {

	user := UserAuth{Login: false, HasThumb: false}
	var err error

	//connect
	fmt.Println("Connecting")
	l := ldap.NewLDAPConnection(ldap_server, ldap_port)
	err = l.Connect()
	if err != nil {
		fmt.Printf("LDAP connectiong error: %v", err)
		user.Err = err
		return user
	}
	defer l.Close()

	//authentification (Bind)
	loginname := account + "@" + domain
	err = l.Bind(loginname, passwd)
	if err != nil {
		fmt.Printf("error: %v", err)
		user.Err = err
		return user
	}
	user.Login = true
	fmt.Print("Authenticated")

	//Search, Get entries and Save entry
	attributes := []string{}
	filter := fmt.Sprintf(
		"(&(objectclass=user)(samaccountname=%s))",
		account,
	)
	search_request := ldap.NewSimpleSearchRequest(
		base_dn,
		2, //ScopeWholeSubtree 2, ScopeSingleLevel 1, ScopeBaseObject 0 ??
		filter,
		attributes,
	)
	sr, _ := l.Search(search_request)
	user.Account = account
	user.Name = sr.Entries[0].GetAttributeValue("name")
	user.Mail = sr.Entries[0].GetAttributeValue("mail")
	user.Thumb = sr.Entries[0].GetAttributeValue("thumbnailphoto")
	if user.Thumb != "" {
		user.HasThumb = true
	}

	return user
}

//func main() {
//	SaAuthCheck("cn=hqj,dc=hz,dc=com","","","192.168.8.1",389,"dc=hz,dc=com")

//}
