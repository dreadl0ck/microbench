import numpy as np
import matplotlib.pyplot as plt

{{ .Data }}

labels = [{{ .Objects }}]
y_pos = np.arange(len(labels))

data=[
{{ .Load }}
]

fig, ax = plt.subplots()
plt.ylabel('Time (ms)')
ax.boxplot(data, labels=labels)

plt.gcf().subplots_adjust(bottom=0.30)
plt.xticks(rotation=45)

#plt.show()
plt.savefig({{ .Out }})