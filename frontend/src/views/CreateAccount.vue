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
                        <ion-card-title>Create an Account</ion-card-title>
                    </ion-card-header>
                    <ion-card-content>
                        <ion-item class="mb-2">
                            <ion-label position="floating">Email Address</ion-label>
                            <ion-input type="email" autocomplete="email" v-model="email" />
                        </ion-item>
                        <ion-item class="mb-4">
                            <ion-label position="floating">Password</ion-label>
                            <ion-input type="password" autocomplete="new-password" v-model="password1" />
                        </ion-item>
                        <ion-item class="mb-4">
                            <ion-label position="floating">Confirm Password</ion-label>
                            <ion-input type="password" autocomplete="new-password" v-model="password2" />
                        </ion-item>
                        <ion-button expand="block" color="primary" class="mb-2" @click="onSubmit">Submit</ion-button>
                        <ion-button expand="block" color="medium" router-link="/login">Back to Login</ion-button>
                    </ion-card-content>
                </ion-card>
            </div>
        </ion-content>

    </ion-page>
</template>

<script>
import {
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonCard,
    IonCardHeader, IonCardTitle, IonCardContent,
    IonItem, IonLabel, IonInput,
    IonButton
} from '@ionic/vue';

export default {
    name: 'CreateAccount',
    components: {
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonCard,
        IonCardHeader, IonCardTitle, IonCardContent,
        IonItem, IonLabel, IonInput,
        IonButton
    },
    data() {
        return {
            email: "",
            password1: "",
            password2: ""
        };
    },
    methods: {
        onSubmit() {
            console.log(this.email, this.password1, this.password2);
        }
    }
}
</script>
