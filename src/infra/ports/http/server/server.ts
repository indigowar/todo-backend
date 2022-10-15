import express, { Express, Request, Response } from 'express';
import Runnable from '../../../../core/infra/runnable';

export default class HTTPServer implements Runnable {
	private app: Express;
	private readonly port: number;

	constructor () {
		this.app = express();
		this.port = 8000;
		// make some preparations
		this.app.get('/', (req: Request, res: Response) => {
			console.log(`got a request from ${req.ip} to ${req.path}`);
			res.send(`hello, from ${req.path}`);
		});
	}

	run (): void {
		this.app.listen(this.port, () => console.log(`running on http://127.0.0.1:${this.port}`));
	}
}
