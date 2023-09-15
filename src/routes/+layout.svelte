<script>
	import Header from './Header.svelte';
	import Footer from './Footer.svelte';
	import './styles.css';
</script>

<div class="app">
	<Header />

	<main>
		<slot />
	</main>

	<Footer />
</div>

<style>
	.app {
		display: flex;
		justify-content: center;
		align-items: center;
		flex-direction: column;
		background-image: radial-gradient(
			circle at center center,
			rgb(15, 177, 180) 0%,
			rgb(15, 177, 180) 3%,
			rgb(255, 255, 255) 3%,
			rgb(255, 255, 255) 11%,
			rgb(255, 255, 255) 11%,
			rgb(255, 255, 255) 100%
		);
		background-size: 73px 73px;
	}

	.app main {
		width: 100%;
	}
</style>
