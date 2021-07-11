import os
import json
import pickle
from .object import Obj, ObjS


def main():
    o1 = Obj("object1", 1)
    print("o1.__dict__ ->", o1.__dict__)
    o2 = Obj("object2", 2)

    objects = ObjS("objects", [])
    objects.Objects.append(o1)
    objects.Objects.append(o2)
    print("objects.ToDict() -> is okay", objects.ToDict())
    #json_dumps = json.dumps(objects.ToDict())
    #print("dumps", json_dumps)
    #pickle_dumps = pickle.dumps(objects.ToDict())
    #print("pickle", pickle_dumps)

    json_file = "./dump.json"
    with open(json_file, "w") as f:
        json.dump(objects.ToDict(), f)
    pickle_file = "./dump.pkl"
    with open(pickle_file, "wb") as f:
        pickle.dump(objects.ToDict(), f)

    restore(json_file, pickle_file)


def restore(json_filename="", pickle_filename=""):
    if not json_filename == "":
        with open(json_filename, "rb") as f:
            json_str = json.load(f)
            restore = ObjS("os", [])
            restore.Restore(json_str)
            os.remove(json_filename)
            print("Restore from json str -> ", restore.ToDict())
    print()
    print("Json diff here")
    print()
    if not pickle_filename == "":
        with open(pickle_filename, "rb") as f:
            pickle_str = pickle.load(f)
            restore = ObjS("os", [])
            restore.Restore(pickle_str)
            os.remove(pickle_filename)
            print("Restore from pickle str -> ", restore.ToDict())


if __name__ == "__main__":
    main()
