T2_cpu = [
	48.85381,
	66.93205,
	72.74762,
	92.93075,
	77.65826,
	54.41868,
	54.31696,
	91.19947,
	61.04778,
	93.27167,
	49.48362,
	77.97944,
	81.36476,
	55.54702,
	57.23942,
	61.81553,
	62.62906,
	49.06423,
	85.82673,
	94.76023,
	95.92509,
	47.91143,
	51.24501,
	70.46086,
	65.67447,
	51.21641,
	76.87557,
	98.80264,
	86.23584,
	53.00187,
	50.28119,
	89.09932,
	44.96466,
	49.91625,
	44.99613,
	62.85774,
	65.13883,
	48.81749,
	48.41926,
	88.23731
]

C3_cpu = [
	48.85381,
	66.93205,
	77.65826,
	54.41868,
	61.04778,
	93.27167,
	81.36476,
	55.54702,
	62.62906,
	49.06423,
	95.92509,
	47.91143,
	65.67447,
	51.21641,
	86.23584,
	53.00187,
	44.96466,
	49.91625,
	65.13883,
	48.81749
]

default_kernel = [
	66.93205,
	72.74762,
	54.41868,
	54.31696,
	93.27167,
	49.48362,
	55.54702,
	57.23942,
	49.06423,
	85.82673,
	47.91143,
	51.24501,
	51.21641,
	76.87557,
	53.00187,
	50.28119,
	49.91625,
	44.99613,
	48.81749,
	48.41926
]

emulated_cpu = [
	57.65581,
	89.31157,
	63.61122,
	61.14589,
	56.12511,
	60.31525,
	49.23990,
	57.67386,
	70.25082,
	68.78634
]

host_cpu = [
	76.74898,
	70.99104,
	56.40613,
	66.40582,
	72.92878,
	102.56819,
	67.62190,
	63.77117,
	90.30249,
	57.60920
]



import statistics as stats
print("script: plots/scripts/mean-hashing-time-sequential.py")
print("stats.mean(T2_cpu):", stats.mean(T2_cpu))
print("stats.mean(C3_cpu):", stats.mean(C3_cpu))
print("stats.mean(default_kernel):", stats.mean(default_kernel))
print("stats.mean(emulated_cpu):", stats.mean(emulated_cpu))
print("stats.mean(host_cpu):", stats.mean(host_cpu))


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
#plt.yticks(np.arange(0, 100, 10))
plt.ylabel('Time (ms)')
plt.title('Mean Hashing Time SHA-256 100 x 1MB')

plt.legend((bar[0], bar[2]), ('QEMU', 'firecracker'))

#plt.gcf().subplots_adjust(bottom=0.30)
#plt.xticks(rotation=45)

plt.savefig('plots/scripts/images/mean-hashing-time-sequential.png')