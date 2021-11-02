<template>
    <ion-page>
            
        <ion-header>
            <ion-toolbar>
                <ion-title>Settings</ion-title>
            </ion-toolbar>
        </ion-header>

        <ion-content :fullscreen="true">
            <ion-item>
                <ion-label position="floating">Email</ion-label>
                <ion-input type="email" v-model="email" readOnly />
            </ion-item>
            <ion-item>
                <ion-label position="floating">First Name</ion-label>
                <ion-input type="text" autocomplete="given-name" v-model="firstName" />
            </ion-item>
            <ion-item>
                <ion-label position="floating">Last Name</ion-label>
                <ion-input type="text" autocomplete="family-name" v-model="lastName" />
            </ion-item>
            <ion-button expand="block" color="primary" @click="doSave" class="mx-3 mt-3">Save</ion-button>
            <ion-button expand="block" color="danger" @click="doLogout" class="mx-3 mt-3">Logout</ion-button>
            <ion-item button @click="openFAQ()" class="mt-3">
                <ion-label>Frequently Asked Questions</ion-label>
                <ion-icon :icon="helpCircle" slot="end"></ion-icon>
            </ion-item>
        </ion-content>

    </ion-page>
</template>

<script>
import { Http } from "@capacitor-community/http";
import vars from "../variables.ts";

import {
    toastController,
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonButton,
    IonItem, IonLabel, IonInput,
    IonIcon
} from '@ionic/vue';
import { helpCircle } from 'ionicons/icons';

export default {
    name: 'Settings',
    components: {
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonButton,
        IonItem, IonLabel, IonInput,
        IonIcon
    },
    data() {
        return {
            email: this.$store.state.user.email,
            firstName: this.$store.state.user.firstName,
            lastName: this.$store.state.user.lastName
        };
    },
    methods: {
        async doSave() {

            const { token } = this.$store.state;
            const response = await Http.patch({
                url: `${vars.backend}/user`,
                headers: { 
                    "Token": token,
                    "content-type": "application/json"
                },
                data: {
                    firstName: this.firstName,
                    lastName: this.lastName
                }
            });

            const { data, status } = response;
            if (status == 200) {

                // show notification that the profile was updated successfully
                const toast = await toastController.create({
                    message: "Information updated successfully.",
                    duration: 3000,
                    position: "top",
                    color: "success"
                });
                toast.present();

                // update user data in global state
                const user = data.data;
                this.$store.dispatch("loggedIn", user);

            } else {

                // failed to update profile, show error notification
                const toast = await toastController.create({
                    message: "Failed to update your information.",
                    duration: 3000,
                    position: "top",
                    color: "danger"
                });
                toast.present();

            }

        },
        async doLogout() {

            // send web request to logout (invalidate current token)
            const { token } = this.$store.state;
            const response = await Http.post({ 
                url: `${vars.backend}/logout`,
                headers: { Token: token }
            });

            // show alert if token was reset in backend
            if (response.status == 200) {
                const toast = await toastController.create({
                    message: "Logged out successfully.",
                    duration: 3000,
                    position: "top",
                    color: "success"
                });
                toast.present();
            }

            // logout locally
            this.$store.dispatch("logout");
            this.$router.push("/auth/login");

        },
        openFAQ() {
            this.$router.push("/faq");
        }
    },
    setup() {
        return { helpCircle };
    }
}
</script>
