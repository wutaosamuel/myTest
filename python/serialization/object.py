from decimal import Decimal


class Obj:
    Name = ""
    ID = -1

    def __init__(self, name="", id=-1):
        self.Name = name
        self.ID = id

    def Equal(self, obj):
        self.Name = obj.Name
        self.ID = obj.ID

    def ToDict(self):
        return {
            "Name": self.Name,
            "ID": self.ID
        }

    def Restore(self, dict={}):
        self.Name = dict["Name"]
        self.ID = dict["ID"]


class ObjS:
    Name = ""
    Objects = []

    def __init__(self, name="", obj_s=[]):
        self.Name = name
        self.Objects = obj_s

    def ToDict(self):
        d = []
        for o in self.Objects:
            obj = Obj()
            obj.Equal(o)
            d.append(obj.ToDict())

        return {
            "Name": self.Name,
            "Objects": d
        }

    def Restore(self, dict={}):
        self.Name = dict["Name"]
        objects_dict = dict["Objects"]
        for d in objects_dict:
            o = Obj()
            o.Restore(d)
            self.Objects.append(o)
