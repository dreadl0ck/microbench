{{.Data}}

import statistics as stats
{{ .Log }}

import matplotlib.pyplot as plt; plt.rcdefaults()
import numpy as np
import matplotlib.pyplot as plt

objects = ({{ .Objects }})

y_pos = np.arange(len(objects))
performance = [
{{ .Load }}
]

bar = plt.bar(y_pos, performance, align='center', alpha=0.5, color=['blue', 'green', 'orange'])
plt.xticks(y_pos, objects)
plt.yticks(np.arange(0, 100, 10))
plt.ylabel('Time (ms)')
plt.title('Mean Hashing Time')

plt.savefig({{ .Out }})