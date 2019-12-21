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

bar = plt.bar(y_pos, performance, align='center', alpha=0.5, color=['blue', 'blue', 'blue', 'blue', 'orange', 'orange'])
plt.xticks(y_pos, objects)
plt.yticks(np.arange(0, 3000, 500))
plt.ylabel('Time (ms)')
plt.title('Mean Kernel Boot Time (Concurrent)')

plt.legend((bar[0], bar[4]), ('QEMU', 'firecracker'))

plt.gcf().subplots_adjust(bottom=0.30)
plt.xticks(rotation=45)

#plt.show()
plt.savefig({{ .Out }})