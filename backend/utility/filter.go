package utility

import (
	"GoRsyncManager/models"
	"fmt"
	"strconv"
)

var PerPageRecord int = 7

func GetFileFilterCmd(index *models.Index) string {
	// path filter
	if index.Path == "" {
		index.Path = "/"
	}

	// pagination calculation
	if index.Page == 0 {
		index.Page = 1
	}
	f := PerPageRecord * index.Page
	t := (f - PerPageRecord) + 1

	// Generate command string
	cmdStr := fmt.Sprintf("ls %[1]s | wc -l;ls %[1]s | sed -n '%[2]d,%[3]dp'", index.Path, t, f)
	return cmdStr
}

func GetTotalPage(totalRecord string) string {
	t, _ := strconv.Atoi(totalRecord)
	totalPage := t / PerPageRecord
	if t%PerPageRecord != 0 {
		totalPage++
	}
	return fmt.Sprint(totalPage)
}
