<script lang="ts">
	import { onMount } from "svelte";
	import type { test } from "../../types";
    let data:test = {message: ""};

    onMount(async () => {
        const text = await fetch("http://localhost:8080/")
        data = await text.json()
        console.log(data)
    })
</script>

<div class="postsListContainer">
    <h1>{data.message}</h1>
</div>
