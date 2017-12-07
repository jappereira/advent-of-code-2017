# input = '''0 2 7 0'''
input = '''14	0	15	12	11	11	3	5	1	6	8	4	9	1	8	4'''

memoryBanks = input.split('\t')
memoryBanks = map(int, memoryBanks)
memoryBanksHistory = [list(memoryBanks)]
redistributionCycles = 0

while True:
    i = memoryBanks.index(max(memoryBanks))
    blocks = memoryBanks[i]
    memoryBanks[i] = 0

    for j in range(blocks):
        i = (i + 1) % len(memoryBanks);
        memoryBanks[i] += 1

    redistributionCycles += 1

    if memoryBanks in memoryBanksHistory:
        break
    else:
        memoryBanksHistory.append(list(memoryBanks))

# 6.1
print(redistributionCycles)

# 6.2
print(len(memoryBanksHistory) - memoryBanksHistory.index(memoryBanks))