package main

type Execution struct {
	Application string `header:"Application"`
	Version     string `header:"Version"`
	TypeExec    string `header:"Type"`
	Date        string `header:"Date"`
	Status      bool   `header:"Passed"`
	User        string `header:"User exec"`
}

func GetExecutions() []Execution {

	var result []Execution

	db, _ := OpenSQL()
	db.Find(&result)
	db.Close()

	return result

}

func SaveExecution(e Execution) error {

	db, _ := OpenSQL()

	if err := db.Create(&e).Error; err != nil {
		ErrorLog.Printf(err.Error())
		CloseSQL(db)
		return err

	}

	CloseSQL(db)
	return nil
}
