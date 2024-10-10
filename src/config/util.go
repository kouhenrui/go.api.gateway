package config

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"reflect"
	"regexp"
)

// ValidateExist 验证参数是否存在
func ValidateExist(a string, b []string) bool {
	for _, s := range b {
		if a == s {
			return true
		}
	}
	return false
}

// GetValidate 反射获取请求字段错误原因
func GetValidate(err error, obj any) error {

	var invalid *validator.InvalidValidationError
	ok := errors.As(err, &invalid)
	if ok {
		fmt.Println("param error:", invalid)
		return invalid
	}
	//反射获取标签的注释
	getObj := reflect.TypeOf(obj)
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		//return errs
		for _, e := range errs {
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				msg := f.Tag.Get("msg")
				return errors.New(msg)
			}
		}
	}
	return err
}

// FuzzyMatch 正则模糊匹配路径
func FuzzyMatch(param string) bool {
	for _, y := range ViperConfig.Service.WhiteUrl {
		if regexp.MustCompile(y).MatchString(param) {

			//fmt.Print("匹配道路进了")
			return true
		}

	}
	return false
}

// Paginate 分页组件
func Paginate(db *gorm.DB, page, pageSize int) *gorm.DB {
	offset := (page - 1) * page

	return db.Offset(offset).Limit(pageSize)
}
