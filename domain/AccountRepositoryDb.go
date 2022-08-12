package domain

import (
	"GoLangApi/errs"
	"GoLangApi/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {

	sqlInsert := "insert into accounts (customer_id ,opening_date ,account_type ,amount ,status )values(?,?,?,?,?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {

		logger.Error("Error while creating new account " + err.Error())

		return nil, errs.NewUnexpectedError("Unexpected error from database")

	}
	id, err := result.LastInsertId()
	if err != nil {

		logger.Error("Error while gettinglast insert id for new account " + err.Error())

		return nil, errs.NewUnexpectedError("Unexpected error from database")

	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil

}

// func NewAccountRepositoryDb(dbClinet *sqlx.DB) AccountRepositoryDb {

// 	client, err := sqlx.Open("mysql", "root:558133Mt*@tcp(localhost:3306)/banking")
// 	if err != nil {
// 		panic(err)
// 	}
// 	// See "Important settings" section.
// 	client.SetConnMaxLifetime(time.Minute * 3)
// 	client.SetMaxOpenConns(10)
// 	client.SetMaxIdleConns(10)

// 	return AccountRepositoryDb{client: client}

// }
func NewAccountRepositoryDb(dbClinet *sqlx.DB) AccountRepositoryDb {

	return AccountRepositoryDb{dbClinet}

}
