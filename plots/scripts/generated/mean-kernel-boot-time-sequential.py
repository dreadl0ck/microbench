firecracker_default_kernel = [
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

qemu_emulated_cpu = [
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

qemu_host_cpu = [
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

firecracker_T2 = [
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

firecracker_C3 = [
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



import statistics as stats
import matplotlib.pyplot as plt; plt.rcdefaults()
import numpy as np
import matplotlib.pyplot as plt

objects = ('qemu host cpu','qemu emulated cpu','firecracker T2','firecracker C3','firecracker default kernel')
y_pos = np.arange(len(objects))
performance = [
	stats.mean(qemu_host_cpu),
	stats.mean(qemu_emulated_cpu),
	stats.mean(firecracker_T2),
	stats.mean(firecracker_C3),
	stats.mean(firecracker_default_kernel)
]

bar = plt.bar(y_pos, performance, align='center', alpha=0.5, color=['blue', 'green', 'orange'])
plt.xticks(y_pos, objects)
plt.yticks(np.arange(0, 1300, 100))
plt.ylabel('Time (ms)')
plt.title('Mean Kernel Boot Time')
#plt.legend((bar[0], bar[1]), ('Single', 'Concurrent'))

plt.gcf().subplots_adjust(bottom=0.30)
plt.xticks(rotation=45)

#plt.show()
plt.savefig('plots/scripts/images/mean-kernel-boot-time-sequential.png')