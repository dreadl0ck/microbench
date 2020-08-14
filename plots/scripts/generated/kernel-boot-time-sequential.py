import numpy as np
import matplotlib.pyplot as plt

firecracker_sequential_default_kernel = [
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

qemu_sequential_emulated = [
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

qemu_sequential = [
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

firecracker_sequential = [
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

firecracker_sequential_C3 = [
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
ax.set_title("Kernel Boot Time")
plt.ylabel('Time (ms)')
ax.boxplot(data, labels=labels)

plt.gcf().subplots_adjust(bottom=0.30)
plt.xticks(rotation=45)

#plt.show()
plt.savefig('plots/scripts/images/kernel-boot-time-sequential.png')