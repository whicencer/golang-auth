import { writable } from 'svelte/store';

interface IAuthStore {
  isAuthorized: boolean;
  me: {
    username: string;
    description: string;
  }
}

export const isAuthorized = writable<IAuthStore>({
  isAuthorized: false,
  me: null
});