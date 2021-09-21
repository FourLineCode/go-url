<script lang="ts">
	import '../app.postcss';
</script>

<body>
	<slot />
</body>

<style global type="postcss">
	body {
		@apply antialiased bg-gray-200 text-gray-900 font-sans;
	}
</style>
