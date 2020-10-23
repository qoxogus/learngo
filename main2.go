package main

// func main() {
// 	//c := make(chan bool)
// 	c := make(chan string)
// 	//people := [2]string{"taehyeon", "nico"}
// 	people := [5]string{"taehyeon", "nico", "dal", "larry", "flynn"}
// 	for _, person := range people {
// 		go isSexy(person, c)
// 	}
// 	// go sexyCount("taehyeon")
// 	// go sexyCount("nico")

// 	// time.Sleep(time.Second * 10)

// 	//result := <-c
// 	//fmt.Println(result) // == fmt.Println(<-c)

// 	// resultOne := <-c
// 	// resultTwo := <-c

// 	// fmt.Println("Wating for message")
// 	// //fmt.Println("Received this message :", <-c)
// 	// //fmt.Println("Received this message :", <-c)//(c)기본적인 채널로부터 메세지를 가져오는거
// 	// fmt.Println("Received this message :", resultOne)
// 	// fmt.Println("Received this message :", resultTwo)

// 	for i:=0; i<len(people);i++{
// 		fmt.Println(<-c)
// 	}

// func main() {
// 	c := make(chan string)
// 	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
// 	for i := 0; i < len(people); i++ {
// 		fmt.Println(<-c)
// 	}
// }

// func isSexy(person string, c chan string) {
// 	time.Sleep(time.Second * 10)
// 	c <- person + " is sexy"
// }

// func isSexy(person string, c chan string) {
// 	time.Sleep(time.Second * 10)
// 	//fmt.Println(person)
// 	c <- person + " is sexy"
// }

// func sexyCount(person string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(person, "is sexy", i)
// 		time.Sleep(time.Second)
// 	}
// }

//메세지를 받는다는건 blocking operation
/*
<- 이건 메세지를 받는것인데 이게 blocking operation
메인함수가 뭘 받기전까진 동작을 멈춘다
for i:=0; i<len(people);i++{
 		fmt.Println(<-c)
	 }
	 fmt.Println(<-c)
	 fmt.Println(<-c)
	 fmt.Println(<-c)
	 fmt.Println(<-c)
	 fmt.Println(<-c)
	 과 같다
*/
//github확인
// func main() {
// 	//c := make(chan bool)
// 	c := make(chan string)
// 	//people := [2]string{"taehyeon", "nico"}
// 	people := [5]string{"taehyeon", "nico", "dal", "larry", "flynn"}
// 	for _, person := range people {
// 		go isSexy(person, c)
// 	}
// 	// go sexyCount("taehyeon")
// 	// go sexyCount("nico")

// 	// time.Sleep(time.Second * 10)

// 	//result := <-c
// 	//fmt.Println(result) // == fmt.Println(<-c)

// 	// resultOne := <-c
// 	// resultTwo := <-c

// 	// fmt.Println("Wating for message")
// 	// //fmt.Println("Received this message :", <-c)
// 	// //fmt.Println("Received this message :", <-c)//(c)기본적인 채널로부터 메세지를 가져오는거
// 	// fmt.Println("Received this message :", resultOne)
// 	// fmt.Println("Received this message :", resultTwo)

// 	for i:=0; i<len(people);i++{
// 		fmt.Println(<-c)
// 	}

// func main() {
// 	c := make(chan string)
// 	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
// 	for i := 0; i < len(people); i++ {
// 		fmt.Println(<-c)
// 	}
// }

// func isSexy(person string, c chan string) {
// 	time.Sleep(time.Second * 10)
// 	c <- person + " is sexy"
// }

// func isSexy(person string, c chan string) {
// 	time.Sleep(time.Second * 10)
// 	//fmt.Println(person)
// 	c <- person + " is sexy"
// }

// func sexyCount(person string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(person, "is sexy", i)
// 		time.Sleep(time.Second)
// 	}
// }

//메세지를 받는다는건 blocking operation
/*
<- 이건 메세지를 받는것인데 이게 blocking operation
메인함수가 뭘 받기전까진 동작을 멈춘다
for i:=0; i<len(people);i++{
 		fmt.Println(<-c)
	 }
	 fmt.Println(<-c)
	 fmt.Println(<-c)
	 fmt.Println(<-c)
	 fmt.Println(<-c)
	 fmt.Println(<-c)
	 과 같다
*/
