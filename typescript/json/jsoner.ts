interface Param {
	Key: string;
	Value: string;
}
class Jsoner {
	Name: string;
	ID: number;
	Params: Param[];
	Array: number[];

	constructor(name?: string, id?: number, params?: Param[]) {
		this.Name = name ? name : "";
		this.ID = id != undefined ? id : 0;
		this.Params = params ? params : [];
		this.Array = [];
	}

	EncodeJson(): string {
		return JSON.stringify(this);
	}

	DecodeJson(text: string) {
		const json: Jsoner = JSON.parse(text);
		this.Equal(json);
	}

	Equal(json: Jsoner) {
		this.Name = json.Name;
		this.ID = json.ID;
		this.Params = json.Params;
		this.Array = json.Array;
	}
}

const json1 = new Jsoner("json1", 1, [{ Key: "key1", Value: "value1" }, { Key: "k", Value: "v" }]);
// let json1String = JSON.stringify(json1);
let jsonString = json1.EncodeJson();
console.log("Json string: ");
// console.log(json1String);
console.log(jsonString);
console.log("After parse");
const json: Jsoner = new Jsoner();
json.DecodeJson(jsonString);
console.log(json);