<template>
    <ion-item v-if="!address">
        <ion-label position="floating">Address</ion-label>
        <ion-input type="text" v-model="query" />
    </ion-item>
    <ion-item v-if="address" @click="resetAddress">
        <ion-label position="floating">Address</ion-label>
        <ion-input type="text" v-model="formattedAddress" readOnly />
    </ion-item>
    <ion-card v-if="results">
        <ion-card-content>
            <ion-item v-for="result in results" :key="result.place_id" color="tertiary" class="addressItem" @click="clickAddress(result)">
                {{ result.formatted_address }}
            </ion-item>
        </ion-card-content>
    </ion-card>
</template>

<script>
import { Http } from "@capacitor-community/http";
import vars from "../variables.ts";

import { 
    IonItem, IonLabel, IonInput,
    IonCard, IonCardContent
} from '@ionic/vue';

import { Geolocation } from '@capacitor/geolocation';

export default {
    name: 'SelectAddress',
    emits: ["address-selected"],
    components: {
        IonItem, IonLabel, IonInput,
        IonCard, IonCardContent
    },
    data() {
        return {
            query: "",
            address: null,
            results: null,
            timeout: null,
            coords: null
        };
    },
    async created() {
        await this.getLocation();
    },
    methods: {
        clickAddress(result) {
            this.address = result;
            this.results = null;
            this.$emit("address-selected", result);
        },
        resetAddress() {
            this.query = "";
            this.address = null;
            this.results = null;
            this.$emit("address-selected", null);
        },
        async getLocation() {
            const response = await Geolocation.getCurrentPosition();        
            this.coords = response.coords;
        }
    },
    watch: {
        query() {

            if (this.timeout != null)
                clearTimeout(this.timeout);

            if (!this.query.length) {
                this.results = null;
                return;
            }

            const params = { query: this.query };
            if (this.coords !== null)
                params.coords = `${this.coords.latitude},${this.coords.longitude}`;

            this.timeout = setTimeout(async () => {

                const { token } = this.$store.state;
                const response = await Http.get({
                    url: `${vars.backend}/places`,
                    headers: { Token: token },
                    params
                });

                const { data, status } = response;
                if (status == 200) {
                    this.results = data.data;
                }

            }, 500);

        }
    },
    computed: {
        formattedAddress() {
            return this.address.formatted_address;
        }
    }
}
</script>

<style scoped>
.addressItem {
    border-radius: 4px; 
    font-size: 14px;
    padding-bottom: 4px;
}
</style>
