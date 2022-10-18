import { TokenInformation, TokenManagerImplementation } from '../../../../core/adapters/auth/token_manager';
import { NewResultErr, NewResultOk, Result } from '../../../../utils/result';
import { JwtPayload, sign, verify } from 'jsonwebtoken';

export default class JwtImplementation implements TokenManagerImplementation {
	public CreatToken (info: TokenInformation, sign_key: string): Result<string> {
		const payload: JwtPayload = {
			sub: info.id,
			exp: info.time,
		};

		return NewResultOk(sign(payload, sign_key));
	}

	public ReadToken (token: string, sign_key: string): Result<TokenInformation> {
		try {
			const payload = verify(token, sign_key);
			if (typeof payload.sub === 'string') {
				const info: TokenInformation = {
					id: payload.sub,
					time: 0,
				};
				return NewResultOk(info);
			}

			return NewResultErr(new Error('token has no subject'));
		} catch (e) {
			return NewResultErr(new Error('reading ended with error'));
		}
	}
}
