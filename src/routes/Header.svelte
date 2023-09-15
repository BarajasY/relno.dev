<header class="header">
	<div class="content">
		<a href="/" class="logo">Relno.dev</a>
		<div class="navbarContent">
			<section>
				<a href="/posts">Publicaciones</a>
<!-- 				<a href="/about">Sobre</a> -->
			</section>
			<section>
				<p>Esp</p>
			</section>
		</div>
	</div>
</header>

<style>
	.header {
		width: 100%;
		height: 70px;
		display: flex;
		justify-content: center;
		align-items: center;
		background-image: repeating-linear-gradient(
				135deg,
				rgba(189, 189, 189, 0.1) 0px,
				rgba(189, 189, 189, 0.1) 2px,
				transparent 2px,
				transparent 4px
			),
			linear-gradient(90deg, rgb(255, 255, 255), rgb(255, 255, 255));
	}

	.content {
		display: flex;
		width: 60%;
		align-items: center;
		justify-content: space-between;
	}

	.logo {
		color: var(--darkblue);
		font-size: 22px;
		font-weight: 700;
		font-family: 'Poppins';
	}

	.navbarContent {
		width: 60%;
		display: flex;
		justify-content: space-between;
	}

	.navbarContent section {
		display: flex;
		justify-content: center;
		align-items: flex-end;
	}

	.navbarContent section p,
	.navbarContent section a {
		margin: 10px;
		color: var(--dblue);
		font-weight: 600;
		transition: .5s;
	}

	.navbarContent section a:before {
		content: "";
		display: block;
		width: 0px;
		height: 2px;
		background: var(--cyan);
		transform: translateY(25px);
		transition: ease .5s;
	}

	.navbarContent section a:hover:before {
		width: 100%;
	}
</style>
