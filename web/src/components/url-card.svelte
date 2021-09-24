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
				<div class="flex items-center space-x-1">
					<span class="font-semibold"
						><svg
							xmlns="http://www.w3.org/2000/svg"
							class="h-4 w-4"
							fill="none"
							viewBox="0 0 24 24"
							stroke="currentColor"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"
							/>
						</svg></span
					>
					<a
						href={`${config.hostUrl}/s/${site.key}`}
						target="_blank"
						class="hover:underline truncate hover:text-green-500"
						>{`${config.hostUrl}/s/${site.key}`}</a
					>
				</div>
				<div class="text-xs flex items-center space-x-2">
					<span class="font-semibold"
						><svg
							xmlns="http://www.w3.org/2000/svg"
							class="h-3 w-3"
							viewBox="0 0 20 20"
							fill="currentColor"
						>
							<path
								fill-rule="evenodd"
								d="M6.672 1.911a1 1 0 10-1.932.518l.259.966a1 1 0 001.932-.518l-.26-.966zM2.429 4.74a1 1 0 10-.517 1.932l.966.259a1 1 0 00.517-1.932l-.966-.26zm8.814-.569a1 1 0 00-1.415-1.414l-.707.707a1 1 0 101.415 1.415l.707-.708zm-7.071 7.072l.707-.707A1 1 0 003.465 9.12l-.708.707a1 1 0 001.415 1.415zm3.2-5.171a1 1 0 00-1.3 1.3l4 10a1 1 0 001.823.075l1.38-2.759 3.018 3.02a1 1 0 001.414-1.415l-3.019-3.02 2.76-1.379a1 1 0 00-.076-1.822l-10-4z"
								clip-rule="evenodd"
							/>
						</svg></span
					>
					<a
						href={site.url}
						target="_blank"
						class="hover:underline truncate hover:text-green-500">{site.url}</a
					>
				</div>
			</div>
			<div class="text-xs text-gray-500">
				<span>Created on: </span>
				<span>{format(new Date(site.created_at), 'dd MMM, yyyy')}</span>
			</div>
		</div>
		<div class="space-x-2 flex items-center">
			<button
				class="bg-green-500 flex items-center space-x-1 hover:bg-green-600 font-semibold text-white px-2 py-1.5"
				><span
					><svg
						xmlns="http://www.w3.org/2000/svg"
						class="h-5 w-5"
						fill="none"
						viewBox="0 0 24 24"
						stroke="currentColor"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
						/>
					</svg></span
				><span>Open</span></button
			>
			<button
				class="bg-indigo-500 flex items-center space-x-1 hover:bg-indigo-600 font-semibold text-white px-2 py-1.5"
			>
				<span
					><svg
						xmlns="http://www.w3.org/2000/svg"
						class="h-5 w-5"
						fill="none"
						viewBox="0 0 24 24"
						stroke="currentColor"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
						/>
					</svg></span
				>
				<span>Update</span></button
			>
			<button
				class="bg-red-500 flex items-center space-x-1 hover:bg-red-600 font-semibold text-white px-2 py-1.5"
				><span
					><svg
						xmlns="http://www.w3.org/2000/svg"
						class="h-5 w-5"
						fill="none"
						viewBox="0 0 24 24"
						stroke="currentColor"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
						/>
					</svg></span
				><span>Delete</span></button
			>
		</div>
	</div>
</div>
