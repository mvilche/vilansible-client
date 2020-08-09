package main

import (
	"os"
	"os/user"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitdB() error {

	db, err := OpenSQL()
	if err == nil {
		if err := db.AutoMigrate(&Execution{}).Error; err != nil {
			ErrorLog.Printf(err.Error())
			CloseSQL(db)
			return err
		}

	} else {

		ErrorLog.Printf(err.Error())
		return err
	}
	CloseSQL(db)

	return nil
}

func OpenSQL() (*gorm.DB, error) {

	u, err := user.Current()

	if err != nil {

		return nil, err
	}

	if _, err := os.Stat(u.HomeDir + "/.vilansible"); os.IsNotExist(err) {
		err := os.Mkdir(u.HomeDir+"/.vilansible", os.ModePerm)
		if err != nil {

			return nil, err
		}
	}

	db, err := gorm.Open("sqlite3", u.HomeDir+"/.vilansible/database.db")
	if err != nil {

		return db, err
	}

	return db, err
}

func CloseSQL(db *gorm.DB) {

	db.Close()

}
