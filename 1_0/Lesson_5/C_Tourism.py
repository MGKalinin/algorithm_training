# C. Туризм
with open('input.txt', 'r') as file:
    N = int(file.readline().strip())
    chain = [list(map(int, file.readline().split())) for i in range(N)]
    # print(N, chain)
    M = int(file.readline().strip())
    track = [list(map(int, file.readline().split())) for k in range(M)]
    # print(M, track)

from dataclasses import dataclass
import typing


@dataclass
class TotalHeight:
    N: int  # длина г/цепи
    M: int  # длина трассы
    chain: list  # г/цепь [[2, 1], [4, 5], [7, 4], [8, 2], [9, 6], [11, 3], [15, 3]]
    track: list  # трасса [[2, 6]]

    def calculating_prefix_amount(self: typing.Any) -> list:
        """
        Метод возвращает префиксную сумму подъёма на горной цепи
        """
        prefix_sum = [0] * (N + 1)
        # print(prefix_sum)
        # подсчёт префиксной суммы подъёмов
        for i in range(1, N):
            if chain[i][1] > chain[i - 1][1]:
                # print(f'{self.chain[i][1]} ')
                prefix_sum[i + 1] = prefix_sum[i] + (chain[i][1] - chain[i - 1][1])
                # print(f'pr', prefix_sum[i])
            else:
                prefix_sum[i+1] = prefix_sum[i]
        # print(prefix_sum)
        return prefix_sum

    def answer(self: typing.Any) -> list:
        """
        Метод возвращает результат подъёма по трассе
        """
        result = []
        summ = self.calculating_prefix_amount()
        for s, f in self.track:
            # print(s, f)
            s -= 1
            f -= 1
            if s > f:
                s, f = f, s
            total_up = summ[f + 1] - summ[s + 1]

            result.append(total_up)
        return result


ex = TotalHeight(N, M, chain, track)
results = ex.answer()
for res in results:
    print(res)



