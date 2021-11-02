<template>
    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-buttons slot="start">
                    <ion-back-button></ion-back-button>
                </ion-buttons>
                <ion-title>Person</ion-title>
            </ion-toolbar>
        </ion-header>

        <ion-content v-if="person" :fullscreen="true">
        
            <div class="w-100 text-center text-2xl mt-5">
                {{ person.firstName }} {{ person.lastName }}
            </div>
            <div class="flex justify-center">
                <div class="w-60 text-center mt-2">
                    {{ person.address.formatted_address }}
                </div>
            </div>

            <div class="mt-6">
                <GoogleMap :placeId="person.address.place_id" width="90%" height="150px" />
            </div>

            <ion-button expand="block" color="success" class="mx-3 mt-6" v-if="canRate" @click="openCreateRating">
                Add Rating
            </ion-button>
            <span v-if="!canRate">
                <ion-button expand="block" color="warning" class="mx-3 mt-6" @click="openEditRating">
                    Edit Rating
                </ion-button>
                <ion-button expand="block" color="danger" class="mx-3 mt-3" @click="deleteRating">
                    Delete Rating
                </ion-button>
            </span>
            
            <div v-if="ratings" class="mt-6">
                <div class="ml-3 mb-2">
                    {{ ratings.length }} total ratings
                    <br />
                    Average Stars: {{ person.avgStars }}
                </div>
                <Rating v-for="rating in ratings" :data="rating" @update-rating="updateRating" :key="rating.ID" />
            </div>

        </ion-content>

    </ion-page>
</template>

<script>
import { Http } from "@capacitor-community/http";
import vars from "../variables.ts";

import Rating from "../components/Rating.vue";
import GoogleMap from "../components/GoogleMap.vue";

import {
    toastController,
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonButtons,
    IonBackButton, IonButton
} from '@ionic/vue';

export default {
    name: 'Person',
    components: {
        Rating, GoogleMap,
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonButtons,
        IonBackButton, IonButton
    },
    data() {
        return {
            person: null,
            ratings: null,
            canRate: false
        };
    },
    async created() {
        await this.loadData();
    },
    methods: {
        async loadData() {

            const { id } = this.$route.params;
            const { token } = this.$store.state;
            if (!id) // prevent bad request from firing
                return;

            // send request to load person data (includes ratings)
            const response = await Http.get({
                url: `${vars.backend}/person/` + id,
                headers: { Token: token }
            });

            const { data, status } = response;
            if (status == 200) {

                // update data in component state
                this.person = data.data.person;
                this.ratings = data.data.ratings;
                this.canRate = data.data.canRate;

            } else {
                
                // alert user that we couldnt load data about this person
                const toast = await toastController.create({
                    message: "Failed to load person.",
                    duration: 3000,
                    position: "top",
                    color: "danger"
                });
                toast.present();

                // go back to previous route
                this.$router.go(-1);

            }

        },
        async deleteRating() {

            const { id } = this.$route.params;
            const { token } = this.$store.state;
            if (!id) // prevent bad request from firing
                return;

            // send request to delete rating
            const response = await Http.post({
                url: `${vars.backend}/person/${id}/deleteRating`,
                headers: { Token: token }
            });

            if (response.status == 200) {

                const toast = await toastController.create({
                    message: "Your rating has been removed.",
                    duration: 3000,
                    position: "top",
                    color: "success"
                });
                toast.present();
                
                await this.loadData();

            } else {

                const toast = await toastController.create({
                    message: "Failed to load delete your rating.",
                    duration: 3000,
                    position: "top",
                    color: "danger"
                });
                toast.present();

            }

        },
        openCreateRating() {
            this.$router.push("/createRating/" + this.$route.params.id);
        },
        openEditRating() {
            this.$router.push("/createRating/" + this.$route.params.id + "?edit=true");
        },
        updateRating(data) {
            this.ratings = this.ratings.map(r => r.ID == data.ID ? data : r);
        }
    },
    watch: {
        async $route (to, from) {
            if (to.path.startsWith("/person/") && from.path.startsWith("/createRating/"))
                await this.loadData();
        }
    }
}
</script>
