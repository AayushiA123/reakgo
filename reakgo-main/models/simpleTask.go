package models

import (
	"log"
	"reakgo/utility"
)

type TaskAddView struct {
	Id           int32
	TaskName     string
	Deadline     int64
	DeadlineUnix string
	Completed    bool
}

func (form TaskAddView) Add(taskname string, deadline int64) {
	utility.Db.MustExec("INSERT INTO addTask (taskname, deadline) VALUES (?, ?)", taskname, deadline)
}

func (form TaskAddView) View() ([]TaskAddView, error) {
	var resultSet []TaskAddView
	//SELECT id, taskname, DATE_FORMAT(FROM_UNIXTIME(deadline),'%e %b %Y') FROM addTask ORDER BY ABS(deadline - UNIX_TIMESTAMP())
	rows, err := utility.Db.Query("select id, taskname,DATE_FORMAT(FROM_UNIXTIME(deadline), '%e %b %Y %h:%i %p') AS 'deadline',completed from addTask ORDER BY UNIX_TIMESTAMP(FROM_UNIXTIME(deadline)) ASC")
	//	rows, err := utility.Db.Query("SELECT id, taskname, DATE_FORMAT(FROM_UNIXTIME(deadline),'%e %b %Y %h:%i %p'), completed FROM addTask ORDER BY ABS(deadline - UNIX_TIMESTAMP())")
	//	rows, err := utility.Db.Query("SELECT id, taskname, DATE_FORMAT(FROM_UNIXTIME(deadline),'%e %b %Y %h:%i %p'), completed FROM addTask ORDER BY ABS(deadline - UNIX_TIMESTAMP())")

	if err != nil {
		log.Println(err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var singleRow TaskAddView
			err = rows.Scan(&singleRow.Id, &singleRow.TaskName, &singleRow.DeadlineUnix, &singleRow.Completed)
			if err != nil {
				log.Println(err)
			} else {
				resultSet = append(resultSet, singleRow)
			}
		}
	}
	return resultSet, err
}

func (form TaskAddView) Delete(taskID int32) error {
	_, err := utility.Db.Exec("DELETE FROM addTask WHERE id=?", taskID)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (form TaskAddView) Complete(taskID int32) error {
	_, err := utility.Db.Exec("UPDATE addTask SET completed = true WHERE id = ?", taskID)
	if err != nil {
		log.Println(err)
	}
	return err
}
