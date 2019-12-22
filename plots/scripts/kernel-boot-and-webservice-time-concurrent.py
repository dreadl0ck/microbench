import matplotlib.pyplot as plt; plt.rcdefaults()
import numpy as np
import matplotlib.pyplot as plt

import numpy as np
import matplotlib.pyplot as plt

{{ .Data }}

labels = ['qemu x10','qemu x10 emulated','qemu x20','qemu x20 emulated','firecracker x10','firecracker x20']
y_pos = np.arange(len(labels))

import statistics as stats

web = [
	stats.mean(qemu_x10_webservice),
	stats.mean(qemu_x10_webservice),
	stats.mean(qemu_x20_webservice),
	stats.mean(qemu_x20_emulated_webservice),
	stats.mean(firecracker_x10_webservice),
	stats.mean(firecracker_x20_webservice)
]

kernel = [
	stats.mean(qemu_x10_kernel_boot),
	stats.mean(qemu_x10_emulated_kernel_boot),
	stats.mean(qemu_x20_kernel_boot),
	stats.mean(qemu_x20_emulated_kernel_boot),
	stats.mean(firecracker_x10_kernel_boot),
	stats.mean(firecracker_x20_kernel_boot)
]

bar2 = plt.bar(y_pos, web, align='center', alpha=0.5,  color=['orange'])
bar1 = plt.bar(y_pos, kernel, align='center', alpha=0.5,  color=['black'])

plt.xticks(y_pos, labels)

#plt.yticks(np.arange(0, 1800, 200))
plt.ylabel('Time (ms)')
plt.title('Stacked Mean Kernel Boot and Web Service Startup Times (Concurrent)')

#plt.legend((bar1[0], bar2[0]), ('Kernel boot time', 'Web service startup time'))
plt.legend((bar2[0], bar1[0]), ('Web service startup time', 'Kernel boot time'))

plt.gcf().subplots_adjust(bottom=0.30)
plt.xticks(rotation=45)
plt.savefig({{ .Out }})