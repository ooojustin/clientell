<template>
    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-title>Search</ion-title>
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
    name: 'Search',
    components: {
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent
    }
}
</script>
