<template>
    <ion-page>

        <ion-header>
            <ion-toolbar>
                <ion-title>Home</ion-title>
            </ion-toolbar>
        </ion-header>

        <ion-content :fullscreen="true">

            <ion-card>
                <ion-card-header class="text-center">

                    <div class="flex justify-center pt-6 pb-8">
                        <img :src="wordmark" class="w-56" />
                    </div>

                    <div>
                        <ion-button expand="block" color="primary" class="mx-3 mt-3 uppercase" @click="$router.push('/tabs/search/name')">
                            <ion-icon :icon="personSharp" class="mr-2" />
                            Search  by Name
                        </ion-button>
                        <ion-button expand="block" color="primary" class="mx-3 mt-3 uppercase" @click="$router.push('/tabs/search/address')">
                            <ion-icon :icon="home" class="mr-2" />
                            Search by Address
                        </ion-button>
                        <ion-button expand="block" color="primary" class="mx-3 mt-3 uppercase" @click="$router.push('/ratings')">
                            <ion-icon :icon="star" class="mr-2" />
                            My Ratings
                        </ion-button>
                    </div>

                </ion-card-header>
            </ion-card>


        </ion-content>

    </ion-page>
</template>

<script>
import {
    IonPage, IonHeader, IonToolbar,
    IonTitle, IonContent, IonCard,
    IonCardHeader, IonButton, IonIcon
} from '@ionic/vue';
import { personSharp, home, star } from 'ionicons/icons';

export default {
    name: 'Home',
    components: {
        IonPage, IonHeader, IonToolbar,
        IonTitle, IonContent, IonCard,
        IonCardHeader, IonButton, IonIcon
        
    },
    computed: {
        wordmark() {
            return require("../assets/img/wordmark.png");
        }
    },
    setup() {
        return { personSharp, home, star };
    }
}
</script>
