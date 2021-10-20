<template>
    <ion-page>
            
        <ion-header>
            <ion-toolbar>
                <ion-title>Settings</ion-title>
            </ion-toolbar>
        </ion-header>

        <ion-content :fullscreen="true">
            <!--TODO-->
            <ion-button expand="block" color="danger" @click="doLogout">
                Logout
            </ion-button>
        </ion-content>

    </ion-page>
</template>

<script>
import { Http } from "@capacitor-community/http";
import vars from "../variables.ts";

import {
    toastController,
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonButton
} from '@ionic/vue';

export default {
    name: 'Settings',
    components: {
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonButton
    },
    methods: {
        async doLogout() {

            // send web request to logout (invalidate current token)
            const { token } = this.$store.state;
            const response = await Http.post({ 
                url: `${vars.backend}/logout`,
                headers: { Token: token }
            });

            // show alert if token was reset in backend
            const { data, status } = response;
            if (status == 200) {
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

        }
    }
}
</script>
