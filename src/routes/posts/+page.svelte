<script lang="ts">
	import { onMount } from 'svelte';
	import type { post } from '../../types';
	let data: post[] = [];

	onMount(async () => {
		const response = await fetch('http://localhost:8080/posts');
		data = await response.json();
	});
</script>

<div class="postsListContainer" style="height: {data.length <= 3 ? "100vh" : "fit-content"}">
	<div class="postsListContent">
        <h1 id="postsListHeader">Publicaciones disponibles</h1>
		{#if data}
			{#each data as post}
				<div class="post">
                    <h1 id="postTitle">{post.post_title}</h1>
                    <p id="postDate">{new Date(post.post_date).toDateString()}</p>
                    <section class="tags">
                        {#each post.post_tags as tag}
                            <p id="postTag">{tag.name}</p>
                        {/each}
                    </section>
                </div>
			{/each}
		{/if}
	</div>
</div>

<style>
    .postsListContainer {
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: flex-start;
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

    .postsListContent {
        margin: 10px;
        height: fit-content;
    }

    #postsListHeader {
        height: 100px;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .post {
        display: flex;
        flex-direction: column;
        padding: 20px;
        background: white;
        width: 500px;
        box-shadow: 0px 0px 10px var(--ddblue);
        border-radius: 10px;
        text-overflow: ellipsis;
        white-space: nowrap;
        overflow: hidden;
        cursor: pointer;
        transition: .2s;
    }

    .post:hover {
        box-shadow: 0px 0px 20px var(--cyan);
    }

    .post h1 {
        font-family: "Raleway";
        font-size: 22px;
        width: 500px;
        text-overflow: ellipsis;
        white-space: nowrap;
        overflow: hidden;
    }

    #postDate {
        margin-top: 5px;
        font-family: "Raleway";
        font-weight: 800;
        color: var(--cyan);
    }

    .tags {
        display: flex;
        justify-content: flex-end;
        margin-top: 20px;
    }

    #postTag {
        text-align: center;
        border-radius: 5px;
        margin: 0px 5px;
        padding: 8px 10px;
        font-family: "Raleway";
        font-weight: 800;
        /* background: var(--lblue); */
        border: 2px solid var(--dddblue);
        transition: .2s;
    }

    .post:hover #postTag {
        color: var(--green);
    }
</style>
