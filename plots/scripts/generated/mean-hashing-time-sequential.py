firecracker_T2 = [
	91.46145,
	122.82453,
	89.34523,
	122.29510,
	58.67076,
	105.04158,
	104.42153,
	58.55728,
	106.47532,
	98.12093,
	104.89656,
	89.17962,
	86.43155,
	88.79501,
	98.89055,
	102.21036,
	54.94133,
	100.87000,
	82.99165,
	57.05992,
	110.07188,
	85.52291,
	72.47333,
	113.69101,
	92.75886,
	101.45224,
	107.40379,
	63.09834,
	103.88068,
	80.19164,
	109.31411,
	67.29307,
	85.94929,
	92.97686,
	66.28031,
	100.26759,
	92.08320,
	90.28110,
	94.17809,
	98.93750
]

firecracker_C3 = [
	91.46145,
	122.82453,
	58.67076,
	105.04158,
	106.47532,
	98.12093,
	86.43155,
	88.79501,
	54.94133,
	100.87000,
	110.07188,
	85.52291,
	92.75886,
	101.45224,
	103.88068,
	80.19164,
	85.94929,
	92.97686,
	92.08320,
	90.28110
]

firecracker_default_kernel = [
	122.82453,
	89.34523,
	105.04158,
	104.42153,
	98.12093,
	104.89656,
	88.79501,
	98.89055,
	100.87000,
	82.99165,
	85.52291,
	72.47333,
	101.45224,
	107.40379,
	80.19164,
	109.31411,
	92.97686,
	66.28031,
	90.28110,
	94.17809
]

qemu_emulated_cpu = [
	79.83308,
	56.86711,
	48.73478,
	82.62705,
	77.75761,
	90.43223,
	65.42166,
	76.74701,
	48.96195,
	82.90538
]

qemu_host_cpu = [
	64.47949,
	85.34255,
	43.27411,
	46.60391,
	70.67962,
	61.71825,
	83.48621,
	76.69465,
	68.02085,
	80.86654
]



import statistics as stats
print(stats.mean(qemu_emulated_cpu))
print(stats.mean(qemu_host_cpu))
print(stats.mean(firecracker_T2))
print(stats.mean(firecracker_C3))
print(stats.mean(firecracker_default_kernel))


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
plt.yticks(np.arange(0, 100, 10))
plt.ylabel('Time (ms)')
plt.title('Mean Hashing Time')

plt.gcf().subplots_adjust(bottom=0.30)
plt.xticks(rotation=45)

plt.savefig('plots/scripts/images/mean-hashing-time-sequential.png')