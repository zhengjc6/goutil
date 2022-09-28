package csvparse

import (
	"encoding/csv"
	"fmt"
	"goutil/strtool"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var typeMap map[string]string

func init() {
	typeMap = map[string]string{
		"int":      "int64",
		"string":   "string",
		"float":    "float64",
		"[]int":    "[]int64",
		"[]string": "[]string",
		"[]float":  "[]float64",
	}
}

func parseCol(colName, colType, colStr string) string {
	switch colType {
	case "int", "float":
		if colStr == "" {
			colStr = "0"
		}
		return fmt.Sprintf("%s:%s", colName, colStr)
	case "string":
		return fmt.Sprintf("%s:\"%s\"", colName, colStr)
	case "[]int", "[]float":
		rexp := regexp.MustCompile("[|]")
		colStr = rexp.ReplaceAllLiteralString(colStr, ",")
		return fmt.Sprintf("%s:%s{%s}", colName, typeMap[colType], colStr)
	case "[]string":
		if colStr == "" {
			return fmt.Sprintf("%s:%s{}", colName, typeMap[colType])
		}
		rexp := regexp.MustCompile("[|]")
		colStr = rexp.ReplaceAllString(colStr, "\",\"")
		return fmt.Sprintf("%s:%s{\"%s\"}", colName, typeMap[colType], colStr)
	default:
		panic("Invalid Column Data " + colStr)
	}
}

func parseFile(filePath, outputDir, pkgName string) (fileNameOnly string, mapName string, err error) {
	fmt.Printf("fullPath = %s\n", filePath)

	fileNameWithSuffix := path.Base(filepath.ToSlash(filePath))

	fileType := path.Ext(fileNameWithSuffix)

	fileNameOnly = strings.TrimSuffix(fileNameWithSuffix, fileType)
	fmt.Printf("fileNameWithSuffix = %s;\nfileType = %s;\nfileNameOnly = %s;\n",
		fileNameWithSuffix, fileType, fileNameOnly)

	file, err := os.Open(filePath)
	if err != nil {
		return fileNameOnly, mapName, err
	}
	defer file.Close()
	reader := csv.NewReader(file)

	var valcomments []string
	var valnames []string
	var valtypes []string
	var ingoreLine bool
	var lineNum int
	var optFlag int
	for fixInfo := 0; fixInfo < 7; fixInfo |= optFlag {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return fileNameOnly, mapName, err
		}
		lineNum++
		optFlag = 0
		for i := 0; i < len(record) && fixInfo <= 3; i++ {
			ingoreLine = false
			col := strtool.ConvertByte2String([]byte(record[i]), "GB18030")
			if i == 0 && col[0] == '#' {
				ingoreLine = true
				break
			}
			col = strings.ToLower(strings.TrimSpace(col))
			fmt.Printf("%s,", col)

			switch fixInfo {
			case 0:
				valcomments = append(valcomments, col)
				optFlag = 1
			case 1:
				valnames = append(valnames, col)
				optFlag = 2
			case 3:
				if _, ok := typeMap[col]; ok {
					valtypes = append(valtypes, col)
				} else {
					return fileNameOnly, mapName, fmt.Errorf("error info: row = %d,col = %d,rawstring = %v", lineNum, i, record[i])
				}
				optFlag = 4
			}
		}
		fmt.Printf("\n")
		if ingoreLine {
			continue
		}
	}
	if len(valcomments) != len(valnames) || len(valnames) != len(valtypes) {
		return fileNameOnly, mapName, fmt.Errorf("lack comment or name or type")
	}

	outFile, err := os.OpenFile(outputDir+"\\"+strings.ToLower(fileNameOnly)+".go", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0744)
	if err != nil {
		return fileNameOnly, mapName, err
	}
	defer outFile.Close()
	//create struct
	outFile.WriteString(fmt.Sprintf("package %s\n", pkgName))
	outFile.WriteString(fmt.Sprintf("type %s struct {\n", fileNameOnly))
	for i := 0; i < len(valcomments); i++ {
		outFile.WriteString(fmt.Sprintf("\t%s\t\t%s\t\t//%s\n", valnames[i], typeMap[valtypes[i]], valcomments[i]))
	}
	outFile.WriteString("}\n")

	mapName = fmt.Sprintf("map%s", fileNameOnly)

	outFile.WriteString(fmt.Sprintf("type %s map[int64]%s\n", mapName, fileNameOnly))

	outFile.WriteString(fmt.Sprintf("func Create%sTable() %s {\n", strtool.FirstUpper(fileNameOnly), mapName))

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
			for i := 0; i < len(record); i++ {
				ingoreLine = false
				col := strtool.ConvertByte2String([]byte(record[i]), "GB18030")
				if i == 0 && col[0] == '#' {
					ingoreLine = true
					break
				}
				if i == 0 {
					outFile.WriteString(fmt.Sprintf("%s:%s{\n", col, fileNameOnly))
				}
				outFile.WriteString(fmt.Sprintf("\t\t\t%s,\n", parseCol(valnames[i], valtypes[i], col)))
			}
			if ingoreLine {
				continue
			}
			outFile.WriteString("\t\t},\n")
		}
	}()
	outFile.WriteString("\t}\n")
	outFile.WriteString("\treturn data\n")
	outFile.WriteString("}\n")
	return fileNameOnly, mapName, nil
}
