import Runnable from '../../../core/infra/runnable';
import { UserService } from '../../../core/services/services';
import HTTPServer from './server/server';

export default class HttpInfra implements Runnable {
  constructor (userSvc: UserService) {
	this.user_service = userSvc;
	this.server = new HTTPServer();
  }

  public run (): void {
	this.server.run();
  }

  private user_service: UserService;
  private server: HTTPServer;
}
