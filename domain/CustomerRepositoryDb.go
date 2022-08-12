package domain

import (
	"GoLangApi/errs"
	"GoLangApi/logger"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll1() ([]Customer, error) {

	client, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	findAllSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {

		logger.Error("Error while querying customer table " + err.Error())
		return nil, err
	}
	customers := make([]Customer, 0)
	err = sqlx.StructScan(rows, &customers)
	if err != nil {

		logger.Error("Error while scaning customer table " + err.Error())
		return nil, err
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

// func NewCustomerRepositoryDb() CustomerRepositoryDb {

// 	client, err := sqlx.Open("mysql", "root:558133Mt*@tcp(localhost:3306)/banking")
// 	if err != nil {
// 		panic(err)
// 	}
// 	// See "Important settings" section.
// 	client.SetConnMaxLifetime(time.Minute * 3)
// 	client.SetMaxOpenConns(10)
// 	client.SetMaxIdleConns(10)

// 	return CustomerRepositoryDb{client: client}

// }
func NewCustomerRepositoryDb(dbClinet *sqlx.DB) CustomerRepositoryDb {

	return CustomerRepositoryDb{dbClinet}

}
func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {

	customerSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id=?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scaning customer table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

	}

	return &c, nil
}
