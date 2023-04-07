import { baseURL } from "../../constants/api"

export const getMe = async (token: string) => {
  const response = await fetch(`${baseURL}/auth/me`, {
    headers: {
      'Authorization': `Bearer ${token}`
    }
  });

  return response.json();
};