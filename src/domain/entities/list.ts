export default class List {
	constructor (id: string, name: string, owner: string, elements?: string[]) {
		this.id = id;
		this.name = name;
		this.owner = owner;
		this._elements = [];
		if (typeof elements !== 'undefined') {
			this._elements = elements;
		}
	}

	public readonly id: string;
	public readonly name: string;
	public readonly owner: string;

	public get elements () {
		return this._elements;
	}

	public add (s: string) {
		this._elements.push(s);
	}

	private readonly _elements: string[];
}
