T2_cpu = [
	2308.50958,
	2342.34196,
	2328.16217,
	2325.80154,
	2283.37688,
	2315.62314,
	2321.66207,
	2298.62009,
	2313.97414,
	2334.21134,
	2328.12113,
	2324.43674,
	2322.27291,
	2345.63835,
	2335.98221,
	2310.37922,
	2284.52232,
	2335.79038,
	2329.57054,
	2310.88430,
	2297.69598,
	2329.90613,
	2327.65408,
	2290.15581,
	2299.94529,
	2339.99714,
	2330.68161,
	2317.77157,
	2322.54290,
	2320.43439,
	2319.77881,
	2306.84751,
	2321.47187,
	2336.10753,
	2332.35128,
	2294.55484,
	2293.27687,
	2334.44313,
	2311.82083,
	2319.89240
]

C3_cpu = [
	2308.50958,
	2342.34196,
	2283.37688,
	2315.62314,
	2313.97414,
	2334.21134,
	2322.27291,
	2345.63835,
	2284.52232,
	2335.79038,
	2297.69598,
	2329.90613,
	2299.94529,
	2339.99714,
	2322.54290,
	2320.43439,
	2321.47187,
	2336.10753,
	2293.27687,
	2334.44313
]

default_kernel = [
	2342.34196,
	2328.16217,
	2315.62314,
	2321.66207,
	2334.21134,
	2328.12113,
	2345.63835,
	2335.98221,
	2335.79038,
	2329.57054,
	2329.90613,
	2327.65408,
	2339.99714,
	2330.68161,
	2320.43439,
	2319.77881,
	2336.10753,
	2332.35128,
	2334.44313,
	2311.82083
]

emulated_cpu = [
	10825.86588,
	10771.42507,
	10577.13343,
	10601.72151,
	10530.66527,
	10602.77075,
	10597.84565,
	10601.20782,
	10530.91069,
	10568.24941
]

host_cpu = [
	10761.49516,
	10711.91262,
	10799.39208,
	10788.80519,
	10792.56301,
	10761.45596,
	10818.13008,
	11537.79621,
	11480.42763,
	10744.87485
]



import statistics as stats
print("script: plots/scripts/mean-shutdown-time-sequential.py")
print("stats.mean(host_cpu):", stats.mean(host_cpu))
print("stats.mean(T2_cpu):", stats.mean(T2_cpu))
print("stats.mean(C3_cpu):", stats.mean(C3_cpu))
print("stats.mean(default_kernel):", stats.mean(default_kernel))
print("stats.mean(emulated_cpu):", stats.mean(emulated_cpu))


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