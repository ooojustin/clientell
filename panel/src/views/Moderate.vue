<template>
    <Navbar />
    <main>
        <div class="container">
            <table class="table">
                <thead>
                    <tr>
                        <th scope="col">#</th>
                        <th scope="col">First</th>
                        <th scope="col">Last</th>
                        <th scope="col">Handle</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <th scope="row">1</th>
                        <td>Justin</td>
                        <td>Garfolo</td>
                        <td>@jstn</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </main>
</template>

<script>
import Navbar from "../components/Navbar.vue";

export default {
    name: 'Moderate',
    components: {
        Navbar
    }
}
</script>

<style scoped>
</style>
