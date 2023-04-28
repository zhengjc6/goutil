package csvparse

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var luaTypeMap map[string]string

func init() {
	luaTypeMap = map[string]string{
		"int":        "integer",
		"list<int>":  "integer[]",
		"list<long>": "integer[]",
	}
}

func parseFile2Lua(filePath, outDir string) error {
	reg := regexp.MustCompile(`\[\]`)
	if reg == nil {
		return errors.New("regexp err")
	}
	_, _, fileNameOnly := parseFileName(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	outFile := fmt.Sprintf("%s/%s.lua", outDir, strings.ToLower(fileNameOnly))
	b := &bytes.Buffer{}
	reader := csv.NewReader(file)
	comments, names, types, emptyCol := getHeadInfo(fileNameOnly, reader)

	//写 emmylua注解
	b.WriteString(fmt.Sprintf("---@class %s_cfg\n", strings.ToLower(fileNameOnly)))
	for i := 0; i < len(names); i++ {
		if _, ok := emptyCol[i]; ok {
			continue
		}
		if val, ok := luaTypeMap[types[i]]; ok {
			b.WriteString(fmt.Sprintf("---@field public %s %s @%s\n", names[i], val, comments[i]))
		}
	}
	b.WriteString("\nlocal empty = {}\n\n")
	b.WriteString("local M = {\n")

	//数据
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		var needtail = true
		for j := 0; j < len(record); j++ {
			val := record[j]
			if j == 0 {
				if len(val) == 0 {
					return errors.New(filePath + " first column must be filled")
				}
				if val[0] == '#' {
					needtail = false
					break
				}
				b.WriteString(fmt.Sprintf("\t[%s] = {\n", val))
				b.WriteString(fmt.Sprintf("\t\t%s = %s,\n", names[j], val))
			} else {
				if _, ok := emptyCol[j]; ok {
					continue
				}
				luaType, ok := luaTypeMap[types[j]]
				if !ok {
					continue
				}
				isarrary := reg.Match([]byte(luaType))
				//println(val, isarrary, luaType)
				if len(val) == 0 {
					if isarrary {
						b.WriteString(fmt.Sprintf("\t\t%s = empty,\n", names[j]))
					} else {
						b.WriteString(fmt.Sprintf("\t\t%s = 0,\n", names[j]))
					}
				} else {
					if isarrary {
						val = strings.Replace(val, "|", ",", -1)
						b.WriteString(fmt.Sprintf("\t\t%s = {%s},\n", names[j], val))
					} else {
						b.WriteString(fmt.Sprintf("\t\t%s = %s,\n", names[j], val))
					}
				}
			}
		}
		if needtail {
			b.WriteString("\t},\n")
		}
	}

	b.WriteString("}\n")
	b.WriteString("return M\n")
	// 写入数据到文件
	return os.WriteFile(outFile, b.Bytes(), os.ModePerm)
}

func Excel2Lua(inputDir, outputDir string, forceClean bool) {
	if forceClean {
		err := os.RemoveAll(outputDir)
		if err != nil {
			panic(err)
		}
	}
	pwd, _ := os.Getwd()
	tmpCsvDir := pwd + "\\csvtmp"
	if err := os.RemoveAll(tmpCsvDir); err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	ExcelToCSVAll(inputDir, tmpCsvDir)
	_, err := os.Stat(outputDir)
	//check dir exists or not
	if err != nil {
		if os.IsNotExist(err) {
			// if dir does not exists then create
			if err = os.MkdirAll(outputDir, 0744); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		dir, _ := os.ReadDir(outputDir)
		if len(dir) > 0 {
			fmt.Printf("the output dir is not exmpty:%s\n", outputDir)
		}
	}
	var files []string

	err = filepath.Walk(tmpCsvDir, visit(&files, &map[string]bool{".csv": true}))

	if err != nil {
		panic(err)
	}
	fmt.Printf("file number = %d\n", len(files))
	for _, file := range files {
		err = parseFile2Lua(file, outputDir)
		if err != nil {
			panic(err)
		}
		fmt.Printf("CSV %s 转换 Lua 成功\n", file)
	}
}
