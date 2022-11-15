// Golang
// r/w yaml

// must be pointer to struct to Unmarshal?
// doesn't need to be pointer to struct to marshal?

package filefun

import (
	"fmt"
	//"os"
	"bufio"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type DirsAndFiles struct {
	DataSpace   string `yaml:"dataSpace"`
	ConfigSpace string `yaml:"configSpace"`
}

type InstanceConfig struct {
	Port          string   `yaml:"port"`
	Host          string   `yaml:"host"`
	BindIP        string   `yaml:"bindIp"`
	CustomOptions []string `yaml:"customOptions"`
}

type Instance struct {
	Installation bool           `yaml:"install"`
	InstanceType string         `yaml:"instanceType"`
	DirsAndFiles DirsAndFiles   `yaml:"dirsAndFiles"`
	ConfigFile   InstanceConfig `yaml:"configFile"`
}

type Person struct {
	Name  string             `yaml:"name"`
	Age   int                `yaml:"age"`
	Tasks []string           `yaml:"tasks"`
	Langs map[string]float32 `yaml:"langs"`
}

type PeopleGroup struct {
	People []Person `yaml:"people"`
}

func ReadInstanceYaml(serverDataFile string, obj *Instance) {
	source1, err := os.ReadFile(serverDataFile)
	if err != nil {
		log.Panic(err)
	}

	if err = yaml.Unmarshal(source1, &obj); err != nil {
		log.Panic(err)
	}
}

func Read2() {

	f := Instance{}
	ReadInstanceYaml("sample.yaml", &f)

	fmt.Println(f.Installation)
	fmt.Println(f.InstanceType)
	fmt.Println(f.DirsAndFiles)
	fmt.Println(f.ConfigFile)
}

func Read1() {

	p1 := &PeopleGroup{}

	//direct read
	content, err := os.ReadFile("t.yaml")
	if err != nil {
		log.Panic(err)
	}

	if err = yaml.Unmarshal(content, &p1); err != nil {
		log.Panic(err)
	}

}

func writeTest1() {

	p1 := Person{
		"John",
		30,
		[]string{"feature/task1", "feature/task2"},
		map[string]float32{"Golang": 19.1, "C++": 10},
	}
	p2 := Person{
		"Mary",
		18,
		[]string{"feature/task4", "feature/task5"},
		map[string]float32{"Java": 34.1, "Ruby": 90},
	}

	peopleList := PeopleGroup{}
	peopleList.People = append(peopleList.People, p1, p2)

	yPeople, err := yaml.Marshal(peopleList)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(yPeople)
	fmt.Println(string(yPeople))

	/*
		//direct write
		if err = os.WriteFile("t.yaml", yPeople, 0644); err != nil{
				log.Panic(err)
		}
	*/

	//buffered write
	file, err := os.OpenFile("t.yaml", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)

	_, err = bufferedWriter.Write(yPeople)
	if err != nil {
		log.Panic(err)
	}

	bufferedWriter.Flush()
}
