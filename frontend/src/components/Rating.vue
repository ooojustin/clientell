<template>
    <ion-card>
        <ion-card-content>
            <span>
                Stars: {{ data.stars }}
            </span>
            <span v-if="data.comment.length > 0">
                Comment:
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
