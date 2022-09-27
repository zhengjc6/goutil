package csvparse

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

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

func ParseDir(dirPath, outputDir, pkgName string) {
	var files []string

	err := filepath.Walk(dirPath, visit(&files))

	if err != nil {
		panic(err)
	}
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err = os.MkdirAll(outputDir, 0744); os.IsExist(err) {
			panic(err)
		}
	} else {
		panic(err)
	}

	for _, file := range files {
		parseFile(file, outputDir, pkgName)
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

func parseFile(filePath, outputDir, pkgName string) {
	fmt.Printf("fullPath = %s\n", filePath)
	//获取文件名称带后缀
	fileNameWithSuffix := path.Base(filepath.ToSlash(filePath))
	//获取文件的后缀(文件类型)
	fileType := path.Ext(fileNameWithSuffix)
	//获取文件名称(不带后缀)
	fileNameOnly := strings.TrimSuffix(fileNameWithSuffix, fileType)
	fmt.Printf("fileNameWithSuffix = %s;\nfileType = %s;\nfileNameOnly = %s;\n",
		fileNameWithSuffix, fileType, fileNameOnly)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	//前三行格式固定
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
			fmt.Println("Error:", err)
			return
		}
		lineNum++
		optFlag = 0
		for i := 0; i < len(record) && fixInfo <= 3; i++ {
			ingoreLine = false
			col := ConvertByte2String([]byte(record[i]), "GB18030")
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
					fmt.Printf("Type Error row = %d,col = %d,rawstring = %v\n", lineNum, i, record[i])
					return
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
		fmt.Printf("Lack comment or name or type")
		return
	}

	outFile, err := os.OpenFile(outputDir+"\\"+strings.ToLower(pkgName)+".go", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0744)
	if err != nil {
		fmt.Println("OpenFils error", err.Error())
		return
	}
	structName := FirstUpper(fileNameOnly)
	//create struct
	outFile.WriteString(fmt.Sprintf("package %s\n", pkgName))
	outFile.WriteString(fmt.Sprintf("type %s struct {\n", structName))
	for i := 0; i < len(valcomments); i++ {
		outFile.WriteString(fmt.Sprintf("\t%s\t%s\n", valnames[i], typeMap[valtypes[i]]))
	}
	outFile.WriteString("}\n")

	mapName := fmt.Sprintf("map%s", structName)

	outFile.WriteString(fmt.Sprintf("type %s map[int64]%s\n", mapName, structName))

	outFile.WriteString(fmt.Sprintf("func Create%sTable() %s {\n", structName, mapName))

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
				col := ConvertByte2String([]byte(record[i]), "GB18030")
				if i == 0 && col[0] == '#' {
					ingoreLine = true
					break
				}
				if i == 0 {
					outFile.WriteString(fmt.Sprintf("%s:%s{\n", col, structName))
				} else {
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
	outFile.WriteString("\treturn data\n")
	outFile.WriteString("}\n")
}
