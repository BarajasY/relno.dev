<header class="header">
    <div class="content">
        <a href="/" class="logo">Relno.dev</a>
        <div class="navbarContent">
            <section>
                <a href="/posts">Publicaciones</a>
                <a href="/about">Sobre</a>
            </section>
            <section>
                <p>Esp</p>
            </section>
        </div>
    </div>
</header>

<style>
    .header{
        width: 100%;
        height: 70px;
        display: flex;
        justify-content: center;
        align-items: center;
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
        font-family: "Poppins";
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

    .navbarContent section p, .navbarContent section a {
        margin: 10px;
        color: var(--dblue);
        font-weight: 600;
    }
</style>
