import { UserService } from './core/services/services';
import UserSvc from './services/user';
import HttpInfra from './infra/ports/http/http';

(async () => {
  const user_service: UserService = new UserSvc();
  const http_infra = new HttpInfra(user_service);
  http_infra.run();
})();
