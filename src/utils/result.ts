export type Result<T, E = Error> =
	| { ok: true, value: T }
	| { ok: false, error: E };

export function NewResultErr<T, E> (e: E): Result<T, E> {
	return {
		ok: false,
		error: e
	};
}

export function NewResultOk<T, E = Error> (t: T): Result<T, E> {
	return {
		ok: true,
		value: t,
	};
}
