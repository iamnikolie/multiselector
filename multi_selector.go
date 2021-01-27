package multiselector

import "github.com/globalsign/mgo/bson"

type MultiSelector interface {
	AddRule(rule interface{})
	Len() int
	ToBson() (bson.M, error)
	ToSql() (string, error)
}
