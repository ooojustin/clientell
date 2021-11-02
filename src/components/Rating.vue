<template>
    <ion-card :class="{ 'user-rating': isUserRating  }">
        <ion-card-content>
            <span>
                <b>Job:</b> {{ data.jobType }}
            </span>
            <br />
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
            <br />
            <span>
                <ion-chip v-if="data.thumbs_up > 0" :color="data.reaction == 'thumbs_up' ? 'primary' : 'medium'">
                    <ion-label>&#128077;&nbsp;{{ data.thumbs_up }}</ion-label>
                </ion-chip>
                <ion-chip v-if="data.thumbs_down > 0" :color="data.reaction == 'thumbs_down' ? 'primary' : 'medium'">
                    <ion-label>&#128078;&nbsp;{{ data.thumbs_down }}</ion-label>
                </ion-chip>
                <ion-chip v-if="data.funny > 0" :color="data.reaction == 'funny' ? 'primary' : 'medium'">
                    <ion-label>&#128514;&nbsp;{{ data.funny }}</ion-label>
                </ion-chip>
                <ion-chip v-if="data.fire > 0" :color="data.reaction == 'fire' ? 'primary' : 'medium'">
                    <ion-label>&#128293;&nbsp;{{ data.fire }}</ion-label>
                </ion-chip>
                <ion-chip v-if="data.heart > 0" :color="data.reaction == 'heart' ? 'primary' : 'medium'">
                    <ion-label>&#10084;&nbsp;{{ data.heart }}</ion-label>
                </ion-chip>
            </span>

            <ion-fab horizontal="end" vertical="bottom" slot="fixed" v-if="!isUserRating && !data.reaction">

                <ion-fab-button color="primary" size="small">
                    <ion-icon :icon="add"></ion-icon>
                </ion-fab-button>

                <ion-fab-list side="top">
                    <ion-fab-button color="medium" @click="addReaction('thumbs_down')">
                        <ion-label>&#128078;</ion-label>
                    </ion-fab-button>
                    <ion-fab-button color="medium" @click="addReaction('thumbs_up')">
                        <ion-label>&#128077;</ion-label>
                    </ion-fab-button>
                </ion-fab-list>

                <ion-fab-list side="start">
                    <ion-fab-button color="medium" @click="addReaction('funny')">
                        <ion-label>&#128514;</ion-label>
                    </ion-fab-button>
                    <ion-fab-button color="medium" @click="addReaction('fire')">
                        <ion-label>&#128293;</ion-label>
                    </ion-fab-button>
                    <ion-fab-button color="medium" @click="addReaction('heart')">
                        <ion-label>&#10084;</ion-label>
                    </ion-fab-button>
                </ion-fab-list>

            </ion-fab>

        </ion-card-content>
    </ion-card>
</template>

<script>
import { Http } from "@capacitor-community/http";
import vars from "../variables.ts";

import {
    toastController,
    IonCard, IonCardContent, IonLabel,
    IonChip, IonFab, IonFabButton,
    IonFabList, IonIcon
} from '@ionic/vue';
import { add } from 'ionicons/icons';

export default {
    name: 'Rating',
    components: {
        IonCard, IonCardContent, IonLabel,
        IonChip, IonFab, IonFabButton,
        IonFabList, IonIcon
    },
    props: {
        data: Object
    },
    methods: {
        async addReaction(reaction) {

            // send web request to create reaction
            const { token } = this.$store.state;
            const response = await Http.post({
                url: `${vars.backend}/rating/${this.data.ID}/react`,
                headers: { Token: token },
                data: { type: reaction }
            });

            // handle success
            const { status, data } = response;
            if (status == 200) {

                // notify user that reaction was left
                const toast = await toastController.create({
                    message: "Your reaction has been recorded.",
                    duration: 3000,
                    position: "top",
                    color: "success"
                });
                toast.present();

                // tell parant component to update rating in the list
                this.$emit("update-rating", data.data.rating);

            }

        }
    },
    computed: {
        isUserRating() {
            const { user } = this.$store.state;
            return this.data.ownerID == user.ID;
        },
        tags() {
            return this.data.tags.split(",").filter(i => i.length > 0);
        }
    },
    setup() {
        return { add };
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
