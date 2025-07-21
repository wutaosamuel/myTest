import os

file_path = os.path.abspath(__file__)
print(file_path)

dir_path = os.path.dirname(os.path.abspath(__file__))
print(dir_path)