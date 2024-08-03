# A. Словарь синонимов
with open('input.txt', 'r') as file:
    num = int(file.readline().strip())
    one = {}
    two = {}
    for i in range(num):
        first, second = file.readline().strip().split()
        one[first] = second
        two[second] = first
    # print(one, two)
    key = file.readline().strip()
    # print(key)


def answer(one, two, key):
    from collections import ChainMap
    ans = ChainMap(one, two)
    return ans[key]


print(answer(one, two, key))
