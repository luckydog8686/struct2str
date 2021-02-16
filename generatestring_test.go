package struct2str

import "testing"

type User struct {
	Name string
	Age int
	Id string
}

func TestGenerateString(t *testing.T) {
	u := User{"TangXiaodong", 100, "0000123"}
	str,err := GenerateString(&u)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)
}