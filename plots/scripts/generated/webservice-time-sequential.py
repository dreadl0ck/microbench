import numpy as np
import matplotlib.pyplot as plt

firecracker_sequential = [
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

firecracker_sequential_C3 = [
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

firecracker_sequential_default_kernel = [
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

qemu_sequential_emulated = [
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

qemu_sequential = [
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