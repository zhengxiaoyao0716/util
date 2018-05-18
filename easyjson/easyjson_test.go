package easyjson

import (
	"fmt"
	"testing"
)

func TestJSON(*testing.T) {
	fmt.Println(ObjectOf(MustParse(" {} ")))
	fmt.Println(ArrayOf(MustParse(" [] ")))
	fmt.Println(StringOf(MustParse(" \"\" ")))
	fmt.Println(BooleanOf(MustParse(" true ")))
	fmt.Println(NumberOf(MustParse(" 0 ")))
	_, err := Parse("qwe")
	fmt.Printf("%#v\n", err)
	_, err = NumberOf(MustParse("{}"))
	fmt.Printf("%#v\n", err)

	jsonObject := MustObjectOf(MustParse("{\"name\": 123}"))
	jsonArray := MustArrayOf(MustParse("[123]"))
	fmt.Println(jsonObject.ValueAt("name"))
	fmt.Println(jsonArray.ValueAt(0))
	fmt.Println(jsonObject.ValueAt("none"))
	fmt.Println(jsonArray.ValueAt(1, 0))
	fmt.Println(jsonObject.NumberAt("name"))
	fmt.Println(jsonArray.StringAt(0))
	fmt.Println(jsonObject.NumberAt("none", 0))
	fmt.Println(jsonArray.StringAt(1, 0))
}
