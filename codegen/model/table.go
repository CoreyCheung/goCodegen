package model

import (
	"strings"
	"bytes"
	"regexp"
	"time"
)

type Table struct {
	Catalog       string `gorm:"column:TABLE_CATALOG;type:varchar(512)"` // 'def'
	Schema        string `gorm:"column:TABLE_SCHEMA;type:varchar(64)"`   // db name
	Name          string `gorm:"column:TABLE_NAME;type:varchar(64)"`     // table name
	Engine        string `gorm:"column:ENGINE;type:varchar(64)"`         // 'InnoDB','MyISAM', 'MEMORY', 'CSV', 'PERFORMANCE_SCHEMA'
	Comment       string `gorm:"column:TABLE_COMMENT;type:varchar(2048)"`
	AutoIncrement uint64 `gorm:"column:AUTO_INCREMENT;type:bigint(21) unsigned"` // if(auto inc)then(max inc id)else(null)
	Collation     string `gorm:"column:TABLE_COLLATION;type:varchar(32)"`        // 'utf8_general_ci'
	RowCount      uint64 `gorm:"column:TABLE_ROWS;type:bigint(21) unsigned"`
	DataLength    uint64 `gorm:"column:DATA_LENGTH;type:bigint(21) unsigned"` // whole table data length
	IndexLength   uint64 `gorm:"column:INDEX_LENGTH;type:bigint(21) unsigned"`

	Columns             []Column
	RemoveTablePrefix   []string  `gorm:"-"`
	PackageName         string    `gorm:"-"`
	CommonPackagePath   string    `gorm:"-"`
	RefPackageName      string    `gorm:"-"`
	HasDatetimeColumn   bool      `gorm:"-"`
	HasTimeColumn       bool      `gorm:"-"`
	HasNumPKOrLong      bool      `gorm:"-"`
	HasProgramKey       bool      `gorm:"-"`
	ProgramKeyName      string    `gorm:"-"`
	Time                time.Time `gorm:"-"`
}

func (Table) TableName() string {
	return "TABLES"
}

func (table Table) HasComment() bool {
	return len(strings.TrimSpace(table.Comment)) > 0
}

func (table Table) CleanName() string {
	cutName := table.Name
	if nil != table.RemoveTablePrefix && len(table.RemoveTablePrefix) > 0 {
		for _, p := range table.RemoveTablePrefix {
			if len(p) == 0 {
				continue
			}
			if p[len(p)-1:] != "_" {
				p = p + "_"
			}
			if strings.Index(cutName, p) == 0 {
				cutName = cutName[len(p):]
			}
		}
	}
	arr := strings.Split(cutName, "_")
	b := bytes.Buffer{}
	for _, ai := range arr {

		b.WriteString(strings.ToUpper(ai[0:1]))
		b.WriteString(ai[1:])
	}
	return b.String()
	//arr := strings.Split(table.Name, "_")
	//b := bytes.Buffer{}
	//for i, ai := range arr {
	//	if len(arr) == 1 || i > 0 {
	//		b.WriteString(strings.ToUpper(ai[0:1]))
	//		b.WriteString(ai[1:])
	//	}
	//}
	//return b.String()
}

func (table Table) ProgramKeyNameLower() string {
	return strings.TrimSpace(strings.ToLower(table.ProgramKeyName))
}

func (table Table) OutputFileName(filePostfix string) string {
	if strings.LastIndex(filePostfix, ".java") == len(filePostfix)-5 {
		return table.CleanName()
	} else {
		str := table.CleanName()
		r := regexp.MustCompile("([A-Z]+)")
		str = r.ReplaceAllString(str, "_$1")
		str = strings.Trim(str, "_")
		str = strings.Replace(str, "__", "_", -1)
		return strings.ToLower(str)
	}
}

func (table Table) LoopColumns() (hasDatetime bool, hasTime bool, hasNumPKOrLong bool, hasNumPK bool, numPKName string) {
	colLen := len(table.Columns)
	if colLen == 0 {
		return false, false, false, false, ""
	}

	upperNameMaxLen := 0
	goTypeMaxLen := 0
	datetimeFlag := false
	timeFlag := false
	numPkCount := 0
	var numPkName string
	for i := 0; i < colLen; i++ {
		table.Columns[i].UpperName = table.Columns[i].GetUpperName()
		if len(table.Columns[i].UpperName) > upperNameMaxLen {
			upperNameMaxLen = len(table.Columns[i].UpperName)
		}
		table.Columns[i].GoType = table.Columns[i].GetGoType()
		table.Columns[i].JavaType = table.Columns[i].GetJavaType()
		table.Columns[i].PythonType = table.Columns[i].GetPythonType()
		gtp := strings.TrimLeft(strings.TrimSpace(table.Columns[i].GoType), "*")
		if gtp == "date.Datetime" || gtp == "date.Date" || gtp == "date.Time" {
			datetimeFlag = true
		}
		if gtp == "time.Time" {
			timeFlag = true
		}
		if len(table.Columns[i].GoType) > goTypeMaxLen {
			goTypeMaxLen = len(table.Columns[i].GoType)
		}
		if strings.TrimSpace(table.Columns[i].GoType) == table.RefPackageName+".Long" {
			table.HasNumPKOrLong = true
		}
		if table.Columns[i].IsProgramKey() {
			numPkCount++
			numPkName = table.Columns[i].UpperName
		}
	}
	if numPkCount == 1 {
		table.HasProgramKey = true
		table.ProgramKeyName = numPkName
	} else if numPkCount > 0 {
		for i := 0; i < colLen; i++ {
			table.Columns[i].HasMultiKey = true
		}
	}
	maxLen := upperNameMaxLen
	if maxLen < goTypeMaxLen {
		maxLen = goTypeMaxLen
	}
	emptyArr := bytes.Buffer{}
	for i := 0; i < maxLen; i++ {
		emptyArr.WriteString(" ")
	}
	empty := emptyArr.String()
	for i := 0; i < colLen; i++ {
		tmp0 := bytes.Buffer{}
		tmp0.WriteString(table.Columns[i].UpperName)
		tmp0.WriteString(empty[0 : upperNameMaxLen-len(table.Columns[i].UpperName)])
		table.Columns[i].UpperName = tmp0.String()
		tmp1 := bytes.Buffer{}
		tmp1.WriteString(table.Columns[i].GoType)
		tmp1.WriteString(empty[0 : goTypeMaxLen-len(table.Columns[i].GoType)])
		table.Columns[i].GoType = tmp1.String()
	}
	return datetimeFlag, timeFlag, table.HasNumPKOrLong, table.HasProgramKey, table.ProgramKeyName
}
