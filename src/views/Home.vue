<template>
    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-title>Home</ion-title>
            </ion-toolbar>
        </ion-header>

        <ion-content :fullscreen="true">
            <ion-card>
                <ion-card-header class="text-center">

                    <div class="flex justify-center pt-6 pb-8">
                        <img :src="wordmark" class="w-56" />
                    </div>

                    <ion-card-subtitle class="font-semibold mb-3">
                        Coming Soon
                    </ion-card-subtitle>

                </ion-card-header>
            </ion-card>
        </ion-content>

    </ion-page>
</template>

<script>
import {
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonCard,
    IonCardHeader, IonCardSubtitle
} from '@ionic/vue';

export default {
    name: 'Home',
    components: {
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonCard,
        IonCardHeader, IonCardSubtitle
    },
    computed: {
        wordmark() {
            return require("../assets/img/wordmark.png");
        }
    }
}
</script>
