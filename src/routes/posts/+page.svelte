<script lang="ts">
	import { onMount } from 'svelte';
	import type { post } from '../../types';
	let data: post[] = [];

	onMount(async () => {
		const response = await fetch('http://localhost:8080/posts');
		data = await response.json();
	});
</script>

<div class="postsListContainer">
	<div class="postsListContent">
        <h1>Publicaciones disponibles</h1>
		{#if data}
			{#each data as post}
				<p>{post.post_content}</p>
			{/each}
		{/if}
	</div>
</div>

<style>
    .postsListContainer {
        width: 100%;
        height: fit-content;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .postsListContent {
        margin: 10px;
        height: fit-content;
    }

    .postsListContent h1:nth-child(1){
        height: 100px;
        display: flex;
        justify-content: center;
        align-items: center;
    }
</style>
