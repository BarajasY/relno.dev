<script>
	import Technologies from './Technologies.svelte';
</script>

<svelte:head>
	<title>Relno.dev</title>
	<meta name="description" content="Relno.dev" />
</svelte:head>

<div class="mainContainer">
	<h1>Tu espacio para aprender nuevas tecnologías y lenguajes</h1>
	<div class="mainContent">
		<section class="mainContentSection1">
			<p>
				En <span>Relno.dev</span> nos encargamos de compartir publicaciones en Español que resulten útiles para
				mejorar tus habilidades dentro del desarrollo <span>web</span>, <span>móvil</span> y de <span>escritorio</span>.
			</p>
		</section>
		<section class="mainContentSection2">
			<a href="/posts">
				<button>Ver Posts</button>
			</a>
		</section>
	</div>
	<Technologies />
</div>

<style>
	.mainContainer {
		width: 100%;
		height: fit-content;
		display: flex;
		justify-content: space-evenly;
		align-items: center;
		flex-direction: column;
		background-image: radial-gradient(
			rgb(15, 177, 180) 0%,
			rgb(15, 177, 180) 3%,
			rgb(255, 255, 255) 3%,
			rgb(255, 255, 255) 11%,
			rgb(255, 255, 255) 11%,
			rgb(255, 255, 255) 100%
		);
		background-size: 73px 73px;
	}

	.mainContent {
		display: flex;
		justify-content: space-between;
		align-items: center;
		width: 70%;
	}

	.mainContainer h1:nth-child(1) {
		font-family: 'Raleway';
		font-weight: 800;
		font-size: 40px;
		padding: 150px;
		text-align: center;
	}

	.mainContentSection1 {
		display: flex;
		justify-content: flex-end;
		text-align: justify;
	}

	.mainContentSection1 p {
		font-weight: 600;
		font-size: 20px;
		height: 100px;
	}

	.mainContentSection1 p span {
		color: var(--cyan);
	}

	section {
		width: 50%;
		height: 400px;
		margin-top: 100px;
	}

	.mainContentSection2 {
		display: flex;
		justify-content: center;
		align-items: flex-start;
	}

	.mainContentSection2 button {
		padding: 15px 20px;
		font-family: 'Raleway';
		font-size: 25px;
		background: white;
		color: var(--bblue);
		box-shadow: 0px 0px 10px black;
		font-weight: 700;
		border-radius: 5px;
		transition: 0.2s;
	}

	.mainContentSection2 button:hover {
		background: var(--black);
		color: var(--bblue);
		cursor: pointer;
		box-shadow: 0px 0px 20px var(--ddblue);
	}
</style>
