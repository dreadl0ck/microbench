T2_cpu = [
	2291.26421,
	2343.77812,
	2305.53904,
	2311.06299,
	2331.43214,
	2346.85258,
	2312.15032,
	2304.59621,
	2296.29995,
	2336.38096,
	2308.60528,
	2308.55577,
	2313.71220,
	2336.00347,
	2323.19398,
	2311.86450,
	2279.76231,
	2313.96348,
	2336.31811,
	2304.46736,
	2306.80928,
	2338.86096,
	2316.77138,
	2321.40447,
	2298.88454,
	2321.19846,
	2317.69327,
	2284.50930,
	2296.59859,
	2330.62205,
	2308.32935,
	2310.95344,
	2309.97688,
	2304.00803,
	2315.54799,
	2289.67137,
	2310.99843,
	2357.18932,
	2301.18490,
	2289.88585
]

C3_cpu = [
	2291.26421,
	2343.77812,
	2331.43214,
	2346.85258,
	2296.29995,
	2336.38096,
	2313.71220,
	2336.00347,
	2279.76231,
	2313.96348,
	2306.80928,
	2338.86096,
	2298.88454,
	2321.19846,
	2296.59859,
	2330.62205,
	2309.97688,
	2304.00803,
	2310.99843,
	2357.18932
]

default_kernel = [
	2343.77812,
	2305.53904,
	2346.85258,
	2312.15032,
	2336.38096,
	2308.60528,
	2336.00347,
	2323.19398,
	2313.96348,
	2336.31811,
	2338.86096,
	2316.77138,
	2321.19846,
	2317.69327,
	2330.62205,
	2308.32935,
	2304.00803,
	2315.54799,
	2357.18932,
	2301.18490
]

emulated_cpu = [
	10638.79766,
	10608.70419,
	10574.32269,
	10568.88791,
	10554.51985,
	10547.09167,
	10556.00174,
	10569.52786,
	10619.52312,
	10568.64811
]

host_cpu = [
	10803.94492,
	10818.61036,
	10742.04889,
	10722.19653,
	10658.34818,
	10739.34372,
	10719.92125,
	10683.85985,
	10789.22607,
	10737.22514
]



import statistics as stats
print("script: plots/scripts/mean-shutdown-time-sequential.py")
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
#plt.yticks(np.arange(0, 2300, 1500))
plt.ylabel('Time (ms)')
plt.title('Mean Shutdown Time')

plt.legend((bar[0], bar[2]), ('QEMU', 'firecracker'))

#plt.gcf().subplots_adjust(bottom=0.20)

#plt.show()
plt.savefig('plots/scripts/images/mean-shutdown-time-sequential.png')