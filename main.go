package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	xmlFile, err := os.Open("xml/ctaf021_fmb.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	var module Module
	xml.Unmarshal(byteValue, &module)

	form := module.FormModule

	fmt.Println(form.Name)
	for _, formTrigger := range form.Trigger {
		triggerText := strings.ReplaceAll(strings.TrimSpace(formTrigger.TriggerText), "&#10;", "\n")
		if triggerText != "" {
			color.Red("Form Trigger: Name=%s", formTrigger.Name)
			color.Cyan("Form Trigger: TriggerText=\n%s\n\n", triggerText)
		}
	}
	for i := 0; i < len(form.Block); i++ {
		if form.Block[i].QueryDataSourceName != "" {
			color.Blue("Block: %s\n", form.Block[i].Name)
			color.Blue("Query Data Source: %s\n", form.Block[i].QueryDataSourceName)
			whereClause := strings.ReplaceAll(form.Block[i].WhereClause, "&#10;", "\n")
			color.Yellow("Where clause: \n%s\n", whereClause)

			for _, item := range form.Block[i].Item {
				if item.ColumnName != "" {
					color.Green("Column: %s \tPrompt: %s \tAttribute:%s",
						item.ColumnName,
						strings.ReplaceAll(item.Prompt, "&#10;", ""), item.ParentName)
				}
				for _, tg := range item.Trigger {
					color.Cyan("-Trigger: %s", tg.Name)
					color.Cyan("-TriggerText: \n%s\n\n", strings.ReplaceAll(tg.TriggerText, "&#10;", "\n"))
				}
			}
		}
	}
}
