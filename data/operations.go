package data

//InsertErrorLog : Insert All Error log record according to alert ID, file name
func InsertErrorLog(alertid, errormsg, fileName string) {
	sqlStr := `INSERT INTO duranz_errorlog (alert_id, error_msg, file_name) VALUES (?, ?, ?)`

	_, err := Db.Exec(sqlStr, alertid, errormsg, fileName)
	if err != nil {
		panic(err)
	}

}
