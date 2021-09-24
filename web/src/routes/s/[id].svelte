<script context="module" lang="ts">
	import type { LoadInput } from '@sveltejs/kit';
	import axios from 'axios';
	import { config } from '../../internal/config';

	export async function load({ page }: LoadInput) {
		try {
			const res = await axios.get(`${config.apiUrl}/site/url/${page.params.id}`);
			const data = res.data;
			return {
				redirect: data.url,
				status: 302,
			};
		} catch (error) {
			return {
				redirect: '/404',
				status: 302,
			};
		}
	}
</script>
