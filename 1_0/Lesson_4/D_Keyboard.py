# D. Клавиатура
from dataclasses import dataclass, field
from typing import List, Any
from collections import Counter


@dataclass
class Keyboard:
    quantity: int  # количество клавиш
    # на клавиатуре
    total_keystrokes: int  # общее
    # количество нажатий клавиш
    quantity_clicks_i: list = field(default_factory=list)  # количество
    # нажатий,выдерживаемых i-ой клавишей
    sequence_pressed: list = field(default_factory=list)  #

    # последовательность нажатых клавиш

    def result(self) -> Any:
        dict_quantity_clicks_i = dict(enumerate(self.quantity_clicks_i,
                                                start=1))
        for key in dict_quantity_clicks_i:
            count = self.sequence_pressed.count(key)
            dict_quantity_clicks_i[key] -= count

        # print(dict_quantity_clicks_i)
        # print(type(dict_quantity_clicks_i))
        # my_dict = {1: -1, 2: 49, 3: 0, 4: 1, 5: -4}
        results = ["NO" if value >= 0 else "YES" for value
                   in dict_quantity_clicks_i.values()]
        # print(results)
        return results


def main():
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
    key = Keyboard(quantity=quantity, total_keystrokes=total_keystrokes,
                   quantity_clicks_i=quantity_clicks_i, sequence_pressed=sequence_pressed)

    answers = key.result()
    for answer in answers:
        print(answer)


if __name__ == "__main__":
    main()
