# B. Номер появления слова
with open('input.txt', 'r') as file:
    text = file.read().split()
    words_count = {}
    result = []
    for word in text:
        if word in words_count:
            # print(f'1', words_count)
            result.append(str(words_count[word]))
            words_count[word] += 1
        else:
            # print(f'2', words_count)
            result.append('0')
            words_count[word] = 1
    print(' '.join(result))
