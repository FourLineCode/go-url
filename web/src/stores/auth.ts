import { writable } from 'svelte/store';

export interface User {
	id: number;
	email: string;
	username: string;
}

export interface AuthState {
	authorized: boolean;
	user?: User;
	token?: string;
}

export const auth = writable<AuthState>({
	authorized: false,
	user: null,
	token: null,
});
