<template>
    <ion-card :class="{ 'user-rating': isUserRating  }">
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
                <br />
                <br />
                <ion-chip v-for="(tag, idx) in tags" :key="idx">
                    <ion-label outline>{{ tag }}</ion-label>
                </ion-chip>
            </span>
        </ion-card-content>
    </ion-card>
</template>

<script>
import {
    IonCard, IonCardContent, IonLabel,
    IonChip
} from '@ionic/vue';

export default {
    name: 'Rating',
    components: {
        IonCard, IonCardContent, IonLabel,
        IonChip
    },
    props: {
        data: Object
    },
    computed: {
        isUserRating() {
            const { user } = this.$store.state;
            return this.data.ownerID == user.ID;
        },
        tags() {
            return this.data.tags.split(",").filter(i => i.length > 0);
        }
    }
}
</script>

<style scoped>
.user-rating {
    border-color: var(--ion-color-primary-shade);
    background-color: var(--ion-color-primary);
    color: var(--ion-color-primary-contrast);
}
</style>
