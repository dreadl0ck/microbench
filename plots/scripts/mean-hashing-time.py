qemu = [
    79.404044,
    81.135267,
    81.719621,
    45.73264,
    89.245779,
    76.068501,
    73.38832,
    49.262559,
    82.260077,
    87.126867

]

fc = [
    96.117267,
    97.526515,
    110.004722,
    94.044048,
    92.982877,
    96.959051,
    46.832791,
    47.361873,
    93.6133,
    98.92226
]

qemu_multi = [
    69.0383,
    195.657,
    108.146,
    344.4,
    122.888,
    432.439,
    499.534,
    68.6498,
    305.47
]

fc_multi = [
    166.741,
    141.47,
    183.080,
    184.403,
    179.617,
    143.62,
    172.534,
    190.416,
    111.444,
    211.460
]

import statistics as stats
print(stats.mean(qemu))
print(stats.mean(qemu_multi))
print(stats.mean(fc))
print(stats.mean(fc_multi))

import matplotlib.pyplot as plt; plt.rcdefaults()
import numpy as np
import matplotlib.pyplot as plt

objects = ('QEMU', 'QEMU x10', 'Firecracker', 'Firecracker x10')
y_pos = np.arange(len(objects))
performance = [
    stats.mean(qemu),
    stats.mean(qemu_multi),
    stats.mean(fc),
    stats.mean(fc_multi)
]

bar = plt.bar(y_pos, performance, align='center', alpha=0.5, color=['orange', 'green', 'orange', 'green'])
plt.xticks(y_pos, objects)
plt.yticks(np.arange(0, 400, 100))
plt.ylabel('Time (ms)')
plt.title('Mean Hashing Time')
# 10 runs
plt.legend((bar[0], bar[1]), ('Single', 'Concurrent'))

plt.show()