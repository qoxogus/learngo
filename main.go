package main

import (
	"fmt"
	"learngo/mydict"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (length int, uppercase string) { //뒷 괄호애들이 리턴할 애들
	//return len(name), strings.ToUpper(name)
	//이름길이, 이름을 대문자로 바꿔주는
	defer fmt.Println("I'm done") //function이 실행되고 난후에 실행되는 코드 defer
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
	//문자의 개수와 상관없이 받아서 출력해주는? ...string
}

func superAdd(numbers ...int) int {
	for index, number := range numbers { //index가 없으면 1,2,3,4,5까지밖에 안나와서 index로 같이 출력해준다
		fmt.Println(index, number)
	}
	return 1
}

func superAdd2(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

func superAdd3(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number //숫자를 주루룩 출력해주고 그 수들의 총합을구하는 함수
	}
	return total
}

func canIDrink(age int) bool {
	//koreanAge := age+2
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canIDrinkswitch(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	//name := "nico"
	//var name string = "nico" == name:="nico"
	//const erum string = "Lynn"
	//const == 변수값 변경불가
	//축약형은 오로지 func안에서만 가능하고 변수에만 적용 가능

	//fmt.Println("Hello world")

	//fmt.Println(multiply(2, 2))

	// totalLenght, up := lenAndUpper("taehyeon")
	// fmt.Println(totalLenght, up)

	// totalLenght, _ := lenAndUpper("taehyeon")
	// fmt.Println(totalLenght)
	//2번째 value 값을 무시하고 실행 결과는 len만 나옴

	//repeatMe("nico", "lynn", "dal", "marl", "flynn")

	//superAdd(1, 2, 3, 4, 5, 6)
	// superAdd2(1, 2, 3, 4, 5, 6)
	// result := superAdd3(1, 2, 3, 4, 5, 6)
	// fmt.Println(result)

	// fmt.Println(canIDrink(18))
	//fmt.Println(canIDrinkswitch(18))

	//#pointer
	// a := 2
	// b := &a //주소
	// *b = 20 //*b가 &a를 가리키고 &a는 a의 주소 결과적으로 *b==a
	// fmt.Println(a)//20

	//#array , slice
	// names := []string{"nico", "lynn", "dal"}
	// names = append(names, "flynn") //배열에 추가한다고 보면됨 index수의 관계없이
	// fmt.Println(names)

	//#map
	// nico := map[string]string{"name": "nico", "age": "12"}
	// //map[key]value
	// //fmt.Println(nico)
	// for key, value := range nico {
	// 	fmt.Println(key, value)
	// }

	//#struct 구조체
	// favFood := []string{"kimchi", "ramen"}
	// nico := person{name: "nico", age: 18, favFood: favFood}
	// fmt.Println(nico.name)

	//#banking
	//account := banking.Account{Owner: "nicolas"}
	// //account.Owner = "pepe"
	//fmt.Println(account)

	//account := account.NewAccount("nico")
	// // fmt.Println(account)
	//account.Deposit(10)
	//fmt.Println(account)
	// //fmt.Println(account.Balance())
	// // err := account.Withdraw(20)
	// // if err != nil {
	// // 	//log.Fatalln(err) //프로그램을 종료시킴
	// // 	fmt.Println(err)
	// // }
	// //fmt.Println(account.Balance(), account.Owner())
	// fmt.Println(account)

	// dictionary := mydict.Dictionary{"first": "First word"}
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// dictionary := mydict.Dictionary{}
	// word := "hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, _ := dictionary.Search(word)
	// fmt.Println("found", word, "definition:", hello)
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	//#2-6 update delete
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	//err := dictionary.Update(baseWord, "Second") //업데이트는 에러도 리턴
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//word, _ := dictionary.Search(baseWord)
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
	fmt.Println(word)

}

// func main1() {
// 	runtime.GOMAXPROCS(runtime.NumCPU()) // CPU 개수를 구한 뒤 사용할 최대 CPU 개수 설정

// 	fmt.Println(runtime.GOMAXPROCS(0)) // 설정 값 출력

// 	s := "Hello, world!"

// 	for i := 0; i < 100; i++ {
// 		go func(n int) { // 익명 함수를 고루틴으로 실행
// 			fmt.Println(s, n)
// 		}(i)
// 	}

// 	fmt.Scanln()
// }
