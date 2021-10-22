<template>
    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-title>Search by name</ion-title>
            </ion-toolbar>
        </ion-header>

        <ion-content :fullscreen="true">
            <ion-item>
                <ion-label position="floating">First Name</ion-label>
                <ion-input type="text" autocomplete="given-name" v-model="firstName" />
            </ion-item>
            <ion-item>
                <ion-label position="floating">Last Name</ion-label>
                <ion-input type="text" autocomplete="family-name" v-model="lastName" />
            </ion-item>
            <ion-button expand="block" color="primary" @click="doSearch" :disabled="disableSearch" class="mx-3 mt-3">Search</ion-button>
            <ion-button expand="block" color="success" router-link="/createPerson" class="mx-3 mt-3">Add Person</ion-button>

            <People :data="people" />

        </ion-content>

    </ion-page>
</template>

<script>
import { Http } from "@capacitor-community/http";
import vars from "../variables.ts";

import People from "../components/People.vue";

import {
    toastController,
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonItem,
    IonLabel, IonInput, IonButton
} from '@ionic/vue';

export default {
    name: 'SearchByName',
    components: {
        People,
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonItem,
        IonLabel, IonInput, IonButton
    },
    data() {
        return {
            people: null,
            firstName: "",
            lastName: ""
        }
    },
    methods: {
        async doSearch() {

            const { token } = this.$store.state;
            const response = await Http.post({
                url: `${vars.backend}/person/search`,
                headers: { Token: token },
                data: {
                    firstName: this.firstName,
                    lastName: this.lastName
                }
            });

            const { data, status } = response;
            if (status == 200) {
                this.people = data.data;
            } else {

                // show alert that search failed, for an unknown reason
                const toast = await toastController.create({
                    message: "Failed to submit search.",
                    duration: 3000,
                    position: "top",
                    color: "danger"
                });
                toast.present();

            }

        }
    },
    computed: {
        disableSearch() {
            return this.firstName.length < 3 || this.lastName.length < 3;
        }
    }
}
</script>
