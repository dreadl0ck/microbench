default_kernel = [
	195.00000,
	195.00000,
	196.00000,
	197.00000,
	197.00000,
	196.00000,
	196.00000,
	197.00000,
	196.00000,
	196.00000,
	196.00000,
	196.00000,
	196.00000,
	196.00000,
	197.00000,
	197.00000,
	197.00000,
	196.00000,
	196.00000,
	197.00000
]

emulated_cpu = [
	215.00000,
	216.00000,
	216.00000,
	216.00000,
	216.00000,
	216.00000,
	216.00000,
	216.00000,
	216.00000,
	216.00000
]

host_cpu = [
	227.00000,
	228.00000,
	229.00000,
	229.00000,
	228.00000,
	228.00000,
	228.00000,
	228.00000,
	228.00000,
	228.00000
]

T2_cpu = [
	213.00000,
	195.00000,
	195.00000,
	213.00000,
	214.00000,
	196.00000,
	197.00000,
	213.00000,
	214.00000,
	197.00000,
	196.00000,
	214.00000,
	213.00000,
	196.00000,
	197.00000,
	213.00000,
	213.00000,
	196.00000,
	196.00000,
	213.00000,
	213.00000,
	196.00000,
	196.00000,
	213.00000,
	214.00000,
	196.00000,
	196.00000,
	213.00000,
	213.00000,
	197.00000,
	197.00000,
	214.00000,
	213.00000,
	197.00000,
	196.00000,
	213.00000,
	214.00000,
	196.00000,
	197.00000,
	213.00000
]

C3_cpu = [
	213.00000,
	195.00000,
	214.00000,
	196.00000,
	214.00000,
	197.00000,
	213.00000,
	196.00000,
	213.00000,
	196.00000,
	213.00000,
	196.00000,
	214.00000,
	196.00000,
	213.00000,
	197.00000,
	213.00000,
	197.00000,
	214.00000,
	196.00000
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
plt.yticks(np.arange(0, 220, 40))
plt.ylabel('Number of log entries')
plt.title('Kernel Log Entries')

plt.legend((bar[0], bar[2]), ('QEMU', 'firecracker'))

#plt.gcf().subplots_adjust(bottom=0.30)
#plt.xticks(rotation=45)

#plt.show()
plt.savefig('plots/scripts/images/kernel-log-entries.png')