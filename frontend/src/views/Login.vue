<template>
    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-title>rateclients</ion-title>
            </ion-toolbar>
        </ion-header>
        
        <ion-content :fullscreen="true">

            <ion-card>
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

        </ion-content>

  </ion-page>
</template>

<script>
import { 
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonInput,
    IonLabel, IonItem, IonCard,
    IonCardHeader, IonCardContent, IonCardTitle,
    IonButton
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
    methods: {
        printCreds() {
            console.log(this.email, this.password);
        }
    }
}
</script>
