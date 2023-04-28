package csvparse

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"metaserver/pkg/strtool"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/tealeg/xlsx"
)

func readExcelFile(path string) (f *xlsx.File) {
	f, err := xlsx.OpenFile(path)
	if err != nil {
		fmt.Printf("excel 文件读取错误 %s\n", path)
		panic(err)
	}
	return f
}

func isEmptyLine(row int, sheet *xlsx.Sheet) bool {
	for col := 0; col < sheet.MaxCol; col++ {
		cell := sheet.Cell(row, col)
		val := cell.Value
		if len(val) != 0 {
			return false
		}
	}
	return true
}

func getHeadInfo(filename string, reader *csv.Reader) (valcomments, valnames, valtypes []string, emptyCol map[int]bool) {
	emptyCol = make(map[int]bool)
	linenumber := 0
	for i := 0; i < 3; i++ {
		linenumber++
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		for j := 0; j < len(record); j++ {
			col := record[j]
			col = strings.ToLower(strings.TrimSpace(col))
			fmt.Printf("%s,", col)

			switch i {
			case 0:
				valcomments = append(valcomments, col)
			case 1:
				valnames = append(valnames, strtool.FirstUpper(col))
				if len(col) == 0 {
					emptyCol[j] = true
				}
			case 2:
				if len(col) == 0 {
					emptyCol[j] = true
					valtypes = append(valtypes, col)
				} else if _, ok := typeMap[col]; ok {
					valtypes = append(valtypes, col)
				} else {
					panic(filename + "," + col + " not define")
				}
			}
		}
		fmt.Printf("\n")
	}
	if linenumber < 3 {
		panic(filename + "at least 3 rows")
	}
	return
}

func parseSheetToCSV(sheet *xlsx.Sheet, toFile string) (err error) {
	b := &bytes.Buffer{}
	rows := sheet.MaxRow
	pattern := `,`
	reg := regexp.MustCompile(pattern)
	keys := make(map[string]bool)
	for i := 0; i < rows; i++ {
		if isEmptyLine(i, sheet) {
			continue
		}
		cols := sheet.MaxCol
		for j := 0; j < cols; j++ {
			cell := sheet.Cell(i, j)
			val := cell.Value
			if j == 0 {
				if _, ok := keys[val]; ok {
					return errors.New("duplicate kes:" + val)
				} else {
					keys[val] = true
				}
			}
			if reg.MatchString(val) {
				return errors.New("invalid char `,`")
			}
			val = strings.Replace(val, "\n", "-", -1)
			//fmt.Println(val)
			b.WriteString(val)
			if j < cols-1 {
				b.WriteString(",")
			}
		}
		b.WriteString("\r\n")
	}
	// 写入数据到文件
	err = os.WriteFile(toFile, b.Bytes(), os.ModePerm)
	return
}

func excelToCSV(filePath, outDir string) {
	file := readExcelFile(filePath)
	_, _, fileNameOnly := parseFileName(filePath)
	outFile := fmt.Sprintf("%s/%s.csv", outDir, fileNameOnly)
	err := parseSheetToCSV(file.Sheets[0], outFile)
	if err != nil {
		println(filePath)
		panic(err)
	}
	fmt.Printf("Excel %s 转换 CSV 成功\n", filePath)
}

func ExcelToCSVAll(inputDir, outputDir string) {
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
		//the dir must be empty,otherwise it needs to be emptied manually
		dir, _ := os.ReadDir(outputDir)
		if len(dir) > 0 {
			panic(fmt.Errorf("please clean the output dir:%s", outputDir))
		}
	}
	var files []string

	err = filepath.Walk(inputDir, visit(&files, &map[string]bool{".xlsx": true}))

	if err != nil {
		panic(err)
	}
	fmt.Printf("file number = %d\n", len(files))
	for _, file := range files {
		excelToCSV(file, outputDir)
	}
}
