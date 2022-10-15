import ValidValue from '../../utils/validValue';

export default class User {
	private readonly _id: string;

	constructor (id: string, name: UserName, password: UserPassword) {
		this._id = id;
		this.name = name;
		this.password = password;
	}

	public name: UserName;
	public password: UserPassword;

	public get id () {
		return this._id;
	}
}

export class UserPassword extends ValidValue<string> {
	constructor (v: string) {
		super(v);
	}

	public valid (v: string): boolean {
		return true;
	}
}

export class UserName extends ValidValue<string> {
	constructor (value: string) {
		super(value);
	}

	public valid (v: string): boolean {
		return true;
	}
}
