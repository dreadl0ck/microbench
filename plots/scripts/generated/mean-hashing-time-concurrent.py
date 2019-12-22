firecracker_x20 = [
	132.29732,
	157.06245,
	309.93979,
	384.04974,
	373.12637,
	238.20949,
	136.03353,
	123.14628,
	208.83135,
	252.36969,
	126.79080,
	148.17366,
	266.51252,
	89.96384,
	219.47702,
	111.77703,
	313.33839,
	160.50905,
	161.11818,
	531.73820,
	302.27080,
	177.98168,
	266.14542,
	170.26839,
	506.31576,
	239.76404,
	305.74444,
	253.75803,
	167.83654,
	142.78564,
	374.68036,
	172.80929,
	167.01410,
	144.47654,
	228.49901,
	148.35313,
	313.32033,
	144.68042,
	160.82415,
	272.06669,
	215.89366,
	125.20595,
	239.90505,
	449.48787,
	208.48992,
	166.99603,
	121.99710,
	168.95548,
	96.79133,
	131.40900,
	138.12581,
	539.20687,
	317.49030,
	188.44615,
	137.69081,
	345.58544,
	264.41744,
	158.22504,
	370.05461,
	123.53943,
	103.10711,
	158.03811,
	348.44416,
	153.11739,
	281.54020,
	105.28376,
	212.69969,
	146.05186,
	196.58556,
	172.98985,
	288.42120,
	166.93423,
	146.21189,
	142.08995,
	225.24380,
	440.89678,
	299.55786,
	138.72586,
	308.70366,
	139.26767
]

qemu_x10_emulated = [
	94.14717,
	156.90176,
	113.15587,
	120.79472,
	90.70383,
	104.04471,
	118.39555,
	131.67722,
	107.08754,
	71.88496
]

qemu_x20_emulated = [
	378.85285,
	469.08973,
	583.81664,
	478.81133,
	286.03118,
	496.60632,
	886.69499,
	376.24134,
	414.71355,
	499.80044,
	510.50934,
	297.13697,
	508.15298,
	143.71336,
	638.46399,
	647.37549,
	289.64443,
	294.99084,
	640.61546,
	766.78585
]

qemu_x10 = [
	182.05452,
	84.64198,
	72.75233,
	75.32137,
	226.06714,
	79.05355,
	77.88387,
	169.59224,
	109.67444,
	88.52244
]

qemu_x20 = [
	195.79982,
	105.65309,
	449.22212,
	356.70131,
	413.75686,
	646.81382,
	1596.23058,
	421.56951,
	311.26853,
	581.86975,
	134.29292,
	566.31754,
	447.35200,
	178.84525,
	261.80951,
	241.25731,
	196.64402,
	564.95767,
	326.32763,
	368.06408
]

firecracker_x10 = [
	135.78019,
	138.65614,
	179.23102,
	134.88292,
	154.27279,
	129.74430,
	180.40123,
	116.56217,
	150.20384,
	120.15418,
	151.59742,
	133.23764,
	162.02727,
	140.62525,
	138.35635,
	121.92696,
	141.50207,
	112.81506,
	146.11274,
	140.43941,
	86.48471,
	95.48688,
	146.68610,
	138.48488,
	120.94616,
	124.28662,
	140.08580,
	161.06567,
	155.96252,
	126.26792,
	149.64719,
	125.76017,
	122.40551,
	101.20659,
	103.36593,
	113.90037,
	108.51243,
	149.67731,
	148.49518,
	122.14131
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