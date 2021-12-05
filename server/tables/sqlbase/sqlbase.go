package sqlbase

import (
	"danteserver/server/pool"
	"database/sql"
	"fmt"
	"gitee.com/yuanxuezhe/dante/comm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func NewDB() (*Dbhandle, error) {
	handle := &Dbhandle{}
	handle.conn = pool.SqlPool.Get().(*sql.DB)
	return handle, nil
}

type Dbhandle struct {
	conn *sql.DB
	sql  string
	stmt *sql.Stmt
	rows *sql.Rows
	log  *log.Logger
}

func (h *Dbhandle) Open(sql string) error {
	var err error
	h.sql = sql
	fmt.Println("Open: ", h.sql)
	return err
}

func (h *Dbhandle) SetParam(args ...interface{}) error {
	var err error
	h.sql = comm.Sprintf(h.sql, args...)
	return err
}

func (h *Dbhandle) Exec() (int, error) {
	var err error
	var i int64
	fmt.Println("Exec: ", h.sql)
	rs, err := h.conn.Exec(h.sql)
	h.conn.Begin()
	if err != nil {
		i = 0
	} else {
		i, err = rs.RowsAffected()
	}
	return int(i), nil
}

func (h *Dbhandle) Query() error {
	var err error
	fmt.Println("Query: ", h.sql)
	h.rows, err = h.conn.Query(h.sql)
	return err
}

func (h *Dbhandle) GetValue(args ...interface{}) error {
	return h.rows.Scan(args...)
}

func (h *Dbhandle) Next() bool {
	return h.rows.Next()
}

func (h *Dbhandle) Close() error {
	var err error
	err = h.rows.Close()
	if err != nil {
		return err
	}
	//err = h.stmt.Close()
	//if err != nil {
	//	return err
	//}
	err = h.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
