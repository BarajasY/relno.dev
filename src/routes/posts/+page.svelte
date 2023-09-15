<script lang="ts">
	import { onMount } from 'svelte';
	import type { post, tag } from '../../types';
	let posts: post[] = [];
	let tags: tag[] = [];

	onMount(async () => {
		const postresponse = await fetch('http://localhost:8080/posts');
		const tagresponse = await fetch('http://localhost:8080/tags');
		posts = await postresponse.json();
		tags = await tagresponse.json();
	});
</script>

<div class="postsListContainer">
	<h1 id="postsListHeader">Lee sobre lo que te apasiona</h1>
	<div class="postsContentWrapper">
		<div class="postsListContent">
			{#if posts}
				<div class="postsContainer">
					{#each posts as post}
						<a href={ "/posts/"+ post.post_title.toLowerCase().split(" ").join("-")} class="postAnchor">
							<div class="post">
								<h1 id="postTitle">{post.post_title}</h1>
								<p id="postDate">{new Date(post.post_date).toDateString()}</p>
								<section class="tags">
									{#each post.post_tags as tag}
										<p id="postTag">{tag.name}</p>
									{/each}
								</section>
							</div>
						</a>
					{/each}
				</div>
			{/if}
		</div>
		<div class="tagsListContent">
			{#if tags}
				<div class="tagsContainer">
					<h1>Tags importantes</h1>
					{#each tags as tag}
						<li>{tag.name}</li>
					{/each}
				</div>
			{/if}
		</div>
	</div>
</div>

<style>
	.postsListContainer {
		width: 100%;
		min-height: 100vh;
		max-height: fit-content;
		display: flex;
		flex-direction: column;
		justify-content: flex-start;
		align-items: center;
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

	.postsContentWrapper {
		display: flex;
		width: 80%;
		justify-content: center;
		align-items: flex-start;
	}

	#postsListHeader {
		height: 100px;
		display: flex;
		justify-content: center;
		align-items: center;
		font-family: 'Raleway';
	}

	.postsContainer {
		padding: 30px;
		border-radius: 5px;
		background: white;
		box-shadow: 0px 0px 15px var(--dddblue);
	}

	.post {
		display: flex;
		flex-direction: column;
		padding: 20px;
		margin: 10px 0px;
		width: 600px;
		border: 2px solid transparent;
		border-radius: 10px;
		text-overflow: ellipsis !important;
		white-space: nowrap !important;
		overflow: hidden !important;
		cursor: pointer;
		transition: 0.2s;
	}

	.post:hover {
		border: 2px solid black;
	}

	.post h1 {
		font-family: 'Raleway';
		font-size: 22px;
		width: 600px;
		text-overflow: ellipsis !important;
		white-space: nowrap !important;
		overflow: hidden !important;
	}

	#postDate {
		margin-top: 5px;
		font-family: 'Raleway';
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
		font-family: 'Raleway';
		font-weight: 800;
		border: 2px solid var(--dddblue);
		transition: 0.2s;
	}

	.post:hover #postTag {
		color: var(--green);
	}

	.tagsListContent {
		margin-top: 10px;
		padding: 20px;
		background: var(--ddddblue);
		border-radius: 5px;
		margin-left: 30px;
	}

	.tagsContainer {
		width: 200px;
	}

	.tagsContainer h1 {
		font-size: 18px;
		text-align: center;
		font-family: 'Raleway';
		color: var(--white);
		font-weight: 900;
	}

	.tagsContainer li {
		font-family: 'Raleway';
		font-weight: 900;
		font-size: 18px;
		margin: 8px 0px;
		color: var(--bblue);
		transition: 0.2s;
	}

	.tagsContainer li:hover {
		color: var(--bgreen);
		cursor: pointer;
	}
</style>
