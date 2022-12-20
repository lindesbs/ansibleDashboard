package main

import (
	"encoding/json"
	"fmt"
	"github.com/inhies/go-bytesize"
	"html/template"
	"io"
	"log"
	"os"
	"time"
)

func init() {
	//temp = template.Must(template.New("index.html").Funcs(template.FuncMap(myFuncMap)).ParseFiles("public-src/index.html"))
	temp = template.Must(template.New("index.html").Funcs(template.FuncMap{
		"calcPerc": func(available int64, total int64) string {
			return fmt.Sprintf("%.1f", 100-100.*(float32(available)/float32(total)))
		},

		"calcSize": func(x1 int64) string {
			return bytesize.New(float64(x1)).String()
		},
	}).ParseFiles(strTemplate))

}

func main() {
	checkInit()

	f, err := os.Open(settings.importPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		jsonFile, err := os.Open(settings.importPath + file.Name())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("opened ", file.Name())

		err = jsonFile.Close()
		if err != nil {
			log.Fatalln(err)
		}

		byteValue, _ := io.ReadAll(jsonFile)

		// we initialize our Users array
		var facts AnsibleFactFull

		// we unmarshal our byteArray which contains our
		// jsonFile's content into 'users' which we defined above
		err = json.Unmarshal(byteValue, &facts)
		if err != nil {
			return
		}
		allServer = append(allServer, facts)

		fmt.Println("*")
	}

	f, err = os.Create(settings.exportPath + "/index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatalln(err)
	}

	ServerFacts.Facts = allServer
	ServerFacts.PageTitle = "ansible Dashboard"
	ServerFacts.CreationDate = time.Now()

	err = temp.Execute(f, ServerFacts)
	if err != nil {
		log.Fatalln(err)
	}
}
