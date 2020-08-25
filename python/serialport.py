import glob
import serial
import platform

def list_serial_ports():
  system_name = platform.system()
  if system_name == "Windows":
    available = []
    for i in range(256):
      try: 
        s = serial.Serial(i)
        available.append(i)
        s.close()
      except serial.SerialException:
        pass
    return available
  elif system_name == "Darwin":
    return glob.glob("/dev/tty*") + glob.glob("/dev/cu*")
  else:
    return glob.glob("/dev/ttyS*") + glob.glob("/dev/ttyUSB*")

if __name__ == "__main__":
  print(list_serial_ports())