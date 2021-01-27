package multiselector

import (
	"errors"
	"strconv"
	"strings"

	"github.com/globalsign/mgo/bson"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrBadRules       = errors.New("bad rules")
	ErrBadFormat      = errors.New("bad format")
)

const (
	OperatorLowerOrEqual   = "<="
	OperatorLower          = "<"
	OperatorGreaterOrEqual = ">="
	OperatorGreater        = ">"
	OperatorEqual          = "=="
	OperatorNotEqual       = "!="
)

type Selector struct {
	rules []string
}

func NewSelector(rules interface{}) (*Selector, error) {
	s := &Selector{}
	switch rules.(type) {
	case []string:
		s.rules = rules.([]string)
	case string:
		s.rules = strings.Split(rules.(string), ",")
	default:
		return nil, ErrBadRules
	}
	return s, nil
}

func NewSelectorDef(rules interface{}) *Selector {
	s, err := NewSelector(rules)
	if err != nil {
		return &Selector{}
	}
	return s
}

func (s *Selector) AddRule(rule interface{}) {
	switch rule.(type) {
	case []string:
		tmp := s.rules
		tmp = append(tmp, rule.([]string)...)
		s.rules = tmp
	case string:
		tmp := s.rules
		tmp = append(tmp, strings.Split(rule.(string), ",")...)
		s.rules = tmp
	default:
		// do nothing
	}
}

func (s Selector) Len() int {
	return len(s.rules)
}

func (s Selector) ToSql() (string, error) {
	// TODO: implement parsing rules to sql
	return "", ErrNotImplemented
}

func (s Selector) ToBson() (bson.M, error) {
	selector := bson.M{}
	for _, r := range s.rules {
		parts := strings.Split(r, " ")
		switch len(parts) {
		case 3:
			key, operator := parts[0], parts[1]
			strValue := strings.Replace(parts[2], "'", "", -1)
			var value interface{}
			if v, err := strconv.ParseInt(strValue, 10, 64); err == nil {
				value = v
			} else if f, err := strconv.ParseFloat(strValue, 64); err == nil {
				value = f
			} else if strValue == "true" {
				value = true
			} else if strValue == "false" {
				value = false
			} else {
				value = strValue
			}
			switch operator {
			case OperatorLowerOrEqual:
				selector[key] = bson.M{"$lte": value}
			case OperatorLower:
				selector[key] = bson.M{"$lt": value}
			case OperatorGreaterOrEqual:
				selector[key] = bson.M{"$gte": value}
			case OperatorGreater:
				selector[key] = bson.M{"$gt": value}
			case OperatorEqual:
				selector[key] = value
			case OperatorNotEqual:
				selector[key] = bson.M{"$ne": value}
			}
		default:
			continue
		}
	}
	return selector, nil
}
