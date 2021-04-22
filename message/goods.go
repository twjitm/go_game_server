package message

type DataArray []interface {
	add()
}

type DataList interface {
	add()
}

type DataFloat []float64

type Name string

func (name *Name) GetNameToEnglish() *Name {
	return name
}

type Goods struct {
	Id      int64
	Code    string
	Pay     float64
	PayNow  float64
	Name    string
	Title   string
	Type    string
	SubType string
	Desc    string
}
