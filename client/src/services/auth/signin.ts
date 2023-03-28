import type { IAuth } from './../../typings/IAuth';

export const signin = async (data: IAuth) => {
  const response = await fetch(`http://127.0.0.1:2000/auth/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
  });

  return response.json();
};