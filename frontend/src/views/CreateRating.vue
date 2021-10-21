<template>
    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-buttons slot="start">
                    <ion-back-button></ion-back-button>
                </ion-buttons>
                <ion-title>Add Rating</ion-title>
            </ion-toolbar>
        </ion-header>

        <ion-content :fullscreen="true">
            <ion-item>
                <ion-label position="floating">Stars</ion-label>
                <ion-select v-model="stars" interface="action-sheet">
                    <ion-select-option value="1">1 &#11088;</ion-select-option>
                    <ion-select-option value="2">2 &#11088;&#11088;</ion-select-option>
                    <ion-select-option value="3">3 &#11088;&#11088;&#11088;</ion-select-option>
                    <ion-select-option value="4">4 &#11088;&#11088;&#11088;&#11088;</ion-select-option>
                    <ion-select-option value="5">5 &#11088;&#11088;&#11088;&#11088;&#11088;</ion-select-option>
                </ion-select>
            </ion-item>
            <ion-item>
                <ion-label position="floating">Comment</ion-label>
                <ion-textarea v-model="comment"></ion-textarea>
            </ion-item>
            <ion-button expand="block" color="primary" @click="doCreate" class="mx-3 mt-3">Create</ion-button>
        </ion-content>

    </ion-page>
</template>

<script>
import {
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonButtons,
    IonBackButton, IonTextarea, IonItem,
    IonLabel, IonSelect, IonSelectOption,
    IonButton
} from '@ionic/vue';

export default {
    name: 'CreateRating',
    components: {
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonButtons,
        IonBackButton, IonTextarea, IonItem,
        IonLabel, IonSelect, IonSelectOption,
        IonButton
    },
    data() {
        return {
            comment: "",
            stars: "5"
        };
    },
    methods: {
        async doCreate() {
            return;
        }
    }
}
</script>
