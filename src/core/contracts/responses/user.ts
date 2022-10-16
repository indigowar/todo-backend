import { Result } from '../../../utils/result';

export type NameResult = Result<string>;
export type AccessResult = Result<string>;
export type RefreshResult = Result<string>;

type BothTokens = { access: string, refresh: string };

export type LoginResult = Result<BothTokens>;
