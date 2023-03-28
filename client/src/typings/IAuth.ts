export interface IAuth {
  nickname: string,
  password: string
}

export interface ISignup extends IAuth {
  description?: string,
}