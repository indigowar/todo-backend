import { Result } from '../../../utils/result';
import Rule from '../../../utils/rule';

export type TokenInformation = { id: string, time: number };

export interface TokenManagerImplementation {
	CreatToken (info: TokenInformation, sign_key: string): Result<string>;

	Verify (token: string, sign_key: string): Result<TokenInformation>;
}

export class TokenManager {
	public NewAccessToken (id: string): Result<string> {
		const exp_time = new Date().setTime(Date.now() + this.access_exp);
		const info: TokenInformation = {
			id: id,
			time: exp_time
		};
		return this.impl.CreatToken(info, this.sign_key);
	}

	public NewRefreshToken (id: string): Result<string> {
		const exp_time = new Date().setTime(Date.now() + this.refresh_exp);
		const info: TokenInformation = {
			id: id,
			time: exp_time
		};

		return this.impl.CreatToken(info, this.sign_key);
	}

	public Verify (token: string): Result<TokenInformation> {
		return this.impl.Verify(token, this.sign_key);
	}

	constructor (impl: TokenManagerImplementation, accessRule: Rule<number>, refreshRule: Rule<number>, signKeyRule: Rule<string>) {
		this.impl = impl;
		this.access_set_rule = accessRule;
		this.refresh_set_rule = refreshRule;
		this.sign_key_set_rule = signKeyRule;

		this.access_exp = 0;
		this.refresh_exp = 0;
		this.sign_key = '';
	}

	set accessExpire (minutes: number) {
		if (this.access_set_rule.isExecutedFor(minutes)) {
			this.access_exp = minutes;
		}
	}

	set refreshExpire (minutes: number) {
		if (this.refresh_set_rule.isExecutedFor(minutes)) {
			this.refresh_exp = minutes;
		}
	}

	set signKey (key: string) {
		if (this.sign_key_set_rule.isExecutedFor(key)) {
			this.sign_key = key;
		}
	}

	// rules for members
	private readonly access_set_rule: Rule<number>;
	private readonly refresh_set_rule: Rule<number>;
	private readonly sign_key_set_rule: Rule<string>;

	// members
	private access_exp: number;
	private refresh_exp: number;
	private sign_key: string;

	// implementation of token manager
	private impl: TokenManagerImplementation;
}
