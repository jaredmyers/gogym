package abstractfactory

import "fmt"

// -- Helper Function -- //
func SetupConstructors(env string) (DatabaseCreator, FilesystemCreator) {
	dbFactory := AbstractFactory("database")
	fsFactory := AbstractFactory("filesystem")

	return dbFactory(env).(DatabaseCreator), fsFactory(env).(FilesystemCreator)

}

// -- Abstract Factory for Factories -- //
func AbstractFactory(fact string) factory {
	switch fact {
	case "database":
		return DatabaseFactory
	case "filesystem":
		return FilesystemFactory
	default:
		return nil
	}
}

type factory func(string) any

// -- Database Factory -- //
func DatabaseFactory(env string) any {
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
	DatabaseCreator interface {
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

// -- FileSystem Factory -- //

func FilesystemFactory(env string) any {
	switch env {
	case "production":
		return &ntfs{files: make(map[string]file)}
	case "development":
		return &ext4{files: make(map[string]file)}
	default:
		return nil
	}
}

type (
	file struct {
		name    string
		content string
	}

	ntfs struct {
		files map[string]file
	}

	ext4 struct {
		files map[string]file
	}
	FilesystemCreator interface {
		CreateFile(string)
		FindFile(string) file
	}
)

func (n *ntfs) CreateFile(path string) {
	file := file{content: "NTFS file", name: path}
	n.files[path] = file
	fmt.Println("ntfs")
}
func (n *ntfs) FindFile(path string) file {
	if _, ok := n.files[path]; !ok {
		return file{}
	}
	return n.files[path]
}

func (e *ext4) CreateFile(path string) {
	file := file{content: "ext4 file", name: path}
	e.files[path] = file
	fmt.Println("ext")
}
func (e *ext4) FindFile(path string) file {
	if _, ok := e.files[path]; !ok {
		return file{}
	}
	return e.files[path]
}
