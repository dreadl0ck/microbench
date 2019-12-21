{{ .Data }}

import statistics as stats
import matplotlib.pyplot as plt; plt.rcdefaults()
import numpy as np
import matplotlib.pyplot as plt

objects = ({{ .Objects }})
y_pos = np.arange(len(objects))
performance = [
{{ .Load }}
]

bar = plt.bar(y_pos, performance, align='center', alpha=0.5, color=['blue', 'blue', 'orange', 'orange', 'orange'])
plt.xticks(y_pos, objects)
#plt.yticks(np.arange(0, 1500, 200))
plt.ylabel('Time (ms)')
plt.title('Mean Web Service Startup Time')

plt.legend((bar[0], bar[2]), ('QEMU', 'firecracker'))

#plt.show()
plt.savefig({{ .Out }})