package main

import (
	"os"
	"io/ioutil"
	"github.com/Sirupsen/logrus"
	"encoding/xml"
)

type Query struct {
	__name []string `xml:"example"`
}

func HandleXML(folders []string,ResultFile *os.File){
	for _, folder := range folders {
		var q Query
		logrus.Info("Opening ", folder)
		f, err := os.Open(folder + "/" + "file.xml")
		if err != nil {
			ResultFile.WriteString(folder+" "+"not found\n")
			continue
		}
		text, err := ioutil.ReadAll(f)
		if err!=nil{
			logrus.Error("Error reading from file")
		}
		f.Close()
		err=xml.Unmarshal(text, &q)
		if err!=nil{
			logrus.Info("xml unmarshaling error: ",err)
		}
		for _, j := range q.__name {
			ResultFile.WriteString(folder + " " + j + "\n")
		}
	}
}

func GetAllFolders(dir string)[]string{
	var folders []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		logrus.Fatal(err)
	}
	for _, file := range files {
		if file.Name()==".idea"{
			continue
		} else if !file.IsDir(){
			continue
		}
		folders = append(folders, file.Name())
	}
	return folders
}

func main() {
	ResultFile, err := os.Create("Result.txt")
	if err != nil {
		logrus.Println("Error creating file Result.txt")
	}
	defer ResultFile.Close()
	
	folders:=GetAllFolders(".")
	HandleXML(folders,ResultFile)
}
