<template>
    <Navbar />
    <main>
        <div class="px-4 py-5 my-5 text-center">
            <h1 class="display-5 fw-bold">ClienTell Staff</h1>
            <div class="col-lg-6 mx-auto">
                <p class="lead mb-4">Panel for staff backend operations.<br/>This will include analytics and post moderation.</p>
                <div class="d-grid gap-2 d-sm-flex justify-content-sm-center">
                </div>
            </div>
        </div>
    </main>
</template>

<script>
import Navbar from "../components/Navbar.vue";

export default {
    name: 'Dashboard',
    components: {
        Navbar
    }
}
</script>

<style scoped>
</style>
