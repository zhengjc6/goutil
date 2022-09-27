package csvparse

import (
	"fmt"
	"goutil/strtool"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func visit(files *[]string) filepath.WalkFunc {
	return func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if path.Ext(filePath) == ".csv" {
			*files = append(*files, filePath)
		}
		return nil
	}
}

func ParseDir(dirPath, outputDir, pkgName string) error {
	_, err := os.Stat(outputDir)
	//check dir exists or not
	if err != nil {
		if os.IsNotExist(err) {
			// if dir does not exists then create
			if err = os.MkdirAll(outputDir, 0744); err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		//the dir must be empty,otherwise it needs to be emptied manually
		dir, _ := ioutil.ReadDir(outputDir)
		if len(dir) > 0 {
			return fmt.Errorf("please lean the output dir:%s", outputDir)
		}
	}

	var files []string

	err = filepath.Walk(dirPath, visit(&files))

	if err != nil {
		panic(err)
	}

	//create package
	outFile, err := os.OpenFile(outputDir+"\\"+strings.ToLower(pkgName)+".go", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0744)
	if err != nil {
		return err
	}
	defer outFile.Close()
	strucName := "DataStruct"
	valNamme := "CSVData"
	outFile.WriteString(fmt.Sprintf("package %s\n\n", pkgName))
	//create struct
	outFile.WriteString(fmt.Sprintf("type %s struct {\n", strucName))

	var fileCreateList []string
	var memberNameList []string
	for _, file := range files {
		fileName, mapName, err := parseFile(file, outputDir, pkgName)
		if err != nil {
			return err
		}
		memberName := strtool.FirstUpper(fileName) + "Table"
		outFile.WriteString(fmt.Sprintf("\t\t%s\t\t%s\n", memberName, mapName))
		fileCreateList = append(fileCreateList, "Create"+memberName)
		memberNameList = append(memberNameList, memberName)
	}
	outFile.WriteString("}\n\n")

	outFile.WriteString(fmt.Sprintf("var  %s %s\n\n", valNamme, strucName))

	outFile.WriteString("func init(){\n")
	for i := 0; i < len(memberNameList); i++ {
		outFile.WriteString(fmt.Sprintf("\t%s.%s = %s()\n", valNamme, memberNameList[i], fileCreateList[i]))
	}

	outFile.WriteString("}\n")
	return nil
}
