<script lang="ts">
	import { onDestroy } from 'svelte';
	import '../app.postcss';
	import { auth, AuthState } from '../stores/auth';

	let authInfo: AuthState;
	const unsubscribe = auth.subscribe((data) => (authInfo = data));

	function logout() {
		window.localStorage.removeItem('auth-token');
		auth.set({
			authorized: false,
			user: null,
			token: null,
		});
	}

	onDestroy(unsubscribe);
</script>

<body>
	<div class="w-full mx-4 lg:w-[1000px] flex justify-between items-center h-16 lg:mx-auto">
		<div class="flex text-4xl italic font-bold">
			<span>GO-</span>
			<span class="underline">URL</span>
		</div>
		<div class="flex items-center space-x-4">
			{#if authInfo.authorized}
				<div class="flex space-x-1 text-xl font-semibold hover:underline">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						class="w-6 h-6"
						viewBox="0 0 20 20"
						fill="currentColor"
					>
						<path
							fill-rule="evenodd"
							d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z"
							clip-rule="evenodd"
						/>
					</svg>
					<span>{authInfo.user.username}</span>
				</div>
				<div
					on:click={logout}
					class="text-sm cursor-pointer hover:text-red-500 hover:underline"
				>
					logout
				</div>
			{:else}
				<div class="font-semibold hover:text-green-500 hover:underline">
					<a href="/login">Login</a>
				</div>
				<div class="font-semibold hover:text-green-500 hover:underline">
					<a href="/register">Register</a>
				</div>
			{/if}
		</div>
	</div>
	<slot />
</body>

<style global type="postcss">
	body {
		@apply antialiased bg-gray-200 text-gray-900 font-sans;
	}
</style>
