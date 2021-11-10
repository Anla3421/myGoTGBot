package rule

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

const (
	rule     = "rule"
	required = "required"
)

var timeKind reflect.Kind

func init() {
	timeKind = reflect.TypeOf(time.Time{}).Kind()
}

// 此為主要rule檢查的進入點 且因為後續valid的function中的input值 皆會轉為 reflet value的type
// 所以在此直接做轉換 方便做遞迴的使用
func Validate(input interface{}) (err error) {
	if timeKind == reflect.Invalid {
		timeKind = reflect.TypeOf(time.Time{}).Kind()
	}
	value := reflect.ValueOf(input)
	err = valid(value)
	return
}

//此為主要做判斷的部分
func valid(value reflect.Value) (err error) {
	valueType := value.Type()

	//此為判斷 是否能正確抓取到type
	if valueType == nil {
		err = errors.New("It is nil")
		return
	}

	// 抓到其種類 是否為pointer 或struct
	valueKind := valueType.Kind()

	// 如果pointer 則做遞迴
	if valueKind == reflect.Ptr {
		value = value.Elem()
		err = valid(value)
		return
	}

	// 如果其種類 非struct類型 則回傳正確 因為比如為string 則初始化時 則其"" 如果其為int 則初始化為0
	if valueKind != reflect.Struct {
		return
	}

	// 判斷其 是否為可用的值
	if !value.IsValid() {
		err = errors.New("the value is not valid")
		return
	}

	// 針對struct中的各個屬性 做loop
	for fieldIndex := 0; fieldIndex < valueType.NumField(); fieldIndex++ {

		// 抓取struct中的屬性
		field := valueType.Field(fieldIndex)

		// 抓取其屬性的struct 這邊作全部轉小寫 避免大小寫問題
		tagStr := strings.ToLower(field.Tag.Get(rule))

		// 確認其有定義相對應的屬性
		if tagStr == "" || tagStr == "-" {
			continue
		}

		// 針對未來此功能的擴充性 先預作用 , 當作不同功能的切割
		tags := strings.Split(tagStr, ",")

		for _, tag := range tags {
			switch tag {
			// 當其有required的tag 則
			case required:
				fieldValue := value.Field(fieldIndex)
				switch fieldKind := field.Type.Kind(); fieldKind {

				// 如果其為struct 則做遞迴
				case reflect.Struct:
					if intErr := valid(fieldValue); intErr != nil {
						err = intErr
						return
					}
				// 如果其為pointer 則判定其是否為nil 否則也做遞迴
				case reflect.Ptr:
					if fieldValue.IsNil() {
						err = errors.New(fmt.Sprintf("the field %s should not nil", field.Name))
						return
					}
					if intErr := valid(fieldValue.Elem()); intErr != nil {
						err = intErr
						return
					}
				// 如果其為string 則回傳錯誤
				case reflect.String:
					if strings.TrimSpace(fieldValue.String()) == "" {
						err = errors.New(fmt.Sprintf("the field %s should have string", field.Name))
						return
					}
				case timeKind:
					if fieldValue.Interface().(time.Time).IsZero() {
						err = errors.New(fmt.Sprintf("the field %s 's time is zero", field.Name))
						return
					}
				}
			}
		}
	}
	return
}
