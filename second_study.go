package main

import "fmt"

//슬라이스 복습
func slice() {
	//슬라이스에는 길이 개념과 용량 개념이 따로다.
	//make함수에서 용량(10 자리)을 생략하면 cap == len인 슬라이스를 만든다.
	s := make([]int, 5, 10)
	fmt.Println(len(s), cap(s))

	//부분 슬라이스: 슬라이스에서 일부를 발췌한 것
	s1 := []int{0, 1, 2, 3, 4, 5}
	s2 = s1[2:5]
	fmt.Println

	//append
	a := []int{1, 2, 3, 4}
	b := append(a, 5, 6) //a 배열에 5, 6을 붙여서 연결
	c := apend(a, b...)  //...을 쓰면 slice끼리 이어붙일 수 있다.

	//append 주의 사항
	p := make([]int, 3, 4) // len: 3 cap: 4인 slice 생성
	p[0] = 10
	p[1] = 20
	p[2] = 30

	q := append(p, 40)  // a에 여분의 용량이 남으므로 내부배열 공유
	r := append(p, 50)  // a에 여분의 용량이 남으므로 내부배열 공유
	ss := append(r, 60) // c에 여분의 용량이 남지 않으므로 새로운 내부배열 할당

	//여기서, b와 c가 a를 같이 공유하므로, c 생성시 값이 덮어씌어짐.
	//그래서 b와 c가 같은 값들을 가짐.

	fmt.Println(p, len(p), cap(p))    // [10 20 30] 3 4
	fmt.Println(q, len(q), cap(q))    // [10 20 30 50] 4 4
	fmt.Println(r, len(r), cap(r))    // [10 20 30 50] 4 4
	fmt.Println(ss, len(ss), cap(ss)) // [10 20 30 50 60] 5 8
}

//map: 해시 테이블 (key-value)
//var 변수명 map[Key타입]Value타입
func map_practice() {
	var id map[int]string //nilmap이다.
	//nilmap은 값을 쓸 수 없다. make함수로 초기화를 하고 써야함
	id = make(map[int]string)
	id[1] = "sadf" //이런식으로 초기화가 가능하다.

	//리터럴 초기화
	trick := map[int]string{
		1: "Sfsd",
		2: "fs",
		3: "sdfs",
	}

	no := trcik[1000] //초기화 되지 않은 인덱스는 nil/0을 리턴
	//특정 키 삭제시
	delete(trick, 2)

	//키 체크하기
	//맵을 참조하게 되면 첫번째 변수에는 값을, 두번쨰 변수에는 그 키에 값이 있는지에 대한 참 거짓을 리턴한다.
	val, exist := trick[2]
	if !exist {
		fmt.Println("none")
	}

	//for 루프에서 쓸때 range를 쓰면 자동으로 키와 값을 리턴해준다.
	for key, val := range id {
		fmt.Println(key, val)
	}
}

//package 관련 설명: http://golang.site/go/article/15-Go-%ED%8C%A8%ED%82%A4%EC%A7%80
//읽어볼 것

//구조체 개념이나 성질은 C와 같다.
//정의 방법.
type fuck struct{
	name string
	age int
}
func use_struct(){
	//구조체 객체 생성하기
	f:=fuck{}

	f.name = "Qrew"
	f.age = 10

	fmt.Println(f)

	//구조체 객체 초기화 방법
	f1 := fuck{"sfdsf", 123}
	f2 := fuck{name:"Sdfs", age:111}

	//구조체 객체 생성2
	f3 := new(fuck)  //new를 이용헌 생성
	// 포인터를 이용한 생성이지만, C에서 처럼 멤버 접근을 위해 ->를 쓸 필요는 없다.



}

//구조체 생성자 -> c++과 같은 개념
//생성과 동시에 초기화가 필요한 경우
type dict struct{
	data map[int]string  //초기화 후에만 사용가능
}
func newDict() *dict{   //생성자. 
	d:=dict{}
	d.data = map[int]string{}

	return &d
}

func main() {

}
