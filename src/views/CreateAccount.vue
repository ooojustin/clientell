<template>
    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-title>rateclients</ion-title>
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
            <ion-item>
                <ion-label position="floating">Email Address</ion-label>
                <ion-input type="email" autocomplete="email" v-model="email" />
            </ion-item>
            <ion-item>
                <ion-label position="floating">Password</ion-label>
                <ion-input type="password" autocomplete="new-password" v-model="password1" />
            </ion-item>
            <ion-item>
                <ion-label position="floating">Confirm Password</ion-label>
                <ion-input type="password" autocomplete="new-password" v-model="password2" />
            </ion-item>
            <ion-button expand="block" color="primary" class="mx-3 mt-3" @click="onSubmit">Submit</ion-button>
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
    IonItem, IonLabel, IonInput
} from '@ionic/vue';

export default {
    name: 'CreateAccount',
    components: {
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonButton,
        IonItem, IonLabel, IonInput
    },
    data() {
        return {
            email: "",
            firstName: "",
            lastName: "",
            password1: "",
            password2: ""
        };
    },
    methods: {
        async onSubmit() {
            
            // make sure password & confirm password match
            if (this.password1 != this.password2) {
                const toast = await toastController.create({
                    message: "Those passwords do not match.",
                    duration: 3000,
                    position: "top",
                    color: "danger"
                });
                toast.present();
                return;
            }

            // send web request to create account
            const response = await Http.post({
                url: `${vars.backend}/create_account`,
                data: {
                    firstName: this.firstName,
                    lastName: this.lastName,
                    email: this.email,
                    password: this.password1
                }
            });

            const { data, status } = response;
            if (status == 200) {

                // show notification that account was created
                const toast = await toastController.create({
                    message: "Account created successfully.",
                    duration: 3000,
                    position: "top",
                    color: "success"
                });
                toast.present();

                // log user into their new account
                const user = data.data;
                this.$store.dispatch("loggedIn", user);

            } else {

                // failed to create account, show error notification
                let message = "Failed to create account.";
                if (Object.prototype.hasOwnProperty.call(data, "error"))
                    message += "\n" + data.error;
                const toast = await toastController.create({
                    message,
                    duration: 3000,
                    position: "top",
                    color: "danger"
                });
                toast.present();

            }

        }
    }
}
</script>
