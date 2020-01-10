package main

import "fmt"


//go 변수들은 초기화 안하면 0 value, bool은 false, string은 ""
//var a int = 10
var f float32 = 11.

var i, j, k int = 1, 2, 3

var b bool = false

//go는 자동 타입 추론 기능이 있다
var q = 1

//함수 내에서는 var을 쓰지 않고  a := 1로 변수 선언이 가능하다

//상수는 const로 선언. 상수도 자동 타입추론이 가능
const t = 1
const s = "fuck"

//const 키워드로 상수값들을 묶어서 저장할 수 있다.
const (
	qwer = iota //iota를 지정하면 qwer부터 0을 저장하고,
	asdf        //asdf부터 1씩 커진 값들을 저장한다
	zxcv
)

//문자열은 '' "" 로 표현한다
// ''은 엔터가 있어도 뉴라인으로 해석x, 그래서 여러줄 표현시 사용
//""은 엔터를 뉴라인으로 해석. 그러므로 여러줄에 걸쳐쓰려면 +를 써야한다

//포인터
var fuck = 10
var pp = &fuck

func jogun() {
	//if문 사용시 조건식은 반드시 부울식이어야한다.
	//숫자 같은 것들은 사용불가
	//else나 else if는 항상 그 조건에 붙여서 사용해야 한다.
	if k == 1 {
		println("one")
	} else if k == 2 {
		println("two")
	}

	//if는 조건식에 간단한 문장을 함께 넣어서 실행 할 수 있다
	//간단한 문장 == 변수선언문   switch, for 등에서도 사용가능
	//단 여기서 선언한 변수는 이 블럭에서만 사용가능
	if ea := i * 2; ea != 0 {
		println(ea)
	}

	//switch에 조건 변수가 없다!!
	//심지어 각 케이스에 조건 식을 쓸 수 있다.
	a := 78
	switch {
	case a >= 90:
		println("A")
	case a >= 80:
		println("B")
	}

	//타입 체킹도 가능하다
	/*	pack := 1
		switch pack.(type) {
		case int:
			println(int)
		case bool:
			println(bool)
		default:
			println(unknown)
		}*/
	//go에서 switch는 break가 있든 없든 모드 케이스에서 항상 break을 한다
	//그러나 fallthrough키워드를 사용하면 그 키워드 밑의 모든 케이스들을 수행하게 됨

	switch a {
	case 1:
		println("1")
		fallthrough
	case 2:
		println("2")
		fallthrough
	case 3:
		println("3")
		break
	default:
		println("sfjasl")
	}
	//입력
	var abc, def int
	fmt.Scan(&abc, &def)
}

func banbok() {
	//for문 기본   go에서는 for i:=1;i<10;i++{ 과 같이 중괄호를 붙여야함
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	//조건식만 있는 for == while
	n := 1
	for n < 100 {
		n *= 2
	}

	//무한 루프 만들기
	for {
		println("Infinite")
	}

	//for range 문
	//배열의 인덱스와 배열 값을 각각 얻어올수 있다.
	//컬렉션 == 배열
	//a := []int {1, 2, 3, 4}
	abcd := []int64{1, 2, 3, 4}
	for idx, val := range abcd {
		println(idx, val)
	}
ONE:
	println("one")

	//goto문 C랑 똑같음   break ONE하면 현재 제어를 ONE으로 넘긴다
	if abcd[1] == 1 {
		goto ONE
	}
}
//함수 기본형 ->  리턴이 여러개가 될 수 있다
func fun(p int) (int, string){
	var a string ="fuck"
	b:=1
	return b , a
}
func msg(a ...string){
	//가변 개수의 동일 타입 파라미터를 갖는 함수의 선언.
	//0~n개의 파라미터를 가진다. ->variadic function
}
//named return param 기능으로 함수짜기
func fun3 (nums ...int) (count int, total int) {
	for _, n := range nums {
        total += n
    }
    count = len(nums)
    return
}

//go에서는 pass by value, pass by reference라는 개념이 있는데, C랑 거의 같다

//익명함수: 함수 명을 갖지 않는 함수
//익명함수는 함수 전체를 변수에 할당하거나, 다른 함수의 파라미터에 직접 정의되어 사용된다.
//ex
func anony_fun(){
	sum:=func(n...int) int{
		s:=0
		for_, i := range n{
			s+=i
		}
		return s
	}
	res :=sum(1, 2, 3, 4, 5)
	println(res)
}

//일급함수 -> 함수는 go에서 기본타입과 동일시 됨
//=> 즉 한 함수는 다른 함수의 파라미터 혹은 리턴으로 사용할 수 있다.
//어지간하면 이렇게 안쓸듯...;; C에 이거 없어도 평생 잘씀
//어 근데 클로저 쓰려면 알아야 되네 하하

//type ==typedef와 같다고 보면 됨
//근데 go는 함수도 자료형이 될 수 있어서 아래처럼 쓸 수 있어
type cal func(int, int) int
func calculator(c cal, a int, b int)int{
	res:=c(a, b) //미리 정의한 함수 원형을 타 메소드에 전달하고 리턴받는 기능
	//Delegate라 한다.
	return res
}

//closure
//closure는 함수 밖의 변수를 참조하는 함수 값
//즉 바깥에 있는 변수를 마치 함수 안으로 끌어들여서 읽고 쓸 수 있다.
//ex 일급함수로 외부 변수 자기꺼 처럼 쓰기
func nextVal() func() int{
	i:=0   //외부 변수
	return func()int{   //일급함수
		i++
		return i     //클로저
	}
}
func nextVAL_exec(){
	next:=nextVal()
	//Go에서는 첫 문자가 대문자로 시작하면 그 패키지를 사용하는 곳에서 접근할 수 있는 exported name이 됨.(private, public)
	fmt.Println(next())  //1
	fmt.Println(next())  //2
	fmt.Println(next())  //3

	another :=nextVal()
	fmt.Println(another) //1
	fmt.Println(another) //2
}

//배열
func array(){
	var a [3]int    //null배열
	var a1 = [3]int{1, 2, 3}  //길이 한정한 배열초기화
	var a2 = [...]int{1, 2, 3}  //초기화 한 갯수만큼 배열크기 정해짐

	//다차원 배열
	var mult = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}	}
	var mult1 [3][4][5]int
	mult1[0][1][2]=100

}
//slice: 동적 배열 처럼 고정된 크기를 미리 지정하지 않아도 됨, 크기 조절도 가능
func slice(){
	var a []int  //슬라이스 변수 선언 방법.
	// 별도의 길이와 용량이 지정되지 않은 것은 nil이라는 키워드로 표현
	a = []int{1, 2, 3}

	b :=make([]int, 3, 4)  //길이3 용량4 슬라이스 변수 선언
	len(b)
	cap(b)   //capacity

}//으어어러어렁 뇌정지

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	a += b
	fmt.Println(a)

}
