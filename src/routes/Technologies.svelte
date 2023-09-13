<script lang="ts">
    const technologies:string[] = ["React", "Rust", "TypeScript", "Dart", "Java", "Flutter", "JavaScript", "Svelte", "Solid.js", "Three.js", "SpringBoot", "SQL"]
</script>

<div class="mainTechnologies">
    <h1>Publicaremos acerca de distintas tecnologías y siempre trataremos de abrirnos a más:</h1>
    <div class="technologiesContainer">
        {#each technologies as tech}
            <h1>{tech}</h1>
        {/each}
    </div>
</div>

<style>
    .mainTechnologies {
        width: 80%;
        display: flex;
        height: fit-content;
        justify-content: center;
        align-items: center;
        flex-direction: column;
		border-radius: 15px;
        margin-bottom: 100px;
		box-shadow: 0px 0px 20px var(--dblue);
		background-image: radial-gradient(
				circle at center center,
				rgb(41, 125, 58) 0%,
				rgb(41, 125, 58) 11%,
				transparent 11%,
				transparent 100%
			),
			linear-gradient(90deg, hsl(37, 14%, 9%), hsl(37, 14%, 9%));
		background-size: 11px 11px;
	}

	.mainTechnologies h1:nth-child(1) {
		font-size: 27px;
		color: var(--white);
        padding: 50px;
	}

    .technologiesContainer {
        width: 100%;
        display: grid;
        grid-template-columns: auto auto auto auto;
        justify-content: center;
        align-items: center;
        gap: 20px;
        column-gap: 50px;
        margin-bottom: 50px;
    }

    .technologiesContainer h1 {
        color: white;
        padding: 10px !important;
        font-size: 18px !important;
        border: 2px solid white;
        height: 30px !important;
        width: 100px !important;
        align-content: center;
        text-align: center;
        border-radius: 6px;
        background: hsl(37, 14%, 9%);
        transition: .2s;
    }

    .technologiesContainer h1:hover {
        color: hsl(37, 14%, 9%);
        background: white;
        box-shadow: 0px 0px 30px var(--lblue);
        cursor: pointer;
    }
</style>
