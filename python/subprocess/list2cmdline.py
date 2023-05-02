import subprocess

lt = subprocess.list2cmdline(["python3", "--version"])
print(lt)