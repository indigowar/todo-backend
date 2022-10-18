import User from '../../../domain/entities/user';
import { Result } from '../../../utils/result';

export default interface UserStorage {
	GetByID (id: string): Promise<Result<User>>;

	GetByName (name: string): Promise<Result<User>>;

	Add (user: User): Promise<Result<boolean>>;

	Delete (id: string): Promise<Result<boolean>>;

	Update (user: User): Promise<Result<boolean>>;
}
