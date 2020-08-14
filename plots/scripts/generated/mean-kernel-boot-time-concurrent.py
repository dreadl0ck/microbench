firecracker_x20 = [
	3554.90400,
	755.61100,
	1819.75800,
	3388.34000,
	760.02600,
	3875.64500,
	2468.07100,
	2455.89900,
	2488.66300,
	3514.74600,
	2779.12800,
	1705.34600,
	2336.01300,
	2094.83800,
	3602.47700,
	1410.17200,
	2597.68900,
	3596.89600,
	3474.84000,
	2619.75300,
	2711.32400,
	3743.93800,
	3190.34300,
	1798.68400,
	2758.68300,
	2875.20500,
	2361.88500,
	844.92200,
	1560.95100,
	3521.15500,
	1740.54100,
	1529.59700,
	814.65000,
	2104.17600,
	3069.22300,
	2504.95900,
	3472.38900,
	884.14700,
	2650.42200,
	2756.08400,
	3593.40100,
	898.76900,
	1437.21900,
	1859.34300,
	869.92600,
	2076.91000,
	2186.29300,
	2246.33900,
	3591.75400,
	2626.09600,
	792.14800,
	2073.09900,
	1752.66200,
	2229.60900,
	808.46700,
	1674.26700,
	858.47700,
	2089.15900,
	2684.21500,
	3350.69400,
	908.58300,
	3637.89600,
	3416.02300,
	2318.66300,
	3430.38900,
	2679.16700,
	705.88000,
	2828.87100
]

qemu_x10_emulated = [
	991.79300,
	1502.57500,
	1560.14300,
	1346.11700,
	1628.59700,
	1746.73600,
	1473.50800,
	1701.27100,
	1024.54100,
	1173.34800
]

qemu_x20_emulated = [
	1898.89000,
	3093.18700,
	3762.62700,
	1985.79800,
	3437.81900,
	2136.57000,
	1602.42000,
	2270.16100,
	2088.60200,
	1669.94800,
	1417.46800,
	2654.15500,
	1324.10200,
	1410.33200,
	2469.50200,
	3145.20900
]

qemu_x10 = [
	1823.44200,
	1067.30400,
	1864.19100,
	1788.42100,
	1773.27000,
	1426.86400,
	1120.89500,
	1748.02200,
	1300.73900,
	1775.35900
]

qemu_x20 = [
	1890.81000,
	4871.14300,
	5480.29900,
	3314.38400,
	3818.85500,
	3949.68300,
	3361.56700,
	1825.03500,
	4658.25700,
	1792.80900,
	1834.52900,
	4444.18100,
	4852.08800,
	2999.03900,
	5111.35000,
	3565.27500
]

firecracker_x10 = [
	834.78400,
	759.35000,
	758.98400,
	796.60900,
	855.59700,
	720.12800,
	780.35200,
	837.40900,
	806.98600,
	738.77200,
	772.68900,
	784.23300,
	820.87000,
	782.18600,
	729.65900,
	809.09600,
	830.62000,
	701.83900,
	723.39900,
	830.52000,
	820.34800,
	780.64800,
	789.78500,
	773.39400,
	774.60400,
	744.55600,
	747.61400,
	806.25300,
	794.27800,
	750.86600,
	724.03100,
	787.14600,
	797.38200,
	809.36400,
	755.32000,
	842.04700,
	783.36600,
	780.51200,
	715.93600,
	786.63100
]



import statistics as stats
print("script: plots/scripts/mean-kernel-boot-time-concurrent.py")
print("stats.mean(firecracker_x10):", stats.mean(firecracker_x10))
print("stats.mean(firecracker_x20):", stats.mean(firecracker_x20))
print("stats.mean(qemu_x10_emulated):", stats.mean(qemu_x10_emulated))
print("stats.mean(qemu_x20_emulated):", stats.mean(qemu_x20_emulated))
print("stats.mean(qemu_x10):", stats.mean(qemu_x10))
print("stats.mean(qemu_x20):", stats.mean(qemu_x20))


import matplotlib.pyplot as plt; plt.rcdefaults()
import numpy as np
import matplotlib.pyplot as plt

objects = ('qemu x10','qemu x10 emulated','qemu x20','qemu x20 emulated','firecracker x10','firecracker x20')
y_pos = np.arange(len(objects))
performance = [
	stats.mean(qemu_x10),
	stats.mean(qemu_x10_emulated),
	stats.mean(qemu_x20),
	stats.mean(qemu_x20_emulated),
	stats.mean(firecracker_x10),
	stats.mean(firecracker_x20)
]

bar = plt.bar(y_pos, performance, align='center', alpha=0.5, color=['blue', 'blue', 'blue', 'blue', 'orange', 'orange'])
plt.xticks(y_pos, objects)
#plt.yticks(np.arange(0, 3000, 500))
plt.ylabel('Time (ms)')
plt.title('Mean Kernel Boot Time (Concurrent)')

plt.legend((bar[0], bar[4]), ('QEMU', 'firecracker'))

plt.gcf().subplots_adjust(bottom=0.30)
plt.xticks(rotation=45)

#plt.show()
plt.savefig('plots/scripts/images/mean-kernel-boot-time-concurrent.png')