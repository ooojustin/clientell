<template>
    <ion-card @click="openPerson">
        <ion-card-header>
            <ion-card-title>
                {{ data.firstName }} {{ data.lastName }}
            </ion-card-title>
        </ion-card-header>
        <ion-card-content>
            {{ data.address.formatted_address }}
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
    },
    methods: {
        openPerson() {
            this.$router.push("/person/" + this.data.ID);
        }
    }
}
</script>

<style scoped>
</style>
