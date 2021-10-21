<template>
    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-title>Search by address</ion-title>
            </ion-toolbar>
        </ion-header>

        <ion-content :fullscreen="true">

            <SelectAddress @address-selected="setAddress" />
            <ion-button expand="block" color="primary" @click="doSearch" :disabled="disableSearch" class="mx-3 mt-3">Search</ion-button>
            <ion-button expand="block" color="success" router-link="/createPerson" class="mx-3 mt-3">Add Person</ion-button>

            <People :data="people" />

        </ion-content>

    </ion-page>
</template>

<script>
import { Http } from "@capacitor-community/http";
import vars from "../variables.ts";

import SelectAddress from '../components/SelectAddress.vue';
import People from "../components/People.vue";

import {
    toastController,
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonButton
} from '@ionic/vue';


export default {
    name: 'SearchByAddress',
    components: {
        SelectAddress, People,
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonButton
    },
    data() {
        return {
            people: null,
            address: null
        };
    },
    methods: {
        async doSearch() {

            const { token } = this.$store.state;

            const response = await Http.post({
                url: `${vars.backend}/person/search`,
                headers: { Token: token },
                data: {
                    address: this.address
                }
            });

            const { data, status } = response;
            if (status == 200) {
                this.people = data.data;
            } else {

                // show alert that search failed, for an unknown reason
                const toast = await toastController.create({
                    message: "Failed to submit search.",
                    duration: 3000,
                    position: "top",
                    color: "danger"
                });
                toast.present();

            }

        },
        setAddress(address) {
            this.address = address;
        }
    },
    computed: {
        disableSearch() {
            return this.address == null;
        }
    }
}
</script>
