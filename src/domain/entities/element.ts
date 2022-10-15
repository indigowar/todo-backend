export default class Element {
	private readonly _id: string;
	private _value: string;
	private _done: boolean;

	constructor (id: string, value: string, done: boolean) {
		this._id = id;
		this._value = value;
		this._done = done;
	}

	public get status () {
		return this._done;
	}

	public toggleStatus () {
		this._done = !this._done;
	}

	public get id () {
		return this._id;
	}

	public get value () {
		return this._value;
	}

	public set value (value: string) {
		this._value = value;
	}
}
