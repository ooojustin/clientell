<template>
    <Navbar />
    <main>
        <div class="container">
            <table class="table">
                <thead>
                    <tr>
                        <th scope="col">ID</th>
                        <th scope="col">Stars</th>
                        <th scope="col">Comment</th>
                        <th scope="col">Actions</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(rating, idx) in ratings" :key="idx">
                        <th scope="row">{{ rating.ID }}</th>
                        <td>{{ rating.stars }}</td>
                        <td>{{ rating.comment }}</td>
                        <td>
                            <div class="btn btn-md btn-success" role="button" @click="onAction(rating, true)">Approve</div>&nbsp;
                            <div class="btn btn-md btn-danger" role="button" @click="onAction(rating, false)">Deny</div>
                        </td>
                    </tr>
                </tbody>
            </table>
            <div class="mt-3">
                As you approve and deny ratings, new ones will appear.
                <br />
                <b>Showing {{ ratings.length }} of {{ count }} results.</b>
            </div>
        </div>
    </main>
</template>

<script>
import Navbar from "../components/Navbar.vue";
import { api } from "../variables.js";

export default {
    name: 'Moderate',
    components: {
        Navbar
    },
    data() {
        return {
            count: 0,
            ratings: null
        };
    },
    created() {
        this.loadRatings();
    },
    methods: {
        async loadRatings() {
            const response = await api.get("/listReviewRatings");
            const { data, status } = response;
            if (status == 200) {
                this.count = data.count;
                this.ratings = data.data;
            }
        },
        async onAction(rating, approve) {
            const action = approve ? "approve" : "deny";
            const response = await api.post(`/reviewRating/${rating.ID}/${action}`);
            if (response.status == 200)
                this.loadRatings();
        }
    }
}
</script>

<style scoped>
</style>
