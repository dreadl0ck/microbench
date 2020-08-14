import numpy as np
import matplotlib.pyplot as plt

firecracker_sequential = [
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

firecracker_sequential_C3 = [
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

firecracker_sequential_default_kernel = [
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

qemu_sequential_emulated = [
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

qemu_sequential = [
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



labels = ['qemu sequential','qemu sequential emulated','firecracker sequential','firecracker sequential C3','firecracker sequential default kernel']
y_pos = np.arange(len(labels))

data=[
	qemu_sequential,
	qemu_sequential_emulated,
	firecracker_sequential,
	firecracker_sequential_C3,
	firecracker_sequential_default_kernel
]

fig, ax = plt.subplots()
ax.set_title("Webservice Startup Time")
plt.ylabel('Time (ms)')
ax.boxplot(data, labels=labels)

plt.gcf().subplots_adjust(bottom=0.35)
plt.xticks(rotation=45)

#plt.show()
plt.savefig('plots/scripts/images/webservice-time-sequential.png')