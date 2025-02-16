package storage

import (
	"crawl-service/models"
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

var file *excelize.File
var rowIndex int = 2

// InitExcelFile Create an Excel file and write the title
func InitExcelFile() {
	file = excelize.NewFile()
	headers := []string{"Title", "Description", "Category", "Sub-category", "URL", "Published Date", "Image URL", "Content", "Hash"}

	for i, h := range headers {
		col := string(rune('A'+i)) + "1"
		file.SetCellValue("Sheet1", col, h)
	}
}

// Save data to Excel file
func SaveExcelFormat(article models.Article, url string) {
	row := fmt.Sprintf("A%d", rowIndex)
	file.SetCellValue("Sheet1", row, article.Title)
	file.SetCellValue("Sheet1", fmt.Sprintf("B%d", rowIndex), article.Description)
	file.SetCellValue("Sheet1", fmt.Sprintf("C%d", rowIndex), article.Category)
	file.SetCellValue("Sheet1", fmt.Sprintf("D%d", rowIndex), article.SubCategory)
	file.SetCellValue("Sheet1", fmt.Sprintf("E%d", rowIndex), url)
	file.SetCellValue("Sheet1", fmt.Sprintf("F%d", rowIndex), article.PublishedDate)
	file.SetCellValue("Sheet1", fmt.Sprintf("G%d", rowIndex), article.ImageURL)
	file.SetCellValue("Sheet1", fmt.Sprintf("H%d", rowIndex), article.Content)
	file.SetCellValue("Sheet1", fmt.Sprintf("I%d", rowIndex), article.Hash)

	rowIndex++
}

// SaveExcelFile Save the Excel file to the current directory
func SaveExcelFile(filename string) {
	err := file.SaveAs(filename)
	if err != nil {
		log.Fatal("Error when saving Excel file:", err)
	}
	fmt.Println("The Excel file has been saved:", filename)
}
