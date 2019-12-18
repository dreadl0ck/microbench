from os import listdir
from os.path import isfile, join
from re import search

def getConcurrent(files):
  concurrentFiles = []
  for f in files:
    if search("(-[0-9]*-*\.log)", f):
      concurrentFiles.append(f)
  return concurrentFiles

def getSingle(files):
  singleFiles = []
  for f in files:
    if search("([0-9]\.[0-9]-*\.log)", f):
      singleFiles.append(f)
  return singleFiles

def getQemu(files):
  qemuFiles = []
  for f in files:
    if search("(qemu.*)", f):
      qemuFiles.append(f)
  return qemuFiles

def getFirecracker(files):
  fcFiles = []
  for f in files:
    if search("(fc.*)", f):
      fcFiles.append(f)
  return fcFiles

files = [f for f in listdir("./") if isfile(join("./", f))]
concurrentLogFiles = getConcurrent(files)
singleLogFiles = getSingle(files)
print(concurrentLogFiles)
print(singleLogFiles)
print(getQemu(files))
print(getFirecracker(files))
