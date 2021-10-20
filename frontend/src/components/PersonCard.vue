<template>
    <ion-card>
        <ion-card-header>
            <ion-card-title>
                {{ data.firstName }} {{ data.lastName }}
            </ion-card-title>
        </ion-card-header>
        <ion-card-content>
            TODO
        </ion-card-content>
    </ion-card>
</template>

<script>
import {
    IonCard, IonCardHeader, IonCardTitle,
    IonCardContent
} from '@ionic/vue';

export default {
    name: 'PersonCard',
    props: {
        data: Object
    },
    components: {
        IonCard, IonCardHeader, IonCardTitle,
        IonCardContent
    }
}
</script>

<style scoped>
</style>
