# C. Туризм
with open('input.txt', 'r') as file:
    N = int(file.readline().strip())
    chain = [tuple(map(int, file.readline().split())) for i in range(N)]
    print(N, chain)
    M = int(file.readline().strip())
    track = [tuple(map(int, file.readline().split())) for k in range(M)]
    print(M, track)

from dataclasses import dataclass


@dataclass
class TotalHeight:
    N: int
    M: int
    chain: list
    track: list

    def calculating_prefix_amount(self):
        prefix_amount = [0] * (self.N + 1)
        for j in range(1, self.N):
            if self.chain[j][1] > self.chain[j - 1][1]:
                prefix_amount[j] = prefix_amount[j-1] + (self.chain[j][1]-self.chain[j-1][1])
            else:
                prefix_amount[j] = prefix_amount[j-1]
        # print(prefix_amount)
        return prefix_amount

    def result(self):



ans = TotalHeight(N, M, chain, track)
print(ans.calculating_prefix_amount())
