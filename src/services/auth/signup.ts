import { baseURL } from "../../constants/api"
import type { ISignup } from "../../typings/IAuth"

export const signup = async (data: ISignup) => {
  const response = await fetch(`${baseURL}/auth/register`, {
    method: 'POST',
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(data)
  })

  return response.json()
}