import ValidValue from "../../utils/validValue";

export default class User {
    private id: string;
    private name: UserName;
    private password: UserPassword;

    constructor(id: string, name: UserName, password: UserPassword) {
        this.id = id;
        this.name = name;
        this.password = password;
    }
}

export class UserPassword extends ValidValue<string> {
    constructor(v: string) {
        super(v);
    }

    public valid(v: string): boolean {
        return true;
    }
}

export class UserName extends ValidValue<string> {
    constructor(value: string) {
        super(value);
    }

    public valid(v: string): boolean {
        return true;
    }
}
