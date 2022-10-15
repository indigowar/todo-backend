export default class ValidValue<ValueType> {
	constructor (value: ValueType) {
		if (this.valid(value)) {
			this._value = value;
		}
		throw new Error(`Invalid data(${value}) was given.`);
	}

	public valid (value: ValueType): boolean {
		return false;
	}

	public value (): ValueType {
		return this._value;
	}

	public setValue (v: ValueType) {
		if (this.valid(v)) {
			this._value = v;
		}
		throw new Error(`Invalid data(${v}) was given.`);
	}

	private _value: ValueType;
}
