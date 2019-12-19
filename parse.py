from os import listdir
from os.path import isfile, join
from re import search

def getStringFromFile(f):
  file_obj = open(f)
  return file_obj.read()

# Get log values
def getKernelBootTime(file):
  with open(file, "r") as f:
    for line in f:
      if search("kernel boot", line):
        f.close()
        delta = str.split(line)[8]
        return delta.replace('delta', '').replace('ms', '').replace('=','')

def getWebServiceStartupTime(file):
  with open(file, "r") as f:
    for line in f:
      if search("time until HTTP reply from webservice", line):
        f.close()
        delta = str.split(line)[10]
        return delta.replace('delta', '').replace('s', '').replace('=','')

def getHashBenchTime(file):
  with open(file, "r") as f:
    for line in f:
      if search("hash loop benchmark", line):
        f.close()
        delta = str.split(line)[7]
        return delta.replace('ms,', '')

# Get files
# @click.command()
# @click.argument('multi')
# def multi(multi):
#   if not multi:
#     getSingle()
#   else:
#     getConcurrent()

def getSingle(files):
  singleFiles = []
  for f in files:
    if search("(-[0-9]*-*\.log)", f):
      singleFiles.append(f)
  return singleFiles

def getConcurrent(files):
  concurrentFiles = []
  for f in files:
    if search("([0-9]\.[0-9]-*\.log)", f):
      concurrentFiles.append(f)
  return concurrentFiles

def getQemu(files):
  qemuFiles = []
  for f in files:
    if search("(qemu.*)", f):
      qemuFiles.append(f)
  return qemuFiles

def getFirecracker(files):
  firecrackerFiles = []
  for f in files:
    if search("(fc.*)", f):
      firecrackerFiles.append(f)
  return firecrackerFiles

files = [f for f in listdir("./") if isfile(join("./", f))]
concurrentLogFiles = getConcurrent(files)
singleLogFiles = getSingle(files)


print(getQemu(getConcurrent(files)))

for f in getQemu(getConcurrent(files)):
  print(getWebServiceStartupTime(f))

