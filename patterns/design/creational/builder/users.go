package builder

import "fmt"

type Builder interface {
	SetName(string) Builder
	SetAge(int) Builder
	SetPhone(string) Builder
	SetAddress(string) Builder
	BuildUser() Userer
}

func GetNewBuilder(builderType string) Builder {
	if builderType == "user" {
		return newUserBuilder()
	}
	return nil
}

// -- Concrete Builder -- //
type userBuilder struct {
	name    string
	age     int
	phone   string
	address string
}

func newUserBuilder() Builder {
	uB := &userBuilder{}
	uB.fillDefaults()
	return uB
}

func (u *userBuilder) fillDefaults() {
	u.name = "nil"
	u.phone = "nil"
	u.address = "nil"
}

func (u *userBuilder) SetName(name string) Builder {
	u.name = name
	return u
}

func (u *userBuilder) SetAge(age int) Builder {
	u.age = age
	return u
}

func (u *userBuilder) SetPhone(phone string) Builder {
	u.phone = phone
	return u
}

func (u *userBuilder) SetAddress(address string) Builder {
	u.address = address
	return u
}

func (u *userBuilder) BuildUser() Userer {
	return &user{
		name:    u.name,
		age:     u.age,
		phone:   u.phone,
		address: u.address,
	}
}

// -- Abstract User -- //
type Userer interface {
	GetInfo()
}

// -- Concrete user -- //
type user struct {
	name    string
	age     int
	phone   string
	address string
}

func (u *user) GetInfo() {
	fmt.Println(u.name)
	fmt.Println(u.age)
	fmt.Println(u.phone)
	fmt.Println(u.address)
}
