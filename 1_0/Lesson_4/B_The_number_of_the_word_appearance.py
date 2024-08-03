# B. Номер появления слова
with open('input.txt', 'r') as file:
    import collections
    temp = []
    for lines in file.readline().split():
        for line in lines.split():
            temp.append(line)
    print(temp)
    res = collections.Counter(temp)
    print(res.values())
    ans = list(int(x) - 1 for x in res.values())
    print(ans)


