package tttt

import (
	"danteserver/server/tables/sqlbase"
	"errors"
	"fmt"
)

type Userinfo struct {
	Userid       int    `json:"userid"`
	Username     string `json:"username"`
	Passwd       string `json:"passwd"`
	Sex          string `json:"sex"`
	Phone        int    `json:"phone"`
	Email        string `json:"email"`
	Status       string `json:"status"`
	Registerdate int    `json:"registerdate"`
}

var TableUserColumns = func(surfix string) string {
	if len(surfix) > 0 {
		surfix = surfix + "."
	}
	return surfix + "userid, " +
		surfix + "username, " +
		surfix + "passwd, " +
		surfix + "sex, " +
		surfix + "phone, " +
		surfix + "email, " +
		surfix + "status, " +
		surfix + "registerdate "
}

func (t *Userinfo) InsertExample() (userinfo *Userinfo, err error) {
	Db, err := sqlbase.NewDB()
	if err != nil {
		return nil, err
	}
	defer Db.Close()
	//strSql := "SELECT " + TableUserColumns("") + " FROM userinfo a where (userid = ? or phone = ? or email = ?) and passwd = ?"
	strSql := "insert into userinfo (  " + TableUserColumns("") + ") values (?,?,?,?,?,?,?,?)"
	err = Db.Open(strSql)
	if err != nil {
		return nil, errors.New("Connection to mysql failed!")
	}

	//err = Db.SetParam(t.Userid, t.Phone, t.Email, t.Passwd)
	err = Db.SetParam(t.Userid,
		t.Username,
		t.Passwd,
		t.Sex,
		t.Phone,
		t.Email,
		t.Status,
		t.Registerdate)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//if Db.Next() {
	//	err = Db.GetValue(&t.Userid,
	//		&t.Username,
	//		&t.Passwd,
	//		&t.Sex,
	//		&t.Phone,
	//		&t.Email,
	//		&t.Status,
	//		&t.Registerdate)
	//
	//	if err != nil {
	//		return nil, err
	//	}
	//} else {
	//	return nil, errors.New("Login failed! Userinfo not exists or passwd is wrong!")
	//}

	return t, nil
}

func (t *Userinfo) QueryExample() (userinfo *Userinfo, err error) {
	Db, err := sqlbase.NewDB()
	if err != nil {
		return nil, err
	}
	defer Db.Close()
	strSql := "SELECT " + TableUserColumns("") + " FROM userinfo where (userid = %1 or phone = %2 or email = %3) and passwd = %4 "
	err = Db.Open(strSql)
	if err != nil {
		return nil, errors.New("Connection to mysql failed!")
	}

	err = Db.SetParam(t.Userid, t.Phone, t.Email, t.Passwd)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = Db.Query()

	if Db.Next() {
		err = Db.GetValue(&t.Userid,
			&t.Username,
			&t.Passwd,
			&t.Sex,
			&t.Phone,
			&t.Email,
			&t.Status,
			&t.Registerdate)

		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Login failed! Userinfo not exists or passwd is wrong!")
	}

	return t, nil
}
