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

