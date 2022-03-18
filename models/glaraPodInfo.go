package models

import "fmt"

type GlaraPodInfo struct {
	PodName        string
	PodLog         string
	OwnerReference string
}

type GlaraPodInfoList struct {
	InfoList []GlaraPodInfo
}

type GlaraPodInfoStack []GlaraPodInfo

func NewGlaraPodInfoStack() *GlaraPodInfoStack {
	return &GlaraPodInfoStack{}
}

func (g *GlaraPodInfoStack) IsEmpty() bool {
	return len(*g) == 0
}

//Push - 스택에 값을 추가하는 함수.
func (s *GlaraPodInfoStack) Push(data GlaraPodInfo) {
	*s = append(*s, data) // 스택 끝(top)에 값을 추가함.
	// fmt.Printf("%s %s Pushed to stack\n", data.PodName, data.PodLog)
}

//Pop - 스택에 값을 제거하고 top위치에 값을 반환하는 함수.
func (s *GlaraPodInfoStack) Pop() GlaraPodInfo {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return GlaraPodInfo{}
	} else {
		top := len(*s) - 1
		data := (*s)[top] // top 위치에 있는 값을 가져 옴
		*s = (*s)[:top]   // 스택에 마지막 데이터 제거함
		return data
	}
}
