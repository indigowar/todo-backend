import { UserService } from '../core/services/services';
import { NamePasswordRequest, UpdateNameRequest, UpdatePasswordRequest } from '../core/contracts/requests/user';
import { TokenRequest } from '../core/contracts/requests/general';
import { StatusResult } from '../core/contracts/responses/general';
import { LoginResult, NameResult, RefreshResult } from '../core/contracts/responses/user';
import { NewResultErr } from '../utils/result';

export default class UserSvc implements UserService {
	Create (request: NamePasswordRequest): Promise<RefreshResult> {
		return Promise.resolve(NewResultErr(new Error('UnImplemented')));
	}

	Delete (request: TokenRequest): Promise<StatusResult> {
		return Promise.resolve(NewResultErr(new Error('UnImplemented')));
	}

	GetName (request: TokenRequest): Promise<NameResult> {
		return Promise.resolve(NewResultErr(new Error('UnImplemented')));
	}

	Login (request: NamePasswordRequest): Promise<LoginResult> {
		return Promise.resolve(NewResultErr(new Error('UnImplemented')));
	}

	UpdateName (request: UpdateNameRequest): Promise<StatusResult> {
		return Promise.resolve(NewResultErr(new Error('UnImplemented')));
	}

	UpdatePassword (request: UpdatePasswordRequest): Promise<StatusResult> {
		return Promise.resolve(NewResultErr(new Error('UnImplemented')));
	}
}
