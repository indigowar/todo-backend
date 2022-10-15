"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
class HTTPServer {
    constructor() {
        this.app = (0, express_1.default)();
        this.port = 8000;
        // make some preparations
        this.app.get('/', (req, res) => {
            console.log(`got a request from ${req.ip} to ${req.path}`);
            res.send(`hello, from ${req.path}`);
        });
    }
    run() {
        this.app.listen(this.port, () => console.log(`running on http://127.0.0.1:${this.port}`));
    }
}
exports.default = HTTPServer;
