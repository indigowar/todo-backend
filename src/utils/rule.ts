export default class Rule<T> {
	private readonly rule: (t: T) => boolean;

	constructor (r: (t: T) => boolean) {
		this.rule = r;
	}

	public isExecutedFor (t: T): boolean {
		return this.rule(t);
	}
}
