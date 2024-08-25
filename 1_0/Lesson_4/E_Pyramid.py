# E. Пирамида

# Для строительства двумерной пирамиды используются прямоугольные блоки, каждый из которых характеризуется
# шириной и высотой.
# Можно поставить один блок на другой, только если ширина верхнего блока строго меньше ширины нижнего
# (блоки нельзя поворачивать). Самым нижним в пирамиде может быть блок любой ширины.
# По заданному набору блоков определите, пирамиду какой наибольшей высоты можно построить из них.
#
# Формат ввода
# В первой строке входных данных задается число N — количество блоков
# В следующих N строках задаются пары натуральных чисел ширина и высота блока соответственно.
#
# Формат вывода
# Выведите одно целое число — максимальную высоту пирамиды.

import typing

# Определение NamedTuple
Blocks = typing.NamedTuple('Blocks', [('width', int), ('height', int)])

# Чтение данных и создание экземпляров Blocks
with open('input.txt', 'r') as file:
    n = int(file.readline().strip())
    blocks = []
    for i in range(n):
        width, height = map(int, file.readline().strip().split())
        block = Blocks(width=width, height=height)
        blocks.append(block)

# print(blocks)
blocks.sort(key=lambda x: (-x.width, x.height))
# print(blocks)
# оставить значения width==width max(height)
blocks_dict = {}
for block in blocks:
    if block.width not in blocks_dict or block.height > blocks_dict[block.width]:
        blocks_dict[block.width] = block.height
# print(blocks_dict)
print(sum(blocks_dict.values()))

