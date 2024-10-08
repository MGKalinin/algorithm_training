# E. OpenCalculator

# В новой программе OpenCalculator появилась новая возможность –
# можно настроить, какие кнопки отображаются,
# а какие – нет. Если кнопка не отображается на экране,
# то ввести соответствующую цифру с клавиатуры или копированием
# из другой программы нельзя. Петя настроил калькулятор так,
# что он отображает только кнопки с цифрами x, y,
# z. Напишите программу, определяющую, сможет ли Петя ввести число N,
# а если нет, то какое минимальное количество
# кнопок надо дополнительно отобразить на экране для его ввода.

# Сначала вводятся три различных числа из диапазона от 0 до 9:
# x, y и z (числа разделяются пробелами).
# Далее вводится целое неотрицательное число N,
# которое Петя хочет ввести в калькулятор. Число N не превышает 10000.

# Выведите, какое минимальное количество кнопок должно быть
# добавлено для того, чтобы можно было ввести число N
# (если число может быть введено с помощью уже имеющихся кнопок,
# выведите 0)


with open('input.txt', 'r') as file:
    numbers = set(file.readline().strip().split())
    digit = set(list(_ for _ in file.readline().strip()))
    # print(numbers, digit)


def answer(numbers, digit):
    res = digit - numbers  # которых есть только в одном
    # print(res)

    if len(res) == 0:
        return 0
    else:
        return len(res)


print(answer(numbers, digit))


