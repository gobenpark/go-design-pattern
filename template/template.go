package template

import "strings"

type MessageRetriever interface {
	Message() string
}

type Template interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}



type TemplateImpl struct {}

func (t *TemplateImpl) first() string{
	return "hello"
}

func (t *TemplateImpl) third() string {
	return "template"
}

func (t *TemplateImpl) ExecuteAlgorithm(m MessageRetriever) string{
	return strings.Join([]string{t.first(),m.Message(),t.third()}," ")
}

type TestStruct struct {
	Template
}

func (m *TestStruct) Message() string{
	return "world"
}


//익명함수를 이용하였지만 인터페이스의 변경이 없었다 이를 해결하려면 ?? 어댑터 패턴

type AnonymousTemplate struct {
}

func (a *AnonymousTemplate) first() string {
	return "hello"
}

func (a *AnonymousTemplate) third() string {
	return "template"
}

func (a *AnonymousTemplate) ExecuteAlgorithm(f func() string) string {
	return strings.Join([]string{a.first(),f(),a.third()}," ")
}


// adapter 패턴 적용


func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &adapter{myFunc: f}
}


type adapter struct {
	myFunc func() string
}

func (a *adapter) Message() string {
	if a.myFunc != nil {
		return a.myFunc()
	}

	return ""
}


// ============ in Go Source Code

type MyList []int

func (m MyList) Len() int {
	return len(m)
}

func (m MyList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MyList) Less(i, j int) bool {
	return m[i] < m[j]
}

var myList MyList = []int{6,4,2,8,1}

fmt.Println(myList)
sort.Sort(myList)
fmt.Println(myList)

