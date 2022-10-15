"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const server_1 = __importDefault(require("./server/server"));
class HttpInfra {
    constructor(userSvc) {
        this.user_service = userSvc;
        this.server = new server_1.default();
    }
    run() {
        this.server.run();
    }
}
exports.default = HttpInfra;
