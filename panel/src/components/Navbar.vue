<template>
    <nav id="nav" class="navbar navbar-expand-md navbar-dark bg-dark mb-4">
        <div class="container-fluid">
            <a class="navbar-brand" href="#">ClienTell</a>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse d-flex justify-content-between" id="navbarCollapse">
                <ul class="navbar-nav me-auto mb-2 mb-md-0">
                    <li class="nav-item">
                        <router-link class="nav-link" to="/dashboard">Dashboard</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link class="nav-link" to="/moderate">Moderate</router-link>
                    </li>
                </ul>
                <button class="btn btn-md btn-danger" type="button" @click="doLogout">Logout</button>
            </div>
        </div>
    </nav>
</template>

<script>
export default {
    name: 'Navbar',
    methods: {
        doLogout() {
            localStorage.removeItem("token");
            this.$toast.warning("You have been logged out.", { position: "top" });
            setTimeout(() => { 
                location.reload();
            }, 1500);
        }
    }
}
</script>

<style scoped>
#nav a.router-link-exact-active {
    color: rgb(0, 148, 255);
    font-weight: bold;
}
</style>
