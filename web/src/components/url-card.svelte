<script lang="ts">
	import { format } from 'date-fns';
	import type { Item } from 'src/routes/dashboard/index.svelte';
	import { config } from '../internal/config';

	export let site: Item;
</script>

<div class="bg-gray-300 w-full py-2 px-4">
	<div class="flex justify-between items-center">
		<div class="space-y-2">
			<div>
				<div>
					<span class="font-semibold">Url: </span>
					<a
						href={`${config.hostUrl}/s/${site.key}`}
						target="_blank"
						class="hover:underline hover:text-green-500"
						>{`${config.hostUrl}/s/${site.key}`}</a
					>
				</div>
				<div class="text-xs">
					<span class="font-semibold">Site: </span>
					<a href={site.url} target="_blank" class="hover:underline hover:text-green-500"
						>{site.url}</a
					>
				</div>
			</div>
			<div class="text-xs text-gray-500">
				<span>Created on: </span>
				<span>{format(new Date(site.created_at), 'dd MMM, yyyy')}</span>
			</div>
		</div>
		<div class="space-x-2">
			<button
				class="bg-indigo-500 hover:bg-indigo-600 font-semibold text-white px-2 py-1.5 text-lg"
				>Update</button
			>
			<button class="bg-red-500 hover:bg-red-600 font-semibold text-white px-2 py-1.5 text-lg"
				>Delete</button
			>
		</div>
	</div>
</div>
