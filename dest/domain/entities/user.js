"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.UserName = exports.UserPassword = void 0;
const validValue_1 = __importDefault(require("../../utils/validValue"));
class User {
    constructor(id, name, password) {
        this.id = id;
        this.name = name;
        this.password = password;
    }
}
exports.default = User;
class UserPassword extends validValue_1.default {
    constructor(v) {
        super(v);
    }
    valid(v) {
        return true;
    }
}
exports.UserPassword = UserPassword;
class UserName extends validValue_1.default {
    constructor(value) {
        super(value);
    }
    valid(v) {
        return true;
    }
}
exports.UserName = UserName;
