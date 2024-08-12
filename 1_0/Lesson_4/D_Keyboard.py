# D. Клавиатура
from collections import Counter
from collections import defaultdict

with open('input.txt', 'r') as file:
    quantity = int(file.readline().strip())  # количество клавиш
    # на клавиатуре
    quantity_clicks_i = list(map(int, file.readline().strip().split()))  # количество
    # нажатий,выдерживаемых i-ой клавишей
    total_keystrokes = int(file.readline().strip())  # общее
    # количество нажатий клавиш
    sequence_pressed = list(map(int, file.readline().strip().split()))  #
    # последовательность нажатых клавиш

    # print(quantity, quantity_clicks_i, total_keystrokes,
    #       sequence_pressed)
    dict_quantity_clicks_i = dict(zip(range(1, len(quantity_clicks_i) + 1), quantity_clicks_i))
    # print(f'словарь номер кнопки: ресурс', dict_quantity_clicks_i)
    dict_sequence_pressed = dict(sorted((Counter(sequence_pressed)).items()))
    # print(f'словарь последовательности нажатий', dict_sequence_pressed)
    result = defaultdict(int)
    # print(result)
    for key in dict_quantity_clicks_i:
        result[key] = dict_quantity_clicks_i[key] - dict_sequence_pressed.get(key)
    # print(dict(result))
    for ans in result.values():
        if ans >= 0:
            print('NO')
        else:
            print('YES')

# if __name__ == "__main__":
#     main()
