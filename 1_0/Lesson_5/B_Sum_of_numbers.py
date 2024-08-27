# B. Сумма номеров
# N машин
# число K
with open('input.txt', 'r') as file:
    N, K = (map(int, file.readline().strip().split()))
    cars = list(map(int, file.readline().strip().split()))

# print(N, K, cars)

prefix_sum = [0] * (N + 1)
# print(prefix_sum)

for i in range(N):
    prefix_sum[i + 1] = prefix_sum[i] + cars[i]
# print(prefix_sum)

count = 0
start = 0
end = 0
curr_sum = 0
while end < N:
    curr_sum = prefix_sum[end + 1] - prefix_sum[start]
    if curr_sum == K:
        count += 1
        start += 1
        end += 1
    elif curr_sum < K:
        end += 1
    else:
        start += 1
print(count)
