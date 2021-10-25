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
            <ion-item>
                <ion-label position="stacked">Tags</ion-label>
                <div class="tag-container">
                    <ion-item v-for="(tag, idx) in allTags" :key="idx">
                        <ion-label>{{ tag }}</ion-label>
                        <ion-checkbox 
                            slot="start" 
                            @update:modelValue="updateTag(tag, $event)"
                            :modelValue="isTagChecked(tag)">
                        </ion-checkbox>
                    </ion-item>
                </div>
            </ion-item>
            <ion-button expand="block" color="primary" @click="doCreate" class="mx-3 mt-3" v-if="!isEditing">Create</ion-button>
            <ion-button expand="block" color="primary" @click="doSave" class="mx-3 mt-3" v-if="isEditing">Save</ion-button>
        </ion-content>

    </ion-page>
</template>

<script>
import { Http } from "@capacitor-community/http";
import vars from "../variables.ts";

import {
    toastController,
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonButtons,
    IonBackButton, IonTextarea, IonItem,
    IonLabel, IonSelect, IonSelectOption,
    IonButton, IonCheckbox
} from '@ionic/vue';

const tags = [
    "polite",
    "rude",
    "test",
    "example",
    "whatever"
];

export default {
    name: 'CreateRating',
    components: {
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonButtons,
        IonBackButton, IonTextarea, IonItem,
        IonLabel, IonSelect, IonSelectOption,
        IonButton, IonCheckbox
    },
    data() {
        return {
            allTags: tags,
            tags: [],
            comment: "",
            stars: "5"
        };
    },
    async created() {
        if (this.isEditing)
            await this.loadExistingRating();
    },
    methods: {
        async loadExistingRating() {

            const { id } = this.$route.params;
            const { token } = this.$store.state;
            const response = await Http.get({
                url: `${vars.backend}/person/${id}/rating`,
                headers: { Token: token },
            });

            const { data, status } = response;
            if (status == 200) {
                this.stars = data.data.stars.toString();
                this.comment = data.data.comment;
                this.tags = data.data.tags.split(",");
            } else {

                const toast = await toastController.create({
                    message: "An error occurred while loading your rating.",
                    duration: 3000,
                    position: "top",
                    color: "danger"
                });
                toast.present();

                this.$router.go(-1);

            }

        },
        async doSave() {

            const { id } = this.$route.params;
            const { token } = this.$store.state;
            const response = await Http.patch({
                url: `${vars.backend}/person/${id}/editRating`,
                headers: { 
                    "Token": token,
                    "content-type": "application/json"
                },
                data: {
                    stars: Number(this.stars),
                    comment: this.comment,
                    tags: this.tags.join()
                }
            });

            if (response.status == 200) {
                
                const toast = await toastController.create({
                    message: "Your rating has been saved.",
                    duration: 3000,
                    position: "top",
                    color: "success"
                });
                toast.present();
                
            } else {

                const toast = await toastController.create({
                    message: "An error occurred while updating your rating.",
                    duration: 3000,
                    position: "top",
                    color: "danger"
                });
                toast.present();
                
            }
            
            this.$router.go(-1);

        },
        async doCreate() {

            const { id } = this.$route.params;
            const { token } = this.$store.state;
            const response = await Http.post({
                url: `${vars.backend}/person/${id}/createRating`,
                headers: { 
                    "Token": token,
                    "content-type": "application/json"
                },
                data: {
                    stars: Number(this.stars),
                    comment: this.comment,
                    tags: this.tags.join()
                }
            });

            const { data, status } = response;
            if (status == 200) {
                
                const toast = await toastController.create({
                    message: "Your rating has been submitted.",
                    duration: 3000,
                    position: "top",
                    color: "success"
                });
                toast.present();
                
            } else {

                let message = "An error occurred while submitting your rating.";
                if (Object.prototype.hasOwnProperty.call(data, "error"))
                    message += "\n" + data.error;
                const toast = await toastController.create({
                    message,
                    duration: 3000,
                    position: "top",
                    color: "danger"
                });
                toast.present();
                
            }
            
            this.$router.go(-1);

        },
        updateTag(tag, checked) {
            if (checked) {
                if (!this.tags.includes(tag))
                    this.tags = this.tags.concat(tag);
            } else {
                this.tags = this.tags.filter(t => t != tag);
            }
        },
        isTagChecked(tag) {
            return this.tags.includes(tag);
        }
    },
    computed: {
        isEditing() {
            const { edit } = this.$route.query;
            return edit !== undefined;
        }
    }
}
</script>

<style scoped>
.tag-container {
    width: 100%;
    max-height: 210px;
    margin-top: 1rem;
    overflow-y: scroll;
}
</style>
