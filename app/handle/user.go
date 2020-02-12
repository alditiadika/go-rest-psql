package handle

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/alditiadika/go-rest-psql/app/model"
	"github.com/alditiadika/go-rest-psql/app/queries"
	"github.com/alditiadika/go-rest-psql/utils"
)

//GetUser func
func GetUser(res http.ResponseWriter, req *http.Request, client *sql.DB) {
	fmt.Println("[GET] /users")
	rows, err := client.Query(queries.GetAllUsers())
	defer rows.Close()
	var users []model.UserModel
	ErrModel := model.ErrorModel{}

	if err != nil {
		ErrModel = model.ErrorModel{
			IsError: true,
			Message: "error when call table",
			Data:    err.Error(),
		}
		utils.ResponseSender(res, req, ErrModel)
		return
	}
	for rows.Next() {
		item := model.UserModel{}
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.IsActive,
			&item.CreatedBy,
			&item.CreatedDate,
			&item.ModifiedBy,
			&item.ModifiedDate,
		)
		if err != nil {
			ErrModel = model.ErrorModel{
				IsError: true,
				Message: "error where scan users item",
				Data:    err.Error(),
			}
			utils.ResponseSender(res, req, ErrModel)
			return
		}
		users = append(users, item)
	}
	err = rows.Err()
	if err != nil {
		ErrModel = model.ErrorModel{
			IsError: true,
			Message: "error where querying",
			Data:    err.Error(),
		}
		utils.ResponseSender(res, req, ErrModel)
		return
	}
	if len(users) == 0 {
		var r [0]int
		utils.ResponseSender(res, req, r)
		return
	}
	utils.ResponseSender(res, req, users)
}

//GetUserByParam func
func GetUserByParam(res http.ResponseWriter, req *http.Request, client *sql.DB) {
	fmt.Println("[GET]", req.URL)
	qsGen, err := queries.GetUserByParam(req)
	if err != nil {
		fmt.Println("error when generate query")
		fmt.Println(err)
	}
	rows, err := client.Query(qsGen)
	defer rows.Close()
	var users []model.UserModel
	ErrModel := model.ErrorModel{}

	if err != nil {
		ErrModel = model.ErrorModel{
			IsError: true,
			Message: "error when call table",
			Data:    err.Error(),
		}
		utils.ResponseSender(res, req, ErrModel)
		return
	}
	for rows.Next() {
		item := model.UserModel{}
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.IsActive,
			&item.CreatedBy,
			&item.CreatedDate,
			&item.ModifiedBy,
			&item.ModifiedDate,
		)
		if err != nil {
			ErrModel = model.ErrorModel{
				IsError: true,
				Message: "error where scan users item",
				Data:    err.Error(),
			}
			utils.ResponseSender(res, req, ErrModel)
			return
		}
		users = append(users, item)
	}
	err = rows.Err()
	if err != nil {
		ErrModel = model.ErrorModel{
			IsError: true,
			Message: "error where querying",
			Data:    err.Error(),
		}
		utils.ResponseSender(res, req, ErrModel)
		return
	}
	if len(users) == 0 {
		var r [0]int
		utils.ResponseSender(res, req, r)
		return
	}
	utils.ResponseSender(res, req, users)
}

//UpdateUser func
func UpdateUser(res http.ResponseWriter, req *http.Request, client *sql.DB) {
	fmt.Println("[PUT]", req.URL)
}
