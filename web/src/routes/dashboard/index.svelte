<script context="module" lang="ts">
	export interface Item {
		id: number;
		key: string;
		url: string;
		created_at: string;
	}

	const composer = writable(false);
</script>

<script lang="ts">
	import { browser } from '$app/env';
	import { goto } from '$app/navigation';
	import axios from 'axios';
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';
	import UrlCard from '../../components/url-card.svelte';
	import { config } from '../../internal/config';
	import { auth, AuthState } from '../../stores/auth';

	let authInfo: AuthState;
	let loading: boolean = true;
	let items: Item[] = [];
	let sortby: 'date-new' | 'date-old' | 'a-z' | 'z-a' = 'date-new';
	let sortedItems: Item[] = [];
	let searchTerm: string = '';
	let filteredItems: Item[] = [];
	let errorMessage = '';
	let createUrlString: string;
	let createErrorMessage = '';
	const unsubscribe = auth.subscribe((data) => (authInfo = data));

	onMount(async () => {
		if (authInfo.authorized) {
			loading = false;
			return;
		}

		const token = window.localStorage.getItem('auth-token');

		try {
			const res = await axios.get(`${config.apiUrl}/user/authorize`, {
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

	$: if (browser && !authInfo.authorized && !loading) {
		goto('/login');
	}

	onMount(async () => {
		items = await getSites();
	});

	async function getSites() {
		try {
			const token = window.localStorage.getItem('auth-token');
			const res = await axios.get(`${config.apiUrl}/site/usersites`, {
				headers: {
					'auth-token': token,
				},
			});
			return res.data;
		} catch (error) {
			errorMessage = error.response.data.error;
		}
	}

	$: sortedItems = items.sort((a, b) => {
		const index = 'http://'.length;
		if (sortby === 'date-new') {
			return Number(new Date(b.created_at)) - Number(new Date(a.created_at));
		} else if (sortby === 'date-old') {
			return Number(new Date(a.created_at)) - Number(new Date(b.created_at));
		} else if (sortby === 'a-z') {
			return a.url.charCodeAt(index) - b.url.charCodeAt(index);
		} else if (sortby === 'z-a') {
			return b.url.charCodeAt(index) - a.url.charCodeAt(index);
		}
	});

	$: filteredItems = sortedItems.filter((item) =>
		item.url.toLowerCase().includes(searchTerm.toLowerCase().trim())
	);

	function filterItems() {
		filteredItems = sortedItems.filter((item) =>
			item.url.toLowerCase().includes(searchTerm.toLowerCase())
		);
	}

	async function createUrl() {
		if (!createUrlString.trim()) return;
		try {
			const token = window.localStorage.getItem('auth-token');
			const urlString =
				createUrlString.startsWith('http://') || createUrlString.startsWith('https://')
					? createUrlString
					: 'http://' + createUrlString;
			const res = await axios.post(
				`${config.apiUrl}/site/create`,
				{
					url: urlString,
				},
				{
					headers: {
						'auth-token': token,
					},
				}
			);
			const data: { success: boolean; site: Item } = res.data;
			if (data.success) {
				const site = data.site;
				const newSite = {
					id: site.id,
					created_at: site.created_at,
					key: site.key,
					url: site.url,
				};
				items = [...items, newSite];
				createUrlString = '';
			} else {
				createErrorMessage = 'Something went wrong';
			}
			composer.set(false);
			// items.push(res.data.site);
		} catch (error) {
			createErrorMessage = error.response.data.error;
		}
	}
</script>

<svelte:head>
	<title>Dashboard | GO-URL Shortener</title>
</svelte:head>

{#if loading}
	<div class="flex justify-center w-screen h-screen">
		<div class="mt-32 space-y-2 text-center">
			<div class="text-4xl italic font-bold">Loading...</div>
		</div>
	</div>
{:else}
	<div class="flex justify-center w-full h-full">
		<div class="w-full lg:w-[800px] mt-4 space-y-4">
			<div class="flex items-center justify-between">
				<form on:submit|preventDefault={filterItems} class="space-x-1">
					<input
						placeholder="Search"
						type="search"
						bind:value={searchTerm}
						class="w-64 px-2 py-1 border border-gray-500 focus:outline-none focus:border-green-500"
					/>
					<button
						type="submit"
						class="px-2 py-1 font-semibold text-white bg-gray-700 hover:bg-gray-800"
						>Search</button
					>
				</form>
				<div>
					<select
						name="sort"
						bind:value={sortby}
						class="w-48 px-2 py-1 border border-gray-500 focus:outline-none focus:border-green-500"
					>
						<option value="date-new">{'Sort by date (newest)'}</option>
						<option value="date-old">{'Sort by date (oldest)'}</option>
						<option value="a-z">Sort by a-z</option>
						<option value="z-a">Sort by z-a</option>
					</select>
					<button
						class="px-2 py-1 font-semibold text-white bg-green-500 hover:bg-green-600"
						on:click={() => composer.update((prev) => !prev)}
					>
						<span class="font-bold">+</span>
						<span>Create new url</span>
					</button>
				</div>
			</div>
			{#if $composer}
				<div class="flex items-center justify-end space-x-2">
					<span class="text-xs text-red-500">{createErrorMessage}</span>
					<form on:submit|preventDefault={createUrl} class="space-x-1">
						<input
							type="text"
							placeholder="Example - http://example.com"
							bind:value={createUrlString}
							on:input={() => (createErrorMessage = '')}
							class="px-2 py-1 border border-gray-500 w-80 focus:outline-none focus:border-green-500"
						/>
						<button
							type="submit"
							class="px-2 py-1 font-semibold text-white bg-green-500 hover:bg-green-600"
							>Submit</button
						>
					</form>
				</div>
			{/if}
			<div class="space-y-2">
				{#if errorMessage}
					<div class="p-4 mt-12 text-2xl font-semibold text-center text-red-500">
						{errorMessage}
					</div>
				{:else}
					{#each filteredItems as item}
						<UrlCard site={item} />
					{/each}
				{/if}
			</div>
		</div>
	</div>
{/if}
