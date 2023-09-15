<script lang="ts">
	export let title: string = 'Relno.dev';
	export let summary: string = 'Tu página favorita para aprender a desarrollar.';
</script>

<svelte:head>
	<title>{title}</title>
	<meta
		name="description"
		content={summary}
	/>
	<meta name="robots" content="index, follow" />
</svelte:head>
