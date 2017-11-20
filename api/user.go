package api

import (
	"net/http"
	"strconv"
	"github.com/s4kibs4mi/emq-am/data"
)

/**
 * := Coded with love by Sakib Sami on 20/11/17.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

func ListUsers(w http.ResponseWriter, r *http.Request) {
	currentPage := 0
	page := r.URL.Query().Get("page")
	n, err := strconv.Atoi(page)
	if err == nil {
		currentPage = n
	}
	user := data.User{}
	users := user.GetUserList(currentPage)
	ServeJSON(w, APIResponse{
		Code: http.StatusOK,
		Data: users,
	}, http.StatusOK)
}
