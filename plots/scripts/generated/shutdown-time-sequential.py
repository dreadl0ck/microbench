import numpy as np
import matplotlib.pyplot as plt

firecracker_sequential = [
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

firecracker_sequential_C3 = [
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

firecracker_sequential_default_kernel = [
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

qemu_sequential_emulated = [
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

qemu_sequential = [
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
ax.set_title('Shutdown Time')
plt.ylabel('Time (ms)')
ax.boxplot(data, labels=labels)

plt.gcf().subplots_adjust(bottom=0.35)
plt.xticks(rotation=45)

#plt.show()
plt.savefig('plots/scripts/images/shutdown-time-sequential.png')