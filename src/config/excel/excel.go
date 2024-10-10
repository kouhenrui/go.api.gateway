package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"unicode"
)

var f *excelize.File

func ExportExcel(sheetName string) {
	f = excelize.NewFile()
	// 假设这是你传入的 user 数据
	users := []map[string]string{
		{"name": "Alice", "age": "25", "city": "New York"},
		{"name": "Bob", "age": "30", "city": "Los Angeles"},
	}

	if len(users) > 0 {
		firstRaw := users[0]
		columns := make([]string, 0, len(firstRaw))
		// 动态生成列名
		//colIndex := 1
		for key, value := range firstRaw {
			colName := toAlphaString(value)
			_ = f.SetCellValue("Sheet1", fmt.Sprintf("%s1", colName), key)
			columns = append(columns, key)
			//colIndex++
		}
		// 动态写入数据
		for i, user := range users {
			for _, key := range columns {
				colName := toAlphaString(key)
				f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", colName, i+2), user[key])
			}
		}
	}

	sheet, _ := f.NewSheet(sheetName)
	f.SetActiveSheet(sheet)

	// 设置响应头以确保返回的文件作为附件下载
	//w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	//w.Header().Set("Content-Disposition", "attachment;filename=report.xlsx")
	//w.Header().Set("File-Name", "report.xlsx")
	//w.Header().Set("Content-Transfer-Encoding", "binary")
	// 将文件保存到 HTTP 响应中
	//if err := f.Write(w); err != nil {
	//	log.Println("Error writing excel file:", err)
	//	http.Error(w, "Unable to write excel file", http.StatusInternalServerError)
	//	return
	//}
}
func toAlphaString(in string) string {
	result := ""
	for _, i2 := range in {
		if unicode.IsLetter(i2) {
			result += string(i2)
		}
	}
	return result
}
