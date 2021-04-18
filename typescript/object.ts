export class Obj {
  Name: string = "object";
  ID: number = -1;

  constructor(name?: string, id?: number) {
    this.Name = name ? name : "";
    if (id != undefined) this.ID = id;
  }
}