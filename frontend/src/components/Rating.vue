<template>
    <ion-card>
        <ion-card-content>
            <span>
                <b>Stars:</b> {{ data.stars }}
            </span>
            <br />
            <br />
            <span v-if="data.comment.length > 0">
                <b>Comment:</b>
                <br />
                {{ data.comment }}
            </span>
        </ion-card-content>
    </ion-card>
</template>

<script>
import {
    IonCard, IonCardContent
} from '@ionic/vue';

export default {
    name: 'Rating',
    components: {
        IonCard, IonCardContent
    },
    props: {
        data: Object
    }
}
</script>

<style scoped>
</style>
