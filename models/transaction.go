package models

import (
	"finance-management-web/config"
	"store/db"
)

type Trasanction struct {
	Id                      int
	Name, Description, Date string
	Value                   float64
}

func CreateTransaction(name, description, date string, value float64) {
	db := config.DatabaseConect()
	createTransaction, err := db.Prepare("insert into transactions (tname, tdescription, tvalue, tdate) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	createTransaction.Exec(name, description, value, date)

	defer db.Close()
}

func SearchTransactions() []Trasanction {
	db := config.DatabaseConect()
	dbTransactions, err := db.Query("select * from transactions order by id asc")

	if err != nil {
		panic(err.Error())
	}

	t := Trasanction{}
	transactions := []Trasanction{}

	for dbTransactions.Next() {
		var id int
		var name, description, date string
		var value float64

		err := dbTransactions.Scan(&id, &name, &description, &value, &date)

		if err != nil {
			panic(err.Error())
		}

		t.Id = id
		t.Name = name
		t.Description = description
		t.Value = value
		t.Date = date

		transactions = append(transactions, t)
	}

	defer db.Close()
	return transactions

}

func DeleteTransaction(id string) {
	db := config.DatabaseConect()

	delete, err := db.Prepare("delete from transactions where id=$1")

	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}

func EditProduct(id string) Trasanction {
	db := db.ConectDatabase()

	transactionDatabase, err := db.Query("select * from transactions where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	transaction := Trasanction{}

	for transactionDatabase.Next() {
		var id int
		var name, description, date string
		var value float64

		err := transactionDatabase.Scan(&id, &name, &description, &value, &date)

		if err != nil {
			panic(err.Error())
		}

		transaction.Id = id
		transaction.Name = name
		transaction.Description = description
		transaction.Value = value
		transaction.Date = date

	}

	defer db.Close()
	return transaction

}

func UpdateTransaction(id int, name, description, date string, value float64) {
	db := db.ConectDatabase()

	update, err := db.Prepare("update transactions set tname=$1, tdescription=$2, tvalue=$3, tdate=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	update.Exec(name, description, value, date, id)

	defer db.Close()
}
