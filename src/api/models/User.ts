interface UserInterface {
  id: number;
  email: string;
  username: string;
  role: string;
  password: string;
}

export class User implements UserInterface {
  id: number;
  email: string;
  username: string;
  role: string;
  password: string;

  constructor(user: any) {
    this.id = user.id;
    this.email = user.email;
    this.username = user.username;
    this.role = user.role;
    this.password = user.password;
  }
}
