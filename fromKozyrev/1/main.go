package main

//Дата: февраль 2025
//Компанию: #Озон
//Грейд: мидл
//
//Задачи присланы участником сообщества. Как прислать свою задачу читайте тут.
//
//Первая задача
//s := "test"
//println(s[0]) // ответ кандидата
//s[0] = "R"
//println(s) // ответ кандидата
////Как поправить код?
//func main() {
//	for i := 0; i < 5; i++ {
//		go func() {
//			fmt.Println(i)
//		}()
//	}
//}
//Вторая задача
//// В чем проблема кода?
//type Storage struct {
//	cache *lru.Cache
//}
//
//func (s *Storage) Set(wh *warehouse.Warehouse) {
//	s.cache.Put(wh.Id, *wh)
//}
//
//func (s *Storage) Get(id types.WarehouseId) *warehouse.Warehouse {
//	item, ok := s.cache.Get(id)
//
//	if ok {
//		if wh, ok := item.(*warehouse.Warehouse); ok {
//			return wh
//		}
//	}
//
//	return nil
//}
//Третья задача
//// user
//id firstname lastname birth
//1 Ivan Petrov 1996-05-01
//2 Anna Petrova 1999-06-01
//3 Anna Petrova 1990-10-02
//
//// purchase
//sku price user_id date
//1 | 5500 1 2021-02-15
//1 | 5700 1 2021-01-15
//2 4000  1 2021-02-14
//3 8000  2 2021-03-01
//4 400    2 2021-03-02
//
//// ban_list
//user_id date_from
//1 2021-03-08
//
//
//1. Вывести уникальные комбинации пользователя и id товара
//для всех покупок, совершенных пользователями до того, как
//их забанили. Отсортировать сначала по имени пользователя, потом по SKU
//2. Найти пользователей, которые совершили покупок на сумму
//больше 5000р. Вывести их имена в формате id пользователя имя фамилия сумма покупок
