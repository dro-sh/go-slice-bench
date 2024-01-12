package main

import "strconv"

type NestedStr struct {
	Name string
	Age  int64
}

type Rule struct {
	Str          string
	Num          int64
	NestedStruct NestedStr
}

func appendSlice(count int) []Rule {
	rules := make([]Rule, 0, count)

	for i := 0; i < count; i++ {
		num := int64(i)
		strNum := strconv.FormatInt(num, 10)

		str := NestedStr{
			Name: "somebody" + strNum,
			Age:  num,
		}

		rule := Rule{
			Str:          "abc" + strNum,
			Num:          num,
			NestedStruct: str,
		}

		rules = append(rules, rule)
	}

	return rules
}

func insertIndexSlice(count int) []Rule {
	rules := make([]Rule, count)

	for i := 0; i < count; i++ {
		num := int64(i)
		strNum := strconv.FormatInt(num, 10)

		str := NestedStr{
			Name: "somebody" + strNum,
			Age:  num,
		}

		rule := Rule{
			Str:          "abc" + strNum,
			Num:          num,
			NestedStruct: str,
		}

		rules[i] = rule
	}

	return rules
}
