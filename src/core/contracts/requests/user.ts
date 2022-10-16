import { TokenRequest } from './general';

export type NamePasswordRequest = { name: string, password: string };

export type UpdateNameRequest = TokenRequest & { name: string };

export type UpdatePasswordRequest = TokenRequest & { name: string };
