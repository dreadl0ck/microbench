firecracker_x20 = [
	6662.04558,
	1464.93841,
	3023.95455,
	5275.09761,
	3865.18206,
	3872.28195,
	1273.15475,
	5791.81434,
	3147.99897,
	4012.15826,
	3956.41905,
	5819.36523,
	5521.42786,
	3911.37648,
	3810.17589,
	5979.85453,
	5893.56611,
	2386.46433,
	4938.19670,
	4781.44906,
	5464.66105,
	4431.16698,
	4608.05328,
	6592.96434,
	5193.94982,
	4178.93364,
	3510.69842,
	3186.62340,
	4966.45678,
	3335.90058,
	3607.61224,
	5271.43227,
	5283.31998,
	1546.03485,
	2806.74006,
	5502.99366,
	4328.78512,
	2490.38942,
	1711.72048,
	6099.75542,
	4982.42622,
	3368.61183,
	4123.85532,
	2346.65670,
	5715.16120,
	4011.36527,
	2075.86513,
	5896.51872,
	2525.98030,
	2425.11249,
	3034.65294,
	1600.17962,
	6024.20327,
	3768.19347,
	4172.28173,
	5959.22453,
	5637.40446,
	4612.68223,
	3177.50506,
	4998.96457,
	1726.65105,
	3520.18092,
	3064.95038,
	5479.43495,
	1029.90202,
	3311.23092,
	4267.41561,
	1439.85255,
	5982.50349,
	4365.75158,
	4481.65483,
	1818.93098,
	6282.92486,
	4200.65042,
	3433.50958,
	6107.99599,
	4522.05002,
	4288.64340,
	1707.02502,
	5223.25839
]

qemu_x10_emulated = [
	1933.22076,
	2073.39720,
	2054.65876,
	1897.79160,
	2054.66048,
	2047.16063,
	2057.93669,
	4594.26131,
	2307.38883,
	2549.90764
]

qemu_x20_emulated = [
	9720.19116,
	9956.20746,
	7383.89662,
	9713.84254,
	8919.88122,
	8648.71127,
	10078.33029,
	5320.92447,
	8109.74281,
	6499.60458,
	6396.48069,
	10666.65209,
	9159.56683,
	7296.32641,
	7550.94801,
	6477.87365,
	8012.21183,
	10840.04977,
	5440.82323,
	9059.51876
]

qemu_x10 = [
	2059.91428,
	1605.60831,
	2076.90703,
	2067.15754,
	2036.47188,
	2091.15891,
	1633.88772,
	2032.32209,
	1996.66400,
	2045.42226
]

qemu_x20 = [
	10147.60486,
	8675.96674,
	8011.81348,
	4302.98623,
	9116.69304,
	9617.45794,
	7979.35393,
	9753.40816,
	7725.50243,
	9172.77131,
	5129.39322,
	8080.44541,
	2127.71210,
	5092.44657,
	7683.83233,
	10011.16776,
	9592.75357,
	6845.09949,
	8012.63131,
	8074.68653
]

firecracker_x10 = [
	1211.56146,
	1043.39741,
	1051.05342,
	1291.94711,
	2109.09652,
	1042.19680,
	1451.66791,
	2710.55414,
	1119.31507,
	1288.31740,
	1152.77327,
	1343.36682,
	2611.47584,
	1083.44683,
	1099.99734,
	1018.66952,
	2440.27908,
	1041.73026,
	1258.02881,
	1038.70825,
	1026.08277,
	1043.05052,
	1033.16352,
	1038.24056,
	2291.90206,
	1037.50839,
	1025.61408,
	1210.35478,
	1018.31279,
	1022.05158,
	1536.31513,
	1005.60613,
	1600.12828,
	1038.48922,
	1220.25308,
	1031.17365,
	1493.16881,
	1042.00201,
	1292.93189,
	1276.62722
]



import statistics as stats
print("script: plots/scripts/mean-webservice-time-concurrent.py")
print("stats.mean(firecracker_x10):", stats.mean(firecracker_x10))
print("stats.mean(firecracker_x20):", stats.mean(firecracker_x20))
print("stats.mean(qemu_x10_emulated):", stats.mean(qemu_x10_emulated))
print("stats.mean(qemu_x20_emulated):", stats.mean(qemu_x20_emulated))
print("stats.mean(qemu_x10):", stats.mean(qemu_x10))
print("stats.mean(qemu_x20):", stats.mean(qemu_x20))


import statistics as stats
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
#plt.yticks(np.arange(0, 8000, 500))
plt.ylabel('Time (ms)')
plt.title('Mean Web Service Startup Time (Concurrent)')

plt.legend((bar[0], bar[4]), ('QEMU', 'firecracker'))

plt.gcf().subplots_adjust(bottom=0.30)
plt.xticks(rotation=45)

#plt.show()
plt.savefig('plots/scripts/images/mean-webservice-time-concurrent.png')