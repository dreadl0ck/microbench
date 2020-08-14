qemu_x10 = [
	195.14518,
	93.65147,
	156.57188,
	149.50182,
	155.30480,
	88.10410,
	89.50312,
	99.13321,
	204.60642,
	172.17208
]

qemu_x20 = [
	289.84107,
	560.51469,
	425.18495,
	739.70125,
	263.82268,
	282.63331,
	727.92277,
	382.75400,
	577.82842,
	411.37526,
	588.12797,
	374.18856,
	393.94855,
	509.07984,
	560.88057,
	271.46601,
	272.71829,
	576.96210,
	874.44772,
	890.37863
]

firecracker_x10 = [
	104.07448,
	99.34839,
	131.20998,
	138.62079,
	86.78961,
	138.65831,
	120.84950,
	121.19064,
	106.94246,
	89.78305,
	131.02459,
	133.71641,
	111.02040,
	142.41236,
	142.69228,
	68.45094,
	117.13368,
	94.88188,
	93.50673,
	110.21706,
	79.75827,
	106.60449,
	124.68322,
	80.06071,
	115.49098,
	105.70085,
	91.78910,
	89.71852,
	65.43178,
	85.41446,
	110.43159,
	68.93732,
	106.87081,
	52.09216,
	138.57745,
	79.73531,
	98.88123,
	121.70279,
	106.57937,
	139.64362
]

firecracker_x20 = [
	156.52121,
	205.97125,
	481.13709,
	147.44884,
	561.44986,
	817.15614,
	288.99171,
	327.18714,
	346.54690,
	107.63186,
	241.10153,
	318.07976,
	331.60826,
	510.25856,
	192.50891,
	451.88238,
	394.30200,
	392.88577,
	328.63356,
	292.25633,
	344.12324,
	274.33473,
	149.66771,
	265.69078,
	305.86843,
	286.99913,
	304.28533,
	699.63628,
	474.74685,
	500.85047,
	559.94130,
	575.74408,
	370.19719,
	155.61765,
	469.82663,
	125.54711,
	384.63699,
	350.50242,
	374.35490,
	246.17883,
	118.25043,
	90.78846,
	168.01235,
	416.05521,
	297.59206,
	588.60189,
	304.47167,
	267.97457,
	267.97723,
	474.05268,
	186.78073,
	302.53314,
	252.83995,
	124.90898,
	214.75820,
	578.54678,
	356.30159,
	379.81670,
	532.03119,
	452.52472,
	130.87572,
	211.01329,
	595.38437,
	199.49937,
	195.59563,
	316.11552,
	236.07661,
	276.57520,
	424.48911,
	588.34366,
	91.47844,
	314.29896,
	464.21225,
	335.96040,
	110.30430,
	210.41962,
	558.80154,
	321.16427,
	128.23410,
	479.97885
]

qemu_x10_emulated = [
	178.36342,
	123.74093,
	229.77334,
	127.03459,
	140.39595,
	214.91456,
	106.77583,
	131.62837,
	238.04944,
	162.64456
]

qemu_x20_emulated = [
	819.50999,
	344.50844,
	665.64176,
	633.60153,
	253.44307,
	512.46019,
	506.13161,
	410.76653,
	329.59584,
	601.54796,
	866.84970,
	431.71351,
	516.38986,
	768.98527,
	825.59214,
	813.34563,
	1030.82185,
	396.30385,
	496.79091,
	371.42032
]



import statistics as stats
print("script: plots/scripts/mean-hashing-time-concurrent.py")
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
#plt.yticks(np.arange(0, 400, 100))
plt.ylabel('Time (ms)')
plt.title('Mean Hashing Time SHA-256 100 x 1MB (Concurrent)')

plt.legend((bar[0], bar[4]), ('QEMU', 'firecracker'))

plt.gcf().subplots_adjust(bottom=0.30)
plt.xticks(rotation=45)

#plt.show()
plt.savefig('plots/scripts/images/mean-hashing-time-concurrent.png')