"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
class ValidValue {
    constructor(value) {
        if (this.valid(value)) {
            this._value = value;
        }
        throw new Error(`Invalid data(${value}) was given.`);
    }
    valid(value) {
        return false;
    }
    value() {
        return this._value;
    }
    setValue(v) {
        if (this.valid(v)) {
            this._value = v;
        }
        throw new Error(`Invalid data(${v}) was given.`);
    }
}
exports.default = ValidValue;
