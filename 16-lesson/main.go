// package main

// import (
// 	"database/sql"
// 	"strconv"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()
// 	f := Filter{}
// 	r.GET("/getAll", f.getAll)
// }

// func (f Filter) getAll(c *gin.Context) {
// 	var (
// 		params = make(map[string]interface{})
// 		arr    []interface{}
// 		limit  string
// 	)

// 	query := `select id, first_name, last_name, age, gender, nation, field, parent_name, city
// 	 	from People where deleted_at != 0 and nation = $3 and gender = $1 and age = $2  `

// 	if len(f.Gender) > 0 {
// 		params["gender"] = f.Gender
// 		filter += " and gender = :gender "
// 	}

// 	if len(f.Nation) > 0 {
// 		params["nation"] = f.Gender
// 		filter += " and nation = :nation "
// 	}

// 	if len(f.Field) > 0 {
// 		params["field"] = f.Gender
// 		filter += " and field = :field "
// 	}

// 	if f.Age > 0 {
// 		params["age"] = f.Gender
// 		filter += " and age = :age "
// 	}

// 	if f.Limit > 0 {
// 		params["limit"] = f.Limit
// 		limit = ` LIMIT :limit`
// 	}

// 	//if f.Offset > 0 {
// 	//	params["offset"] = (f.Offset - 1) * f.Limit
// 	//	offset = ` OFFSET :offset`
// 	//}
// 	//
// 	query = query + filter + limit

// 	query, arr = ReplaceQueryParams(query, params)
// 	db := &sql.DB{}
// 	db.Query(query, arr...)

// }

// type Filter struct {
// 	Age           int
// 	Gender        string
// 	Nation        string
// 	Field         string
// 	Limit, Offset int
// }

// func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
// 	var (
// 		i    int = 1
// 		args []interface{}
// 	)

// 	for k, v := range params {
// 		namedQuery = "where nation = $3 and gender = $1 and age = $2 "
// 		if k != "" && strings.Contains(namedQuery, ":"+k) {
// 			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
// 			args = append(args, v)

// 			i++
// 		}
// 	}

// 	return namedQuery, args
// }
