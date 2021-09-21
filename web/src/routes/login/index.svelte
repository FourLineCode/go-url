<script lang="ts">
	let error = '';

	function setError(msg: string) {
		error = msg;
	}
</script>

<svelte:head>
	<title>Login | GO-URL Shortener</title>
</svelte:head>

<div class="flex justify-center w-screen h-screen">
	<div class="flex justify-center w-full p-2 mt-32 space-y-2">
		<div class="w-1/5 p-2 space-y-2 h-96">
			<div class="text-4xl font-bold text-center">Login</div>
			<form action="" class="space-y-4">
				<input
					type="text"
					placeholder="Email"
					class="w-full p-2 border border-gray-500 focus:outline-none focus:border-green-500"
				/>
				<input
					type="password"
					placeholder="Password"
					class="w-full p-2 border border-gray-500 focus:outline-none focus:border-green-500"
				/>
				<div>
					{#if error}
						<div class="text-center text-red-500">{error}</div>
					{/if}
					<button
						type="submit"
						class="w-full p-2 font-semibold text-white bg-green-500 hover:bg-green-600"
						on:click|preventDefault={() => setError('Invalid credentials')}
						>Login</button
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
