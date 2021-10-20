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
                        <ion-item class="mb-2">
                            <ion-label position="floating">Email Address</ion-label>
                            <ion-input type="email" autocomplete="email" v-model="email" />
                        </ion-item>
                        <ion-item class="mb-4">
                            <ion-label position="floating">Password</ion-label>
                            <ion-input type="password" autocomplete="current-password" v-model="password" />
                        </ion-item>
                        <ion-button expand="block" color="primary" class="mb-2" @click="doLogin">Submit</ion-button>
                        <ion-button expand="block" color="medium" router-link="/create_account">Create Account</ion-button>
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
            // automatically redirect user when logged in
            if (val)
                this.$router.push("/tabs/");
        }
    },
    mounted() {
        this.$store.dispatch("restoreLogin");
    },
    methods: {
        async doLogin() {

            // send web request to log user in with credentials
            const { email, password } = this;
            const response = await Http.post({
                url: `${vars.backend}/login`,
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
