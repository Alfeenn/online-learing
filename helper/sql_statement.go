package helper

import "fmt"

func SQLStatement(host, port, user, password, dbname string) string {

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user, password, host, port, dbname)

}
