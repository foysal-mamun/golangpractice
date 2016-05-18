package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

type Countries struct {
	iso            string `PK`
	name           string
	printable_name string
	iso3           string
	numcode        int
}

func LearnORM() {
	db, err := sql.Open("mysql", "root:dpde4t@/test?charset=utf8")
	checkErr(err)
	defer db.Close()

	orm := beedb.New(db)

	var saveone Countries
	saveone.iso = "nn"
	saveone.name = "Numa"
	saveone.printable_name = "Faezah Numa"
	saveone.iso3 = "NM"
	saveone.numcode = 1222

	orm.Save(&saveone)
}

func LearnBasic() {

	db, err := sql.Open("mysql", "root:dpde4t@/test?charset=utf8")
	checkErr(err)
	defer db.Close()

	// stmt, err := db.Prepare("INSERT bf_countries SET iso=?, name=?, printable_name=?, iso3=?, numcode=?")
	// checkErr(err)

	// res, err := stmt.Exec("fn", "Foysal", "Foysal Mamun", "FM", "1221")
	// checkErr(err)

	// id, err := res.LastInsertId()
	// checkErr(err)

	stmt, err := db.Prepare("UPDATE bf_countries SET printable_name=? WHERE numcode=?")
	checkErr(err)

	res, err := stmt.Exec("Mamun1", 1221)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	rows, err := db.Query("SELECT * FROM bf_countries WHERE numcode=1221 LIMIT 5")
	checkErr(err)

	for rows.Next() {
		var iso string
		var name string
		var printable_name string
		var iso3 string
		var numcode int
		err = rows.Scan(&iso, &name, &printable_name, &iso3, &numcode)
		checkErr(err)
		fmt.Println(iso)
		fmt.Println(name)
		fmt.Println(printable_name)
		fmt.Println(iso3)
		fmt.Println(numcode)
	}

}

func main() {
	// LearnBasic()
	LearnORM()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
