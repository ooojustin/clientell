<template>
    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-title>Search by name</ion-title>
            </ion-toolbar>
        </ion-header>

        <ion-content :fullscreen="true">
            <ion-item>
                <ion-label position="floating">First Name</ion-label>
                <ion-input type="text" autocomplete="given-name" v-model="firstName" />
            </ion-item>
            <ion-item>
                <ion-label position="floating">Last Name</ion-label>
                <ion-input type="text" autocomplete="family-name" v-model="lastName" />
            </ion-item>
            <ion-button expand="block" color="primary" @click="doSearch" class="mx-3 mt-3">Search</ion-button>
        </ion-content>

    </ion-page>
</template>

<script>
import {
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonItem,
    IonLabel, IonInput, IonButton
} from '@ionic/vue';

export default {
    name: 'SearchByName',
    components: {
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonItem,
        IonLabel, IonInput, IonButton
    },
    data() {
        return {
            firstName: "",
            lastName: ""
        }
    },
    methods: {
        async doSearch() {
            // todo
            console.log(this.firstName, this.lastName);
        }
    }
}
</script>
