<template>
    <div class="signin-wrapper">
        <main class="form-signin">
            <form @submit.prevent="onSubmit">

                <h1 class="h3 mb-3 fw-normal">Please sign in</h1>

                <div class="form-floating">
                    <input type="email" class="form-control" id="email" v-model="email">
                    <label for="email">Email address</label>
                </div>
                <div class="form-floating">
                    <input type="password" class="form-control" id="password" v-model="password">
                    <label for="password">Password</label>
                </div>

                <button class="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
                <p class="mt-5 mb-3 text-muted">Copyright &copy; 2021 STING DGI LLC</p>

            </form>
        </main>
    </div>
</template>

<script>
import axios from "axios";
import vars from "../variables.js";
import { mapGetters } from "vuex";

export default {
    name: 'Login',
    data() {
        return {
            email: "",
            password: ""
        };
    },
    computed: {
        ...mapGetters(["isAuthenticated"])
    },
    watch: {
        isAuthenticated(val) {
            // automatically redirect user when logged in
            if (val)
                this.$router.push("/dashboard");
        }
    },
    mounted() {
        // try to login from token in localStorage
        this.$store.dispatch("restoreLogin");
    },
    methods: {
        async onSubmit() {

            const { email, password } = this;
            const response = await axios.post(`${vars.backend}/login`, { email, password });

            const { data, status } = response;
            if (status == 200) {
                const user = data.data;
                this.$store.dispatch("loggedIn", user);
            }

        }
    }
}
</script>

<style scoped>
div.signin-wrapper {
    display: flex;
    align-items: center;
    padding-top: 40px;
    padding-bottom: 40px;
}

.form-signin {
    width: 100%;
    max-width: 330px;
    padding: 15px;
    margin: auto;
}

.form-signin .form-floating:focus-within {
    z-index: 2;
}

.form-signin input[type="email"] {
    margin-bottom: -1px;
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0;
}

.form-signin input[type="password"] {
    margin-bottom: 10px;
    border-top-left-radius: 0;
    border-top-right-radius: 0;
}
</style>
