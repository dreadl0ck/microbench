import numpy as np
import matplotlib.pyplot as plt

qemu_x20_emulated = [
	2227.71900,
	2153.67300,
	2077.32500,
	2189.70900,
	2513.94900,
	1968.49400,
	1425.81800,
	1969.73200,
	2824.23300,
	1491.00600,
	1524.03400,
	3898.57800,
	2205.78200,
	987.00000,
	2075.55500,
	3543.20400
]

qemu_x10 = [
	1808.02000,
	1785.44400,
	1814.50600,
	959.30600,
	1823.86800,
	1843.04300,
	1786.47600,
	841.46600,
	957.47000,
	943.35600
]

qemu_x20 = [
	1769.69200,
	845.00600,
	1625.41800,
	4957.51400,
	2736.37800,
	2949.16500,
	3076.66200,
	3058.14400,
	1718.07500,
	2936.87400,
	1807.19800,
	2653.38900,
	2798.30300,
	1835.52200,
	2750.62900,
	3426.62700,
	1811.71300,
	2707.24100,
	2060.75600,
	1809.65300
]

firecracker_x10 = [
	827.23400,
	867.32700,
	796.71300,
	722.12400,
	873.04600,
	814.83500,
	808.52500,
	844.53200,
	860.44900,
	836.74100,
	862.53100,
	544.05400,
	897.59100,
	842.71700,
	1400.19400,
	865.74200,
	901.30400,
	859.36200,
	865.08900,
	922.34000,
	919.82700,
	822.72400,
	1309.68200,
	665.55000,
	870.66500,
	836.81700,
	852.01200,
	901.95100,
	884.81700,
	843.94900,
	889.80500,
	704.54400,
	831.81700,
	840.66300,
	812.49400,
	903.74200,
	659.04900,
	854.26500,
	835.07000,
	833.29000
]

firecracker_x20 = [
	1204.44700,
	824.36700,
	1787.69900,
	1199.23300,
	844.43200,
	1576.95100,
	938.50900,
	824.74000,
	905.54900,
	866.83300,
	516.06300,
	815.73100,
	620.01900,
	1987.05500,
	839.47400,
	2013.45600,
	883.38600,
	1250.76400,
	1616.05000,
	965.29800,
	1256.70200,
	772.98100,
	827.40200,
	1395.90000,
	954.03600,
	871.77600,
	864.34800,
	1691.44000,
	910.21600,
	1296.94100,
	879.29400,
	1119.59600,
	608.66900,
	926.33000,
	838.47700,
	2257.68700,
	641.90100,
	2187.78900,
	861.54300,
	1878.58600,
	1619.62000,
	1222.85200,
	818.45300,
	1227.04300,
	1883.21900,
	1088.39700,
	833.05900,
	1967.74700,
	915.59300,
	856.05100,
	844.39000,
	868.45700,
	1533.19900,
	761.30700,
	1539.75100,
	1756.33600,
	904.74100,
	600.85500,
	1262.55500,
	1430.29800,
	2427.80300,
	1012.97500,
	785.19900,
	1771.76600,
	2764.79800,
	1159.09700,
	838.79000,
	1711.41900,
	879.12100,
	1163.59200,
	1050.61500,
	1010.61800,
	1410.67500,
	912.64300,
	833.47100,
	1408.59000,
	895.04800
]

qemu_x10_emulated = [
	1635.29500,
	1340.00000,
	1421.46400,
	895.74900,
	1691.70600,
	1237.65600,
	1379.18500,
	1551.78300,
	1609.93400,
	1390.90000
]



labels = ['qemu x10','qemu x20','qemu x10 emulated','qemu x20 emulated','firecracker x10','firecracker x20']
y_pos = np.arange(len(labels))

data=[
	qemu_x10,
	qemu_x20,
	qemu_x10_emulated,
	qemu_x20_emulated,
	firecracker_x10,
	firecracker_x20
]

fig, ax = plt.subplots()
ax.set_title("Kernel Boot Time (Concurrent)")
plt.ylabel('Time (ms)')
ax.boxplot(data, labels=labels)

plt.gcf().subplots_adjust(bottom=0.30)
plt.xticks(rotation=45)

#plt.show()
plt.savefig('plots/scripts/images/kernel-boot-time-concurrent.png')