<template>
    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-buttons slot="start">
                    <ion-back-button></ion-back-button>
                </ion-buttons>
                <ion-title>Create Person</ion-title>
            </ion-toolbar>
        </ion-header>

        <ion-content :fullscreen="true">
            <ion-item>
                <ion-label position="floating">First Name</ion-label>
                <ion-input type="text" v-model="firstName" />
            </ion-item>
            <ion-item>
                <ion-label position="floating">Last Name</ion-label>
                <ion-input type="text" v-model="lastName" />
            </ion-item>
            <ion-button expand="block" color="primary" @click="doCreate" class="mx-3 mt-3">Create</ion-button>
        </ion-content>

    </ion-page>
</template>

<script>
import { Http } from "@capacitor-community/http";
import vars from "../variables.ts";

import {
    toastController,
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonButtons,
    IonBackButton, IonItem, IonLabel,
    IonInput, IonButton
} from '@ionic/vue';

export default {
    name: 'Home',
    components: {
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonButtons,
        IonBackButton, IonItem, IonLabel,
        IonInput, IonButton
    },
    data() {
        return {
            firstName: "",
            lastName: ""
        };
    },
    methods: {
        async doCreate() {
            
            const { token } = this.$store.state;
            const response = await Http.post({
                url: `${vars.backend}/person/create`,
                headers: { Token: token },
                data: {
                    firstName: this.firstName,
                    lastName: this.lastName
                }
            });

            const { data, status } = response;
            if (status == 200) {

                // show alert that a new person has been created
                const toast = await toastController.create({
                    message: "New person has been created.",
                    duration: 3000,
                    position: "top",
                    color: "success"
                });
                toast.present();

            } else {

                // show alert that we failed to create the person
                const toast = await toastController.create({
                    message: "Failed to create new person.",
                    duration: 3000,
                    position: "top",
                    color: "danger"
                });
                toast.present();

            }
    
            // go back to previous route
            this.$router.go(-1);

        }
    }
}
</script>
