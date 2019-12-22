C3_cpu = [
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

default_kernel = [
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

emulated_cpu = [
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

host_cpu = [
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

T2_cpu = [
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