import matplotlib.pyplot as plt; plt.rcdefaults()
import numpy as np
import matplotlib.pyplot as plt

import numpy as np
import matplotlib.pyplot as plt

{{ .Data }}

labels = ['host cpu','emulated cpu','T2 cpu','C3 cpu','default kernel']
y_pos = np.arange(len(labels))

import statistics as stats

web = [
	stats.mean(host_cpu_webservice),
	stats.mean(emulated_cpu_webservice),
	stats.mean(T2_cpu_webservice),
	stats.mean(C3_cpu_webservice),
	stats.mean(default_kernel_webservice),
]

kernel = [
	stats.mean(host_cpu_kernel_boot),
	stats.mean(emulated_cpu_kernel_boot),
	stats.mean(T2_cpu_kernel_boot),
	stats.mean(C3_cpu_kernel_boot),
	stats.mean(default_kernel_kernel_boot),
]

bar2 = plt.bar(y_pos, web, align='center', alpha=0.5,  color=['blue'])
bar1 = plt.bar(y_pos, kernel, align='center', alpha=0.5,  color=['green'])

plt.xticks(y_pos, labels)

#plt.yticks(np.arange(0, 1800, 200))
plt.ylabel('Time (ms)')
plt.title('Stacked Mean Kernel Boot and Web Service Startup Times')

plt.legend((bar1[0], bar2[0]), ('Kernel boot time', 'Web service startup time'))

plt.gcf().subplots_adjust(bottom=0.30)
plt.xticks(rotation=45)
plt.savefig({{ .Out }})