<header class="header">
    <div class="content">
        <h1>Relno.dev</h1>
        <div class="navbarContent">
            <section>
                <p>Lenguajes</p>
                <p>Tecnologias</p>
                <p>Sobre</p>
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

    .content h1:nth-child(1) {
        color: var(--sblue);
        font-size: 22px;
    }

    .navbarContent {
        width: 60%;
        display: flex;
        justify-content: space-between;
    }

    .navbarContent section {
        display: flex;
    }

    .navbarContent section p {
        margin: 10px;
        color: var(--dblue);
        font-weight: 600;
    }
</style>
