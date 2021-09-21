<script lang="ts">
	import { goto } from '$app/navigation';
	import axios from 'axios';
	import { onDestroy } from 'svelte';
	import { auth, AuthState } from '../../stores/auth';

	let authInfo: AuthState;
	const unsubscribe = auth.subscribe((data) => (authInfo = data));

	let email: string = '';
	let password: string = '';

	let status: string;

	async function login() {
		try {
			const res = await axios.post('http://localhost:5000/user/login', {
				email,
				password,
			});
			const data = res.data;

			if (data.success) {
				auth.set({
					authorized: true,
					token: data.token,
					user: {
						id: data.user.id,
						email: data.user.email,
						username: data.user.username,
					},
				});
				window.localStorage.setItem('auth-token', data.token);
				goto('/dashboard');
			}
		} catch (error) {
			status = error.response.data.error;
		}
	}

	onDestroy(unsubscribe);
</script>

<svelte:head>
	<title>Login | GO-URL Shortener</title>
</svelte:head>

<div class="flex justify-center w-screen h-screen">
	<div class="flex justify-center w-full p-2 mt-32 space-y-2">
		<div class="w-1/5 min-w-[300px] p-2 space-y-2 h-96">
			<div class="text-4xl font-bold text-center">Login</div>
			<form action="" class="space-y-4">
				<input
					type="text"
					placeholder="Email"
					bind:value={email}
					class="w-full p-2 border border-gray-500 focus:outline-none focus:border-green-500"
				/>
				<input
					type="password"
					placeholder="Password"
					bind:value={password}
					class="w-full p-2 border border-gray-500 focus:outline-none focus:border-green-500"
				/>
				<div>
					{#if status}
						<div class="text-xs text-center text-red-500">
							{status}
						</div>
					{/if}
					<button
						type="submit"
						class="w-full p-2 font-semibold text-white bg-green-500 hover:bg-green-600"
						on:click|preventDefault={login}>Login</button
					>
				</div>
				<div class="text-center">
					<span>Don't have an account? </span>
					<span class="font-semibold text-green-500 hover:underline"
						><a href="/register">register</a></span
					>
				</div>
			</form>
		</div>
	</div>
</div>
