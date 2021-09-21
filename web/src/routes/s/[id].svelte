<script context="module" lang="ts">
	import type { LoadInput } from '@sveltejs/kit';

	export async function load({ page }: LoadInput) {
		return {
			props: {
				id: page.params.id,
			},
		};
	}
</script>

<script lang="ts">
	export let id: string;
</script>

<svelte:head>
	<title>{id} | GO-URL Shortener</title>
</svelte:head>

<div class="flex justify-center w-screen h-screen">
	<div class="mt-32 space-y-2 text-center">
		<div class="text-4xl italic font-bold">Site - {id}</div>
	</div>
</div>
