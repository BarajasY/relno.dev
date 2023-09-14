<script>
    import Header from "./Header.svelte";
    import Footer from "./Footer.svelte";
    import "./styles.css";
</script>

<div class="app">
    <Header />

    <main>
        <slot/>
    </main>

    <Footer/>
</div>

<style>
    .app {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
    }

    .app main {
        width: 100%;
    }
</style>
