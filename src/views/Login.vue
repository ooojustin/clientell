<template>
    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-title>Login</ion-title>
            </ion-toolbar>
        </ion-header>
        
        <ion-content :fullscreen="true">
            
            <div class="flex justify-center pt-14 pb-16">
                <img :src="wordmark" class="w-56" />
            </div>

            <ion-item>
                <ion-label position="floating">Email Address</ion-label>
                <ion-input type="email" autocomplete="email" v-model="email" />
            </ion-item>
            <ion-item>
                <ion-label position="floating">Password</ion-label>
                <ion-input type="password" autocomplete="current-password" v-model="password" />
            </ion-item>
            <ion-button expand="block" color="primary" class="mx-3 mt-3" @click="doLogin">Submit</ion-button>
            
        </ion-content>

    </ion-page>
</template>

<script>
import { Http } from "@capacitor-community/http";
import vars from "../variables.ts";

import { 
    toastController,
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonInput,
    IonLabel, IonItem, IonButton 
} from '@ionic/vue';


export default {
    name: 'Login',
    components: { 
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonInput,
        IonLabel, IonItem, IonButton 
    },
    data() {
        return {
            email: "",
            password: ""
        };
    },
    computed: {
        wordmark() {
            return require("../assets/img/wordmark.png");
        }
    },
    methods: {
        async doLogin() {

            // send web request to log user in with credentials
            const { email, password } = this;
            const response = await Http.post({
                url: `${vars.backend}/login`,
                headers: { "content-type": "application/json" },
                data: { email, password }
            }); 

            const { data, status } = response;
            if (status == 200) {

                // show notification that the user was logged in successfully
                const toast = await toastController.create({
                    message: "Logged in successfully.",
                    duration: 3000,
                    position: "top",
                    color: "success"
                });
                toast.present();

                // log user into their account
                const user = data.data;
                this.$store.dispatch("loggedIn", user);

            } else {

                // failed to login, show error notification
                const toast = await toastController.create({
                    message: "Failed to login with those credentials.",
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
