package meta

type Structs []Struct

func ToStructs[S []T, T any](value S) Structs {
	var s Structs
	for _, v := range value {
		str, err := ToStruct(v)
		if err != nil {
			return nil
		}
		s = append(s, str)
	}
	return s
}

func (s Structs) TagName(keys ...string) []string {
	var names []string
	for _, str := range s {
		names = append(names, str.TagName(keys...))
	}
	return names
}

func (s Structs) Identifiers() []string {
	var identifiers []string
	for _, str := range s {
		identifiers = append(identifiers, str.Identifier())
	}
	return identifiers
}

func (s Structs) ToStructMap() map[string]Struct {
	structmap := map[string]Struct{}
	for _, str := range s {
		structmap[str.Name] = str
	}
	return structmap
}

// func ToStructsWithData[S []T, T any](value S) map[string]StructWithData {
// 	swd := map[string]StructWithData{}
// ToStructWithData(value)
// }

func DevTime[S []T, T any](value S) []Struct {
	str, err := ToStruct(value)
	if err != nil {
		return nil
	}

	childStructs := str.Fields().WithTagTrue("struct")
	dataMap := map[string]Data{
		str.Name: Data{},
	}
	for _, child := range childStructs {
		dataMap[child.Name] = Data{}
	}
	// for _, stri := range str.
	return nil
}

// func ExtractDataByName(str Struct, childNames ...string) []Data {
// 	data := map[string]Data{}
// 	str.Fields().ByNames(childNames...)

// }

func (s Structs) ExtractDataByName(names ...string) map[string]Data {
	data := map[string]Data{}
	for _, str := range s {
		strData := Data(ToSlice(str.Value.Interface()))
		data[str.Name] = append(data[str.Name], strData)
		for _, child := range str.Fields().ByNames(names...) {
			childData := Data(ToSlice(child.Value.Interface()))
			data[child.Name] = append(data[child.Name], childData...)
		}
	}
	return data
}
