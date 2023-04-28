package csvparse

import "os"

func Excel2Go(inputDir, outputDir, pkgname string, forceClean bool) {
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
	ParseDir(tmpCsvDir, outputDir, pkgname)
}
