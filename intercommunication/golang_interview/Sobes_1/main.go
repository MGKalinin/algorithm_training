package main

import "fmt"

//
//==========================================
//Задача 1
//Что выведет код? Исправить все проблемы
//==========================================

//func main() {
//	ch := make(chan int) //добавить буфер канала 3
//	wg := &sync.WaitGroup{}
//	wg.Add(3)
//
//	// горутина закрытия канала после завершения всех писателей
//	go func() {
//		wg.Wait()
//		close(ch)
//	}()
//
//	for i := 0; i < 3; i++ {
//		go func(v int) {
//			defer wg.Done()
//			ch <- v * v // 0,1,2
//		}(i)
//	}
//	//wg.Wait()
//	//close(ch) //закрыть канал
//	var sum int
//	for v := range ch {
//		sum += v
//	}
//	fmt.Printf("result: %d\n", sum) //0,1,4
//}

//
//==========================================
//Задача 2
//Что выведет код?
//==========================================
//
//func main() {
//	v := []int{3, 4, 1, 2, 5}
//	ap(v)
//	sr(v)
//	//fmt.Println(v) //[1 2 3 4 5]
//}
//
//func ap(arr []int) {
//	arr = append(arr, 10) //v := []int{3, 4, 1, 2, 5} ,10
//	//fmt.Println(arr)
//}
//
//func sr(arr []int) {
//	sort.Ints(arr) //1 2 3 4 5 - ответ
//	//fmt.Println(arr)
//
//}

//==========================================
//Задача 3
//Что выведет код?
//==========================================

type impl struct{}

type I interface {
	C()
}

func (*impl) C() {}

func A() I {
	return nil

}
func B() I {
	var ret *impl
	return ret
}

func main() {
	a := A()
	b := B()
	fmt.Println(a == b)
}

//==========================================
//Задача 4
//Что выведет код? Должны выводится все значения
//==========================================
//
//func main() {
//a := 5000
//for i := 0; i < a; i++ {
//go fmt.Println(i)
//}
//}
