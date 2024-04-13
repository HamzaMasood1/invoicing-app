package backend

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Name struct {
	ID        int
	FirstName string
	LastName  string
}

func AddName(db *sql.DB, newName Name) {

	stmt, _ := db.Prepare("INSERT INTO names (firstName, lastName) VALUES (?, ?)")
	stmt.Exec(newName.FirstName, newName.LastName)
	defer stmt.Close()

	fmt.Printf("Added %v %v \n", newName.FirstName, newName.LastName)
}

func GetAllNames(db *sql.DB) []Name {
	rows, err := db.Query("SELECT ID, firstName, lastName FROM names")
	if err != nil {
		fmt.Println("oh no")
	}
	defer rows.Close()

	var names []Name

	for rows.Next() {
		var name Name
		if err := rows.Scan(&name.ID, &name.FirstName, &name.LastName); err != nil {
			fmt.Println("oh no")
		}
		names = append(names, name)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("oh no")
	}

	return names
}

//func searchForPerson(db *sql.DB, searchString string) []person {
//
//	rows, err := db.Query("SELECT id, first_name, last_name, email, ip_address FROM people WHERE first_name like '%" + searchString + "%' OR last_name like '%" + searchString + "%'")
//
//	defer rows.Close()
//
//	err = rows.Err()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	people := make([]person, 0)
//
//	for rows.Next() {
//		ourPerson := person{}
//		err = rows.Scan(&ourPerson.id, &ourPerson.first_name, &ourPerson.last_name, &ourPerson.email, &ourPerson.ip_address)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		people = append(people, ourPerson)
//	}
//
//	err = rows.Err()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return people
//}
//
//
//func updatePerson(db *sql.DB, ourPerson person) int64 {
//
//	stmt, err := db.Prepare("UPDATE people set first_name = ?, last_name = ?, email = ?, ip_address = ? where id = ?")
//	checkErr(err)
//	defer stmt.Close()
//
//	res, err := stmt.Exec(ourPerson.first_name, ourPerson.last_name, ourPerson.email, ourPerson.ip_address, ourPerson.id)
//	checkErr(err)
//
//	affected, err := res.RowsAffected()
//	checkErr(err)
//
//	return affected
//}
//
//func deletePerson(db *sql.DB, idToDelete string) int64 {
//
//	stmt, err := db.Prepare("DELETE FROM people where id = ?")
//	checkErr(err)
//	defer stmt.Close()
//
//	res, err := stmt.Exec(idToDelete)
//	checkErr(err)
//
//	affected, err := res.RowsAffected()
//	checkErr(err)
//
//	return affected
//
//}
