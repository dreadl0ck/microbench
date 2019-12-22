import matplotlib.pyplot as plt; plt.rcdefaults()
import numpy as np
import matplotlib.pyplot as plt

import numpy as np
import matplotlib.pyplot as plt

T2_cpu_kernel_boot = [
	927.48300,
	886.07500,
	724.23700,
	677.63500,
	827.18300,
	869.90400,
	863.57600,
	366.07800,
	913.84300,
	851.78400,
	861.95500,
	894.33300,
	506.69100,
	863.79500,
	878.91900,
	428.19200,
	300.70300,
	882.32500,
	863.36400,
	563.08900,
	908.47500,
	867.60700,
	873.27500,
	234.84300,
	926.29900,
	877.94900,
	862.74200,
	526.01100,
	901.96100,
	859.73000,
	875.10100,
	747.47000,
	295.35300,
	874.85900,
	867.29300,
	525.60300,
	918.10900,
	865.40500,
	849.13900,
	915.19300
]

T2_cpu_webservice = [
	1190.23436,
	1253.17204,
	1201.01271,
	1222.61022,
	1031.47671,
	1257.38374,
	1259.45029,
	1033.03587,
	1203.33090,
	1011.06872,
	1029.60918,
	1199.01270,
	1160.92970,
	1020.95356,
	1260.95425,
	1198.52139,
	1009.99655,
	1253.76479,
	1019.99951,
	1022.17308,
	1997.25129,
	1250.21300,
	1035.67160,
	1159.91482,
	1204.60626,
	1266.71832,
	1255.39149,
	1023.08462,
	1187.27939,
	1030.35450,
	1257.55789,
	1007.45253,
	1182.59821,
	1038.79605,
	1021.24892,
	1189.53622,
	1197.17705,
	1264.89341,
	1008.33589,
	2006.94145
]

emulated_cpu_kernel_boot = [
	1040.95300,
	1145.67800,
	1139.67400,
	1051.04500,
	1059.45200,
	1067.81100,
	1063.05900,
	1089.26000,
	1152.37600,
	1033.63700
]

host_cpu_webservice = [
	1106.60577,
	1174.30180,
	2039.68959,
	2042.32914,
	1144.12149,
	1141.54589,
	1167.05263,
	1147.44122,
	1112.47571,
	1170.27711
]

C3_cpu_kernel_boot = [
	927.48300,
	886.07500,
	827.18300,
	869.90400,
	913.84300,
	851.78400,
	506.69100,
	863.79500,
	300.70300,
	882.32500,
	908.47500,
	867.60700,
	926.29900,
	877.94900,
	901.96100,
	859.73000,
	295.35300,
	874.85900,
	918.10900,
	865.40500
]

C3_cpu_webservice = [
	1190.23436,
	1253.17204,
	1031.47671,
	1257.38374,
	1203.33090,
	1011.06872,
	1160.92970,
	1020.95356,
	1009.99655,
	1253.76479,
	1997.25129,
	1250.21300,
	1204.60626,
	1266.71832,
	1187.27939,
	1030.35450,
	1182.59821,
	1038.79605,
	1197.17705,
	1264.89341
]

default_kernel_kernel_boot = [
	886.07500,
	724.23700,
	869.90400,
	863.57600,
	851.78400,
	861.95500,
	863.79500,
	878.91900,
	882.32500,
	863.36400,
	867.60700,
	873.27500,
	877.94900,
	862.74200,
	859.73000,
	875.10100,
	874.85900,
	867.29300,
	865.40500,
	849.13900
]

default_kernel_webservice = [
	1253.17204,
	1201.01271,
	1257.38374,
	1259.45029,
	1011.06872,
	1029.60918,
	1020.95356,
	1260.95425,
	1253.76479,
	1019.99951,
	1250.21300,
	1035.67160,
	1266.71832,
	1255.39149,
	1030.35450,
	1257.55789,
	1038.79605,
	1021.24892,
	1264.89341,
	1008.33589
]

emulated_cpu_webservice = [
	1450.47179,
	2043.29865,
	2062.87465,
	1361.20997,
	1394.29882,
	1380.92312,
	1382.55454,
	1391.98473,
	2050.92573,
	1340.94562
]

host_cpu_kernel_boot = [
	844.09100,
	860.92600,
	1857.71200,
	1862.11400,
	853.24500,
	870.04100,
	860.37100,
	853.40300,
	856.47700,
	841.76000
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