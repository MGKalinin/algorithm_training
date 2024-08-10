# C. Самое частое слово
with open('input.txt', 'r') as file:
    from collections import Counter

    text = file.read().split()
    # print(text)
    res = Counter(text)
    # print(res)
    # print(res.items())
    max_count = max(res.values())
    # print(max_count)
    most_comm_word = [word for word, count in res.items() if count == max_count]
    # print(most_comm_word)
    print(min((most_comm_word)))
