package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client	*sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllSql := "SELECT * FROM customers"

	rows, err := d.client.Query(findAllSql)

	if err != nil {
		log.Println("Error while quering customer table: " + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer

		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)

		if err != nil {
			log.Println("Error while scanning customers: " + err.Error())
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, error) {
	findCustomer := "SELECT * FROM customers WHERE customer_id = ?"

	row := d.client.QueryRow(findCustomer, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)

	if err != nil {
		log.Println("Error while scanning customer: " + err.Error())
		return nil, err
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:codecamp@tcp(localhost:3307)/banking")

	if err != nil {
		panic(err)
	}

	client.SetConnMaxIdleTime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetConnMaxIdleTime(10)

	return CustomerRepositoryDb{client}
}