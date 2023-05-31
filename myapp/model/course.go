package model

import postgres "myapp/datastore/postgres"

// course
type Course struct {
	CId        string `json:"cid"`
	CourseName string `json:"coursename"`
}

const (
	// course
	queryInsertcourse  = "INSERT INTO course(cid,coursename) VALUES($1,$2)"
	queryGetUsercourse = "select cid,coursename FROM course WHERE cid=$1;"
	queryUpdatecourse  = "UPDATE course SET cid=$1, coursename=$2 WHERE cid=$3 RETURNING cid"
	queryDeletecourse  = "DELETE FROM course WHERE cid=$1;"
)

// for course
func (c *Course) Create() error {
	_, err := postgres.Db.Exec(queryInsertcourse, c.CId, c.CourseName)
	return err

	// return nil
}

func (c *Course) Read() error {
	row := postgres.Db.QueryRow(queryGetUsercourse, c.CId) // s.std = $1
	err := row.Scan(&c.CId, &c.CourseName)
	return err

}
func (c *Course) Update(old_coId string) error {
	err := postgres.Db.QueryRow(queryUpdatecourse, c.CId, c.CourseName, old_coId).Scan(&c.CId)
	return err
}

func (c *Course) Delete() error {
	if _, err := postgres.Db.Exec(queryDeletecourse, c.CId); err != nil {
		return err
	}
	return nil
}

// getallcour
func GetALLCourses() ([]Course, error) { // no particular receiver because we need a set of students not particular
	rows, err := postgres.Db.Query("SELECT * FROM Course;") // directly passing from here
	// error handling
	if err != nil {
		return nil, err
	}
	//    slice of type student
	courses := []Course{} // empty slice of type student

	// iterate rows
	for rows.Next() { // next iterates rows one by one
		var c Course
		dbErr := rows.Scan(&c.CId, &c.CourseName) // scan - store the values in s
		if dbErr != nil {
			return nil, dbErr
		}
		courses = append(courses, c) // adding data in []student using append
	}

	rows.Close()
	return courses, nil
}
