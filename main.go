package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/inhies/go-bytesize"
	"html"
	"html/template"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
	"time"
)

const (
	layoutsDir   = "templates/layouts"
	templatesDir = "templates"
	extension    = "/*.html.tmpl"
)

var (
	//go:embed templates/* templates/layouts/*
	files          embed.FS
	templates      map[string]*template.Template
	cssFileContent []byte

	//go:embed templates/css/bootstrap.css
	cssFile embed.FS
)

func LoadTemplates() error {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	tmplFiles, err := fs.ReadDir(files, templatesDir)
	if err != nil {
		return err
	}

	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() {
			continue
		}

		var mapKey = strings.Split(tmpl.Name(), ".")[0]
		pt, err := template.New(tmpl.Name()).Funcs(template.FuncMap{
			"calcPerc": calcPerc,
			"calcSize": calcSize,
		}).ParseFS(files, templatesDir+"/"+tmpl.Name(), layoutsDir+extension)
		if err != nil {
			return err
		}

		templates[mapKey] = pt
	}

	cssFileContent, _ = cssFile.ReadFile("templates/css/bootstrap.css")
	return nil
}

func calcPerc(available int64, total int64) string {
	return fmt.Sprintf("%.1f", 100-100.*(float32(available)/float32(total)))
}

func calcSize(x1 int64) string {
	return bytesize.New(float64(x1)).String()
}

func init() {
	checkInit()
	err := LoadTemplates()
	if err != nil {
		fmt.Println(err)
	}

}

func main() {

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

		/*
			fmt.Println(facts.AnsibleFacts.AnsibleLvm.Lvs)
			fmt.Println(facts.AnsibleFacts.AnsibleLvm.Pvs)
			fmt.Println(facts.AnsibleFacts.AnsibleLvm.Vgs)
		*/

	}

	fmt.Println("Generate output in " + settings.exportPath)
	f, err = os.Create(settings.exportPath + "/index.html")
	if err != nil {
		log.Fatalln(err)
	}

	ServerFacts.Facts = allServer
	ServerFacts.PageTitle = "ansible Dashboard"
	ServerFacts.CreationDate = time.Now()

	ServerFacts.CssFile = html.UnescapeString(string(cssFileContent))

	temp := templates["index"]
	if temp == nil {
		log.Fatalln(templates)
	}
	fmt.Println("Template loaded")

	err = temp.Execute(f, ServerFacts)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("executed")
	err = f.Close()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Finish")
}
