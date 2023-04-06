export const checkUserNameTaken = async (userName: string): Promise<boolean> => {
  try {
    const response = await fetch(`http://127.0.0.1:2000/users/${userName}`);
    const data = await response.json();
    return data.userExists;
  } catch (error) {
    return false;
  }
};