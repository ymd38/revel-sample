package controllers

import (
	"database/sql"
	"security-cop/app/models"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/modules/db/app"
	r "github.com/revel/revel"
)

var (
	Dbm *gorp.DbMap
)

func InitDB() {
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	Dbm.TraceOn("[gorp]", r.INFO)
	Dbm.AddTable(models.Issue{}).SetKeys(true, "Id")
	Dbm.AddTable(models.User{}).SetKeys(true, "Id")
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
