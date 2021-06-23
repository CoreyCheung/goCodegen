package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"codegen/model"

	"github.com/cbroglie/mustache"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func main() {
	// 1.读取配置文件
	viper.SetDefault("database.db_host", "root:@") // 默认连接127.0.0.1:3306数据库root用户无密码
	viper.SetDefault("database.db_name", "test")
	viper.SetConfigName("codegen")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	dbHost := viper.GetString("database.db_host")
	schema := viper.GetString("database.db_name")
	jobs := viper.GetStringSlice("codegen.jobs")
	if nil == jobs || len(jobs) == 0 {
		panic("Empty job list")
	}

	// 2.读取Mysql表结构
	db, err := gorm.Open("mysql", dbHost+"/information_schema?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	if err = db.DB().Ping(); err != nil {
		fmt.Println(err)
		return
	}

	db.LogMode(true)
	//db.DB().SetMaxIdleConns(10)
	//db.DB().SetMaxOpenConns(500)
	//db.DB().SetConnMaxLifetime(60000)

	var tables []model.Table
	db.Find(&tables, &model.Table{Schema: schema})

	var columns []model.Column
	db.Find(&columns, &model.Column{Schema: schema})

	for _, confPrefix := range jobs {
		fmt.Println("我进来了")
		render(tables, columns, confPrefix)
	}
}

func render(tables []model.Table, columns []model.Column, confPrefix string) {
	removeTablePrefix := viper.GetStringSlice(confPrefix + ".remove_table_prefix")
	excludeTables := viper.GetStringSlice(confPrefix + ".exclude_tables")
	excludeCols := viper.GetStringSlice(confPrefix + ".exclude_cols")
	packageName := viper.GetString(confPrefix + ".package_name")
	commonPackagePath := viper.GetString(confPrefix + ".common_package_path")
	refPackageName := viper.GetString(confPrefix + ".ref_package_name")
	tplPath := viper.GetString(confPrefix + ".tpl_path")
	outputPath := viper.GetString(confPrefix + ".output_path")
	filePrefix := viper.GetString(confPrefix + ".file_prefix")
	filePostfix := viper.GetString(confPrefix + ".file_postfix")

	var tableArr []model.Table
	for _, table := range tables {
		flag := true
		if nil != excludeTables && len(excludeTables) > 0 {
			for _, p := range excludeTables {
				if p == table.Name {
					flag = false
					break
				}
			}
		}
		if flag {
			tableArr = append(tableArr, table)
		}
	}

	var colArr []model.Column
	for _, col := range columns {
		flag := true
		if nil != excludeCols && len(excludeCols) > 0 {
			for _, p := range excludeCols {
				if p == col.ColumnName {
					flag = false
					break
				}
			}
		}
		if flag {
			colArr = append(colArr, col)
		}
	}

	mapperColumns := func(table model.Table, columns []model.Column) []model.Column {
		var cols []model.Column
		for i := 0; i < len(columns); i++ {
			col := columns[i]
			if table.Name == col.Name {
				col.ColumnComment = strings.Replace(col.ColumnComment, "\n", " ", -1)
				col.TableCollation = table.Collation
				col.RefPackageName = table.RefPackageName
				cols = append(cols, col)
			}
		}
		return cols
	}

	// 3.渲染模板，生成代码，输出文件
	fmt.Println("tplPath:", tplPath)
	tpl, _ := mustache.ParseFile(tplPath)
	fmt.Println("要生成代码了")
	os.MkdirAll(outputPath, os.ModePerm)
	for i := 0; i < len(tableArr); i++ {
		table := tableArr[i]
		table.RemoveTablePrefix = removeTablePrefix
		table.CommonPackagePath = commonPackagePath
		table.RefPackageName = refPackageName
		table.Comment = strings.Replace(table.Comment, "\n", " ", -1)

		table.Columns = mapperColumns(table, colArr)
		hasDatetime, hasTime, hasNumPKOrLong, hasProgramKey, programKeyName := table.LoopColumns()
		fmt.Println("programkeyName:=======", programKeyName)
		table.HasDatetimeColumn = hasDatetime
		table.HasTimeColumn = hasTime
		table.HasNumPKOrLong = hasNumPKOrLong
		table.HasProgramKey = hasProgramKey
		table.ProgramKeyName = programKeyName
		table.PackageName = packageName
		renderAndWriteFile(tpl, &table, outputPath, filePrefix+table.OutputFileName(filePostfix)+filePostfix)
	}
}

func renderAndWriteFile(tpl *mustache.Template, table *model.Table, outputPath string, fileName string) {
	if len(outputPath) > 0 && tpl != nil {
		table.Time = time.Now()
		modelStr, err := tpl.Render(table)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Printf("%s Model:\n%s\n", table.Name, modelStr)
		f, err := os.Create(outputPath + fileName)
		defer f.Close()
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.WriteString(modelStr)
		if err != nil {
			fmt.Println(err)
		}
	}
}
