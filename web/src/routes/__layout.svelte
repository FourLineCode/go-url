<script>
	import '../app.postcss';
</script>

<body>
	<slot />
</body>

<style global type="postcss">
	body {
		@apply antialiased bg-gray-900 text-white;
	}
</style>
