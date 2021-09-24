<script context="module" lang="ts">
	export interface Item {
		id: number;
		key: string;
		url: string;
		created_at: string;
	}
</script>

<script lang="ts">
	import { browser } from '$app/env';
	import { goto } from '$app/navigation';
	import axios from 'axios';
	import { onDestroy, onMount } from 'svelte';
	import UrlCard from '../../components/url-card.svelte';
	import { config } from '../../internal/config';
	import { auth, AuthState } from '../../stores/auth';

	let authInfo: AuthState;
	let loading: boolean = true;
	let errorMessage = '';
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
		try {
		} catch (error) {
			errorMessage = error.response.data.message;
		}
	});

	onDestroy(unsubscribe);

	let items: Item[] = [
		{
			id: 2,
			created_at: '2021-09-24T09:56:53.847153385+06:00',
			key: 'zvnP5Npx6dGXyLudB6rPEg',
			url: 'http://sjdnjsd.sdkns',
		},
		{
			id: 4,
			created_at: '2021-09-24T09:58:22.751357728+06:00',
			key: 'UnSC4oDyA7UR9aJzepWv7K',
			url: 'http://akmalakmal.com',
		},
		{
			id: 6,
			created_at: '2021-09-24T10:13:56.671464378+06:00',
			key: '8d4RzQosgZxT3Rdb9PZSxk',
			url: 'http://kekw.com',
		},
		{
			id: 2,
			created_at: '2021-09-24T09:56:53.847153385+06:00',
			key: 'zvnP5Npx6dGXyLudB6rPEg',
			url: 'http://sjdnjsd.sdkns',
		},
		{
			id: 4,
			created_at: '2021-09-24T09:58:22.751357728+06:00',
			key: 'UnSC4oDyA7UR9aJzepWv7K',
			url: 'http://akmalakmal.com',
		},
		{
			id: 6,
			created_at: '2021-09-24T10:13:56.671464378+06:00',
			key: '8d4RzQosgZxT3Rdb9PZSxk',
			url: 'http://kekw.com',
		},
		{
			id: 2,
			created_at: '2021-09-24T09:56:53.847153385+06:00',
			key: 'zvnP5Npx6dGXyLudB6rPEg',
			url: 'http://sjdnjsd.sdkns',
		},
		{
			id: 4,
			created_at: '2021-09-24T09:58:22.751357728+06:00',
			key: 'UnSC4oDyA7UR9aJzepWv7K',
			url: 'http://akmalakmal.com',
		},
		{
			id: 6,
			created_at: '2021-09-24T10:13:56.671464378+06:00',
			key: '8d4RzQosgZxT3Rdb9PZSxk',
			url: 'http://kekw.com',
		},
		{
			id: 2,
			created_at: '2021-09-24T09:56:53.847153385+06:00',
			key: 'zvnP5Npx6dGXyLudB6rPEg',
			url: 'http://sjdnjsd.sdkns',
		},
		{
			id: 4,
			created_at: '2021-09-24T09:58:22.751357728+06:00',
			key: 'UnSC4oDyA7UR9aJzepWv7K',
			url: 'http://akmalakmal.com',
		},
		{
			id: 6,
			created_at: '2021-09-24T10:13:56.671464378+06:00',
			key: '8d4RzQosgZxT3Rdb9PZSxk',
			url: 'http://kekw.com',
		},
	];
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
			<div class="flex justify-between items-center">
				<form class="space-x-1">
					<input
						placeholder="Search"
						type="search"
						class="py-1 px-2 w-64 focus:outline-none border border-gray-500 focus:border-green-500"
					/>
					<button
						type="submit"
						class="bg-gray-700 hover:bg-gray-800 font-semibold text-white py-1 px-2"
						>Search</button
					>
				</form>
				<div>
					<select
						name="sort"
						class="py-1 px-2 w-32 focus:outline-none border border-gray-500 focus:border-green-500"
					>
						<option value="date">Sort by date</option>
						<option value="a-z">Sort by a-z</option>
						<option value="z-a">Sort by z-a</option>
					</select>
					<button
						class="bg-green-500 hover:bg-green-600 font-semibold text-white py-1 px-2"
					>
						<span class="font-bold">+</span>
						<span>Create new url</span>
					</button>
				</div>
			</div>
			<div class="space-y-2">
				{#if errorMessage}
					<div class="text-center p-4 mt-12 font-semibold text-2xl text-red-500">
						{errorMessage}
					</div>
				{:else}
					{#each items as item}
						<UrlCard site={item} />
					{/each}
				{/if}
			</div>
		</div>
	</div>
{/if}
