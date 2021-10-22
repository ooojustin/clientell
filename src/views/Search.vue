<template>
    <ion-page>
        <ion-tabs>
            <ion-router-outlet></ion-router-outlet>
            <ion-tab-bar slot="top">
                
                <ion-tab-button tab="searchName" href="/tabs/search/name">
                    <ion-icon :icon="personSharp" />
                    <ion-label>By Name</ion-label>
                </ion-tab-button>

                <ion-tab-button tab="searchAddress" href="/tabs/search/address">
                    <ion-icon :icon="home" />
                    <ion-label>By Address</ion-label>
                </ion-tab-button>

            </ion-tab-bar>
        </ion-tabs>
    </ion-page>
</template>

<script>
import {
    IonPage, IonTabs, IonTabBar, 
    IonTabButton, IonIcon, IonLabel, 
    IonRouterOutlet
} from '@ionic/vue';
import { personSharp, home } from 'ionicons/icons';

export default {
    name: 'Search',
    components: {
        IonPage, IonTabs, IonTabBar, 
        IonTabButton, IonIcon, IonLabel, 
        IonRouterOutlet
    },
    setup() {
        return { personSharp, home };
    }
}
</script>
