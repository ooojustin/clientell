<template>
    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-title>rateclients</ion-title>
            </ion-toolbar>
        </ion-header>
        
        <ion-content :fullscreen="true">
            <div class="h-full flex items-center">
                <ion-card class="w-full" style="margin-bottom: 10rem;">
                    <ion-card-header>
                        <ion-card-title>Login</ion-card-title>
                    </ion-card-header>
                    <ion-card-content>
                        <ion-item>
                            <ion-label position="floating">Email Address</ion-label>
                            <ion-input type="email" autocomplete="email" v-model="email" />
                        </ion-item>
                        <ion-item style="margin-top: 10px; margin-bottom: 10px;">
                            <ion-label position="floating">Password</ion-label>
                            <ion-input type="password" autocomplete="current-password" v-model="password" />
                        </ion-item>
                        <ion-button expand="block" color="primary" @click="printCreds">Submit</ion-button>
                    </ion-card-content>
                </ion-card>
            </div>
        </ion-content>

  </ion-page>
</template>

<script>
import { Http } from "@capacitor-community/http";
import { mapGetters } from "vuex";
import vars from "../variables.ts";

import { 
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonInput,
    IonLabel, IonItem, IonCard,
    IonCardHeader, IonCardContent, IonCardTitle,
    IonButton, toastController
} from '@ionic/vue';


export default {
    name: 'Login',
    components: { 
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonInput,
        IonLabel, IonItem, IonCard,
        IonCardHeader, IonCardContent, IonCardTitle,
        IonButton
    },
    data() {
        return {
            email: "",
            password: ""
        };
    },
    computed: {
        ...mapGetters(["isAuthenticated"])
    },
    watch: {
        isAuthenticated(val) {
            if (val)
                this.$router.push("/tabs/");
        }
    },
    mounted() {
        this.$store.dispatch("restoreLogin");
    },
    methods: {
        async printCreds() {

            const { email, password } = this;
            const response = await Http.post({
                url: `${vars.backend}/login`,
                data: { email, password }
            }); 

            const { data, status } = response;
            if (status == 200) {

                const toast = await toastController.create({
                    message: "Logged in successfully.",
                    duration: 3000,
                    position: "top",
                    color: "success"
                });
                toast.present();

                const user = data.data;
                this.$store.dispatch("loggedIn", user);

            } else {

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
