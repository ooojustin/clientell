<template>
    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-title>Home</ion-title>
            </ion-toolbar>
        </ion-header>

        <ion-content :fullscreen="true">
            <ion-card>
                <ion-card-header>
                    <ion-card-title>rateclients</ion-card-title>
                    <ion-card-subtitle>coming soon...</ion-card-subtitle>
                </ion-card-header>
            </ion-card>
        </ion-content>

    </ion-page>
</template>

<script>
import {
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonCard,
    IonCardHeader, IonCardTitle, IonCardSubtitle
} from '@ionic/vue';

export default {
    name: 'Home',
    components: {
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonCard,
        IonCardHeader, IonCardTitle, IonCardSubtitle
    }
}
</script>
