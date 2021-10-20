<template>
    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-title>Search by address</ion-title>
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
    IonTitle, IonContent
} from '@ionic/vue';

export default {
    name: 'SearchByAddress',
    components: {
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent
    }
}
</script>
