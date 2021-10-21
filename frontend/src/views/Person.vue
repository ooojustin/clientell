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
            <ion-button expand="block" color="success" class="mx-3 mt-6" @click="openCreateRating">
                Add Rating
            </ion-button>

            <div v-if="ratings" class="mt-6">
                <Rating v-for="rating in ratings" :data="rating" :key="rating.ID" />
            </div>

        </ion-content>

    </ion-page>
</template>

<script>
import { Http } from "@capacitor-community/http";
import vars from "../variables.ts";

import Rating from "../components/Rating.vue";

import {
    toastController,
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonButtons,
    IonBackButton, IonButton
} from '@ionic/vue';

export default {
    name: 'Person',
    components: {
        Rating,
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonButtons,
        IonBackButton, IonButton
    },
    data() {
        return {
            person: null,
            ratings: null
        };
    },
    async created() {
        await this.loadData();
    },
    methods: {
        async loadData() {

            const { id } = this.$route.params;
            const { token } = this.$store.state;

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
        openCreateRating() {
            this.$router.push("/createRating/" + this.$route.params.id);
        }
    }
}
</script>
