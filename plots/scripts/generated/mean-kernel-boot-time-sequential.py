T2_cpu = [
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

C3_cpu = [
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

default_kernel = [
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

emulated_cpu = [
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

host_cpu = [
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



import statistics as stats
print("script: plots/scripts/mean-kernel-boot-time-sequential.py")
print("stats.mean(T2_cpu):", stats.mean(T2_cpu))
print("stats.mean(C3_cpu):", stats.mean(C3_cpu))
print("stats.mean(default_kernel):", stats.mean(default_kernel))
print("stats.mean(emulated_cpu):", stats.mean(emulated_cpu))
print("stats.mean(host_cpu):", stats.mean(host_cpu))


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
#plt.yticks(np.arange(0, 1300, 100))
plt.ylabel('Time (ms)')
plt.title('Mean Kernel Boot Time')

plt.legend((bar[0], bar[2]), ('QEMU', 'firecracker'))

#plt.gcf().subplots_adjust(bottom=0.20)

#plt.show()
plt.savefig('plots/scripts/images/mean-kernel-boot-time-sequential.png')