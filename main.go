package main

import (
	"encoding/json"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type GroupsTree struct {
	Id         int    `json:"id"`
	GroupName  string `json:"groupName"`
	GroupLevel int    `json:"groupLevel"`
	EmpId      string `json:"emp_id"`
}

func write() {
	f := excelize.NewFile()

	index := f.NewSheet("Sheet1")

	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "GroupName")
	f.SetCellValue("Sheet1", "C1", "GroupLevel")
	f.SetCellValue("Sheet1", "D1", "EmpID")

	// set trang hoat donog
	f.SetActiveSheet(index)

	// save xlsx file by the given path
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		panic(err)
	}

}

func read() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		panic(err)
	}

	// get value from cell by given works sheet name and axis
	cell := f.GetCellValue("Sheet1", "B2")
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println(cell)

	//  get all value
	rows := f.GetRows("Sheet1")
	for _, colCell := range rows {
		fmt.Print(" \t", colCell)
	}
}

func main() {
	myStruct := GroupsTree{
		Id:         1,
		GroupName:  "G-r",
		GroupLevel: 0,
		EmpId:      "0",
	}

	structToExcel(myStruct)

	// write()
	// read()
}

func structToExcel(myStruct GroupsTree) {
	idByte, err := json.Marshal(myStruct.Id)
	if err != nil {
		panic(err)
	}
	GNByte, err := json.Marshal(myStruct.GroupName)
	if err != nil {
		panic(err)
	}

	GLByte, err := json.Marshal(myStruct.GroupLevel)
	if err != nil {
		panic(err)
	}

	empID, err := json.Marshal(myStruct.EmpId)
	if err != nil {
		panic(err)
	}

	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		panic(err)
	}

	f.SetCellValue("Sheet1", "A2", string(idByte))
	f.SetCellValue("Sheet1", "B2", string(GNByte))
	f.SetCellValue("Sheet1", "C2", string(GLByte))
	f.SetCellValue("Sheet1", "D2", string(empID))

	// save xlsx file by the given path
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		panic(err)
	}
}
