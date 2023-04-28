package csvparse

import (
	"encoding/csv"
	"fmt"
	"io"
	"metaserver/pkg/strtool"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var typeMap map[string]string
var regCol *regexp.Regexp

func init() {
	typeMap = map[string]string{
		"int":           "int64",
		"string":        "string",
		"float":         "float64",
		"list<long>":    "[]int64",
		"list<int>":     "[]int64",
		"list<string>":  "[]string",
		"list<float64>": "[]float64",
	}
	regCol = regexp.MustCompile(`\|`)
}

func parseCol(colName, colType, colStr string) string {
	switch colType {
	case "int", "float":
		if colStr == "" {
			colStr = "0"
		}
		return fmt.Sprintf("%s:%s", colName, colStr)
	case "string":
		return fmt.Sprintf("%s:`%s`", colName, colStr)
	case "list<int>", "list<float>", "list<long>", "list<string>":
		if colStr == "" {
			return fmt.Sprintf("%s:%s{}", colName, typeMap[colType])
		}
		colStr = regCol.ReplaceAllLiteralString(colStr, ",")
		return fmt.Sprintf("%s:%s{%s}", colName, typeMap[colType], colStr)
	// case "list<string>":
	// 	if colStr == "" {
	// 		return fmt.Sprintf("%s:%s{}", colName, typeMap[colType])
	// 	}
	// 	rexp := regexp.MustCompile("[|]")
	// 	colStr = rexp.ReplaceAllString(colStr, "\",\"")
	// 	return fmt.Sprintf("%s:%s{\"%s\"}", colName, typeMap[colType], colStr)
	default:
		panic("Invalid Column Data " + colStr)
	}
}

func parseFileName(filePath string) (fileNameWithSuffix, fileType, fileNameOnly string) {

	fileNameWithSuffix = path.Base(filepath.ToSlash(filePath))

	fileType = path.Ext(fileNameWithSuffix)

	fileNameOnly = strings.TrimSuffix(fileNameWithSuffix, fileType)
	return
}

func parseFile(filePath, outputDir, pkgName string) (fileNameOnly string, mapName string, err error) {
	fmt.Printf("fullPath = %s\n", filePath)
	fileNameWithSuffix, fileType, fileNameOnly := parseFileName(filePath)

	fmt.Printf("fileNameWithSuffix = %s;\nfileType = %s;\nfileNameOnly = %s;\n",
		fileNameWithSuffix, fileType, fileNameOnly)
	file, err := os.Open(filePath)
	if err != nil {
		return fileNameOnly, mapName, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	valcomments, valnames, valtypes, emptyCol := getHeadInfo(fileNameOnly, reader)
	var invalidCol = len(valnames)
	outFile, err := os.OpenFile(outputDir+"\\"+strings.ToLower(fileNameOnly)+".go", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0744)
	if err != nil {
		return fileNameOnly, mapName, err
	}
	defer outFile.Close()
	//create struct
	outFile.WriteString(fmt.Sprintf("package %s\n", pkgName))
	FirstUpperFileName := strtool.FirstUpper(fileNameOnly)
	outFile.WriteString(fmt.Sprintf("type %s struct {\n", FirstUpperFileName))
	for i := 0; i < invalidCol; i++ {
		outFile.WriteString(fmt.Sprintf("\t%s\t\t%s\t\t//%s\n", strtool.FirstUpper(valnames[i]), typeMap[valtypes[i]], valcomments[i]))
	}
	outFile.WriteString("}\n")

	mapName = fmt.Sprintf("Map%s", FirstUpperFileName)

	outFile.WriteString(fmt.Sprintf("type %s map[int64]%s\n", mapName, FirstUpperFileName))

	outFile.WriteString(fmt.Sprintf("func Create%sTable() *%s {\n", FirstUpperFileName, mapName))

	outFile.WriteString(fmt.Sprintf("\tdata := %s{\n", mapName))
	//continue read csv and fill data
	func() {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println("Error:", err)
				return
			}
			outFile.WriteString("\t\t")
			ingoreLine := false
			for i := 0; i < invalidCol; i++ {
				if _, ok := emptyCol[i]; ok {
					continue
				}
				col := record[i]
				//col := strtool.ConvertByte2String([]byte(record[i]), "GB18030")
				if i == 0 && col[0] == '#' {
					ingoreLine = true
					break
				}
				if i == 0 {
					outFile.WriteString(fmt.Sprintf("%s:%s{\n", col, FirstUpperFileName))
				}
				if valnames[i] != "" {
					outFile.WriteString(fmt.Sprintf("\t\t\t%s,\n", parseCol(valnames[i], valtypes[i], col)))
				}
			}
			if ingoreLine {
				continue
			}
			outFile.WriteString("\t\t},\n")
		}
	}()
	outFile.WriteString("\t}\n")
	outFile.WriteString("\treturn &data\n")
	outFile.WriteString("}\n")
	return FirstUpperFileName, mapName, nil
}
