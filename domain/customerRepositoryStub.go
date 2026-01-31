package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
		{
			Id: "101",
			Name: "Naruto",
			City: "Konoha Gakure",
			Zipcode: "12345",
			DateofBirth: "2000-01-01",
			Status: "1",
		},
		{
			Id: "102",
			Name: "Sasuke",
			City: "Konoha Gakure",
			Zipcode: "12345",
			DateofBirth: "2000-01-01",
			Status: "1",
		},
	}

	return CustomerRepositoryStub{customers}
}