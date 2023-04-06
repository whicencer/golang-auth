export const getMe = async (token: string) => {
  const response = await fetch('http://127.0.0.1:2000/auth/me', {
    headers: {
      'Authorization': `Bearer ${token}`
    }
  });

  return response.json();
};