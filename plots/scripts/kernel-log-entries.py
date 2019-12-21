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
plt.yticks(np.arange(0, 220, 40))
plt.ylabel('Number of log entries')
plt.title('Kernel Log Entries')

plt.legend((bar[0], bar[2]), ('QEMU', 'firecracker'))

#plt.gcf().subplots_adjust(bottom=0.30)
#plt.xticks(rotation=45)

#plt.show()
plt.savefig({{ .Out }})