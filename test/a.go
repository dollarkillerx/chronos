package main

import (
	"fmt"
	"log"
	"reflect"
	"time"
)

type Student struct {
	Id   int
	Name string
}

func (s Student) Hello() {
	fmt.Println("我是一个学生")
}

func test1(b interface{}) {
	//s := Student{Id: 1, Name: "咖啡色的羊驼"}
	//var b interface{}
	// 获取目标对象
	t := reflect.TypeOf(b)
	log.Println(t.Kind())
	// .Name()可以获取去这个类型的名称
	fmt.Println("这个类型的名称是:", t.Name())

	// 获取目标对象的值类型
	v := reflect.ValueOf(b)
	// .NumField()来获取其包含的字段的总数
	for i := 0; i < t.NumField(); i++ {
		// 从0开始获取Student所包含的key
		key := t.Field(i)

		// 通过interface方法来获取key所对应的值
		value := v.Field(i).Interface()

		fmt.Printf("第%d个字段是：%s:%v = %v \n", i+1, key.Name, key.Type, value)
	}

	// 通过.NumMethod()来获取Student里头的方法
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("第%d个方法是：%s:%v\n", i+1, m.Name, m.Type)
	}
}

type ABACEnforce struct {
	Name  string
	Time  int64
	Count int64
}

func main() {
	a := ABACEnforce{
		Name:  "CPX",
		Time:  time.Now().UnixNano(),
		Count: 100,
	}

	val, err := GetStructVal(a, "Name", "Time", "Count")
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(val)

	//test1(a)
}

type StructItem struct {
	Data interface{}
	Type string
}

func GetStructVal(st interface{}, field ...string) (result map[string]StructItem, err error) {
	result = map[string]StructItem{}

	t := reflect.TypeOf(st)
	exs := sliceToSet(field)
	if t.Kind().String() != "struct" {
		return nil, fmt.Errorf("Not Struct")
	}

	v := reflect.ValueOf(st)
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i)
		value := v.Field(i).Interface()

		_, ex := exs[key.Name]
		if ex {
			result[key.Name] = StructItem{
				Data: value,
				Type: key.Type.String(),
			}
		}
	}

	return result, nil
}

func sliceToSet(field []string) map[string]bool {
	output := map[string]bool{}
	for _, v := range field {
		output[v] = true
	}

	return output
}
