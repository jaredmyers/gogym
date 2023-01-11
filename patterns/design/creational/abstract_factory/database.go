package abstractfactory

import "fmt"

func DatabaseFactory(env string) Database {
	switch env {
	case "production":
		return &mongoDB{
			database: make(map[string]string),
		}
	case "developement":
		return &sqlite{
			database: make(map[string]string),
		}
	default:
		return nil
	}
}

type (
	Database interface {
		GetData(string) string
		PutData(string, string)
	}
	mongoDB struct {
		database map[string]string
	}
	sqlite struct {
		database map[string]string
	}
)

func (mdb *mongoDB) GetData(query string) string {
	if _, ok := mdb.database[query]; !ok {
		return ""
	}
	fmt.Println("MongoDB")
	return mdb.database[query]
}

func (sql *sqlite) GetData(query string) string {
	if _, ok := sql.database[query]; !ok {
		return ""
	}
	fmt.Println("sqlite")
	return sql.database[query]
}

func (mdb *mongoDB) PutData(query, data string) {
	mdb.database[query] = data
}

func (sql *sqlite) PutData(query, data string) {
	sql.database[query] = data
}
