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
		fmt.Println("opened ", file.Name())

		byteValue, errReadAll := io.ReadAll(jsonFile)
		if errReadAll != nil {
			log.Fatalln(errReadAll)
		}

		err = jsonFile.Close()
		if err != nil {
			log.Fatalln(err)
		}

		var facts AnsibleFactFull
		err = json.Unmarshal(byteValue, &facts)
		if err != nil {
			log.Fatalln(err)
		}
		allServer = append(allServer, facts)

		fmt.Println(facts.AnsibleFacts.AnsibleLvm.Lvs)
		fmt.Println(facts.AnsibleFacts.AnsibleLvm.Pvs)
		fmt.Println(facts.AnsibleFacts.AnsibleLvm.Vgs)

	}

	fmt.Println("Generate output in " + settings.exportPath)
	f, err = os.Create(settings.exportPath + "/index.html")
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

	err = f.Close()
	if err != nil {
		log.Fatalln(err)
	}

}
