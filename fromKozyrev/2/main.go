package main

// TODO писать код в песочнице
// TODO 1.fan-in - намудрить запрос разными горутинами и т.п.
// TODO 2.semaphore
// TODO 3.worker pool
// TODO 4.использовать контекст отмены после двух 200 ок
//Озон Платформа
//
//Дата: февраль 2025
//Грейд: мидл
//#Озон
//
//Задачи присланы участником сообщества. Как прислать свою задачу читайте тут.
//
//1️⃣ Скрининг
//👉 Дано x, определить, является ли x степенью двойки.
//💡 Решение: x & (x - 1) == 0 – если выполняется, значит x степень двойки.
//2️⃣ Техническое собеседование
//📌 Задача 1:
//Дан список URL, нужно синхронно пройтись по нему и:
//✅ Вывести "ОК", если статус-код 200.
//❌ Вывести "не ОК", если статус-код  не 200 или произошла ошибка.
//
//📌 Задача 2:
//Та же проверка, но теперь асинхронно/в параллель, при этом ничего не должно потеряться.
//urls := []string{
//"https://yandex.ru",
//"https://habr.com",
//"https://lenta.ru",
//"https://ria.ru",
//"https://tass.ru",
//}
