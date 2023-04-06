import type { ISignup } from "../../typings/IAuth"

export const signup = async (data: ISignup) => {
  const response = await fetch('http://127.0.0.1:2000/auth/register', {
    method: 'POST',
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(data)
  })

  return response.json()
}