<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import type { post } from '../../../types';
	import MetaTags from '../../MetaTags.svelte';
	let post: post;

	onMount(async () => {
		let formattedParam = $page.params.post.split('-').join(' ');
		let startContent = document.getElementById("startContent");
		let domParser = new DOMParser();
		const postResponse = await fetch(`http://localhost:8080/posts/${formattedParam}`);
		post = await postResponse.json();
		//converts html string to html elements.
		const dom = domParser.parseFromString(post.post_content, "text/html")
		//appends html elements to element with id startContent
		startContent?.appendChild(dom.body)

	});
</script>

{#if post}
	<MetaTags title={post.post_title} summary={post.post_summary} />
{/if}
<div class="postContainer">
	<div class="postContent">
		{#if post}
			<h1>{post.post_title}</h1>
		{/if}
		<div class="startContent" id="startContent"></div>
	</div>
</div>

<style>
	.postContainer {
		width: 100%;
		min-height: 100vh;
		max-height: fit-content;
		display: flex;
		justify-content: center;
		align-items: flex-start;
	}

	.postContent {
		margin-top: 30px;
	}
</style>
