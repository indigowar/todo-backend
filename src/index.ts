import {UserService} from "./core/services/services";
import UserSvc from "./services/user";
import HttpInfra from "./infra/ports/http/http";

(async () => {
    let user_service: UserService = new UserSvc();
    let http_infra = new HttpInfra(user_service);
    http_infra.run();
})();