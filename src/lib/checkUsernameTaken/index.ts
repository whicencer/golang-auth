import { baseURL } from "../../constants/api";

export const checkUserNameTaken = async (userName: string): Promise<boolean> => {
  try {
    const response = await fetch(`${baseURL}/users/${userName}`);
    const data = await response.json();
    return data.userExists;
  } catch (error) {
    return false;
  }
};