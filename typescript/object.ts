export class Obj {
  Name: string = "object";
  ID: number = -1;
  Contents: string[] = [];
  OBJ: Obj[] = [];

  constructor(name?: string, id?: number, contents?: string[], obj?: Obj[]) {
    this.Name = name ? name : "";
    if (id != undefined) this.ID = id;
    this.Contents = contents ? contents : [];
    this.OBJ = obj ? obj : [];
  }

  Equal(obj: Obj) {
    this.Name = obj.Name;
    this.ID = obj.ID;
    this.Contents = obj.Contents;
    this.OBJ = obj.OBJ;
  }

  IsEqual(obj: Obj): boolean {
    if (this.Name != obj.Name)
      return false;
    if (this.ID != obj.ID)
      return false;
    for (let i = 0; i < this.Contents.length; i++) {
      if (this.Contents[i] != obj.Contents[i])
        return false;
    }

    return true;
  }
}