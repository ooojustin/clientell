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

            <div>
                <ion-button expand="block" color="primary" class="mx-3 mt-3 uppercase" @click="$router.push('/tabs/search/name')">
                    Search  by Name
                </ion-button>
                <ion-button expand="block" color="primary" class="mx-3 mt-3 uppercase" @click="$router.push('/tabs/search/address')">
                    Search by Address
                </ion-button>
                <ion-button expand="block" color="primary" class="mx-3 mt-3 uppercase" @click="$router.push('/ratings')">
                    My Reviews
                </ion-button>
            </div>

        </ion-content>

    </ion-page>
</template>

<script>
import {
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonCard,
    IonCardHeader, IonCardSubtitle, IonButton
} from '@ionic/vue';

export default {
    name: 'Home',
    components: {
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonCard,
        IonCardHeader, IonCardSubtitle, IonButton
    },
    computed: {
        wordmark() {
            return require("../assets/img/wordmark.png");
        }
    }
}
</script>
