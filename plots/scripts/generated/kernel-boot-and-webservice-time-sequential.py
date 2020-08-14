import matplotlib.pyplot as plt; plt.rcdefaults()
import numpy as np
import matplotlib.pyplot as plt

import numpy as np
import matplotlib.pyplot as plt

default_kernel_kernel_boot = [
	809.67400,
	653.23300,
	797.03800,
	806.10400,
	795.63200,
	763.22300,
	820.67500,
	778.95400,
	807.32700,
	789.19600,
	786.12000,
	790.14600,
	794.05000,
	797.56900,
	789.26200,
	788.23500,
	801.87500,
	808.71600,
	796.51200,
	804.35000
]

default_kernel_webservice = [
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

emulated_cpu_kernel_boot = [
	1120.01600,
	1081.14700,
	1132.67100,
	1138.79800,
	1150.83300,
	1083.49400,
	1138.03700,
	1082.97900,
	1075.31700,
	1083.02800
]

emulated_cpu_webservice = [
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

host_cpu_kernel_boot = [
	848.15500,
	857.38500,
	843.29600,
	842.25100,
	853.88500,
	846.86300,
	855.40700,
	843.80400,
	845.97000,
	853.70900
]

T2_cpu_webservice = [
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

C3_cpu_kernel_boot = [
	823.27400,
	809.67400,
	843.16300,
	797.03800,
	844.86400,
	795.63200,
	830.17300,
	820.67500,
	841.22400,
	807.32700,
	878.80700,
	786.12000,
	857.73400,
	794.05000,
	878.16300,
	789.26200,
	843.24000,
	801.87500,
	823.76900,
	796.51200
]

C3_cpu_webservice = [
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

host_cpu_webservice = [
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

T2_cpu_kernel_boot = [
	823.27400,
	809.67400,
	653.23300,
	564.89600,
	843.16300,
	797.03800,
	806.10400,
	874.29800,
	844.86400,
	795.63200,
	763.22300,
	815.65700,
	830.17300,
	820.67500,
	778.95400,
	743.22400,
	841.22400,
	807.32700,
	789.19600,
	794.15800,
	878.80700,
	786.12000,
	790.14600,
	809.86700,
	857.73400,
	794.05000,
	797.56900,
	830.90300,
	878.16300,
	789.26200,
	788.23500,
	855.53800,
	843.24000,
	801.87500,
	808.71600,
	840.63900,
	823.76900,
	796.51200,
	804.35000,
	820.21700
]



labels = ['qemu host cpu','qemu emulated cpu','firecracker T2 cpu','firecracker C3 cpu','firecracker default kernel']
y_pos = np.arange(len(labels))

import statistics as stats

web = [
	stats.mean(host_cpu_webservice),
	stats.mean(emulated_cpu_webservice),
	stats.mean(T2_cpu_webservice),
	stats.mean(C3_cpu_webservice),
	stats.mean(default_kernel_webservice),
]

kernel = [
	stats.mean(host_cpu_kernel_boot),
	stats.mean(emulated_cpu_kernel_boot),
	stats.mean(T2_cpu_kernel_boot),
	stats.mean(C3_cpu_kernel_boot),
	stats.mean(default_kernel_kernel_boot),
]

bar2 = plt.bar(y_pos, web, align='center', alpha=0.5,  color=['orange'])
bar1 = plt.bar(y_pos, kernel, align='center', alpha=0.5,  color=['black'])

plt.xticks(y_pos, labels)

#plt.yticks(np.arange(0, 1800, 200))
plt.ylabel('Time (ms)')
plt.title('Stacked Mean Kernel Boot and Web Service Startup Times')

plt.legend((bar2[0], bar1[0]), ('Web service startup time', 'Kernel boot time'))

plt.gcf().subplots_adjust(bottom=0.30)
plt.xticks(rotation=45)
plt.savefig('plots/scripts/images/kernel-boot-and-webservice-time-sequential.png')