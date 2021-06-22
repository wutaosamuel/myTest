from decimal import Decimal


class Obj:
    Name = ""
    ID = -1

    def __init__(self, name="", id=-1):
        if name == "" or id == -1:
            return False
        self.Name = name
        self.ID = id
