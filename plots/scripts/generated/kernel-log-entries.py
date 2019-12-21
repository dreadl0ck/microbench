firecracker_T2 = [
	213000.00000,
	195000.00000,
	195000.00000,
	213000.00000,
	214000.00000,
	196000.00000,
	197000.00000,
	213000.00000,
	214000.00000,
	197000.00000,
	196000.00000,
	214000.00000,
	213000.00000,
	196000.00000,
	197000.00000,
	213000.00000,
	213000.00000,
	196000.00000,
	196000.00000,
	213000.00000,
	213000.00000,
	196000.00000,
	196000.00000,
	213000.00000,
	214000.00000,
	196000.00000,
	196000.00000,
	213000.00000,
	213000.00000,
	197000.00000,
	197000.00000,
	214000.00000,
	213000.00000,
	197000.00000,
	196000.00000,
	213000.00000,
	214000.00000,
	196000.00000,
	197000.00000,
	213000.00000
]

firecracker_C3 = [
	213000.00000,
	195000.00000,
	214000.00000,
	196000.00000,
	214000.00000,
	197000.00000,
	213000.00000,
	196000.00000,
	213000.00000,
	196000.00000,
	213000.00000,
	196000.00000,
	214000.00000,
	196000.00000,
	213000.00000,
	197000.00000,
	213000.00000,
	197000.00000,
	214000.00000,
	196000.00000
]

firecracker_default_kernel = [
	195000.00000,
	195000.00000,
	196000.00000,
	197000.00000,
	197000.00000,
	196000.00000,
	196000.00000,
	197000.00000,
	196000.00000,
	196000.00000,
	196000.00000,
	196000.00000,
	196000.00000,
	196000.00000,
	197000.00000,
	197000.00000,
	197000.00000,
	196000.00000,
	196000.00000,
	197000.00000
]

qemu_emulated_cpu = [
	215000.00000,
	216000.00000,
	216000.00000,
	216000.00000,
	216000.00000,
	216000.00000,
	216000.00000,
	216000.00000,
	216000.00000,
	216000.00000
]

qemu_host_cpu = [
	227000.00000,
	228000.00000,
	229000.00000,
	229000.00000,
	228000.00000,
	228000.00000,
	228000.00000,
	228000.00000,
	228000.00000,
	228000.00000
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

bar = plt.bar(y_pos, performance, align='center', alpha=0.5, color=['orange', 'green', 'orange', 'green'])
plt.xticks(y_pos, objects)
plt.yticks(np.arange(0, 200, 50))
plt.ylabel('Number of log entries')
plt.title('Kernel Log Entries')
#plt.legend((bar[0], bar[1]), ('Single', 'Concurrent'))

plt.gcf().subplots_adjust(bottom=0.30)
plt.xticks(rotation=45)

#plt.show()
plt.savefig('plots/scripts/images/kernel-log-entries.png')