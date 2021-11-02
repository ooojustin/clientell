<template>
    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-buttons slot="start">
                    <ion-back-button></ion-back-button>
                </ion-buttons>
                <ion-title>Ratings</ion-title>
            </ion-toolbar>
        </ion-header>

        <ion-content :fullscreen="true">
            <div v-if="ratings">
                <div v-for="(rating, idx) in ratings" :key="idx">
                    <Rating :data="rating" />
                </div>
            </div>

        </ion-content>

    </ion-page>
</template>

<script>
import { Http } from "@capacitor-community/http";
import vars from "../variables.ts";

import Rating from "../components/Rating.vue";

import {
    toastController, loadingController,
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonButtons,
    IonBackButton
} from '@ionic/vue';

export default {
    name: 'UserRatings',
    components: {
        Rating,
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonButtons,
        IonBackButton
    },
    data() {
        return {
            ratings: null
        };
    },
    async created() {
        await this.loadData();
    },
    methods: {
        async loadData() {

            // only run code when user ratings tab is current path
            if (this.$router.currentRoute._value.path != "/ratings")
                return;

            // present loading overlay
            const loading = await loadingController.create({ message: 'Loading ratings...' });
            await loading.present();

            // load list of ratings
            const { token } = this.$store.state;
            const response = await Http.get({
                url: `${vars.backend}/listRatings`,
                headers: { Token: token }
            });
            
            // hide loading overlay
            loading.dismiss();

            // handle response data
            const { status, data } = response;
            if (status == 200) {
                this.ratings = data.data;
            } else {

                const toast = await toastController.create({
                    message: "An error occurred while loading your ratings.",
                    duration: 3000,
                    position: "top",
                    color: "danger"
                });
                toast.present();

                this.$router.go(-1);

            }

        }
    }
}
</script>

<style scoped>
</style>
