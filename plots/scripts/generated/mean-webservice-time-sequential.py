host_cpu = [
	1157.95299,
	1165.36658,
	1152.09570,
	1218.73698,
	1179.61424,
	1185.44472,
	1203.34863,
	1190.19420,
	1175.66090,
	1155.82302
]

T2_cpu = [
	1016.42953,
	1024.95370,
	1067.08935,
	1997.75441,
	2305.34122,
	1011.93757,
	1036.41593,
	1278.90710,
	1015.08059,
	1170.91840,
	1010.84990,
	2148.78671,
	1699.22946,
	1037.96780,
	1010.78237,
	1008.55516,
	1012.26030,
	1025.59098,
	1182.20765,
	1354.10523,
	2074.62211,
	1018.95668,
	1010.50159,
	1014.94541,
	1029.40083,
	1020.62301,
	1150.55952,
	1299.17002,
	1277.86156,
	1018.02097,
	1019.23645,
	1042.75358,
	1017.59681,
	1025.73531,
	1038.73871,
	1015.92998,
	1017.66057,
	1024.79558,
	1035.53808,
	2127.56267
]

C3_cpu = [
	1016.42953,
	1024.95370,
	2305.34122,
	1011.93757,
	1015.08059,
	1170.91840,
	1699.22946,
	1037.96780,
	1012.26030,
	1025.59098,
	2074.62211,
	1018.95668,
	1029.40083,
	1020.62301,
	1277.86156,
	1018.02097,
	1017.59681,
	1025.73531,
	1017.66057,
	1024.79558
]

default_kernel = [
	1024.95370,
	1067.08935,
	1011.93757,
	1036.41593,
	1170.91840,
	1010.84990,
	1037.96780,
	1010.78237,
	1025.59098,
	1182.20765,
	1018.95668,
	1010.50159,
	1020.62301,
	1150.55952,
	1018.02097,
	1019.23645,
	1025.73531,
	1038.73871,
	1024.79558,
	1035.53808
]

emulated_cpu = [
	2032.75584,
	1378.70338,
	1468.89976,
	1478.86437,
	2031.78867,
	1403.51850,
	2049.36518,
	1386.62287,
	1406.38919,
	1384.86904
]



import statistics as stats
print("script: plots/scripts/mean-webservice-time-sequential.py")
print("stats.mean(C3_cpu):", stats.mean(C3_cpu))
print("stats.mean(default_kernel):", stats.mean(default_kernel))
print("stats.mean(emulated_cpu):", stats.mean(emulated_cpu))
print("stats.mean(host_cpu):", stats.mean(host_cpu))
print("stats.mean(T2_cpu):", stats.mean(T2_cpu))


import statistics as stats
import matplotlib.pyplot as plt; plt.rcdefaults()
import numpy as np
import matplotlib.pyplot as plt

objects = ('host cpu','emulated cpu','T2 cpu','C3 cpu','default kernel')
y_pos = np.arange(len(objects))
performance = [
	stats.mean(host_cpu),
	stats.mean(emulated_cpu),
	stats.mean(T2_cpu),
	stats.mean(C3_cpu),
	stats.mean(default_kernel)
]

bar = plt.bar(y_pos, performance, align='center', alpha=0.5, color=['blue', 'blue', 'orange', 'orange', 'orange'])
plt.xticks(y_pos, objects)
#plt.yticks(np.arange(0, 1500, 200))
plt.ylabel('Time (ms)')
plt.title('Mean Web Service Startup Time')

plt.legend((bar[0], bar[2]), ('QEMU', 'firecracker'))

#plt.show()
plt.savefig('plots/scripts/images/mean-webservice-time-sequential.png')