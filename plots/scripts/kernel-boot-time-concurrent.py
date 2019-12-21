import numpy as np
import matplotlib.pyplot as plt

{{ .Data }}

labels = [{{ .Objects }}]
y_pos = np.arange(len(labels))

data=[
{{ .Load }}
]

fig, ax = plt.subplots()
ax.set_title("Kernel Boot Time (Concurrent)")
plt.ylabel('Time (ms)')

#plt.yticks(np.arange(0, 4000, 500))
ax.boxplot(data, labels=labels)

plt.gcf().subplots_adjust(bottom=0.30)
plt.xticks(rotation=45)

#plt.show()
plt.savefig({{ .Out }})