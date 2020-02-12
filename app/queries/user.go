package queries

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/alditiadika/go-rest-psql/utils"
)

//GetAllUsers queries
func GetAllUsers() string {
	return `
		SELECT * FROM master_user mu
	`
}

type e error

//GetUserByParam
func GetUserByParam(req *http.Request) (string, e) {
	queryReturn := "select * from master_user mu where 1=1 "
	u, err := url.Parse(req.URL.String())
	if err != nil {
		return "parse url err", err
	}
	operator := ""
	params := u.Query()
	for key, v := range params {
		rv := v[0]
		raw := strings.Split(v[0], "|")
		if len(raw) == 0 || len(raw) == 1 {
			operator = "eq"
		} else {
			operator = raw[1]
			rv = raw[0]
		}
		value := strings.Split(rv, ",")
		if len(value) > 0 {
			fieldType := databaseFieldType(key)
			switch fieldType {
			case "numeric":
				queryReturn += utils.NumericQuery(key, value, operator)
				break
			case "string":
				queryReturn += utils.StringQuery(key, value, operator)
				break
			case "boolean":
				queryReturn += utils.BooleanQuery(key, value[0])
				break
			case "date":
				break
			default:
				break
			}
		}
	}
	return queryReturn, nil
}

func databaseFieldType(field string) string {
	switch field {
	case "id":
		return "numeric"
	case "name":
		return "string"
	case "is_active":
		return "boolean"
	case "created_date":
		return "date"
	case "created_by":
		return "string"
	case "modified_date":
		return "date"
	case "modified_by":
		return "string"
	default:
		return ""
	}
}
