import { UserService } from '../core/services/services';
import { NamePasswordRequest, UpdateNameRequest, UpdatePasswordRequest } from '../core/contracts/requests/user';
import { TokenRequest } from '../core/contracts/requests/general';
import { StatusResult } from '../core/contracts/responses/general';
import { LoginResult, NameResult, RefreshResult } from '../core/contracts/responses/user';
import { NewResultErr } from '../utils/result';
import UserStorage from '../core/adapters/storages/user_storage';
import { TokenManagerImplementation } from '../core/adapters/auth/token_manager';

export default class UserSvc implements UserService {
	private storage: UserStorage;
	private token_manager: TokenManagerImplementation;

	constructor (storage: UserStorage, tm: TokenManagerImplementation) {
		this.storage = storage;
		this.token_manager = tm;
	}

	public Create (request: NamePasswordRequest): Promise<RefreshResult> {
		return Promise.resolve(NewResultErr(new Error('Not Implemented')));
	}

	public Delete (request: TokenRequest): Promise<StatusResult> {
		return Promise.resolve(NewResultErr(new Error('Not Implemented')));
	}

	public GetName (request: TokenRequest): Promise<NameResult> {
		return Promise.resolve(NewResultErr(new Error('Not Implemented')));
	}

	public Login (request: NamePasswordRequest): Promise<LoginResult> {
		return Promise.resolve(NewResultErr(new Error('Not Implemented')));
	}

	public UpdateName (request: UpdateNameRequest): Promise<StatusResult> {
		return Promise.resolve(NewResultErr(new Error('Not Implemented')));
	}

	public UpdatePassword (request: UpdatePasswordRequest): Promise<StatusResult> {
		return Promise.resolve(NewResultErr(new Error('Not Implemented')));
	}
}
