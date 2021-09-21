<script lang="ts">
	import { goto } from '$app/navigation';
	import axios from 'axios';
	import { onDestroy, onMount } from 'svelte';
	import { auth, AuthState } from '../../stores/auth';

	let authInfo: AuthState;
	let loading: boolean = true;
	const unsubscribe = auth.subscribe((data) => (authInfo = data));

	onMount(async () => {
		if (authInfo.authorized) {
			loading = false;
			return;
		}

		const token = window.localStorage.getItem('auth-token');

		try {
			const res = await axios.get('http://localhost:5000/user/authorize', {
				headers: {
					'auth-token': token,
				},
			});
			const data = res.data;

			if (data.authorized) {
				auth.set({
					authorized: true,
					token: token,
					user: {
						id: data.user.id,
						email: data.user.email,
						username: data.user.username,
					},
				});
			}
		} catch (error) {
			goto('/login');
		}
		loading = false;
	});

	onDestroy(unsubscribe);
</script>

<svelte:head>
	<title>Dashboard | GO-URL Shortener</title>
</svelte:head>

{#if !loading}
	<div class="flex justify-center w-screen h-screen">
		<div class="mt-32 space-y-2 text-center">
			<div class="text-4xl italic font-bold">
				<span>Welcome to </span>
				<span class="underline">GO-URL Shortener</span>
				<span>Dashboard</span>
			</div>
		</div>
	</div>
{:else}
	<div class="flex justify-center w-screen h-screen">
		<div class="mt-32 space-y-2 text-center">
			<div class="text-4xl italic font-bold">Loading...</div>
		</div>
	</div>
{/if}
