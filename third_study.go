package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

//구조체의 메서드
type rect struct {
	width, height int
}

//()부분을 리시버라 하여 어떤 구조체의 메소드인지를 결정한다.
func (r rect) area() int { //여기서 r은 입력파라미터에 불과.
	return r.width * r.height
}

//Value Receiver V.S. Pointer Receiver
//리시버는 구조체의 데이터를 복사하여 전달하는 value receiver와
//구조체의 포인터만을 전달하는 Pointer receiver로 구분된다.
//call by value 와 call by reference와 같은 개념으로 생각하면 된다.
//메소드내에서 필드 값 변경이 이루어질 때 value면 호출자의 데이터가 변경되지 않지만,
//pointer receiver라면 변경된다.
func (r *rect) area2() int {
	r.width++ //Pointer receiver이므로 호출자의 값의 변경이 이뤄진다.
	return r.width * r.height
}

//인터페이스: 구조체는 필드들의 집합. 인터페이스는 메서드들의 집합.
//정의 형태는 구조체와 비슷하다. 그러나 구조체와 같이 정의가 이루어지게 된다.
//인턴페이스도 여타 다른 함수들과 마찬가지로 함수의 파라미터가 될 수 있다.
type Shape interface {
	area_interface() float64
	param() float64
}

//인터페이스를 사용할 구조체 선언
type circle struct {
	radius float64
}

//circle의 인터페이스 함수 구현
func (c circle) area_interface() float64 {
	return c.radius * c.radius * math.Pi
}
func (c circle) param() float64 {
	return 2 * math.Pi * c.radius
}

//위에서 선언한 rect의 인터페이스 구현
func (r rect) area_interface() float64 {
	return r.width * r.height
}
func (r rect) param() float64 {
	return 2 * (r.width + r.height)
}

//interface type: 빈 인터페이스를 부르는 말
//메서드가 없는 빈 인터페이스 -> 모든 type을 나타낼 수 있다.(go의 모든 type은 0개이상의 메서드를 구현하므로)
func inter_type_test() {
	var x interface{}
	x = 1
	x = "Tom"
	x = 7.1

	fmt.Println(x)

	test := x //x, test모두 dynamic type. 단 test는 포인터 주소
	//x.(T) 와 같은 방식으로 x가 T타입인지 확인해주는 방식. T가 아니면 run time err
	test2 := x.(float64) //test2는 float64 타입.  type assertion이라 함

}

//에러처리
//go는 내장타입으로 error interface 타입을 가짐.
//아래와 같이 생겼다.
/*
type error interface{
	Error() string
}
이를 커스터마이징 하여 커스텀 에러 타입을 만들 수 있다.
혹은 각 라이브러리별로 Error를 구현한 경우도 있다.
*/

func test_error() {
	f, err := os.Open("C:not_exist.txt")
	if err != nil {
		log.Fatal(err.Error()) //os라이브러리의 Open 함수는 결과와 에러를 함께리턴
		//이때 에러가 nil이 아니라면 Open함수에서 구현한 Error함수를 사용하여 에러메시지를 볼 수 있다.
		//참고: log.Fatal() -> 해당 메시지를 출력하고, os.Exit(1)을 호출하여 프로그램을 종료시킴
	}
	fmt.Println(f.Name())
}

//커스텀에러와 시스템 에러를 구별하여 별도의 에러처리 구현
func error_exec() {
	res, err := your_func()

	switch err.(type) {
	default: //no err
		fmt.Println("no err")
	case MyErr:
		log.Print("log my error")
		//err.Error()를 커스터마이징했다면 사용가능
	case error:
		log.Fatal(err.Error())
	}
}

//defer 키워드:이 키워드가 붙은 특정 문장 혹은 함수는
//defer를 호출하는 함수가 리턴 되기 직전에 실행되도록 강제된다.
//주로 함수 종료 직전에 cleaning 용도로 사용 -> 파일 입출력동안 어떤 에러가 발생해도 파일을 닫게끔 만듬
func defer_test() {
	f, err := os.Open("1.txt")
	if err != nil {
		panic(err)
	}

	// main 마지막에 파일 close 실행
	defer f.Close()

	// 파일 읽기
	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))
}

//panic 키워드: 현재 함수를 즉시 멈추고
//현재 함수 안에 있는 defer 함수들을 모두 실행시킨 후 함수를 리턴함
func panic_starter() {
	openFile("Invalid.txt")
	println("Done") //이 문장은 실행 안됨
}

func panic_fun(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	// 파일 close 실행됨
	defer f.Close()
}

//recover 함수:  panic 함수에 의한 패닉 상태를 다시 정상 상태로 되돌리는 함수
//이거 예제 모르겠음;;
func recover_caller() {
	openFile("1.txt")
	println("Done") // 이 문장 실행됨
}

func recover_test(fn string) {
	// defer 함수. panic 호출시 실행됨
	defer func() { //4
		if r := recover(); r != nil { //panic을 제거함  5
			fmt.Println("OPEN ERROR", r) //실행됨.   6
		}
	}()

	f, err := os.Open(fn) //1
	if err != nil {       //2
		panic(err) //3
	}

	// 파일 close 실행됨
	defer f.Close() //7?
}

func main() {
	/*
		//메소드 호출법
		rectangle := rect{1, 2} //구조체 객체 초기화하기
		area := rectangle.area()
		fmt.Println(area)

		area2 := rectangle.area2() //여기서 rectangle 객체의 width 필드값이 증가.
		fmt.Println(area2)
		fmt.Println(rectangle.width, rectangle.height)
	*/
	recover_caller()
}
