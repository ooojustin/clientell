<template>

    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-buttons slot="start">
                    <ion-back-button></ion-back-button>
                </ion-buttons>
                <ion-title>FAQ</ion-title>
            </ion-toolbar>
        </ion-header>


        <ion-content>
            <div>
                <ion-list>
                    <ion-item v-for="(qa, idx) in questions" :key="idx">
                        <ion-label text-wrap>
                            <h3>Question #{{ idx + 1 }}</h3>
                            <h2 class="question">{{ qa.question }}</h2>
                            <h3>{{ qa.answer }}</h3>
                        </ion-label>
                    </ion-item>
                </ion-list>
            </div>
        </ion-content>

    </ion-page>

</template>

<script>
import {
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonButtons,
    IonBackButton, IonList, IonItem,
    IonLabel
} from '@ionic/vue';

const questions = [
    {
        question: "Lorem Ipsum",
        answer: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam efficitur risus vel turpis auctor, vel sollicitudin leo facilisis."
    },
    {
        question: "Lorem Ipsum",
        answer: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam efficitur risus vel turpis auctor, vel sollicitudin leo facilisis."
    },
    {
        question: "Lorem Ipsum",
        answer: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam efficitur risus vel turpis auctor, vel sollicitudin leo facilisis."
    },
];

export default {
    name: 'FAQ',
    components: {
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonButtons,
        IonBackButton, IonList, IonItem,
        IonLabel
    },
    data() {
        return {
            questions
        };
    }
}
</script>

<style scoped>
h2.question {
    font-weight: 700;
}
</style>
