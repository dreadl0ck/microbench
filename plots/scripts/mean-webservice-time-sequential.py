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

plt.bar(y_pos, performance, align='center', alpha=0.5, color=['blue', 'green', 'orange'])
plt.xticks(y_pos, objects)
plt.yticks(np.arange(0, 1500, 200))
plt.ylabel('Time (ms)')
plt.title('Mean Web Service Startup Time')

plt.gcf().subplots_adjust(bottom=0.30)
plt.xticks(rotation=45)

#plt.show()
plt.savefig({{ .Out }})