package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	// 写文件
	f := excelize.NewFile()
	index := f.NewSheet("sheet2")
	f.SetCellValue("sheet2", "A2", "helloworld")
	f.SetCellValue("sheet1", "B2", 110)
	f.SetActiveSheet(index)
	if err := f.SaveAs("file.xlsx"); err != nil {
		fmt.Println(err)
	}

	//读文件
	f1, err := excelize.OpenFile("cw.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	cell := f1.GetCellValue("流水", "E3")
	fmt.Println(cell)
	rows := f1.GetRows("流水")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
