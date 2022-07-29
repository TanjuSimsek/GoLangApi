package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}
func NewCustomerRepositoryStub() CustomerRepositoryStub {

	customers := []Customer{

		{Id: "1001", Name: "TANJU", City: "Istanbul", Zipcode: "34500", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "1002", Name: "elif", City: "Istanbul", Zipcode: "34500", DateofBirth: "2000-01-01", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}

}
