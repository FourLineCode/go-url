<script lang="ts">
	import axios from 'axios';

	let email: string = '';
	let username: string = '';
	let password: string = '';

	let status: string;

	async function register() {
		try {
			const res = await axios.post('http://localhost:5000/user/register', {
				email,
				username,
				password,
			});
			const data = res.data;

			if (data.success) {
				status = JSON.stringify(data.user, null, 2);
			}
		} catch (error) {
			status = JSON.stringify(error, null, 2);
		}
	}
</script>

<svelte:head>
	<title>Register | GO-URL Shortener</title>
</svelte:head>

<div class="flex justify-center w-screen h-screen">
	<div class="flex justify-center w-full p-2 mt-32 space-y-2">
		<div class="w-1/5 p-2 space-y-2 h-96">
			<div class="text-4xl font-bold text-center">Register</div>
			<form action="" class="space-y-4">
				<input
					type="text"
					placeholder="Email"
					bind:value={email}
					class="w-full p-2 border border-gray-500 focus:outline-none focus:border-green-500"
				/>
				<input
					type="text"
					placeholder="Username"
					bind:value={username}
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
						<div class="text-center text-red-500">
							<pre>{status}</pre>
						</div>
					{/if}
					<button
						type="submit"
						class="w-full p-2 font-semibold text-white bg-green-500 hover:bg-green-600"
						on:click|preventDefault={register}>Register</button
					>
				</div>
				<div class="text-center">
					<span>Have an account? </span>
					<span class="font-semibold text-green-500 hover:underline"
						><a href="/login">login</a></span
					>
				</div>
			</form>
		</div>
	</div>
</div>
