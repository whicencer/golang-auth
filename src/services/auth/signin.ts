import { baseURL } from '../../constants/api';
import type { IAuth } from './../../typings/IAuth';

export const signin = async (data: IAuth) => {
  const response = await fetch(`${baseURL}/auth/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
  });

  return response.json();
};