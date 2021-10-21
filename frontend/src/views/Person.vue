<template>
    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-buttons slot="start">
                    <ion-back-button></ion-back-button>
                </ion-buttons>
                <ion-title>Person</ion-title>
            </ion-toolbar>
        </ion-header>

        <ion-content :fullscreen="true">
            <!-- TODO -->
        </ion-content>

    </ion-page>
</template>

<script>
import {
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonButtons,
    IonBackButton
} from '@ionic/vue';

export default {
    name: 'Person',
    components: {
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonButtons,
        IonBackButton
    },
    created() {
        console.log("ID", this.$route.params.id);
    }
}
</script>
