import { NamePasswordRequest, UpdateNameRequest, UpdatePasswordRequest } from '../contracts/requests/user';
import { TokenRequest } from '../contracts/requests/general';
import { LoginResult, NameResult, RefreshResult } from '../contracts/responses/user';
import { StatusResult } from '../contracts/responses/general';

export interface UserService {
	Create (request: NamePasswordRequest): Promise<RefreshResult>;

	Delete (request: TokenRequest): Promise<StatusResult>;

	GetName (request: TokenRequest): Promise<NameResult>;

	UpdatePassword (request: UpdatePasswordRequest): Promise<StatusResult>;

	UpdateName (request: UpdateNameRequest): Promise<StatusResult>;

	Login (request: NamePasswordRequest): Promise<LoginResult>;
}
