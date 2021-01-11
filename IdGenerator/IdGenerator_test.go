package IdGenerator

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	n := 0
	for n < 10 {
		id, err := RandomIdGenerator{}.Generate()
		if err != nil {
			t.Errorf("the error is: %v", err)
		}
		if id == "" {
			t.Errorf("id is empty: %v \n", id)
		}
		//t.Logf("the id is: %v", id)
		//fmt.Printf("the id is: %v.\n", id)
		n++
	}
	t.Deadline()
}

func TestGetLastSubstrSplittedByDot(t *testing.T) {
	random := RandomIdGenerator{}
	actualSubstr := random.GetLastSubstrSplittedByDot("field1.field2.field3")
	if actualSubstr != "field3" {
		t.Errorf("must be field3, %s.\v", actualSubstr)
	}
	actualSubstr = random.GetLastSubstrSplittedByDot("field1")
	if actualSubstr != "field1" {
		t.Errorf("must be field1, %s.\v", actualSubstr)
	}
	actualSubstr = random.GetLastSubstrSplittedByDot("field1#field2$field3")
	if actualSubstr != "field1#field2$field3" {
		t.Errorf("must be field1#field2$field3, %s.\v", actualSubstr)
	}
	actualSubstr = random.GetLastSubstrSplittedByDot("")
	if actualSubstr != "" {
		t.Errorf("must be \"\", %s.\v", actualSubstr)
	}
	t.Deadline()
}

func TestGenerateRandomAlphameric(t *testing.T) {
	random := RandomIdGenerator{}
	actualRandomString := random.GenerateRandomAlphameric(6)
	if actualRandomString == "" {
		t.Errorf("must be not empty, %s.\n", actualRandomString)
	}
	if len(actualRandomString) != 6 {
		t.Errorf("length must be 6, %d.\n", len(actualRandomString))
	}
	//for _, v := range actualRandomString {
	//	c := byte(v)
	//	if !('0' <= c && c <= '9') || !('a' <= c && c <= 'z') || !('A' <= c && c <= 'Z') {
	//		t.Errorf("must be in [0-9|a-z|A-Z], %c.\n", c)
	//		break
	//	}
	//}
	//actualRandomString = random.GenerateRandomAlphameric(0)
	//if actualRandomString == "" {
	//	t.Errorf("must be not empty, %s.\n", actualRandomString)
	//}
	//actualRandomString = random.GenerateRandomAlphameric(-1)
	//if actualRandomString == "" {
	//	t.Errorf("must be not empty, %s.\n", actualRandomString)
	//}
	t.Deadline()
}