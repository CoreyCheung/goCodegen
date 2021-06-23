package model

import (
	"strings"
	"bytes"
	"fmt"
)

type Column struct {
	Catalog        string `gorm:"column:TABLE_CATALOG;type:varchar(512)"` // 'def'
	Schema         string `gorm:"column:TABLE_SCHEMA;type:varchar(64)"`   // db name
	Name           string `gorm:"column:TABLE_NAME;type:varchar(64)"`     // table name
	ColumnName     string `gorm:"column:COLUMN_NAME;type:varchar(64)"`
	ColumnDefault  string `gorm:"column:COLUMN_DEFAULT;type:longtext"`
	DataType       string `gorm:"column:DATA_TYPE;type:varchar(64)"` // 'varchar'
	ColumnType     string `gorm:"column:COLUMN_TYPE;type:longtext"`  // 'varchar(64)'
	ColumnComment  string `gorm:"column:COLUMN_COMMENT;type:varchar(1024)"`
	Length         uint64 `gorm:"column:CHARACTER_OCTET_LENGTH;type:bigint(21) unsigned"` // 192 if string type & if utf8 = 64 * 3
	Precision      uint64 `gorm:"column:NUMERIC_PRECISION;type:bigint(21) unsigned"`      // 21  if number type
	Scale          uint64 `gorm:"column:NUMERIC_SCALE;type:bigint(21) unsigned"`          // 0   if number type
	Charset        string `gorm:"column:CHARACTER_SET_NAME;type:varchar(32)"`             // 'utf8'
	Collation      string `gorm:"column:COLLATION_NAME;type:varchar(32)"`                 // 'utf8_general_ci'
	Nullable       string `gorm:"column:IS_NULLABLE;type:varchar(3)"`                     // 'YES' or 'NO'
	ColumnKey      string `gorm:"column:COLUMN_KEY;type:varchar(3)"`                      // 'PRI' or 'MUL'
	Extra          string `gorm:"column:EXTRA;type:varchar(30)"`                          // include auto_increment & etc.
	Position       uint64 `gorm:"column:ORDINAL_POSITION;type:bigint(21) unsigned"`
	UpperName      string `gorm:"-"`
	GoType         string `gorm:"-"`
	JavaType       string `gorm:"-"`
	PythonType     string `gorm:"-"`
	TableCollation string `gorm:"-"`
	RefPackageName string `gorm:"-"`
	HasMultiKey    bool   `gorm:"-"`
}

func (Column) TableName() string {
	return "COLUMNS"
}

func (column Column) IsNotNull() bool {
	return strings.EqualFold(column.Nullable, "NO")
}

func (column Column) IsPK() bool {
	return strings.EqualFold(column.ColumnKey, "PRI")
	// return strings.EqualFold(column.ColumnKey, "PRI") || (column.Position == 1 && strings.Contains(strings.ToLower(column.TableCollation), "utf8mb4"))
}

func (column Column) IsAutoIncrement() bool {
	return strings.Contains(strings.ToLower(column.Extra), "auto_increment")
}

func (column Column) IsProgramKey() bool {
	return column.IsPK() && !column.IsAutoIncrement() &&
		(strings.TrimSpace(column.GoType) == column.RefPackageName+".Long" ||
			strings.TrimSpace(column.GoType) == "int64" || strings.TrimSpace(column.GoType) == "uint64")
}

func (column Column) ColumnHasComment() bool {
	return len(strings.TrimSpace(column.ColumnComment)) > 0
}

func (column Column) HasLength() bool {
	return strings.TrimSpace(column.GoType) == "string" && column.Length > 0
}

func (column Column) GetNumericLength() string {
	if column.IsNotNum() || (column.Precision <= 0 && column.Scale <=0) {
		return ""
	} else if column.Scale > 0 {
		if column.Precision <= 0 {
			column.Precision = 0
		}
		return fmt.Sprintf("%d, %d", column.Precision, column.Scale)
	} else {
		return fmt.Sprintf("%d", column.Precision)
	}
}

func (column Column) GetTrimUpperColumnType() string {
	return strings.TrimSpace(strings.ToUpper(column.ColumnType))
}

func (column Column) GetTrimUpperName() string {
	return strings.TrimSpace(column.GetUpperName())
}

func (column Column) GetUpperName() string {
	if column.UpperName == "" {
		arr := strings.Split(column.ColumnName, "_")
		b := bytes.Buffer{}
		for _, ai := range arr {
			b.WriteString(strings.ToUpper(ai[0:1]))
			b.WriteString(ai[1:])
		}
		column.UpperName = b.String()
	}
	name := strings.ToLower(column.UpperName)
	if name == "id" {
		column.UpperName = "ID"
	}
	if name == "createtime" || name == "createdtime" ||
		name == "createon" || name == "createdon" ||
		name == "createat" || name == "createdat" {
		column.UpperName = "CreatedAt"
	}
	if name == "updatetime" || name == "updatedtime" ||
		name == "lastupdatetime" || name == "lastupdatedtime" ||
		name == "updateon" || name == "updatedon" ||
		name == "updateat" || name == "updatedat" {
		column.UpperName = "UpdatedAt"
	}
	if name == "deletetime" || name == "deletedtime" ||
		name == "removetime" || name == "removedtime" ||
		name == "deleteon" || name == "deletedon" ||
		name == "deleteat" || name == "deletedat" ||
		name == "removeon" || name == "removedon" ||
		name == "removeat" || name == "removedat" {
		column.UpperName = "DeletedAt"
	}
	return column.UpperName
}

func (column Column) LowerName() string {
	arr := strings.Split(column.ColumnName, "_")
	b := bytes.Buffer{}
	for i, ai := range arr {
		if i == 0 {
			b.WriteString(strings.ToLower(ai[0:1]))
		} else {
			b.WriteString(strings.ToUpper(ai[0:1]))
		}
		b.WriteString(ai[1:])
	}
	str := b.String()
	name := strings.ToLower(str)
	if name == "createtime" || name == "createdtime" ||
		name == "createon" || name == "createdon" ||
		name == "createat" || name == "createdat" {
		str = "createdAt"
	}
	if name == "updatetime" || name == "updatedtime" ||
		name == "lastupdatetime" || name == "lastupdatedtime" ||
		name == "updateon" || name == "updatedon" ||
		name == "updateat" || name == "updatedat" {
		str = "updatedAt"
	}
	if name == "deletetime" || name == "deletedtime" ||
		name == "removetime" || name == "removedtime" ||
		name == "deleteon" || name == "deletedon" ||
		name == "deleteat" || name == "deletedat" ||
		name == "removeon" || name == "removedon" ||
		name == "removeat" || name == "removedat" {
		str = "deletedAt"
	}
	return str
}

func (column Column) IsUpdated() bool {
	return column.LowerName() == "updatedAt"
}

func (column Column) GetPythonType() string {
	if column.PythonType == "" {
		switch strings.ToLower(column.DataType) {
		case "varchar", "nvarchar", "char", "nchar":
			column.PythonType = "String";
		case "text", "longtext":
			column.PythonType = "Text";
		case "bigint":
			column.PythonType = "BigInteger";
		case "int":
			column.PythonType = "Integer";
		case "smallint":
			column.PythonType = "SmallInteger";
		case "tinyint":
			if strings.Contains(strings.ToLower(column.ColumnType), "tinyint(1)") {
				column.PythonType = "Boolean"
			} else {
				column.PythonType = "SmallInteger"
			}
		case "json":
			column.PythonType = "JSON";
		case "bit":
			column.PythonType = "Boolean";
		case "real", "numeric", "smallmoney", "money", "decimal", "double":
			column.PythonType = "Numeric";
		case "float":
			column.PythonType = "Float";
		case "image", "binary", "varbinary":
			column.PythonType = "BINARY"
		case "smalldatetime", "datetime", "datetime2", "timestamp":
			column.PythonType = "DateTime"
		case "date":
			column.PythonType = "Date"
		case "time":
			column.PythonType = "Time"
		default:
			column.PythonType = "String"
		}
	}
	return column.PythonType
}

func (column Column) GetJavaType() string {
	if column.JavaType == "" {
		switch strings.ToLower(column.DataType) {
		case "varchar", "nvarchar", "text", "longtext", "char", "nchar":
			column.JavaType = "String";
		case "bigint":
			column.JavaType = "Long";
		case "int":
			column.JavaType = "Integer";
		case "smallint":
			column.JavaType = "Short";
		case "tinyint":
			if strings.Contains(strings.ToLower(column.ColumnType), "tinyint(1)") {
				column.JavaType = "Boolean"
			} else {
				column.JavaType = "Byte"
			}
		case "bit":
			column.JavaType = "Boolean";
		case "real", "numeric", "smallmoney", "money", "decimal", "double":
			column.JavaType = "Double";
		case "float":
			column.JavaType = "Float";
		case "image", "binary", "varbinary":
			column.JavaType = "Byte[]"
		case "smalldatetime", "datetime", "datetime2", "timestamp":
			column.JavaType = "LocalDateTime"
		case "date":
			column.JavaType = "LocalDate"
		case "time":
			column.JavaType = "LocalTime"
		default:
			column.JavaType = "String"
		}
	}
	return column.JavaType
}

func (column Column) GetGoType() string {
	if column.GoType == "" {
		unsigned := func(t string) string {
			if strings.Contains(strings.ToLower(column.ColumnType), "unsigned") {
				return "u" + t
			} else {
				return t
			}
		}
		switch strings.ToLower(column.DataType) {
		case "varchar", "nvarchar", "text", "longtext", "char", "nchar":
			column.GoType = "string";
		case "bigint":
			// column.GoType =  unsigned("int64");
			if strings.Contains(strings.ToLower(column.ColumnType), "unsigned") {
				return "uint64"
			} else {
				return column.RefPackageName + ".Long"
			}
		case "int":
			column.GoType = unsigned("int");
		case "smallint":
			column.GoType = unsigned("int16");
		case "tinyint":
			if strings.Contains(strings.ToLower(column.ColumnType), "tinyint(1)") {
				column.GoType = "bool"
			} else {
				column.GoType = "int8"
			}
		case "bit":
			column.GoType = "bool";
		case "real", "numeric", "smallmoney", "money", "decimal", "double":
			column.GoType = "float64";
		case "float":
			column.GoType = "float32";
		case "image", "binary", "varbinary":
			column.GoType = "[]byte"
		case "smalldatetime", "datetime", "datetime2", "timestamp":
			if strings.TrimSpace(column.UpperName) == "DeletedAt" {
				column.GoType = "*date.Datetime"
			} else {
				column.GoType = "date.Datetime"
			}
		case "date":
			if strings.TrimSpace(column.UpperName) == "DeletedAt" {
				column.GoType = "*date.Date"
			} else {
				column.GoType = "date.Date"
			}
		case "time":
			if strings.TrimSpace(column.UpperName) == "DeletedAt" {
				column.GoType = "*date.Time"
			} else {
				column.GoType = "date.Time"
			}
		default:
			column.GoType = "string"
		}
	}
	return column.GoType
}

var numTypes = []string{"int64", "int", "int16", "bool", "byte", "float64", "float32", "uint64", "uint", "uint16", "int8", "uint8", "[RefPackageName].Long"}

func (column Column) IsNotNum() bool {
	isNum := false
	for _, str := range numTypes {
		tmp := strings.Replace(str, "[RefPackageName]", column.RefPackageName, -1)
		if tmp == strings.TrimSpace(column.GoType) {
			isNum = true
			break
		}
	}
	return !isNum
}
